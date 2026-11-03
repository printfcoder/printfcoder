#!/usr/bin/env node
// Build a paste-ready syndication bundle from one authored post.
//
// The post Markdown stays the single source of truth. This script only:
//   1. strips frontmatter,
//   2. rewrites diagram figure placeholders to committed PNG URLs,
//   3. prepends a canonical-source header,
//   4. appends series index + GitHub links (juejin full repost),
//      or produces a short index topic (v2ex forum post).
//
// Usage:
//   node tools/publish/build-bundle.mjs <slug> --platform juejin [--locale zh]
//   node tools/publish/build-bundle.mjs <slug> --platform v2ex   [--locale zh]
//
// Output: .publish/<platform>-<slug>.<locale>.md (gitignored).
// Nothing is uploaded — publishing is always a manual, human-confirmed step.

import { existsSync, mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';

const here = dirname(fileURLToPath(import.meta.url));
const appRoot = resolve(here, '..', '..');

const SITE = 'https://me.shux.me';
const GITHUB = 'https://github.com/peers-labs/peers-touch';

function parseArgs(argv) {
  const [slug, ...rest] = argv;
  const args = { slug, platform: null, locale: 'zh' };
  for (let i = 0; i < rest.length; i += 1) {
    if (rest[i] === '--platform') args.platform = rest[i + 1];
    if (rest[i] === '--locale') args.locale = rest[i + 1];
  }
  if (!slug || !['juejin', 'v2ex'].includes(args.platform)) {
    console.error('Usage: build-bundle.mjs <slug> --platform juejin|v2ex [--locale zh|en]');
    process.exit(1);
  }
  if (!['zh', 'en'].includes(args.locale)) {
    console.error(`Unsupported locale: ${args.locale}`);
    process.exit(1);
  }
  return args;
}

function splitFrontmatter(raw) {
  const match = raw.match(/^---\r?\n([\s\S]*?)\r?\n---\r?\n?([\s\S]*)$/);
  if (!match) throw new Error('frontmatter not found');
  return { fm: match[1], body: match[2].trim() };
}

const field = (fm, name) => {
  const m = fm.match(new RegExp(`^${name}:\\s*(.+?)\\s*$`, 'm'));
  return m ? m[1].trim().replace(/^['"]|['"]$/g, '') : null;
};

function rewriteFigures(markdown, locale) {
  return markdown.replace(
    /<figure[^>]*class="diagram-embed"[^>]*data-slug="([^"]+)"[^>]*>\s*<figcaption>([\s\S]*?)<\/figcaption>\s*<\/figure>/g,
    (_full, slug, caption) =>
      `![${caption.trim()}](${SITE}/architecture-assets/${slug}.${locale}.png)`,
  );
}

function loadSeries(slug) {
  const registry = JSON.parse(
    readFileSync(join(appRoot, 'src', 'data', 'series.json'), 'utf8'),
  );
  return registry.find((s) => s.slug === slug) || null;
}

function postUrl(slug, locale) {
  return locale === 'zh' ? `${SITE}/zh/posts/${slug}/` : `${SITE}/posts/${slug}/`;
}

function blogUrl(locale) {
  return locale === 'zh' ? `${SITE}/zh/blog` : `${SITE}/blog`;
}

function juejinBundle({ fm, body, slug, locale }) {
  const title = field(fm, 'title');
  const seriesSlug = field(fm, 'series');
  const lines = [];

  lines.push(
    `> 本文首发于个人站点：[《${title}》](${postUrl(slug, locale)})。` +
      '掘金版保留全部正文与图示；后续篇目在个人站点更新，转载请注明出处。',
    '',
    body,
  );

  if (seriesSlug) {
    const series = loadSeries(seriesSlug);
    if (series) {
      lines.push('', '---', '', `**《${series.title.zh}》系列**`, '');
      series.parts.forEach((part, i) => {
        const current = part.slug === slug;
        const label = `${i + 1}. ${part.title.zh}${current ? '（本文）' : ''}`;
        lines.push(current ? `**${label}**` : label);
      });
    }
  }

  lines.push(
    '',
    '---',
    '',
    `系列目录与后续更新：[me.shux.me 博客](${blogUrl(locale)})  `,
    `项目源码：[Peers Touch on GitHub](${GITHUB})`,
  );

  return lines.join('\n');
}

function v2exBundle({ fm, slug, locale }) {
  const title = field(fm, 'title');
  const description = field(fm, 'description');
  const seriesSlug = field(fm, 'series');
  const series = seriesSlug ? loadSeries(seriesSlug) : null;

  const lines = [];
  lines.push(
    `我在写一个系列《${series?.title.zh ?? 'Agent 是如何工作的'}》，刚发了第一篇：[《${title}》](${postUrl(slug, locale)})。`,
    '',
    description,
    '',
  );

  if (series) {
    lines.push('计划篇目：', '');
    series.parts.forEach((part, i) => lines.push(`${i + 1}. ${part.title.zh}`));
    lines.push('');
  }

  lines.push(
    `全文、架构图（可交互）和后续更新都在个人站点：${blogUrl(locale)}  `,
    `研究样本是两个开源仓库，项目本身在 GitHub：${GITHUB}`,
    '',
    '欢迎对内容和论点拍砖。',
  );

  return lines.join('\n');
}

const { slug, platform, locale } = parseArgs(process.argv.slice(2));
const postPath = join(appRoot, 'src', 'content', 'posts', `${slug}.${locale}.md`);
if (!existsSync(postPath)) {
  console.error(`Post not found: ${postPath}`);
  process.exit(1);
}

const raw = readFileSync(postPath, 'utf8');
const { fm, body } = splitFrontmatter(raw);
if (field(fm, 'draft') === 'true') {
  console.error('Refusing to bundle a draft.');
  process.exit(1);
}

const rewritten = rewriteFigures(body, locale);
const bundle =
  platform === 'juejin'
    ? juejinBundle({ fm, body: rewritten, slug, locale })
    : v2exBundle({ fm, body: rewritten, slug, locale });

const outDir = join(appRoot, '.publish');
mkdirSync(outDir, { recursive: true });
const outPath = join(outDir, `${platform}-${slug}.${locale}.md`);
writeFileSync(outPath, `${bundle}\n`);

const title = field(fm, 'title');
console.log(`bundle written: ${outPath}`);
console.log(`title: ${title}`);
if (platform === 'v2ex') {
  console.log('suggested node: /go/create (程序员) — review node rules before posting');
}
