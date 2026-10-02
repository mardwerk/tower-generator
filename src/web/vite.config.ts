import { defineConfig } from 'vite';
import tailwindcss from '@tailwindcss/vite';
import { labAPI } from './server.js';

export default defineConfig({
  root: 'src/web',
  plugins: [
    tailwindcss(),
    {
      name: 'local-lab',
      configureServer(server) {
        server.middlewares.use(labAPI);
      },
      configurePreviewServer(server) {
        server.middlewares.use(labAPI);
      },
    },
  ],
  server: { port: 5173, strictPort: true },
  preview: { port: 5173, strictPort: true },
  build: { outDir: 'dist', emptyOutDir: true },
});
