// Runs the web client's *.test.ts files with node:test. Each file exports
// `tests`, named async cases; esbuild bundles it first, as for the build.
// Usage: node src/web/test.mjs
import { readdir } from 'node:fs/promises';
import { fileURLToPath } from 'node:url';
import { test } from 'node:test';
import { build } from 'esbuild';

const web = fileURLToPath(new URL('./', import.meta.url));
const files = (await readdir(web, { recursive: true }))
  .filter((file) => file.endsWith('.test.ts') && !file.startsWith('dist'))
  .sort();
for (const file of files) {
  const bundle = await build({
    entryPoints: [web + file],
    bundle: true,
    format: 'esm',
    platform: 'node',
    target: 'node22',
    write: false,
  });
  const source = `data:text/javascript;base64,${Buffer.from(bundle.outputFiles[0].text).toString('base64')}`;
  const { tests } = await import(source);
  for (const [name, run] of Object.entries(tests)) test(`${file}: ${name}`, run);
}
