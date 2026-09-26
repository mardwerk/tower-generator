package provider

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestOpenRouterModelsUseSelectedEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/models" || r.Header.Get("Authorization") != "Bearer test-key" {
			t.Errorf("request path or authorization missing")
		}
		_, _ = w.Write([]byte(`{"data":[{"id":"openai/example","name":"Example","supported_parameters":["reasoning"]},{"id":"other/basic","name":"Basic","supported_parameters":[]}]}`))
	}))
	defer server.Close()
	models, err := OpenRouterModels(context.Background(), "test-key", server.URL, server.Client())
	if err != nil || len(models) != 2 || models[0].ID != "openai/example" ||
		!reflect.DeepEqual(models[0].Reasoning, []string{"none", "low", "medium", "high"}) ||
		!reflect.DeepEqual(models[1].Reasoning, []string{"none"}) {
		t.Fatalf("models %+v: %v", models, err)
	}
}

func TestCodexModelsUseInstalledCatalog(t *testing.T) {
	path := filepath.Join(t.TempDir(), "codex")
	script := "#!/bin/sh\n[ \"$1 $2\" = \"debug models\" ] || exit 1\nprintf '%s' '{\"models\":[{\"slug\":\"gpt-6-luna\",\"display_name\":\"Luna\",\"visibility\":\"list\",\"supported_reasoning_levels\":[{\"effort\":\"medium\"},{\"effort\":\"high\"}]},{\"slug\":\"hidden\",\"visibility\":\"hide\"}]}'\n"
	if err := os.WriteFile(path, []byte(script), 0o700); err != nil {
		t.Fatal(err)
	}
	models, err := CodexModels(context.Background(), path)
	if err != nil || len(models) != 1 || models[0].ID != "gpt-6-luna" ||
		!reflect.DeepEqual(models[0].Reasoning, []string{"medium", "high"}) {
		t.Fatalf("models %+v: %v", models, err)
	}
}
