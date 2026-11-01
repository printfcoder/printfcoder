// Single source of truth for site identity and URL helpers.
export const SITE = {
  name: 'Shu Xian',
  handle: 'printfcoder',
  role: 'Backend & platform engineer',
  location: 'Shenzhen, China',
  email: 'i@shux.me',
  description:
    'Shu Xian (printfcoder) — backend and platform engineer building cloudless, decentralized computing.',
  github: 'https://github.com/printfcoder',
  peersTouch: 'https://github.com/peers-labs/peers-touch',
} as const;

// Astro injects the configured `base` through BASE_URL ("/printfcoder" here).
// Every internal href goes through this helper so a domain/base change is one-line.
export function withBase(path: string): string {
  const base = import.meta.env.BASE_URL.replace(/\/$/, '');
  const normalized = path.startsWith('/') ? path : `/${path}`;
  return `${base}${normalized}`;
}
