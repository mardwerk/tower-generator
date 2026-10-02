import { execFile } from 'node:child_process';
import { fileURLToPath } from 'node:url';
import type { Connect } from 'vite';

let writing = false;
const project = fileURLToPath(new URL('../../', import.meta.url));
const routes = [
  '/api/wiki',
  '/api/entry',
  '/api/categories',
  '/api/collect',
  '/api/save',
  '/api/review',
];

function operation(action: string, input: Record<string, unknown> = {}): Promise<unknown> {
  const args = ['-B', '-m', 'src.wiki', '--json', action];
  if (['show', 'save', 'review'].includes(action)) args.push(String(input.key ?? ''));
  if (action === 'list') args.push('--query', String(input.query ?? ''));
  if (['collect', 'save'].includes(action)) args.push('--input', '-');
  if (action === 'review' && input.revision)
    args.push('--expected-revision', String(input.revision));
  return new Promise((resolve, reject) => {
    const child = execFile(
      'python3',
      args,
      { cwd: project, timeout: 180_000, maxBuffer: 8_000_000 },
      (error, stdout) => {
        try {
          const result = JSON.parse(stdout) as { error?: string };
          if (error || result.error)
            reject(new Error(result.error || 'The research command failed.'));
          else resolve(result);
        } catch {
          reject(
            new Error(
              'The research command failed. Check the terminal and installed Python requirements.',
            ),
          );
        }
      },
    );
    child.stdin!.end(JSON.stringify(input));
  });
}

export const labAPI: Connect.NextHandleFunction = (request, response, next) => {
  const url = new URL(request.url ?? '/', 'http://localhost');
  if (!routes.includes(url.pathname)) return next();
  response.setHeader('Content-Type', 'application/json');
  response.setHeader('Cache-Control', 'no-store');
  const send = (status: number, value: unknown) => {
    response.statusCode = status;
    response.end(JSON.stringify(value));
  };
  let host: URL;
  try {
    host = new URL(`http://${request.headers.host}`);
  } catch {
    send(403, { error: 'The Wiki accepts local requests only.' });
    return;
  }
  if (
    !['localhost', '127.0.0.1', '[::1]'].includes(host.hostname) ||
    (request.headers.origin && request.headers.origin !== host.origin)
  ) {
    send(403, { error: 'The Wiki accepts local requests only.' });
    return;
  }
  if (request.method === 'GET') {
    const actions: Record<string, string> = {
      '/api/wiki': 'list',
      '/api/entry': 'show',
      '/api/categories': 'categories',
    };
    const action = actions[url.pathname];
    if (!action) {
      send(405, { error: 'Use POST for this operation.' });
      return;
    }
    void operation(action, {
      key: url.searchParams.get('key'),
      query: url.searchParams.get('query') ?? '',
    })
      .then((value) => send(200, value))
      .catch((error: Error) => send(400, { error: error.message }));
    return;
  }
  const actions: Record<string, string> = {
    '/api/collect': 'collect',
    '/api/save': 'save',
    '/api/review': 'review',
  };
  const action = actions[url.pathname];
  if (request.method !== 'POST' || !action) {
    send(405, { error: 'Method not allowed.' });
    return;
  }
  if (request.headers['content-type']?.split(';')[0] !== 'application/json') {
    send(415, { error: 'Send JSON for this operation.' });
    return;
  }
  if (writing) {
    send(409, { error: 'Another save is running. Try again when it finishes.' });
    return;
  }
  writing = true;
  let submitted = false;
  let size = 0;
  const chunks: Buffer[] = [];
  request.on('data', (chunk: Buffer) => {
    size += chunk.length;
    if (size <= 600_000) chunks.push(chunk);
  });
  request.on('aborted', () => {
    if (!submitted) writing = false;
  });
  request.on('error', () => {
    if (!submitted) writing = false;
  });
  request.on('end', () => {
    submitted = true;
    void (async () => {
      if (size > 600_000) {
        send(413, { error: 'Request is too large.' });
        return;
      }
      const input = JSON.parse(Buffer.concat(chunks).toString('utf8')) as Record<string, unknown>;
      if (!input || typeof input !== 'object' || Array.isArray(input))
        throw new Error('Expected a JSON object.');
      send(200, await operation(action, input));
    })()
      .catch((error: Error) => send(400, { error: error.message }))
      .finally(() => {
        writing = false;
      });
  });
};
