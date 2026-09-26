// Character lookups through AuthoringJobs against a scripted API. No model
// call: research, library and prepare are answered from fixtures.
import type { LibraryEntry, Sources } from '../../api/contract.js';
import type { api } from '../../api/client.js';
import { emptyRequest, type Revision } from '../../api/artifacts.js';
import {
  AuthoringJobs,
  lookupFor,
  type AuthoringJob,
  type Choice,
  type Lookup,
} from './authoring-jobs.js';

type Call = { endpoint: string; body: unknown };

const character = { name: 'Monkey D. Luffy', work: 'One Piece', scope: '' };
const sources = (id: string, text: string): Sources => ({
  schemaVersion: '1',
  kind: 'sources',
  query: 'Luffy',
  retrievedAt: '2026-09-01T00:00:00Z',
  character,
  documents: [
    {
      id,
      kind: 'source',
      text,
      origin: { location: `https://example.org/${id}`, access: 'retrieved', note: null },
    },
  ],
});
const earlier = sources('earlier', 'Saved reference.');
const extended: Sources = {
  ...sources('fresh', 'Fresh reference.'),
  documents: [...sources('fresh', 'Fresh reference.').documents, ...earlier.documents],
};
const entry = (id: string, work: string): LibraryEntry => ({
  id,
  savedAt: '2026-09-01T00:00:00Z',
  kind: 'sources',
  character: { ...character, work },
  artifactId: id,
  path: `${id}.json`,
  query: 'Luffy',
});
const pages: Choice[] = [
  { id: 11, name: 'Monkey D. Luffy', description: 'One Piece' },
  { id: 12, name: 'Luffy (film)', description: 'Film character' },
];

/** A scripted API: research answers choices until a page is chosen. */
function scripted(entries: LibraryEntry[], ambiguous: boolean) {
  const calls: Call[] = [];
  const respond = (endpoint: string, body: unknown): unknown => {
    switch (endpoint) {
      case 'library':
        return { directory: 'library', entries };
      case 'library/load':
        return { artifact: earlier };
      case 'research': {
        const lookup = body as { choice?: number; previous?: Sources };
        if (ambiguous && lookup.choice === undefined) return { kind: 'choices', choices: pages };
        return lookup.previous ? extended : sources('fresh', 'Fresh reference.');
      }
      case 'library/save':
        return (body as { artifact: Sources }).artifact;
      case 'prepare':
        return { kind: 'prepared', request: { ...emptyRequest(), character } };
    }
    throw new Error(`unexpected ${endpoint}`);
  };
  const fake = async (endpoint: string, body?: unknown) => {
    calls.push({ endpoint, body });
    return respond(endpoint, body);
  };
  let jobs: AuthoringJob[] = [];
  const manager = new AuthoringJobs({
    api: fake as typeof api,
    changed: (next) => (jobs = next),
    revision: () => undefined,
    complete: async () => undefined,
    needsProvider: () => undefined,
  });
  const revision: Revision = {
    id: 'r1',
    label: 'Luffy',
    createdAt: '',
    request: emptyRequest(),
    artifact: null,
  };
  return {
    calls,
    job: () => jobs[0]!,
    lookup: (lookup: Omit<Lookup, 'name'>) =>
      manager.start(revision, { remaining: false, lookup: { name: 'Luffy', ...lookup } }),
    bodies: (endpoint: string) => calls.filter((call) => call.endpoint === endpoint),
  };
}

function equal(actual: unknown, expected: unknown, what: string) {
  const [a, b] = [JSON.stringify(actual), JSON.stringify(expected)];
  if (a !== b) throw new Error(`${what}:\n  got      ${a}\n  expected ${b}`);
}

export const tests: Record<string, () => Promise<void>> = {
  async 'an ambiguous refresh extends the saved Sources with the chosen page'() {
    const api = scripted([entry('saved', 'One Piece')], true);
    await api.lookup({ refresh: true });
    const offered = api.job().choices;
    equal(api.job().state, 'waiting', 'first lookup');
    equal(
      offered.map((choice) => [choice.id, choice.sourcesId, choice.refresh]),
      [
        [11, 'saved', true],
        [12, 'saved', true],
      ],
      'page choices carry the saved Sources and refresh',
    );
    await api.lookup(lookupFor(offered[0]!));
    equal(
      api.bodies('research').map((call) => call.body),
      [
        { name: 'Luffy', previous: earlier },
        { name: 'Luffy', choice: 11, previous: earlier },
      ],
      'research payloads',
    );
    equal(api.bodies('library/load').length, 2, 'saved Sources loaded for each request');
    equal(
      api.bodies('library/save').map((call) => call.body),
      [{ artifact: extended }],
      'saved Sources',
    );
    equal(api.bodies('prepare')[0]!.body, { sources: extended }, 'prepared Sources');
    equal(api.job().state, 'finished', 'second lookup');
  },
  async 'an unambiguous refresh extends the saved Sources'() {
    const api = scripted([entry('saved', 'One Piece')], false);
    await api.lookup({ refresh: true });
    equal(
      api.bodies('research').map((call) => call.body),
      [{ name: 'Luffy', previous: earlier }],
      'research payload',
    );
    equal(
      api.bodies('library/save').map((call) => call.body),
      [{ artifact: extended }],
      'saved Sources',
    );
    equal(api.job().state, 'finished', 'lookup');
  },
  async 'a page chosen without saved Sources skips the library'() {
    const api = scripted([], true);
    await api.lookup({});
    const offered = api.job().choices;
    equal(offered, pages, 'page choices are unchanged');
    equal(lookupFor(offered[0]!), { choice: 11 }, 'chosen lookup');
    await api.lookup(lookupFor(offered[0]!));
    equal(
      api.calls.slice(2).map((call) => call.endpoint),
      ['research', 'library/save', 'prepare'],
      'second request order',
    );
    equal(api.bodies('research')[1]!.body, { name: 'Luffy', choice: 11 }, 'research payload');
  },
  async 'a saved entry (negative) reuses its Sources without research'() {
    const api = scripted([entry('first', 'One Piece'), entry('second', 'Film')], false);
    await api.lookup({});
    const offered = api.job().choices;
    equal(
      offered.map((choice) => [choice.id, choice.sourcesId ?? null, choice.refresh ?? false]),
      [
        [-1, 'first', false],
        [-2, 'second', false],
        [0, null, true],
      ],
      'saved choices',
    );
    equal(lookupFor(offered[0]!), { sourcesId: 'first' }, 'chosen lookup');
    await api.lookup(lookupFor(offered[0]!));
    equal(api.bodies('research').length, 0, 'research calls');
    equal(api.bodies('prepare')[0]!.body, { sources: earlier }, 'prepared Sources');
  },
  async 'Find references again (zero) asks which saved Sources to extend'() {
    const api = scripted([entry('first', 'One Piece'), entry('second', 'Film')], false);
    await api.lookup({});
    const again = api.job().choices.find((choice) => choice.id === 0)!;
    equal(lookupFor(again), { refresh: true }, 'chosen lookup');
    await api.lookup(lookupFor(again));
    const offered = api.job().choices;
    equal(
      offered.map((choice) => [choice.id, choice.sourcesId, choice.refresh]),
      [
        [-1, 'first', true],
        [-2, 'second', true],
      ],
      'saved choices to extend',
    );
    equal(lookupFor(offered[0]!), { sourcesId: 'first', refresh: true }, 'chosen lookup');
    await api.lookup(lookupFor(offered[0]!));
    equal(
      api.bodies('research').map((call) => call.body),
      [{ name: 'Luffy', previous: earlier }],
      'research payload',
    );
  },
};
