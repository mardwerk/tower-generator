// This experiment tests process-kill recovery, not a product execution policy.
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const schemaVersion = 1

var steps = []string{"identity", "synthesis"}
var operationPattern = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9_.-]{0,127}$`)

type options struct {
	store, provider, operation, pauseAfter, pauseStep, marker string
}

type attempt struct {
	SchemaVersion int    `json:"schema_version"`
	OperationID   string `json:"operation_id"`
	Step          string `json:"step"`
	Status        string `json:"status"`
}

type checkpoint struct {
	SchemaVersion int    `json:"schema_version"`
	OperationID   string `json:"operation_id"`
	Step          string `json:"step"`
	Result        string `json:"result"`
}

type state struct {
	SchemaVersion    int               `json:"schema_version"`
	OperationID      string            `json:"operation_id"`
	Status           string            `json:"status"`
	Results          map[string]string `json:"results"`
	ReservedAttempts []attempt         `json:"reserved_attempts"`
	Error            string            `json:"error,omitempty"`
}

func main() {
	var o options
	flag.StringVar(&o.store, "store", "", "Experiment operation store")
	flag.StringVar(&o.provider, "provider", "", "Fake provider base URL")
	flag.StringVar(&o.operation, "operation", "", "Operation ID")
	flag.StringVar(&o.pauseAfter, "pause-after", "", "Pause after checkpoint or provider")
	flag.StringVar(&o.pauseStep, "pause-step", "identity", "Step at which to pause")
	flag.StringVar(&o.marker, "marker", "", "Marker written when the pause is reached")
	flag.Parse()
	s, err := execute(o)
	if err != nil {
		if s.Status != "spend_uncertain" {
			s.Status = "failed"
		}
		s.Error = err.Error()
	}
	if encodeErr := json.NewEncoder(os.Stdout).Encode(s); encodeErr != nil {
		fmt.Fprintln(os.Stderr, encodeErr)
		os.Exit(1)
	}
	if err != nil {
		os.Exit(1)
	}
}

func execute(o options) (state, error) {
	s := state{SchemaVersion: schemaVersion, OperationID: o.operation, Status: "running", Results: map[string]string{}, ReservedAttempts: []attempt{}}
	if o.store == "" || o.provider == "" || !operationPattern.MatchString(o.operation) {
		return s, errors.New("--store, --provider and a safe --operation ID are required")
	}
	if o.pauseAfter != "" && (o.pauseAfter != "checkpoint" && o.pauseAfter != "provider" || o.marker == "" || o.pauseStep != "identity" && o.pauseStep != "synthesis") {
		return s, errors.New("pause requires checkpoint|provider, identity|synthesis and --marker")
	}
	dir := filepath.Join(o.store, o.operation)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return s, err
	}
	statePath := filepath.Join(dir, "state.json")
	var previous state
	if found, err := readJSON(statePath, &previous); err != nil {
		return s, err
	} else if found && (previous.SchemaVersion != schemaVersion || previous.OperationID != o.operation) {
		return s, errors.New("unsupported or mismatched state record")
	}
	// Records and checkpoints are authoritative even when state.json lags a write.
	for _, step := range steps {
		var a attempt
		attemptFound, err := readJSON(filepath.Join(dir, step+".attempt.json"), &a)
		if err != nil {
			return s, err
		}
		var c checkpoint
		checkpointFound, err := readJSON(filepath.Join(dir, step+".checkpoint.json"), &c)
		if err != nil {
			return s, err
		}
		if checkpointFound && !attemptFound {
			return s, errors.New("checkpoint has no reserved attempt")
		}
		if attemptFound {
			if a.SchemaVersion != schemaVersion || a.OperationID != o.operation || a.Step != step || a.Status != "in_flight" {
				return s, errors.New("unsupported or mismatched attempt record")
			}
			if step == "synthesis" {
				if _, ok := s.Results["identity"]; !ok {
					return s, errors.New("synthesis attempt precedes the identity checkpoint")
				}
			}
			if checkpointFound {
				if c.SchemaVersion != schemaVersion || c.OperationID != o.operation || c.Step != step {
					return s, errors.New("unsupported or mismatched checkpoint")
				}
				a.Status = "completed"
				s.Results[step] = c.Result
			} else {
				a.Status = "spend_uncertain"
				s.Status = "spend_uncertain"
			}
			s.ReservedAttempts = append(s.ReservedAttempts, a)
		}
	}
	if s.Status == "spend_uncertain" {
		err := errors.New("an interrupted paid attempt has no checkpoint; automatic retry is refused")
		s.Error = err.Error()
		if writeErr := atomicJSON(statePath, s, false); writeErr != nil {
			return s, writeErr
		}
		return s, err
	}
	if err := atomicJSON(statePath, s, false); err != nil {
		return s, err
	}
	client := &http.Client{Timeout: 30 * time.Second}
	for _, step := range steps {
		if _, completed := s.Results[step]; completed {
			continue
		}
		a := attempt{schemaVersion, o.operation, step, "in_flight"}
		if err := atomicJSON(filepath.Join(dir, step+".attempt.json"), a, true); err != nil {
			return s, err
		}
		s.ReservedAttempts = append(s.ReservedAttempts, a)
		if err := atomicJSON(statePath, s, false); err != nil {
			return s, err
		}
		c, err := paidStep(client, o, step, s.Results["identity"])
		if err != nil {
			s.Status = "spend_uncertain"
			s.ReservedAttempts[len(s.ReservedAttempts)-1].Status = "spend_uncertain"
			s.Error = err.Error()
			if writeErr := atomicJSON(statePath, s, false); writeErr != nil {
				return s, writeErr
			}
			return s, err
		}
		if err := pause(o, "provider", step); err != nil {
			return s, err
		}
		if err := atomicJSON(filepath.Join(dir, step+".checkpoint.json"), c, true); err != nil {
			return s, err
		}
		if err := pause(o, "checkpoint", step); err != nil {
			return s, err
		}
		s.Results[step] = c.Result
		s.ReservedAttempts[len(s.ReservedAttempts)-1].Status = "completed"
		if err := atomicJSON(statePath, s, false); err != nil {
			return s, err
		}
	}
	s.Status = "completed"
	return s, atomicJSON(statePath, s, false)
}

func paidStep(client *http.Client, o options, step, identityResult string) (checkpoint, error) {
	payload := map[string]string{"operation_id": o.operation, "step": step}
	if step == "synthesis" {
		payload["input"] = identityResult
	}
	input, err := json.Marshal(payload)
	if err != nil {
		return checkpoint{}, err
	}
	response, err := client.Post(strings.TrimRight(o.provider, "/")+"/paid-step", "application/json", bytes.NewReader(input))
	if err != nil {
		return checkpoint{}, err
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return checkpoint{}, fmt.Errorf("provider returned HTTP %d; paid outcome is uncertain", response.StatusCode)
	}
	data, err := io.ReadAll(io.LimitReader(response.Body, 65537))
	if err != nil {
		return checkpoint{}, err
	}
	if len(data) > 65536 {
		return checkpoint{}, errors.New("provider response exceeds the experiment limit")
	}
	var output struct {
		OperationID string  `json:"operation_id"`
		Step        string  `json:"step"`
		Result      *string `json:"result"`
	}
	if err := json.Unmarshal(data, &output); err != nil {
		return checkpoint{}, err
	}
	if output.OperationID != o.operation || output.Step != step || output.Result == nil {
		return checkpoint{}, errors.New("provider returned a mismatched operation or step")
	}
	return checkpoint{schemaVersion, output.OperationID, output.Step, *output.Result}, nil
}

func readJSON(path string, value any) (bool, error) {
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, json.Unmarshal(data, value)
}

func atomicJSON(path string, value any, immutable bool) error {
	if immutable {
		if _, err := os.Stat(path); err == nil {
			return fmt.Errorf("immutable record already exists: %s", filepath.Base(path))
		} else if !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	file, err := os.CreateTemp(filepath.Dir(path), ".record-*")
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())
	if _, err = file.Write(append(data, '\n')); err == nil {
		err = file.Sync()
	}
	closeErr := file.Close()
	if err != nil {
		return err
	}
	if closeErr != nil {
		return closeErr
	}
	// One runner owns each operation store. This is not a concurrency lock.
	return os.Rename(file.Name(), path)
}

func pause(o options, boundary, step string) error {
	if o.pauseAfter != boundary || o.pauseStep != step {
		return nil
	}
	if err := os.WriteFile(o.marker, []byte(boundary+":"+step+"\n"), 0600); err != nil {
		return err
	}
	for {
		time.Sleep(time.Hour)
	}
}
