import { useEffect, useRef, useState } from 'react';
import { ArrowLeft, BookOpen, Download, Library, Plus, Upload, UserRound } from 'lucide-react';
import type { CharacterSource, SourceEntry } from '../types.js';
import { Alert } from '../ui/alert.js';
import { Badge } from '../ui/badge.js';
import { Button } from '../ui/button.js';
import { Disclosure } from '../ui/disclosure.js';
import { Field } from '../ui/field.js';
import { IconButton } from '../ui/icon-button.js';
import { Input, Textarea } from '../ui/input.js';
import { cn, download } from '../ui/utils.js';
import { api } from './api.js';
import { SourceEvidence } from './sources.js';
import { Tower } from './tower.js';

type View = 'create' | 'library' | 'tower';
const views = [
  { value: 'create' as const, label: 'Create', icon: Plus },
  { value: 'library' as const, label: 'Library', icon: Library },
  { value: 'tower' as const, label: 'Tower', icon: UserRound },
];
const emptyForm = { name: '', work: '', scope: '', urls: '', supplied: '', notes: '' };

export function App() {
  const [view, setView] = useState<View>('create');
  const [form, setForm] = useState(emptyForm);
  const [entries, setEntries] = useState<SourceEntry[]>([]);
  const [opened, setOpened] = useState<CharacterSource>();
  const [chosen, setChosen] = useState<CharacterSource>();
  const [draft, setDraft] = useState('');
  const [filter, setFilter] = useState('');
  const [busy, setBusy] = useState(false);
  const [error, setError] = useState('');
  const inputFile = useRef<HTMLInputElement>(null);

  async function refresh() {
    setEntries((await api<{ sources: SourceEntry[] }>('/api/sources')).sources);
  }
  useEffect(() => {
    void refresh().catch((failure: Error) => setError(failure.message));
  }, []);
  async function run(task: () => Promise<void>) {
    setBusy(true);
    setError('');
    try {
      await task();
    } catch (failure) {
      setError(failure instanceof Error ? failure.message : String(failure));
    } finally {
      setBusy(false);
    }
  }
  function change(key: keyof typeof form, value: string) {
    setForm({ ...form, [key]: value });
  }
  function show(next: View) {
    if (busy) return;
    setView(next);
    setError('');
    if (next === 'library') void run(refresh);
  }
  useEffect(() => {
    document.title = view === 'tower' ? 'Usopp · Tower Generator' : 'Tower Generator';
  }, [view]);
  function refreshResearch() {
    if (!opened) return;
    setForm({
      ...opened.character,
      urls: opened.documents
        .map((document) => document.url)
        .filter(Boolean)
        .join('\n'),
      supplied: opened.documents.find((document) => document.id === 'supplied')?.text ?? '',
      notes: opened.notes,
    });
    setView('create');
  }
  const matches = entries.filter((entry) =>
    `${entry.character.name} ${entry.character.work}`.toLowerCase().includes(filter.toLowerCase()),
  );

  return (
    <>
      <header className="sticky top-0 z-40 grid min-h-[70px] grid-cols-[minmax(0,1fr)_auto] items-center gap-1.5 border-b border-border bg-canvas px-3.5 py-2.5 sm:grid-cols-[minmax(0,1fr)_auto_minmax(0,1fr)] sm:gap-3 sm:px-6 sm:py-0">
        <a
          href="#"
          onClick={(event) => {
            event.preventDefault();
            show('create');
          }}
          className="col-start-1 row-start-1 flex items-center gap-2.5 text-foreground hover:no-underline"
        >
          <img
            src="/mardwerk.png"
            alt="Mardwerk"
            width={30}
            height={36}
            className="object-contain"
          />
          <span className="font-mono text-[13px] font-semibold tracking-tight">
            tower-generator
          </span>
        </a>
        <nav
          aria-label="Workspace"
          className="col-span-full row-start-2 flex items-center justify-center gap-3.5 sm:col-span-1 sm:col-start-2 sm:row-start-1 sm:gap-2.5"
        >
          {views.map(({ value, label, icon: Icon }) => (
            <IconButton
              key={value}
              label={label}
              disabled={busy}
              aria-current={view === value ? 'page' : undefined}
              className={cn(view === value && 'bg-accent text-foreground')}
              onClick={() => show(value)}
            >
              <Icon className="size-[19px]" />
            </IconButton>
          ))}
        </nav>
        <div className="col-start-2 row-start-1 justify-self-end sm:col-start-3">
          <IconButton
            label="Import character source"
            disabled={busy}
            onClick={() => inputFile.current?.click()}
          >
            <Upload />
          </IconButton>
        </div>
      </header>
      <input
        ref={inputFile}
        type="file"
        accept=".json,application/json"
        hidden
        onChange={(event) => {
          const file = event.target.files?.[0];
          if (file)
            void run(async () => {
              if (file.size > 600_000) throw new Error('The source file is too large.');
              const source = await api<CharacterSource>(
                '/api/import-source',
                JSON.parse(await file.text()),
              );
              setOpened(source);
              await refresh();
              setView('library');
            });
          event.target.value = '';
        }}
      />
      {view === 'tower' ? (
        <Tower draft={draft} onDraft={setDraft} source={chosen} onBusy={setBusy} />
      ) : (
        <main className="px-[18px] py-10 sm:px-6 sm:py-[70px]">
          <div className={cn('mx-auto', view === 'create' ? 'max-w-[680px]' : 'max-w-[1000px]')}>
            <h1 className="mb-6 text-[1.6rem] font-semibold tracking-tight">
              {view === 'create' ? 'Create a Tower' : opened ? opened.character.name : 'Library'}
            </h1>
            {error && <Alert>{error}</Alert>}
            {view === 'create' ? (
              <>
                <form
                  onSubmit={(event) => {
                    event.preventDefault();
                    void run(async () => {
                      const source = await api<CharacterSource>('/api/research', {
                        character: { name: form.name, work: form.work, scope: form.scope },
                        urls: form.urls
                          .split('\n')
                          .map((url) => url.trim())
                          .filter(Boolean),
                        supplied: form.supplied,
                        notes: form.notes,
                      });
                      setOpened(source);
                      await refresh();
                      setView('library');
                    });
                  }}
                >
                  <div className="grid gap-x-3 sm:grid-cols-2">
                    <Field label="Character">
                      <Input
                        required
                        autoComplete="off"
                        placeholder="Usopp"
                        value={form.name}
                        disabled={busy}
                        onChange={(event) => change('name', event.target.value)}
                      />
                    </Field>
                    <Field label="Work">
                      <Input
                        required
                        placeholder="One Piece"
                        value={form.work}
                        disabled={busy}
                        onChange={(event) => change('work', event.target.value)}
                      />
                    </Field>
                  </div>
                  <Field label="Canon scope">
                    <Input
                      required
                      placeholder="Manga through Dressrosa"
                      value={form.scope}
                      disabled={busy}
                      onChange={(event) => change('scope', event.target.value)}
                    />
                  </Field>
                  <Field label="Source URLs, one per line">
                    <Textarea
                      rows={3}
                      placeholder="Official character page, manga reference or technique overview"
                      value={form.urls}
                      disabled={busy}
                      onChange={(event) => change('urls', event.target.value)}
                    />
                  </Field>
                  <Disclosure bare title="Supplied evidence and research notes">
                    <Field label="Evidence">
                      <Textarea
                        rows={4}
                        placeholder="Paste source passages with their references."
                        value={form.supplied}
                        disabled={busy}
                        onChange={(event) => change('supplied', event.target.value)}
                      />
                    </Field>
                    <Field label="Research notes">
                      <Textarea
                        rows={3}
                        placeholder="Summarize the evidence and its limits. Keep Tower mechanics in the design."
                        value={form.notes}
                        disabled={busy}
                        onChange={(event) => change('notes', event.target.value)}
                      />
                    </Field>
                  </Disclosure>
                  <div className="mt-5 flex flex-wrap gap-2">
                    <Button type="submit" variant="primary" disabled={busy}>
                      {busy ? 'Collecting…' : 'Research and save'}
                    </Button>
                    <Button
                      disabled={busy}
                      onClick={() => {
                        setOpened(undefined);
                        show('library');
                      }}
                    >
                      <Library /> Use saved research
                    </Button>
                  </div>
                </form>
                <p className="mt-5 text-xs text-muted-foreground">
                  Research is saved once and can be reused across Tower designs and Profiles. Review
                  the retained evidence before treating it as canon.
                </p>
                <Button variant="ghost" size="sm" className="mt-3" onClick={() => show('tower')}>
                  Open the Usopp example
                </Button>
              </>
            ) : opened ? (
              <>
                <div className="mb-5 flex flex-wrap items-center gap-2 text-xs text-muted-foreground">
                  <span>{opened.character.work}</span>
                  <Badge>{opened.documents.length} documents</Badge>
                </div>
                <p className="mb-4 text-[13px] text-muted-foreground">{opened.character.scope}</p>
                <div className="mb-5 flex flex-wrap gap-2">
                  <Button disabled={busy} onClick={() => setOpened(undefined)}>
                    <ArrowLeft /> Library
                  </Button>
                  <Button
                    onClick={() =>
                      download(
                        `${opened.id}.json`,
                        JSON.stringify(opened, null, 2),
                        'application/json',
                      )
                    }
                  >
                    <Download /> Export source
                  </Button>
                  <Button disabled={busy} onClick={refreshResearch}>
                    Update research
                  </Button>
                  {opened.character.name.toLowerCase() === 'usopp' &&
                    opened.character.work.toLowerCase() === 'one piece' && (
                      <Button
                        variant="primary"
                        onClick={() => {
                          setChosen(opened);
                          show('tower');
                        }}
                      >
                        Use for Usopp
                      </Button>
                    )}
                </div>
                <SourceEvidence source={opened} />
              </>
            ) : (
              <>
                <Field label="Find a character">
                  <Input
                    placeholder="Character or work"
                    value={filter}
                    onChange={(event) => setFilter(event.target.value)}
                  />
                </Field>
                <div className="mt-6 grid gap-3 sm:grid-cols-2 lg:grid-cols-3">
                  {matches.map((entry) => (
                    <button
                      key={entry.id}
                      disabled={busy}
                      className="min-w-0 rounded-lg border border-border bg-card p-4 text-left transition-colors hover:border-input hover:bg-accent focus-visible:ring-2 focus-visible:ring-ring/60"
                      onClick={() =>
                        void run(async () =>
                          setOpened(
                            await api<CharacterSource>(
                              `/api/source?id=${encodeURIComponent(entry.id)}`,
                            ),
                          ),
                        )
                      }
                    >
                      <div className="flex items-center gap-2 text-[13px]">
                        <BookOpen className="size-4 text-muted-foreground" />
                        {entry.character.name}
                      </div>
                      <p className="mt-2 text-xs text-muted-foreground">{entry.character.work}</p>
                      <p className="mt-2 line-clamp-2 text-xs text-muted-foreground">
                        {entry.character.scope}
                      </p>
                      <p className="mt-3 text-xs text-muted-foreground">
                        {entry.documents} evidence documents
                      </p>
                    </button>
                  ))}
                </div>
                {!matches.length && (
                  <p className="py-10 text-sm text-muted-foreground">No saved characters match.</p>
                )}
              </>
            )}
          </div>
        </main>
      )}
    </>
  );
}
