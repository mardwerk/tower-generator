package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func TestRecordedIdentityFeedsSynthesisAfterContinue(t *testing.T) {
	var suite testsuite.WorkflowTestSuite
	env := suite.NewTestWorkflowEnvironment()
	env.RegisterActivity(PaidStep)
	steps := []string{}
	env.OnActivity(PaidStep, mock.Anything, mock.Anything).Return(func(_ context.Context, input PaidInput) (string, error) {
		steps = append(steps, input.Step)
		if input.Step == "identity" {
			require.Empty(t, input.Previous)
			return "identity-response-unique", nil
		}
		require.Equal(t, "identity-response-unique", input.Previous)
		return "synthesis-from-identity-response-unique", nil
	})
	env.RegisterDelayedCallback(func() {
		value, err := env.QueryWorkflow("status")
		require.NoError(t, err)
		var status Status
		require.NoError(t, value.Get(&status))
		require.Equal(t, "waiting", status.State)
		require.True(t, status.Checkpoint)
		require.Equal(t, []string{"identity"}, status.CompletedSteps)
		require.Equal(t, []string{"identity"}, steps)
		env.SignalWorkflow("continue", true)
	}, time.Second)
	env.ExecuteWorkflow(Research, Input{Operation: "test", PauseAfter: "checkpoint", PauseStep: "identity"})
	require.NoError(t, env.GetWorkflowError())
	var result Status
	require.NoError(t, env.GetWorkflowResult(&result))
	require.Equal(t, "completed", result.State)
	require.Equal(t, []string{"identity", "synthesis"}, steps)
	require.Equal(t, "synthesis-from-identity-response-unique", result.Results["synthesis"])
	env.AssertExpectations(t)
}
