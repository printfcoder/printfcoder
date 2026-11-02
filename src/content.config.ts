import { defineCollection, z } from 'astro:content';
import { glob } from 'astro/loaders';

// Learning notes — short readings and study cards grouped by category.
const notes = defineCollection({
  loader: glob({ pattern: '**/*.md', base: './src/content/notes' }),
  schema: z.object({
    title: z.string(),
    description: z.string().optional(),
    pubDate: z.coerce.date(),
    category: z.enum(['tech', 'history', 'philosophy']),
    tags: z.array(z.string()).default([]),
    source: z.string().optional(),
    lang: z.enum(['en', 'zh']).default('zh'),
    draft: z.boolean().default(false),
  }),
});

// Blog — long-form essays.
const posts = defineCollection({
  loader: glob({ pattern: '**/*.md', base: './src/content/posts' }),
  schema: z.object({
    title: z.string(),
    description: z.string().optional(),
    pubDate: z.coerce.date(),
    updated: z.coerce.date().optional(),
    lang: z.enum(['en', 'zh']).default('en'),
    draft: z.boolean().default(false),
  }),
});

export const collections = { notes, posts };
