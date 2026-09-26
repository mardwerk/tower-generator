package server

import (
	"context"
	"errors"
	"strings"
	"sync"

	"github.com/mardwerk/unit-generator/src/cli/internal/provider"
	s "github.com/mardwerk/unit-generator/src/cli/internal/schema"
	"github.com/mardwerk/unit-generator/src/cli/internal/unit"
)

// KeyState says whether a key is set and where it came from. It never
// holds the key; the hint is a masked fragment of a long key.
type KeyState struct {
	Configured bool    `json:"configured"`
	Source     string  `json:"source"` // env, env-file, settings or none
	Hint       *string `json:"hint"`
}

// ImageState is the image model and whether it can run.
type ImageState struct {
	Model string `json:"model"`
	Ready bool   `json:"ready"`
}

// ProviderState is the server's model connection as the client sees it.
type ProviderState struct {
	Provider  string     `json:"provider"`
	Model     string     `json:"model"`
	Reasoning string     `json:"reasoning"`
	Ready     bool       `json:"ready"`
	Images    ImageState `json:"images"`
	Key       KeyState   `json:"key"`
	Message   string     `json:"message"`
}

// connection is the server's one model connection. The key stays in this
// process and is never part of an artifact.
type connection struct {
	mu                                     sync.Mutex
	env                                    provider.Environment
	key, keySource                         string
	provider, model, reasoning, imageModel string
	active                                 unit.Model
}

var settingsSchema = s.StrictObject(
	s.F("provider", s.Enum("openrouter", "codex")),
	s.F("apiKey", s.Optional(s.String().Trim().Min(1).Max(4096))),
	s.F("model", s.Optional(s.String().Trim().Min(1).Max(200))),
	s.F("reasoning", s.Optional(s.String().Trim().Min(1).Max(20))),
	s.F("imageModel", s.Optional(s.String().Trim().Min(1).Max(200))),
)

func newConnection(env provider.Environment, providerName, model string) (*connection, error) {
	c := &connection{env: env, provider: "openrouter"}
	c.key, c.keySource = env.Get("OPENROUTER_API_KEY")
	c.imageModel = env.Value("OPENROUTER_IMAGE_MODEL")
	if c.imageModel == "" {
		c.imageModel = provider.DefaultImageModel
	}
	settings := s.NewObject().Set("provider", providerName)
	if model != "" {
		settings.Set("model", model)
	}
	if _, err := c.configure(settings); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *connection) state() ProviderState {
	c.mu.Lock()
	defer c.mu.Unlock()
	ready := c.provider == "codex" || c.key != ""
	state := ProviderState{
		Provider:  c.provider,
		Model:     c.model,
		Reasoning: c.reasoning,
		Ready:     ready,
		Images:    ImageState{Model: c.imageModel, Ready: c.key != ""},
		Key:       KeyState{Configured: c.key != "", Source: c.keySource, Hint: provider.KeyHint(c.key)},
	}
	switch {
	case c.provider == "codex":
		if c.model == "" {
			state.Message = "Uses your local Codex configuration and login."
		} else {
			state.Message = "Using " + c.model + " with " + firstNonempty(c.reasoning, "configured") + " reasoning through local Codex."
		}
	case ready:
		state.Message = "Using " + c.model + ". The API key stays on this local server."
	default:
		state.Message = "Add an OpenRouter API key to use free models, or select Local Codex."
	}
	return state
}

// configure validates new settings completely before applying any of them.
func (c *connection) configure(value any) (ProviderState, error) {
	var settings struct {
		Provider   string  `json:"provider"`
		APIKey     *string `json:"apiKey"`
		Model      *string `json:"model"`
		Reasoning  *string `json:"reasoning"`
		ImageModel *string `json:"imageModel"`
	}
	if err := s.ParseInto(settingsSchema, value, &settings); err != nil {
		return ProviderState{}, err
	}
	c.mu.Lock()
	key := c.key
	if settings.APIKey != nil {
		key = *settings.APIKey
	}
	model := ""
	if settings.Model != nil {
		model = *settings.Model
	} else if settings.Provider == "openrouter" {
		model = c.env.Value("OPENROUTER_MODEL")
		if model == "" {
			model = provider.FreeModel
		}
	} else {
		model = c.env.Value("CODEX_MODEL")
	}
	reasoning := ""
	if settings.Reasoning != nil {
		reasoning = *settings.Reasoning
	} else if settings.Provider == "openrouter" {
		reasoning = c.env.Value("OPENROUTER_REASONING")
	} else {
		reasoning = c.env.Value("CODEX_REASONING")
	}
	imageModel := c.imageModel
	if settings.ImageModel != nil {
		imageModel = *settings.ImageModel
	}
	c.mu.Unlock()
	if key != "" && strings.Contains(model, key) {
		return ProviderState{}, errors.New("The model name must not contain an API key.")
	}
	// Validate the image model without contacting the provider.
	if _, err := provider.NewImages(provider.ImageOptions{APIKey: key, Model: imageModel}); err != nil {
		return ProviderState{}, err
	}
	var client unit.Model
	var err error
	if settings.Provider == "openrouter" {
		client, err = provider.NewOpenRouter(provider.OpenRouterOptions{APIKey: key, Model: model, Reasoning: reasoning})
	} else {
		client, err = provider.NewCodex(provider.CodexOptions{Model: model, Reasoning: reasoning})
	}
	if err != nil {
		return ProviderState{}, err
	}
	c.mu.Lock()
	c.key = key
	if settings.APIKey != nil {
		c.keySource = "settings"
	}
	if c.key == "" {
		c.keySource = provider.SourceNone
	}
	c.provider, c.model, c.reasoning, c.imageModel, c.active = settings.Provider, model, reasoning, imageModel, client
	c.mu.Unlock()
	return c.state(), nil
}

type modelCatalog struct {
	Models           []provider.ModelChoice `json:"models"`
	DefaultModel     string                 `json:"defaultModel"`
	DefaultReasoning string                 `json:"defaultReasoning"`
}

func (c *connection) models(ctx context.Context, name string) (modelCatalog, error) {
	c.mu.Lock()
	key := c.key
	c.mu.Unlock()
	if name == "openrouter" {
		models, err := provider.OpenRouterModels(ctx, key, "", nil)
		return modelCatalog{Models: models, DefaultModel: firstNonempty(c.env.Value("OPENROUTER_MODEL"), provider.FreeModel), DefaultReasoning: firstNonempty(c.env.Value("OPENROUTER_REASONING"), "none")}, err
	}
	if name == "codex" {
		models, err := provider.CodexModels(ctx, "")
		return modelCatalog{Models: models, DefaultModel: c.env.Value("CODEX_MODEL"), DefaultReasoning: c.env.Value("CODEX_REASONING")}, err
	}
	return modelCatalog{}, errors.New("Provider must be openrouter or codex.")
}

func firstNonempty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

func (c *connection) client() unit.Model {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.active
}

func (c *connection) images() (*provider.Images, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return provider.NewImages(provider.ImageOptions{APIKey: c.key, Model: c.imageModel})
}
