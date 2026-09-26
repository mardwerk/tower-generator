package provider

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

// ModelChoice is one model offered by the selected provider. Reasoning levels
// are the values that provider reports, or its supported request values.
type ModelChoice struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Reasoning []string `json:"reasoning"`
}

// OpenRouterModels reads the public model catalog. The configured key is used
// so the catalog can reflect the account's available endpoint.
func OpenRouterModels(ctx context.Context, key, baseURL string, client *http.Client) ([]ModelChoice, error) {
	if baseURL == "" {
		baseURL = OpenRouterURL
	}
	if client == nil {
		client = &http.Client{Timeout: 10 * time.Second}
	}
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, strings.TrimRight(baseURL, "/")+"/models", nil)
	if err != nil {
		return nil, errors.New("Could not request the OpenRouter model catalog.")
	}
	if key != "" {
		req.Header.Set("Authorization", "Bearer "+key)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Could not reach the OpenRouter model catalog.")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("OpenRouter did not return its model catalog.")
	}
	var result struct {
		Data []struct {
			ID                  string   `json:"id"`
			Name                string   `json:"name"`
			SupportedParameters []string `json:"supported_parameters"`
		} `json:"data"`
	}
	if err := json.NewDecoder(io.LimitReader(resp.Body, 5_000_000)).Decode(&result); err != nil {
		return nil, errors.New("OpenRouter returned an invalid model catalog.")
	}
	choices := []ModelChoice{}
	for _, item := range result.Data {
		if !modelID.MatchString(item.ID) {
			continue
		}
		reasoning := []string{"none"}
		for _, parameter := range item.SupportedParameters {
			if parameter == "reasoning" {
				reasoning = []string{"none", "low", "medium", "high"}
				break
			}
		}
		choices = append(choices, ModelChoice{ID: item.ID, Name: item.Name, Reasoning: reasoning})
	}
	return choices, nil
}

// CodexModels asks the installed Codex CLI for its current model catalog.
// Its debug command refreshes the catalog through the user's existing login.
func CodexModels(ctx context.Context, executable string) ([]ModelChoice, error) {
	if executable == "" {
		executable = "codex"
	}
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, executable, "debug", "models")
	output, err := cmd.Output()
	if err != nil || len(output) > 5_000_000 {
		return nil, errors.New("Could not read the local Codex model catalog.")
	}
	var result struct {
		Models []struct {
			Slug       string `json:"slug"`
			Name       string `json:"display_name"`
			Visibility string `json:"visibility"`
			Levels     []struct {
				Effort string `json:"effort"`
			} `json:"supported_reasoning_levels"`
		} `json:"models"`
	}
	if json.Unmarshal(output, &result) != nil {
		return nil, errors.New("Codex returned an invalid model catalog.")
	}
	choices := []ModelChoice{}
	for _, item := range result.Models {
		if item.Visibility != "list" || item.Slug == "" {
			continue
		}
		choice := ModelChoice{ID: item.Slug, Name: item.Name}
		for _, level := range item.Levels {
			choice.Reasoning = append(choice.Reasoning, level.Effort)
		}
		choices = append(choices, choice)
	}
	return choices, nil
}
