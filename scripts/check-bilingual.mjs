#!/usr/bin/env node
// Merge gate for the unified posts collection.
//
// Language pairs
//   Every *published* post must exist as a matched pair:
//     src/content/posts/<slug>.en.md
//     src/content/posts/<slug>.zh.md
//   and each file's `lang:` frontmatter must agree with its suffix.
//   Drafts (`draft: true`) are exempt.
//
// Organization model
//   tags   — open but registered in src/data/tags.json (en+zh display names)
//   series — optional, must exist in src/data/series.json; when present
//            series_order (positive integer) is required and must agree
//            across the en/zh pair
//
// Exits non-zero with a report so the pull_request workflow blocks merges.

import { readdirSync, readFileSync, existsSync } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { dirname, join } from 'node:path';

const here = dirname(fileURLToPath(import.meta.url));
const repoRoot = join(here, '..');
const CONTENT_DIR = join(repoRoot, 'src', 'content', 'posts');

const seriesRegistry = JSON.parse(
  readFileSync(join(repoRoot, 'src/data/series.json'), 'utf8'),
);
const tagRegistry = JSON.parse(
  readFileSync(join(repoRoot, 'src/data/tags.json'), 'utf8'),
);

const seriesSlugs = new Set(seriesRegistry.map((s) => s.slug));
const tagSlugs = new Set(Object.keys(tagRegistry));

const failures = [];

// Registry integrity.
const seriesParts = new Map(); // series slug -> Set(part slug)
for (const s of seriesRegistry) {
  for (const lang of ['en', 'zh']) {
    if (!s.title?.[lang] || !s.description?.[lang]) {
      failures.push(`[registry] series "${s.slug}" missing ${lang} title/description`);
    }
  }
  if (s.parts !== undefined) {
    if (!Array.isArray(s.parts) || s.parts.length === 0) {
      failures.push(`[registry] series "${s.slug}" parts must be a non-empty array`);
      continue;
    }
    const partSlugs = new Set();
    for (const part of s.parts) {
      if (partSlugs.has(part.slug)) {
        failures.push(`[registry] series "${s.slug}" has duplicate part slug "${part.slug}"`);
      }
      partSlugs.add(part.slug);
      for (const lang of ['en', 'zh']) {
        if (!part.title?.[lang]) {
          failures.push(
            `[registry] series "${s.slug}" part "${part.slug}" missing ${lang} title`,
          );
        }
      }
      const outline = part.outline;
      if (outline !== undefined) {
        for (const lang of ['en', 'zh']) {
          if (!Array.isArray(outline[lang]) || outline[lang].length === 0) {
            failures.push(
              `[registry] series "${s.slug}" part "${part.slug}" outline missing ${lang} entries`,
            );
          } else if (outline[lang].some((line) => typeof line !== 'string' || !line.trim())) {
            failures.push(
              `[registry] series "${s.slug}" part "${part.slug}" outline ${lang} has an empty entry`,
            );
          }
        }
        if (
          Array.isArray(outline?.en) &&
          Array.isArray(outline?.zh) &&
          outline.en.length !== outline.zh.length
        ) {
          failures.push(
            `[registry] series "${s.slug}" part "${part.slug}" outline en/zh length differs (${outline.en.length} vs ${outline.zh.length})`,
          );
        }
      }
    }
    seriesParts.set(s.slug, partSlugs);
  }
}
for (const [tag, names] of Object.entries(tagRegistry)) {
  for (const lang of ['en', 'zh']) {
    if (!names?.[lang]) failures.push(`[registry] tag "${tag}" missing ${lang} name`);
  }
}

function frontmatter(raw) {
  const match = raw.match(/^---\r?\n([\s\S]*?)\r?\n---/);
  return match ? match[1] : '';
}
const field = (fm, name) => {
  const m = fm.match(new RegExp(`^${name}:\\s*(.+?)\\s*$`, 'm'));
  return m ? m[1].trim() : null;
};
const isDraft = (fm) => /^draft:\s*true\s*$/m.test(fm);
const declaredLang = (fm) => field(fm, 'lang')?.replace(/['"]/g, '') ?? null;
const parseList = (raw) =>
  raw
    ? raw
        .replace(/^\[/, '')
        .replace(/\]$/, '')
        .split(',')
        .map((s) => s.trim().replace(/['"]/g, ''))
        .filter(Boolean)
    : [];

const pairs = new Map(); // slug -> { en: meta, zh: meta }
const drafts = new Set();

if (existsSync(CONTENT_DIR)) {
  for (const file of readdirSync(CONTENT_DIR).filter((f) => f.endsWith('.md'))) {
    const match = file.match(/^(.+)\.(en|zh)\.md$/);
    const fm = frontmatter(readFileSync(join(CONTENT_DIR, file), 'utf8'));

    if (!match) {
      failures.push(`[posts] ${file}: filename must end with '.en.md' or '.zh.md'`);
      continue;
    }

    const [, slug, suffixLang] = match;
    const meta = {
      lang: declaredLang(fm),
      tags: parseList(field(fm, 'tags')),
      series: field(fm, 'series')?.replace(/['"]/g, '') ?? null,
      seriesOrder: field(fm, 'series_order'),
      authorship: field(fm, 'authorship')?.replace(/['"]/g, '') ?? null,
    };

    if (isDraft(fm)) {
      drafts.add(`${slug}.${suffixLang}`);
      continue;
    }

    if (meta.lang !== suffixLang) {
      failures.push(
        `[posts] ${file}: frontmatter lang is '${meta.lang ?? 'missing'}' but suffix says '${suffixLang}'`,
      );
    }

    if (meta.authorship !== 'human' && meta.authorship !== 'human-ai') {
      failures.push(
        `[posts] ${file}: authorship must be 'human' or 'human-ai' (got '${meta.authorship ?? 'missing'}')`,
      );
    }

    for (const tag of meta.tags) {
      if (!tagSlugs.has(tag)) {
        failures.push(
          `[posts] ${file}: tag '${tag}' is not registered in src/data/tags.json`,
        );
      }
    }

    if (meta.series) {
      if (!seriesSlugs.has(meta.series)) {
        failures.push(
          `[posts] ${file}: series '${meta.series}' is not registered in src/data/series.json`,
        );
      }
      if (meta.seriesOrder === null) {
        failures.push(
          `[posts] ${file}: series '${meta.series}' requires a positive integer series_order`,
        );
      }
    }
    if (!meta.series && meta.seriesOrder !== null) {
      failures.push(`[posts] ${file}: series_order set without series`);
    }

    if (!pairs.has(slug)) pairs.set(slug, {});
    pairs.get(slug)[suffixLang] = meta;
  }
}

for (const [slug, langs] of pairs) {
  for (const lang of ['en', 'zh']) {
    if (!langs[lang]) {
      failures.push(
        `[posts] ${slug}.${lang}.md is missing — published content requires both English and Chinese. ` +
          `Mark it 'draft: true' if the translation is not ready.`,
      );
    }
  }
  if (langs.en && langs.zh) {
    if (langs.en.series !== langs.zh.series || langs.en.seriesOrder !== langs.zh.seriesOrder) {
      failures.push(
        `[posts] ${slug}: series/series_order differ between en (${langs.en.series ?? '-'} #${langs.en.seriesOrder ?? '-'}) and zh (${langs.zh.series ?? '-'} #${langs.zh.seriesOrder ?? '-'})`,
      );
    }
    // A post that belongs to a catalogued series must itself be catalogued
    // as one of its parts, or the series TOC cannot link to it.
    const s = langs.en.series;
    if (s && seriesParts.has(s) && !seriesParts.get(s).has(slug)) {
      failures.push(
        `[posts] ${slug}: belongs to series "${s}" but is not registered in its parts catalog (src/data/series.json)`,
      );
    }
  }
}

// Long-form page copy: same pairing rule as posts, no tags/series semantics.
const COPY_DIR = join(repoRoot, 'src', 'content', 'copy');
const copyPairs = new Map();
if (existsSync(COPY_DIR)) {
  for (const file of readdirSync(COPY_DIR).filter((f) => f.endsWith('.md'))) {
    const match = file.match(/^(.+)\.(en|zh)\.md$/);
    if (!match) {
      failures.push(`[copy] ${file}: filename must end with '.en.md' or '.zh.md'`);
      continue;
    }
    const [, slug, suffixLang] = match;
    const fm = frontmatter(readFileSync(join(COPY_DIR, file), 'utf8'));
    if (declaredLang(fm) !== suffixLang) {
      failures.push(
        `[copy] ${file}: frontmatter lang is '${declaredLang(fm) ?? 'missing'}' but suffix says '${suffixLang}'`,
      );
    }
    if (!copyPairs.has(slug)) copyPairs.set(slug, new Set());
    copyPairs.get(slug).add(suffixLang);
  }
  for (const [slug, langs] of copyPairs) {
    for (const lang of ['en', 'zh']) {
      if (!langs.has(lang)) failures.push(`[copy] ${slug}.${lang}.md is missing — page copy requires both languages.`);
    }
  }
}

if (drafts.size > 0) {
  console.log(`Skipping ${drafts.size} draft file(s): ${[...drafts].join(', ')}`);
}

if (failures.length > 0) {
  console.error(`\n✖ Content gate failed with ${failures.length} issue(s):\n`);
  for (const failure of failures) console.error(`  - ${failure}`);
  process.exit(1);
}

console.log('✓ Content gate passed: language pairs, tags, and series are consistent.');
