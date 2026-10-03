package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

const taskQueue = "recovery-evaluation"

type Input struct {
	Provider   string `json:"provider"`
	Operation  string `json:"operation"`
	Marker     string `json:"marker,omitempty"`
	PauseAfter string `json:"pause_after,omitempty"`
	PauseStep  string `json:"pause_step,omitempty"`
}

type Status struct {
	Operation      string            `json:"operation"`
	State          string            `json:"state"`
	Step           string            `json:"step"`
	Checkpoint     bool              `json:"checkpoint"`
	Results        map[string]string `json:"results"`
	CompletedSteps []string          `json:"completed_steps"`
}

type PaidInput struct {
	Input    Input
	Step     string
	Previous string
}

func Research(ctx workflow.Context, input Input) (Status, error) {
	status := Status{Operation: input.Operation, State: "running", Results: map[string]string{}, CompletedSteps: []string{}}
	if err := workflow.SetQueryHandler(ctx, "status", func() (Status, error) { return status, nil }); err != nil {
		return status, err
	}
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout:    15 * time.Second,
		ScheduleToCloseTimeout: 60 * time.Second,
		RetryPolicy:            &temporal.RetryPolicy{InitialInterval: 100 * time.Millisecond, MaximumInterval: 100 * time.Millisecond, MaximumAttempts: 2},
	})
	for _, step := range []string{"identity", "synthesis"} {
		status.Step = step
		var result string
		if err := workflow.ExecuteActivity(ctx, PaidStep, PaidInput{Input: input, Step: step, Previous: status.Results["identity"]}).Get(ctx, &result); err != nil {
			status.State = "failed"
			return status, err
		}
		status.Results[step] = result
		status.CompletedSteps = append(status.CompletedSteps, step)
		if input.PauseAfter == "checkpoint" && input.PauseStep == step {
			// Activity.Get has returned its recorded completion before this wait.
			status.State, status.Checkpoint = "waiting", true
			var proceed bool
			workflow.GetSignalChannel(ctx, "continue").Receive(ctx, &proceed)
			status.State, status.Checkpoint = "running", false
		}
	}
	status.State = "completed"
	return status, nil
}

func PaidStep(ctx context.Context, input PaidInput) (string, error) {
	body, err := json.Marshal(map[string]string{"operation_id": input.Input.Operation, "step": input.Step, "input": input.Previous})
	if err != nil {
		return "", err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, input.Input.Provider+"/paid-step", bytes.NewReader(body))
	if err != nil {
		return "", temporal.NewNonRetryableApplicationError("invalid provider URL", "configuration", err)
	}
	req.Header.Set("Content-Type", "application/json")
	response, err := (&http.Client{Timeout: 10 * time.Second}).Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("provider HTTP %d", response.StatusCode)
	}
	var paid struct {
		Operation string `json:"operation_id"`
		Step      string `json:"step"`
		Result    string `json:"result"`
	}
	if err := json.NewDecoder(io.LimitReader(response.Body, 1024*1024)).Decode(&paid); err != nil {
		return "", err
	}
	if paid.Operation != input.Input.Operation || paid.Step != input.Step {
		return "", temporal.NewNonRetryableApplicationError("provider returned wrong operation or step", "provider-response", nil)
	}
	if input.Input.PauseAfter == "provider" && input.Input.PauseStep == input.Step && activity.GetInfo(ctx).Attempt == 1 {
		if err := os.MkdirAll(filepath.Dir(input.Input.Marker), 0700); err != nil {
			return "", temporal.NewNonRetryableApplicationError("cannot create marker directory", "configuration", err)
		}
		if err := os.WriteFile(input.Input.Marker, []byte("provider succeeded before Activity completion\n"), 0600); err != nil {
			return "", temporal.NewNonRetryableApplicationError("cannot create marker", "configuration", err)
		}
		// Do not heartbeat or report completion. The driver kills this Worker;
		// the Service times out this attempt and permits exactly one retry.
		<-ctx.Done()
		return "", ctx.Err()
	}
	return paid.Result, nil
}

func run() error {
	if len(os.Args) < 2 {
		return errors.New("usage: recovery-temporal worker|start|status|history|continue [flags]")
	}
	command := os.Args[1]
	flags := flag.NewFlagSet(command, flag.ContinueOnError)
	address := flags.String("address", "127.0.0.1:7233", "Temporal Service address")
	var input Input
	flags.StringVar(&input.Provider, "provider", "", "fake provider base URL")
	flags.StringVar(&input.Operation, "operation", "", "Workflow ID")
	flags.StringVar(&input.Marker, "marker", "", "provider failpoint marker file")
	flags.StringVar(&input.PauseAfter, "pause-after", "", "checkpoint or provider")
	flags.StringVar(&input.PauseStep, "pause-step", "identity", "identity or synthesis")
	if err := flags.Parse(os.Args[2:]); err != nil {
		return err
	}
	if command != "worker" && input.Operation == "" {
		return errors.New("--operation is required")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	c, err := client.DialContext(ctx, client.Options{HostPort: *address, Namespace: "default"})
	if err != nil {
		return err
	}
	defer c.Close()
	write := func(value any) error { return json.NewEncoder(os.Stdout).Encode(value) }
	switch command {
	case "worker":
		w := worker.New(c, taskQueue, worker.Options{})
		w.RegisterWorkflow(Research)
		w.RegisterActivity(PaidStep)
		return w.Run(worker.InterruptCh())
	case "start":
		if input.Provider == "" || (input.PauseAfter == "provider" && input.Marker == "") {
			return errors.New("--provider is required; provider failpoint also requires --marker")
		}
		if input.PauseAfter != "" && input.PauseAfter != "checkpoint" && input.PauseAfter != "provider" {
			return errors.New("--pause-after must be checkpoint or provider")
		}
		if input.PauseStep != "identity" && input.PauseStep != "synthesis" {
			return errors.New("--pause-step must be identity or synthesis")
		}
		_, err = c.ExecuteWorkflow(ctx, client.StartWorkflowOptions{ID: input.Operation, TaskQueue: taskQueue}, Research, input)
		if err != nil {
			return err
		}
		return write(map[string]string{"operation": input.Operation, "state": "started"})
	case "status":
		value, err := c.QueryWorkflow(ctx, input.Operation, "", "status")
		if err != nil {
			return err
		}
		var status Status
		if err := value.Get(&status); err != nil {
			return err
		}
		return write(status)
	case "continue":
		if err := c.SignalWorkflow(ctx, input.Operation, "", "continue", true); err != nil {
			return err
		}
		return write(map[string]string{"operation": input.Operation, "state": "signaled"})
	case "history":
		iterator := c.GetWorkflowHistory(ctx, input.Operation, "", false, enums.HISTORY_EVENT_FILTER_TYPE_ALL_EVENT)
		events := []map[string]any{}
		for iterator.HasNext() {
			event, err := iterator.Next()
			if err != nil {
				return err
			}
			entry := map[string]any{"event_id": event.EventId, "event_type": event.EventType.String()}
			if scheduled := event.GetActivityTaskScheduledEventAttributes(); scheduled != nil {
				entry["activity_type"] = scheduled.ActivityType.Name
			}
			events = append(events, entry)
		}
		return write(map[string]any{"operation": input.Operation, "events": events})
	default:
		return fmt.Errorf("unknown command %q", command)
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
