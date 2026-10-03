package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func fakeProvider(t *testing.T, calls *[]string) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/paid-step" || r.Method != http.MethodPost {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		var input map[string]string
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			t.Errorf("request: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		*calls = append(*calls, input["step"])
		result := "saved identity"
		if input["step"] == "synthesis" {
			if input["input"] == "" {
				t.Error("synthesis did not receive the saved identity")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			result = "synthesis of " + input["input"]
		}
		if err := json.NewEncoder(w).Encode(map[string]string{"operation_id": input["operation_id"], "step": input["step"], "result": result}); err != nil {
			t.Errorf("response: %v", err)
		}
	}))
}

func TestCompletedOperationDoesNotRepeatPaidCalls(t *testing.T) {
	var calls []string
	provider := fakeProvider(t, &calls)
	defer provider.Close()
	o := options{store: t.TempDir(), provider: provider.URL, operation: "repeat"}
	first, err := execute(o)
	if err != nil || first.Status != "completed" {
		t.Fatalf("first execution: %+v, %v", first, err)
	}
	second, err := execute(o)
	if err != nil || second.Status != "completed" || !reflect.DeepEqual(first.Results, second.Results) {
		t.Fatalf("repeat execution: %+v, %v", second, err)
	}
	if !reflect.DeepEqual(calls, steps) {
		t.Fatalf("paid calls repeated: %v", calls)
	}
}

func TestCheckpointRecoversWithoutAnUpdatedStateRecord(t *testing.T) {
	var calls []string
	provider := fakeProvider(t, &calls)
	defer provider.Close()
	o := options{store: t.TempDir(), provider: provider.URL, operation: "checkpoint"}
	dir := filepath.Join(o.store, o.operation)
	if err := os.Mkdir(dir, 0700); err != nil {
		t.Fatal(err)
	}
	if err := atomicJSON(filepath.Join(dir, "identity.attempt.json"), attempt{schemaVersion, o.operation, "identity", "in_flight"}, true); err != nil {
		t.Fatal(err)
	}
	saved := checkpoint{schemaVersion, o.operation, "identity", "identity from the previous process"}
	path := filepath.Join(dir, "identity.checkpoint.json")
	if err := atomicJSON(path, saved, true); err != nil {
		t.Fatal(err)
	}
	before, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	s, err := execute(o)
	if err != nil || s.Status != "completed" || s.Results["identity"] != saved.Result || s.Results["synthesis"] != "synthesis of "+saved.Result {
		t.Fatalf("checkpoint recovery: %+v, %v", s, err)
	}
	after, err := os.ReadFile(path)
	if err != nil || !reflect.DeepEqual(before, after) {
		t.Fatalf("checkpoint was changed: %v", err)
	}
	if !reflect.DeepEqual(calls, []string{"synthesis"}) {
		t.Fatalf("identity was called again: %v", calls)
	}
}

func TestUnknownAttemptRefusesAllProviderWork(t *testing.T) {
	var calls []string
	provider := fakeProvider(t, &calls)
	defer provider.Close()
	o := options{store: t.TempDir(), provider: provider.URL, operation: "unknown"}
	dir := filepath.Join(o.store, o.operation)
	if err := os.Mkdir(dir, 0700); err != nil {
		t.Fatal(err)
	}
	if err := atomicJSON(filepath.Join(dir, "identity.attempt.json"), attempt{schemaVersion, o.operation, "identity", "in_flight"}, true); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 2; i++ {
		s, err := execute(o)
		if err == nil || s.Status != "spend_uncertain" || len(s.ReservedAttempts) != 1 || s.ReservedAttempts[0].Status != "spend_uncertain" {
			t.Fatalf("unknown attempt recovery: %+v, %v", s, err)
		}
	}
	if len(calls) != 0 {
		t.Fatalf("unknown attempt caused paid work: %v", calls)
	}
}
