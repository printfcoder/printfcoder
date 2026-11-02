import { defineConfig } from 'astro/config';
import tailwindcss from '@tailwindcss/vite';
import sitemap from '@astrojs/sitemap';

// Custom domain served from the repository root; public/CNAME binds me.shux.me.
// Falling back to github.io would require `base: '/printfcoder'` instead.
export default defineConfig({
  site: 'https://me.shux.me',
  integrations: [sitemap()],
  vite: {
    plugins: [tailwindcss()],
  },
});
