import { useEffect, useState } from 'react';
import {
  ArrowLeft,
  BookOpen,
  Download,
  FilePenLine,
  Plus,
  RotateCcw,
  SlidersHorizontal,
} from 'lucide-react';
import type { Categories, WikiEntry, WikiSummary } from '../types.js';
import { Alert } from '../ui/alert.js';
import { Badge } from '../ui/badge.js';
import { Button } from '../ui/button.js';
import { Disclosure } from '../ui/disclosure.js';
import { Field } from '../ui/field.js';
import { IconButton } from '../ui/icon-button.js';
import { Input, Textarea } from '../ui/input.js';
import { cn, download, safeUrl } from '../ui/utils.js';
import { api } from './api.js';
import { Topbar } from './topbar.js';

const emptyForm = { name: '', work: '', scope: '', urls: '', supplied: '', summary: '' };

export function App() {
  const [view, setView] = useState<'wiki' | 'create'>('create');
  const [entries, setEntries] = useState<WikiSummary[]>([]);
  const [reading, setReading] = useState(true);
  const [warnings, setWarnings] = useState<string[]>([]);
  const [entry, setEntry] = useState<WikiEntry>();
  const [form, setForm] = useState(emptyForm);
  const [refresh, setRefresh] = useState(false);
  const [inputsOpen, setInputsOpen] = useState(false);
  const [query, setQuery] = useState('');
  const [draft, setDraft] = useState('');
  const [editing, setEditing] = useState(false);
  const [busy, setBusy] = useState(false);
  const [error, setError] = useState('');
  const [status, setStatus] = useState('');
  const [categories, setCategories] = useState<Categories>();
  const [evidence, setEvidence] = useState<string>();
  const [openSources, setOpenSources] = useState<string[]>([]);
  useEffect(() => {
    if (evidence)
      document
        .getElementById(`passage-${evidence.replace('/', '-')}`)
        ?.scrollIntoView({ block: 'center' });
  }, [evidence]);

  async function loadList() {
    setReading(true);
    try {
      const result = await api<{ entries: WikiSummary[]; warnings: string[] }>('/api/wiki');
      setEntries(result.entries);
      setWarnings(result.warnings);
    } finally {
      setReading(false);
    }
  }
  useEffect(() => {
    void Promise.all([loadList(), api<Categories>('/api/categories').then(setCategories)]).catch(
      (failure: Error) => setError(failure.message),
    );
  }, []);
  useEffect(() => {
    document.title = `${view === 'create' ? 'Create' : 'Wiki'} · Tower Generator`;
  }, [view]);
  async function run(task: () => Promise<void>) {
    setBusy(true);
    setError('');
    setStatus('');
    try {
      await task();
    } catch (failure) {
      setError(failure instanceof Error ? failure.message : String(failure));
    } finally {
      setBusy(false);
    }
  }
  function open(value: WikiEntry) {
    setEntry(value);
    setDraft(value.markdown);
    setEditing(false);
    setEvidence(undefined);
    setOpenSources([]);
    setView('wiki');
  }
  function change(key: keyof typeof form, value: string) {
    setForm({ ...form, [key]: value });
  }
  function showWiki() {
    setView('wiki');
    setError('');
    setStatus('');
    void run(loadList);
  }
  function showCreate() {
    setView('create');
    setInputsOpen(false);
    setError('');
    setStatus('');
  }
  function updateEvidence() {
    if (!entry) return;
    setForm({
      name: entry.metadata.name,
      work: entry.metadata.work,
      scope: entry.metadata.scope,
      urls: entry.documents
        .map((source) => source.url)
        .filter(Boolean)
        .join('\n'),
      supplied: '',
      summary: '',
    });
    setRefresh(false);
    setInputsOpen(true);
    setView('create');
    setError('');
    setStatus('');
  }
  const matches = entries.filter((item) =>
    `${item.name} ${item.work} ${item.aliases.join(' ')}`
      .toLowerCase()
      .includes(query.toLowerCase()),
  );
  const dirty = !!entry && draft !== entry.markdown;
  const researchRevision =
    entry && form.name === entry.metadata.name && form.work === entry.metadata.work
      ? entry.revision
      : undefined;

  return (
    <>
      <Topbar view={view} busy={busy} onWiki={showWiki} onCreate={showCreate} />
      <main className="px-[18px] py-10 sm:px-6 sm:py-[60px]">
        <div className={cn('mx-auto', view === 'create' ? 'max-w-[680px]' : 'max-w-[1000px]')}>
          <h1 className="mb-6 text-[1.6rem] font-semibold tracking-tight">
            {view === 'create' ? 'Create' : entry ? entry.metadata.name : 'Wiki'}
          </h1>
          {error && <Alert>{error}</Alert>}
          {status && (
            <p role="status" className="mb-4 text-xs text-success">
              {status}
            </p>
          )}
          {view === 'create' ? (
            <form
              onSubmit={(event) => {
                event.preventDefault();
                if (dirty) {
                  setError('Save or discard the edited entry before collecting more evidence.');
                  return;
                }
                const manual = !!(form.urls.trim() || form.supplied.trim());
                if (manual && (!form.work.trim() || !form.scope.trim())) {
                  setInputsOpen(true);
                  setError('Add the series and canon scope for the supplied references.');
                  return;
                }
                void run(async () => {
                  const result = await api<{
                    entry: WikiEntry;
                    fetched?: number;
                    reused?: number;
                    newRecords?: number;
                    updatedRecords?: number;
                  }>(
                    manual ? '/api/collect' : '/api/research',
                    manual
                      ? {
                          character: { name: form.name, work: form.work, scope: form.scope },
                          urls: form.urls
                            .split('\n')
                            .map((url) => url.trim())
                            .filter(Boolean),
                          supplied: form.supplied,
                          summary: form.summary,
                          refresh,
                          revision: researchRevision,
                        }
                      : { name: form.name, series: form.work, scope: form.scope },
                  );
                  open(result.entry);
                  setForm(emptyForm);
                  setRefresh(false);
                  setInputsOpen(false);
                  await loadList();
                  setStatus(
                    manual
                      ? `Evidence saved. ${result.fetched} sources updated, ${result.reused} reused.`
                      : `Research saved. ${result.newRecords} new records, ${result.updatedRecords} refined drafts.`,
                  );
                });
              }}
            >
              <div className="flex flex-wrap items-end gap-2 sm:flex-nowrap">
                <Field
                  label="Character name"
                  className="my-0 min-w-0 flex-1 basis-full sm:basis-auto"
                >
                  <Input
                    required
                    autoComplete="off"
                    className="h-11 text-base"
                    placeholder="Usopp"
                    value={form.name}
                    disabled={busy}
                    onChange={(event) => change('name', event.target.value)}
                  />
                </Field>
                <Button
                  type="submit"
                  variant="primary"
                  size="lg"
                  className="flex-1 sm:flex-none"
                  disabled={busy}
                >
                  {busy ? 'Researching…' : 'Research'}
                </Button>
              </div>
              <div className="mt-2 flex flex-wrap items-center gap-1">
                <Button
                  variant="ghost"
                  size="xs"
                  disabled={busy}
                  aria-expanded={inputsOpen}
                  aria-controls="research-inputs"
                  onClick={() => setInputsOpen(!inputsOpen)}
                >
                  <SlidersHorizontal /> Research inputs
                </Button>
                <Button variant="ghost" size="xs" disabled={busy} onClick={showWiki}>
                  <BookOpen /> Wiki
                </Button>
              </div>
              {inputsOpen && (
                <div
                  id="research-inputs"
                  className="mt-4 rounded-xl border border-border bg-card p-5"
                >
                  <p className="text-xs text-muted-foreground">
                    Research finds English references automatically. Add a series or canon scope to
                    narrow it, or supply your own references below.
                  </p>
                  <Field label="Series">
                    <Input
                      placeholder="One Piece"
                      value={form.work}
                      disabled={busy}
                      onChange={(event) => change('work', event.target.value)}
                    />
                  </Field>
                  <Field label="Canon scope">
                    <Input
                      placeholder="Manga through Dressrosa"
                      value={form.scope}
                      disabled={busy}
                      onChange={(event) => change('scope', event.target.value)}
                    />
                  </Field>
                  <Field label="Source URLs, one per line">
                    <Textarea
                      rows={3}
                      placeholder="Up to 10 relevant English sources"
                      value={form.urls}
                      disabled={busy}
                      onChange={(event) => change('urls', event.target.value)}
                    />
                  </Field>
                  <label className="my-4 flex items-start gap-2 text-xs text-muted-foreground">
                    <input
                      type="checkbox"
                      checked={refresh}
                      disabled={busy}
                      onChange={(event) => setRefresh(event.target.checked)}
                    />
                    Fetch saved URLs again
                  </label>
                  <Disclosure bare title="Supplied evidence and summary">
                    <Field label="Evidence">
                      <Textarea
                        rows={4}
                        placeholder="Paste source passages with their origin when pages cannot be fetched."
                        value={form.supplied}
                        disabled={busy}
                        onChange={(event) => change('supplied', event.target.value)}
                      />
                    </Field>
                    <Field label="Summary">
                      <Textarea
                        rows={3}
                        placeholder="Findings, source limits and open questions. Leave empty to preserve a saved summary."
                        value={form.summary}
                        disabled={busy}
                        onChange={(event) => change('summary', event.target.value)}
                      />
                    </Field>
                  </Disclosure>
                </div>
              )}
            </form>
          ) : entry ? (
            <>
              <div className="mb-3 flex flex-wrap items-center gap-2 text-xs text-muted-foreground">
                <span>{entry.metadata.work}</span>
                <Badge variant={entry.metadata.status === 'reviewed' ? 'success' : 'warning'}>
                  {entry.metadata.status}
                </Badge>
                <span>
                  {entry.documents.length} source{entry.documents.length === 1 ? '' : 's'}
                </span>
                <span>
                  {entry.metadata.abilities.length} record
                  {entry.metadata.abilities.length === 1 ? '' : 's'}
                </span>
              </div>
              <p className="mb-3 text-[13px] text-muted-foreground">{entry.metadata.scope}</p>
              <p className="mb-5 break-all font-mono text-xs text-muted-foreground">
                wiki/{entry.key}/README.md
              </p>
              {entry.reviewStale && (
                <Alert>The files changed since review. Review this draft again.</Alert>
              )}
              <div className="mb-6 flex flex-wrap gap-2">
                <Button
                  disabled={busy}
                  onClick={() => {
                    if (dirty) {
                      setError('Save the edited entry before opening another character.');
                      return;
                    }
                    setEntry(undefined);
                    setEditing(false);
                    void run(loadList);
                  }}
                >
                  <ArrowLeft /> Wiki
                </Button>
                <Button disabled={busy} onClick={() => setEditing(!editing)}>
                  <FilePenLine />
                  {editing ? 'Read entry' : 'Edit entry'}
                </Button>
                <Button disabled={busy || dirty} onClick={updateEvidence}>
                  Collect more evidence
                </Button>
                <Button
                  disabled={busy}
                  onClick={() => download('README.md', draft, 'text/markdown')}
                >
                  <Download /> Export Markdown
                </Button>
                <Button
                  variant="primary"
                  disabled={
                    busy ||
                    dirty ||
                    !entry.metadata.abilities.length ||
                    entry.metadata.status === 'reviewed'
                  }
                  onClick={() =>
                    void run(async () => {
                      open(
                        await api<WikiEntry>('/api/review', {
                          key: entry.key,
                          revision: entry.revision,
                        }),
                      );
                      await loadList();
                      setStatus(
                        'Manual review recorded. Classification tags retain their own review status.',
                      );
                    })
                  }
                >
                  Mark reviewed
                </Button>
              </div>
              {editing ? (
                <>
                  <Field label="Entry Markdown and YAML metadata">
                    <Textarea
                      aria-label="Entry Markdown"
                      rows={24}
                      className="font-mono text-xs"
                      value={draft}
                      disabled={busy}
                      onChange={(event) => setDraft(event.target.value)}
                    />
                  </Field>
                  <div className="my-4 flex flex-wrap gap-2">
                    <Button
                      variant="primary"
                      disabled={busy || !dirty}
                      onClick={() =>
                        void run(async () => {
                          open(
                            await api<WikiEntry>('/api/save', {
                              key: entry.key,
                              markdown: draft,
                              revision: entry.revision,
                            }),
                          );
                          await loadList();
                          setStatus('Draft saved.');
                        })
                      }
                    >
                      Save draft
                    </Button>
                    <Button
                      disabled={busy || !dirty}
                      onClick={() => {
                        setDraft(entry.markdown);
                        setStatus('Editor restored to the opened entry.');
                      }}
                    >
                      Discard edits
                    </Button>
                  </div>
                  <p className="text-xs text-muted-foreground">
                    Each ability, trait or equipment record needs a source and passage ID. Saving
                    changes resets the entry to draft.
                  </p>
                </>
              ) : (
                <>
                  <p className="mb-6 whitespace-pre-wrap text-[13px] leading-relaxed [overflow-wrap:anywhere]">
                    {entry.body
                      .replace(/^# [^\n]*\n?/, '')
                      .replace(/<!-- research:(?:begin|end) -->/g, '')
                      .trim()}
                  </p>
                  <div className="space-y-3">
                    {entry.metadata.abilities.map((ability) => (
                      <div
                        key={ability.id}
                        className="rounded-lg border border-border p-4 text-[13px]"
                      >
                        <div className="flex flex-wrap items-center gap-2">
                          <span>{ability.name}</span>
                          <Badge>{ability.kind}</Badge>
                        </div>
                        <p className="mt-3 leading-relaxed">{ability.description}</p>
                        <p className="mt-2 text-xs text-muted-foreground">{ability.scope}</p>
                        {ability.limitations && (
                          <p className="mt-2 text-xs text-muted-foreground">
                            {ability.limitations}
                          </p>
                        )}
                        <div className="mt-3 flex flex-wrap gap-1.5">
                          <Badge>{ability.classification?.delivery ?? 'unknown'}</Badge>
                          {(ability.classification?.functions ?? ['unknown']).map((tag) => (
                            <Badge key={tag}>{tag}</Badge>
                          ))}
                          <Badge
                            variant={
                              ability.classification?.status === 'reviewed' ? 'success' : 'warning'
                            }
                          >
                            {ability.classification?.status ?? 'suggested'} tags
                          </Badge>
                        </div>
                        <div className="mt-3 flex flex-wrap gap-2">
                          {ability.evidence.map((reference) => (
                            <button
                              key={`${reference.source}/${reference.passage}`}
                              className="break-all text-left text-xs text-link hover:underline"
                              onClick={() => {
                                const id = reference.source.slice(8, -3);
                                setOpenSources((current) => [...new Set([...current, id])]);
                                setEvidence(`${id}/${reference.passage}`);
                              }}
                            >
                              {reference.source}#{reference.passage}
                            </button>
                          ))}
                        </div>
                      </div>
                    ))}
                  </div>
                  {!entry.metadata.abilities.length && (
                    <p className="my-5 text-xs text-muted-foreground">
                      Evidence is saved. Edit the entry to add cited abilities, traits and equipment
                      before manual review.
                    </p>
                  )}
                </>
              )}
              <Disclosure bare title="Classification categories">
                <p className="mb-3 text-xs text-muted-foreground">
                  Tags describe sourced capabilities. Game rules and adaptations belong to a later
                  consumer. Suggested tags need manual review.
                </p>
                {categories &&
                  ['delivery', 'functions'].map((group) => (
                    <ul key={group} className="my-3 space-y-2 text-xs text-muted-foreground">
                      {Object.entries(categories[group as 'delivery' | 'functions']).map(
                        ([tag, description]) => (
                          <li key={tag}>
                            <span className="text-foreground">{tag}</span>: {description}
                          </li>
                        ),
                      )}
                    </ul>
                  ))}
              </Disclosure>
              {entry.documents.map((source) => (
                <Disclosure
                  key={source.id}
                  bare
                  title={source.title}
                  open={openSources.includes(source.id)}
                  onOpenChange={(value) => {
                    setOpenSources((current) =>
                      value
                        ? [...new Set([...current, source.id])]
                        : current.filter((id) => id !== source.id),
                    );
                    if (!value && evidence?.startsWith(`${source.id}/`)) setEvidence(undefined);
                  }}
                >
                  {source.url && safeUrl(source.url) && (
                    <a
                      href={source.url}
                      target="_blank"
                      rel="noreferrer"
                      className="break-all text-xs"
                    >
                      {source.url}
                    </a>
                  )}
                  <p className="my-3 text-xs text-muted-foreground">
                    {source.access === 'retrieved'
                      ? `Retrieved ${source.retrievedAt?.slice(0, 10)}`
                      : 'Supplied evidence'}
                    {source.excerpt ? ' · Retained excerpt' : ''}
                  </p>
                  <div className="max-h-[500px] space-y-4 overflow-auto">
                    {source.passages.map((passage) => (
                      <div
                        key={passage.id}
                        id={`passage-${source.id}-${passage.id}`}
                        className={cn(
                          'rounded-md p-2',
                          evidence === `${source.id}/${passage.id}` && 'bg-accent',
                        )}
                      >
                        <p className="mb-2 break-all font-mono text-[11px] text-muted-foreground">
                          {passage.id}
                        </p>
                        {source.retainedPassages?.some(
                          (retained) => retained.id === passage.id,
                        ) && (
                          <p className="mb-2 text-xs text-warning">
                            Retained from an earlier retrieval because a finding cites it.
                          </p>
                        )}
                        <p className="whitespace-pre-wrap text-xs leading-relaxed [overflow-wrap:anywhere]">
                          {passage.text}
                        </p>
                      </div>
                    ))}
                  </div>
                </Disclosure>
              ))}
            </>
          ) : (
            <>
              <div className="flex items-end gap-2">
                <Field label="Find a character" className="min-w-0 flex-1">
                  <Input
                    placeholder="Character or anime/manga"
                    value={query}
                    onChange={(event) => setQuery(event.target.value)}
                  />
                </Field>
                <IconButton
                  label="Reload local Wiki"
                  disabled={busy}
                  className="mb-3"
                  onClick={() => void run(loadList)}
                >
                  <RotateCcw />
                </IconButton>
              </div>
              {warnings.map((warning) => (
                <Alert key={warning}>{warning}</Alert>
              ))}
              <div className="mt-6 grid gap-3 sm:grid-cols-2 lg:grid-cols-3">
                {matches.map((item) => (
                  <button
                    key={item.key}
                    disabled={busy}
                    className="min-w-0 rounded-lg border border-border bg-card p-4 text-left transition-colors hover:bg-accent focus-visible:ring-2 focus-visible:ring-ring/60"
                    onClick={() =>
                      void run(async () =>
                        open(
                          await api<WikiEntry>(`/api/entry?key=${encodeURIComponent(item.key)}`),
                        ),
                      )
                    }
                  >
                    <div className="flex flex-wrap items-center gap-2 text-[13px]">
                      <BookOpen className="size-4 text-muted-foreground" />
                      {item.name}
                      <Badge variant={item.status === 'reviewed' ? 'success' : 'warning'}>
                        {item.status}
                      </Badge>
                    </div>
                    <p className="mt-2 text-xs text-muted-foreground">{item.work}</p>
                    <p className="mt-2 line-clamp-2 text-xs text-muted-foreground">{item.scope}</p>
                    <p className="mt-3 text-xs text-muted-foreground">
                      {item.sources} source{item.sources === 1 ? '' : 's'} · {item.abilities} record
                      {item.abilities === 1 ? '' : 's'}
                    </p>
                  </button>
                ))}
              </div>
              {reading ? (
                <p role="status" className="my-8 text-sm text-muted-foreground">
                  Reading local Wiki…
                </p>
              ) : (
                !matches.length && (
                  <p className="my-8 text-sm text-muted-foreground">
                    {entries.length
                      ? 'No characters match.'
                      : 'The local Wiki is empty. Collect evidence to start an entry.'}
                  </p>
                )
              )}
              <Button className="mt-6" disabled={busy} onClick={showCreate}>
                <Plus /> Create
              </Button>
            </>
          )}
        </div>
      </main>
    </>
  );
}
