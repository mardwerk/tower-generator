import { execFile } from 'node:child_process';
import { join } from 'node:path';
import type { Connect } from 'vite';
import { defaults, readExample, readJSON } from './example.js';

let writing = false;
const project = join(defaults, '..');
const routes = [
  '/api/example',
  '/api/tower',
  '/api/sources',
  '/api/source',
  '/api/research',
  '/api/import-source',
  '/api/build',
];

function operation(action: string, input: unknown): Promise<unknown> {
  return new Promise((resolve, reject) => {
    const child = execFile(
      'python3',
      ['-B', '-m', 'src.lab', action],
      { cwd: project, timeout: 120_000, maxBuffer: 8_000_000 },
      (error, stdout) => {
        try {
          const result = JSON.parse(stdout) as { error?: string };
          if (error || result.error)
            reject(new Error(result.error || 'The local operation failed.'));
          else resolve(result);
        } catch {
          reject(new Error('The local operation failed. Check the terminal.'));
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
    send(403, { error: 'The Lab accepts local requests only.' });
    return;
  }
  if (
    !['localhost', '127.0.0.1', '[::1]'].includes(host.hostname) ||
    (request.headers.origin && request.headers.origin !== host.origin)
  ) {
    send(403, { error: 'The Lab accepts local requests only.' });
    return;
  }
  if (request.method === 'GET') {
    void (async () => {
      if (url.pathname === '/api/sources') return operation('list', {});
      if (url.pathname === '/api/source')
        return operation('load', { id: url.searchParams.get('id') });
      if (!['/api/example', '/api/tower'].includes(url.pathname))
        throw new Error('Use POST for this operation.');
      const example = readExample();
      if (url.pathname === '/api/example') return example;
      const name = url.searchParams.get('state');
      if (!example.states.some((state) => state.name === name)) {
        send(404, { error: 'Unknown Tower state.' });
        return;
      }
      return readJSON(join(defaults, 'game-data/Towers', example.design.id, `${name}.json`));
    })()
      .then((value) => {
        if (!response.writableEnded) send(200, value);
      })
      .catch((error: Error) =>
        send(url.pathname === '/api/example' ? 503 : 400, {
          error:
            url.pathname === '/api/example'
              ? 'The Tower is not prepared. Run python3 setup.py.'
              : error.message,
        }),
      );
    return;
  }
  const actions: Record<string, string> = {
    '/api/research': 'collect',
    '/api/import-source': 'import',
    '/api/build': 'build',
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
      const input: unknown = JSON.parse(Buffer.concat(chunks).toString('utf8'));
      send(200, await operation(action, input));
    })()
      .catch((error: Error) => send(400, { error: error.message }))
      .finally(() => {
        writing = false;
      });
  });
};
