import { defineConfig } from 'astro/config';
import tailwindcss from '@tailwindcss/vite';
import sitemap from '@astrojs/sitemap';

// The profile repository serves Pages from https://printfcoder.github.io/printfcoder/.
// A custom domain or a printfcoder.github.io deployment repo only changes `site`/`base`.
export default defineConfig({
  site: 'https://printfcoder.github.io',
  base: '/printfcoder',
  integrations: [sitemap()],
  vite: {
    plugins: [tailwindcss()],
  },
});
