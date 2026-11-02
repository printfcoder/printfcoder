#!/usr/bin/env node
// Build-time generator for the diagram readers embedded in articles.
//
// Source of truth:  architecture/<slug>.<locale>.json
//                   (archify JSON, authored and reviewed in git)
// Generated output: public/architecture-assets/<slug>.<locale>.html
//                   (self-contained readers, gitignored; Astro copies public/
//                   into dist/)
//
// Readers are deliberately NOT committed: each is a large (~0.8MB) generated
// artifact produced by the vendored archify renderer (tools/archify).
// Runs automatically before `astro build`; can also run directly:
//   node tools/render-architecture.mjs

import { execFileSync } from 'node:child_process';
import {
  mkdtempSync,
  readdirSync,
  readFileSync,
  rmSync,
  mkdirSync,
  writeFileSync,
  unlinkSync,
} from 'node:fs';
import { tmpdir } from 'node:os';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';

const here = dirname(fileURLToPath(import.meta.url));
const appRoot = resolve(here, '..');
const sourceDir = join(appRoot, 'architecture');
const outputDir = join(appRoot, 'public', 'architecture-assets');
const archifyBin = join(appRoot, 'tools', 'archify', 'bin', 'archify.mjs');

// Diagram types this site currently authors. Keep this list aligned with the
// renderers vendored under tools/archify/renderers.
const supportedTypes = new Set([
  'architecture',
  'sequence',
  'workflow',
  'dataflow',
  'lifecycle',
]);

const candidates = readdirSync(sourceDir)
  .filter((name) => /\.(en|zh)\.json$/.test(name))
  .sort();

if (candidates.length === 0) {
  console.error(`No diagram candidates found in ${sourceDir}`);
  process.exit(1);
}

// Only readers are regenerated: exported PNGs are committed and must survive
// a CI build (which has no Chrome to re-export them). Clean stale *.html.
mkdirSync(outputDir, { recursive: true });
for (const stale of readdirSync(outputDir).filter((f) => f.endsWith('.html'))) {
  unlinkSync(join(outputDir, stale));
}
const tmpDir = mkdtempSync(join(tmpdir(), 'archify-candidates-'));

let rendered = 0;
try {
  for (const name of candidates) {
    const candidate = JSON.parse(readFileSync(join(sourceDir, name), 'utf8'));
    const diagramType = candidate.diagram_type || 'architecture';
    if (!supportedTypes.has(diagramType)) {
      throw new Error(`${name} uses unsupported diagram_type: ${diagramType}`);
    }
    const outName = name.replace(/\.json$/, '.html');
    const sanitizedPath = join(tmpDir, name);
    writeFileSync(sanitizedPath, JSON.stringify(candidate));
    execFileSync(
      process.execPath,
      [
        archifyBin,
        'render',
        diagramType,
        sanitizedPath,
        join(outputDir, outName),
        '--quality',
        'showcase',
      ],
      { stdio: ['ignore', 'ignore', 'inherit'] },
    );
    rendered += 1;
  }
} finally {
  rmSync(tmpDir, { recursive: true, force: true });
}

console.log(
  `rendered ${rendered} diagram readers -> public/architecture-assets`,
);
