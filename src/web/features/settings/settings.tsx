import { useEffect, useRef, useState } from 'react';
import { KeyRound } from 'lucide-react';
import type { KeyState, ModelCatalog, ProviderState } from '../../api/contract.js';
import type { GenerationLibrary } from '../library/library.js';
import { api } from '../../api/client.js';
import { Alert } from '../../ui/alert.js';
import { Button } from '../../ui/button.js';
import { Modal } from '../../ui/dialog.js';
import { Disclosure } from '../../ui/disclosure.js';
import { Field } from '../../ui/field.js';
import { Input } from '../../ui/input.js';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '../../ui/select.js';
import { cn } from '../../ui/utils.js';

const keySources: Record<KeyState['source'], string> = {
  'env-file': 'from .env',
  env: 'from the OPENROUTER_API_KEY environment variable',
  settings: 'entered in Settings',
  none: '',
};

const reasoningLabels: Record<string, string> = {
  none: 'None',
  minimal: 'Minimal',
  low: 'Low',
  medium: 'Medium',
  high: 'High',
  xhigh: 'Xhigh',
  max: 'Max',
  ultra: 'Ultra',
};

const codexDefaultModel = '__codex_default__';

/** The provider state the server reports, refreshed whenever Settings closes. */
export function useProvider(settingsOpen: boolean): ProviderState | null {
  const [state, setState] = useState<ProviderState | null>(null);
  useEffect(() => {
    if (settingsOpen) return;
    let active = true;
    void api<ProviderState>('provider')
      .then((next) => {
        if (active) setState(next);
      })
      .catch(() => undefined);
    return () => {
      active = false;
    };
  }, [settingsOpen]);
  return state;
}

/** The masked key in use, visible without opening Settings. */
export function KeyStatus({ state, onOpen }: { state: ProviderState | null; onOpen: () => void }) {
  if (!state) return null;
  const { key } = state;
  const missing = !key.configured && state.provider === 'openrouter';
  return (
    <Button
      id="key-status"
      variant="ghost"
      size="xs"
      className={cn(
        'max-w-[190px] font-mono',
        missing && 'font-sans text-warning hover:text-warning',
      )}
      title={`${state.model || 'Codex configuration'}${state.reasoning ? ` (${state.reasoning} reasoning)` : ''} via ${
        state.provider === 'openrouter' ? 'OpenRouter' : 'Local Codex'
      }. ${
        key.configured
          ? `OpenRouter key ${keySources[key.source]}.`
          : 'No OpenRouter key configured.'
      } Open Settings to change them.`}
      onClick={onOpen}
    >
      <KeyRound aria-hidden="true" />
      <span className="sr-only">OpenRouter key: </span>
      <span className="hidden truncate sm:inline">
        {key.configured ? (key.hint ?? 'set') : 'No API key'}
      </span>
    </Button>
  );
}

/** Keys are submitted to the local server and never included in saved authoring work. */
export function Settings({
  disabled,
  library,
  onClose,
}: {
  disabled: boolean;
  library: GenerationLibrary;
  onClose: () => void;
}) {
  const [provider, setProvider] = useState<ProviderState['provider']>('openrouter');
  const [model, setModel] = useState('');
  const [reasoning, setReasoning] = useState('');
  const [models, setModels] = useState<ModelCatalog['models']>([]);
  const [modelMessage, setModelMessage] = useState('');
  const [imageModel, setImageModel] = useState('');
  const [key, setKey] = useState('');
  const [keyState, setKeyState] = useState<ProviderState['key'] | null>(null);
  const [ready, setReady] = useState(false);
  const [message, setMessage] = useState('Loading provider...');
  const [error, setError] = useState('');
  const [pending, setPending] = useState(true);
  const [directory, setDirectory] = useState(library.directory);
  const [folderMessage, setFolderMessage] = useState('');
  const modelRequest = useRef(0);
  const selectedModel = models.find((entry) => entry.id === model);
  const reasoningLevels = selectedModel?.reasoning.length
    ? selectedModel.reasoning
    : !model || models.length === 0
      ? provider === 'codex'
        ? ['minimal', 'low', 'medium', 'high', 'xhigh', 'max', 'ultra']
        : ['none', 'low', 'medium', 'high']
      : reasoning
        ? [reasoning]
        : [];
  async function loadModels(name: ProviderState['provider'], reset: boolean) {
    const request = ++modelRequest.current;
    setModelMessage('Loading models...');
    try {
      const catalog = await api<ModelCatalog>('models', { provider: name });
      if (request !== modelRequest.current) return;
      setModels([...catalog.models].sort((a, b) => a.id.localeCompare(b.id)));
      setModelMessage(
        `${catalog.models.length} models from ${name === 'codex' ? 'Codex' : 'OpenRouter'}.`,
      );
      if (reset) {
        const nextModel = catalog.defaultModel;
        setModel(nextModel);
        const offered = catalog.models.find((entry) => entry.id === nextModel);
        const preferred = catalog.defaultReasoning || (name === 'codex' ? 'medium' : 'none');
        setReasoning(
          offered?.reasoning.length && !offered.reasoning.includes(preferred)
            ? (offered.reasoning[0] ?? preferred)
            : preferred,
        );
      }
    } catch (error) {
      if (request !== modelRequest.current) return;
      setModels([]);
      setModelMessage(error instanceof Error ? error.message : String(error));
    }
  }
  useEffect(() => {
    let active = true;
    void api<ProviderState>('provider')
      .then((state) => {
        if (active) {
          setProvider(state.provider);
          setModel(state.model);
          setReasoning(state.reasoning || (state.provider === 'codex' ? 'medium' : 'none'));
          setImageModel(state.images.model);
          setKeyState(state.key);
          setReady(state.ready);
          setMessage(state.message);
          void loadModels(state.provider, false);
        }
      })
      .catch((error) => {
        if (active) setError(String(error));
      })
      .finally(() => {
        if (active) setPending(false);
      });
    return () => {
      active = false;
    };
  }, []);
  useEffect(() => setDirectory(library.directory), [library.directory]);
  useEffect(() => {
    if (selectedModel?.reasoning.length && !selectedModel.reasoning.includes(reasoning)) {
      setReasoning(
        selectedModel.reasoning.includes('medium')
          ? 'medium'
          : (selectedModel.reasoning[0] ?? reasoning),
      );
    }
  }, [selectedModel, reasoning]);
  async function saveProvider() {
    setPending(true);
    setError('');
    try {
      const state = await api<ProviderState>('provider', {
        provider,
        ...(imageModel.trim() ? { imageModel: imageModel.trim() } : {}),
        ...(key.trim() ? { apiKey: key.trim() } : {}),
        ...(model.trim() ? { model: model.trim() } : {}),
        ...(reasoning ? { reasoning } : {}),
      });
      setKey('');
      setKeyState(state.key);
      setReady(state.ready);
      setMessage(state.message);
      setModel(state.model);
      setReasoning(state.reasoning);
      setImageModel(state.images.model);
      void loadModels(state.provider, false);
    } catch (error) {
      setError(error instanceof Error ? error.message : String(error));
    } finally {
      setPending(false);
    }
  }
  async function saveFolder() {
    setPending(true);
    setError('');
    try {
      await library.configure(directory);
      setFolderMessage('Library folder changed. Existing files stay in their original folder.');
    } catch (error) {
      setError(error instanceof Error ? error.message : String(error));
    } finally {
      setPending(false);
    }
  }
  return (
    <Modal title="Settings" id="settings-dialog" onClose={onClose}>
      <fieldset disabled={disabled || pending} id="provider-settings" className="min-w-0">
        <h3 className="mb-1 text-sm font-semibold">Model provider</h3>
        <Field label="Connection">
          <Select
            value={provider}
            onValueChange={(next) => {
              const value = next as ProviderState['provider'];
              setProvider(value);
              setModel('');
              setReasoning('');
              setModels([]);
              void loadModels(value, true);
              setKey('');
              setReady(false);
              setMessage(
                value === 'openrouter'
                  ? 'Free models still require an API key.'
                  : 'Uses your existing local Codex login.',
              );
            }}
          >
            <SelectTrigger id="provider-select">
              <SelectValue />
            </SelectTrigger>
            <SelectContent>
              <SelectItem
                value="openrouter"
                description="Free models with a key, paid models by ID"
              >
                OpenRouter
              </SelectItem>
              <SelectItem value="codex" description="Your existing Codex login">
                Local Codex
              </SelectItem>
            </SelectContent>
          </Select>
        </Field>
        {
          <>
            <p id="current-key" className="font-mono text-xs text-muted-foreground">
              {keyState?.configured
                ? `Current key: ${keyState.hint ?? 'set (too short to show a fragment)'} (${
                    keySources[keyState.source]
                  })`
                : 'No OpenRouter key configured.'}
            </p>
            <Field label="OpenRouter API key">
              <Input
                type="password"
                autoComplete="off"
                value={key}
                placeholder={
                  keyState?.configured ? 'Enter a new key to replace it' : 'Enter an API key'
                }
                onChange={(e) => setKey(e.target.value)}
              />
            </Field>
          </>
        }
        <Disclosure title="Model" defaultOpen>
          <Field label="Model">
            <Select
              value={model || (provider === 'codex' ? codexDefaultModel : undefined)}
              onValueChange={(value) => {
                const next = value === codexDefaultModel ? '' : value;
                setModel(next);
                if (!next) {
                  setReasoning('medium');
                  return;
                }
                const offered = models.find((entry) => entry.id === next);
                if (offered?.reasoning.length && !offered.reasoning.includes(reasoning)) {
                  setReasoning(
                    offered.reasoning.includes('medium')
                      ? 'medium'
                      : (offered.reasoning[0] ?? reasoning),
                  );
                }
              }}
            >
              <SelectTrigger
                id="model-select"
                disabled={provider === 'openrouter' && models.length === 0 && !model}
              >
                <SelectValue placeholder="Loading models..." />
              </SelectTrigger>
              <SelectContent>
                {provider === 'codex' && (
                  <SelectItem
                    value={codexDefaultModel}
                    description="Use your local Codex configuration"
                  >
                    Codex default
                  </SelectItem>
                )}
                {model && !selectedModel && (
                  <SelectItem
                    value={model}
                    description="Current setting; not in the fetched catalog"
                  >
                    {model}
                  </SelectItem>
                )}
                {models.map((entry) => (
                  <SelectItem
                    key={entry.id}
                    value={entry.id}
                    description={entry.name === entry.id ? undefined : entry.name}
                  >
                    {entry.id}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </Field>
          <p className="text-xs text-muted-foreground" role="status" aria-live="polite">
            {modelMessage}
          </p>
          <fieldset className="my-3 min-w-0">
            <legend className="mb-1.5 text-xs text-muted-foreground">Reasoning level</legend>
            <div className="flex flex-wrap gap-1.5">
              {reasoningLevels.map((level) => (
                <label key={level} className="cursor-pointer">
                  <input
                    type="radio"
                    name="reasoning-level"
                    value={level}
                    checked={reasoning === level}
                    onChange={() => setReasoning(level)}
                    className="peer sr-only"
                  />
                  <span className="inline-flex h-8 items-center rounded-md border border-input bg-background px-3 text-xs text-foreground transition-colors hover:bg-accent peer-checked:border-primary peer-checked:bg-accent peer-checked:text-accent-foreground peer-focus-visible:ring-2 peer-focus-visible:ring-ring/60">
                    {reasoningLabels[level] ?? level}
                  </span>
                </label>
              ))}
            </div>
            <p className="mt-1.5 text-xs text-muted-foreground">
              {selectedModel?.reasoning.length
                ? 'Only levels available for this model are shown. Higher levels generally use more time and tokens.'
                : 'Choose a listed model to see its supported levels. Higher levels generally use more time and tokens.'}
            </p>
          </fieldset>
        </Disclosure>
        <Disclosure title="Image generation">
          <Field label="OpenRouter image model">
            <Input
              value={imageModel}
              onChange={(e) => setImageModel(e.target.value)}
              placeholder="meta/muse-image"
            />
          </Field>
          <p className="text-xs text-muted-foreground">
            Muse Image is the default, listed at about $0.01 per image. Images use your OpenRouter
            key even when Unit drafting uses Codex. Every image requires confirmation from its icon
            dialog.
          </p>
        </Disclosure>
        <Button variant="primary" onClick={() => void saveProvider()}>
          Save provider
        </Button>
        <p className="mt-2 text-xs text-muted-foreground" role="status">
          {message}
        </p>
      </fieldset>
      <fieldset disabled={disabled || pending} className="mt-6 min-w-0 border-t border-border pt-5">
        <h3 className="mb-1 text-sm font-semibold">Local library</h3>
        <Field label="Library folder">
          <Input
            className="font-mono text-xs"
            value={directory}
            onChange={(e) => setDirectory(e.target.value)}
            spellCheck={false}
          />
        </Field>
        <p className="text-xs text-muted-foreground">
          Completed generations are saved here automatically. Changing this folder does not move
          existing files.
        </p>
        <Button className="mt-3" onClick={() => void saveFolder()}>
          Use folder
        </Button>
        <p className="mt-2 text-xs text-muted-foreground" role="status">
          {folderMessage}
        </p>
      </fieldset>
      {error && <Alert>{error}</Alert>}
    </Modal>
  );
}
