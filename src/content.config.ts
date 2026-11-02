import { defineCollection, z } from 'astro:content';
import { glob } from 'astro/loaders';

// One collection for all writing. Organization is purely tag-driven:
//   fixed tags  — 学习/科技/历史/哲学 (registered in src/data/tags.json)
//   topic tags  — agent, harness, ... (open, but registered for bilingual names)
//   series      — optional ordered reading sequence, with series_order
const posts = defineCollection({
  // Keep the authored "<slug>.<locale>.md" convention visible in entry ids:
  // the default id generator strips the inner dot, which would break the
  // ".en/.zh" pair-key logic used by routes, feeds, and listing views.
  loader: glob({
    pattern: '**/*.md',
    base: './src/content/posts',
    generateId: ({ entry }) => entry.replace(/\.md$/, ''),
  }),
  schema: z.object({
    title: z.string(),
    description: z.string().optional(),
    pubDate: z.coerce.date(),
    updated: z.coerce.date().optional(),
    tags: z.array(z.string()).default([]),
    series: z.string().optional(),
    series_order: z.number().int().positive().optional(),
    lang: z.enum(['en', 'zh']).default('en'),
    // Who wrote it. Mandatory so authorship is never left ambiguous:
    //   human     — written entirely by the author
    //   human-ai  — human/AI collaboration (research, drafting, or editing
    //               by an AI; the author verifies and owns every claim)
    authorship: z.enum(['human', 'human-ai']),
    draft: z.boolean().default(false),
  }),
});

// Long-form page copy lives here as Markdown, not as quoted strings in code.
// Same "<key>.<locale>.md" pairing convention as posts, e.g. home-intro.en.md.
// The body is the copy itself; only `lang` is carried in frontmatter.
const copy = defineCollection({
  loader: glob({
    pattern: '**/*.md',
    base: './src/content/copy',
    generateId: ({ entry }) => entry.replace(/\.md$/, ''),
  }),
  schema: z.object({
    lang: z.enum(['en', 'zh']),
  }),
});

export const collections = { posts, copy };
