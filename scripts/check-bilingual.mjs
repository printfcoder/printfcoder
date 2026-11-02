#!/usr/bin/env node
// Merge gate for content collections.
//
// Every *published* note/post must exist as a matched language pair:
//   src/content/<collection>/<slug>.en.md
//   src/content/<collection>/<slug>.zh.md
// and each file's `lang:` frontmatter must agree with its suffix.
// Drafts (`draft: true`) and unrelated files are exempt.
//
// Exits non-zero with a human-readable report when a pair is missing,
// so the pull_request workflow can block merges to main.

import { readdirSync, readFileSync, existsSync } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { dirname, join } from 'node:path';

const here = dirname(fileURLToPath(import.meta.url));
const repoRoot = join(here, '..');
const COLLECTIONS = ['notes', 'posts'];

const failures = [];

function frontmatter(raw) {
  const match = raw.match(/^---\r?\n([\s\S]*?)\r?\n---/);
  return match ? match[1] : '';
}

function isDraft(fm) {
  return /^draft:\s*true\s*$/m.test(fm);
}

function declaredLang(fm) {
  const m = fm.match(/^lang:\s*['"]?(\w+)['"]?\s*$/m);
  return m ? m[1] : null;
}

for (const collection of COLLECTIONS) {
  const dir = join(repoRoot, 'src', 'content', collection);
  if (!existsSync(dir)) continue;

  const files = readdirSync(dir).filter((f) => f.endsWith('.md'));
  const pairs = new Map(); // slug -> Set<lang>
  const drafts = new Set();

  for (const file of files) {
    const match = file.match(/^(.+)\.(en|zh)\.md$/);
    const raw = readFileSync(join(dir, file), 'utf8');
    const fm = frontmatter(raw);

    if (!match) {
      failures.push(
        `[${collection}] ${file}: filename must end with '.en.md' or '.zh.md' (matched pairs by slug)`,
      );
      continue;
    }

    const [, slug, suffixLang] = match;

    if (isDraft(fm)) {
      drafts.add(`${slug}.${suffixLang}`);
      continue;
    }

    const lang = declaredLang(fm);
    if (lang !== suffixLang) {
      const shown = lang ?? 'missing';
      failures.push(
        `[${collection}] ${file}: frontmatter lang is '${shown}' but suffix says '${suffixLang}'`,
      );
    }

    if (!pairs.has(slug)) pairs.set(slug, new Set());
    pairs.get(slug).add(suffixLang);
  }

  for (const [slug, langs] of pairs) {
    for (const lang of ['en', 'zh']) {
      if (!langs.has(lang)) {
        failures.push(
          `[${collection}] ${slug}.${lang}.md is missing — published content requires both English and Chinese. ` +
            `Mark it 'draft: true' if the translation is not ready.`,
        );
      }
    }
  }

  if (drafts.size > 0) {
    console.log(`[${collection}] skipping ${drafts.size} draft file(s): ${[...drafts].join(', ')}`);
  }
}

if (failures.length > 0) {
  console.error(`\n✖ Bilingual check failed with ${failures.length} issue(s):\n`);
  for (const failure of failures) console.error(`  - ${failure}`);
  console.error(
    '\nPolicy: every published note/post must ship in both English (.en.md) and Chinese (.zh.md).',
  );
  process.exit(1);
}

console.log('✓ Bilingual content check passed: every published entry has an en/zh pair.');
