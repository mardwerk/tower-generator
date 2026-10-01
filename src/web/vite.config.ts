import { join } from 'node:path';
import { defineConfig } from 'vite';
import type { Connect } from 'vite';
import tailwindcss from '@tailwindcss/vite';
import { defaults, readExample, readJSON } from './example.js';

const exampleAPI: Connect.NextHandleFunction = (request, response, next) => {
  const url = new URL(request.url ?? '/', 'http://localhost');
  if (!['/api/example', '/api/tower'].includes(url.pathname)) return next();
  response.setHeader('Content-Type', 'application/json');
  response.setHeader('Cache-Control', 'no-store');
  if (request.method !== 'GET') {
    response.statusCode = 405;
    response.end(JSON.stringify({ error: 'This example is read-only.' }));
    return;
  }
  try {
    const example = readExample();
    if (url.pathname === '/api/example') {
      response.end(JSON.stringify(example));
    } else {
      const name = url.searchParams.get('state');
      if (!example.states.some((state) => state.name === name)) {
        response.statusCode = 404;
        response.end(JSON.stringify({ error: 'Unknown Tower state.' }));
        return;
      }
      response.end(
        JSON.stringify(
          readJSON(join(defaults, 'game-data/Towers', example.design.id, `${name}.json`)),
        ),
      );
    }
  } catch {
    response.statusCode = 503;
    response.end(
      JSON.stringify({ error: 'The example is not prepared. Run python3 setup.py, then reload.' }),
    );
  }
};

export default defineConfig({
  root: 'src/web',
  plugins: [
    tailwindcss(),
    {
      name: 'local-example',
      configureServer(server) {
        server.middlewares.use(exampleAPI);
      },
      configurePreviewServer(server) {
        server.middlewares.use(exampleAPI);
      },
    },
  ],
  server: { port: 5173, strictPort: true },
  preview: { port: 5173, strictPort: true },
  build: { outDir: 'dist', emptyOutDir: true },
});
