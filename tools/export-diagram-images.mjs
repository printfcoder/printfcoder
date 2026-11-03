#!/usr/bin/env node
// Export static PNGs from the rendered diagram readers.
//
// Readers (public/architecture-assets/*.html) are self-contained pages whose
// inline SVG carries an authored viewBox. External platforms (Juejin, etc.)
// cannot embed the HTML readers, so we screenshot each reader in headless
// Chrome at its authored size, 2x for retina sharpness.
//
// Unlike the HTML readers, PNGs are committed (see .gitignore): they must be
// reachable at https://me.shux.me/architecture-assets/<slug>.<locale>.png
// regardless of whether CI has Chrome installed.
//
// Usage: node tools/export-diagram-images.mjs
// Override the browser with CHROME_BIN.

import { execFileSync } from 'node:child_process';
import { existsSync, readdirSync, readFileSync, statSync, unlinkSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath, pathToFileURL } from 'node:url';

const here = dirname(fileURLToPath(import.meta.url));
const appRoot = resolve(here, '..');
const assetsDir = join(appRoot, 'public', 'architecture-assets');
const SCALE = 2;

const chromeBin =
  process.env.CHROME_BIN ||
  [
    '/Applications/Google Chrome.app/Contents/MacOS/Google Chrome',
    '/Applications/Chromium.app/Contents/MacOS/Chromium',
    '/usr/bin/google-chrome',
    '/usr/bin/chromium',
    '/usr/bin/chromium-browser',
  ].find((p) => existsSync(p));

if (!chromeBin) {
  console.error(
    'Chrome/Chromium not found. Install Chrome or set CHROME_BIN, then rerun.',
  );
  process.exit(1);
}

const readers = readdirSync(assetsDir).filter((f) => f.endsWith('.html'));
if (readers.length === 0) {
  console.error(`No readers found in ${assetsDir}. Run the diagram renderer first.`);
  process.exit(1);
}

// Drop PNGs whose reader was renamed or removed.
const liveStems = new Set(readers.map((f) => f.replace(/\.html$/, '')));
for (const png of readdirSync(assetsDir).filter((f) => f.endsWith('.png'))) {
  if (!liveStems.has(png.replace(/\.png$/, ''))) {
    unlinkSync(join(assetsDir, png));
    console.log(`removed stale ${png}`);
  }
}

let exported = 0;
for (const reader of readers) {
  const htmlPath = join(assetsDir, reader);
  const html = readFileSync(htmlPath, 'utf8');
  const viewBox = html.match(/<svg[^>]*viewBox="0 0 (\d+(?:\.\d+)?) (\d+(?:\.\d+)?)"/);
  if (!viewBox) {
    throw new Error(`${reader}: cannot read SVG viewBox, export aborted`);
  }
  const width = Math.ceil(Number(viewBox[1]));
  const height = Math.ceil(Number(viewBox[2]));
  const pngPath = htmlPath.replace(/\.html$/, '.png');

  execFileSync(
    chromeBin,
    [
      '--headless=new',
      '--disable-gpu',
      '--hide-scrollbars',
      `--force-device-scale-factor=${SCALE}`,
      `--window-size=${width},${height}`,
      `--screenshot=${pngPath}`,
      `${pathToFileURL(htmlPath).href}?theme=light&embed=1`,
    ],
    { stdio: ['ignore', 'ignore', 'inherit'], timeout: 60_000 },
  );

  if (!existsSync(pngPath) || statSync(pngPath).size === 0) {
    throw new Error(`${reader}: screenshot was not produced`);
  }
  exported += 1;
  console.log(`exported ${reader.replace(/\.html$/, '.png')} (${width}x${height} @${SCALE}x)`);
}

console.log(`exported ${exported} diagram PNGs -> public/architecture-assets`);
