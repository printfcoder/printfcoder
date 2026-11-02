import rss from '@astrojs/rss';
import { getCollection } from 'astro:content';
import type { APIContext } from 'astro';
import { SITE, categoryLabel, type Locale } from './i18n';

export async function buildFeed(locale: Locale, context: APIContext) {
  const [posts, notes] = await Promise.all([
    getCollection('posts', (p) => !p.data.draft && p.data.lang === locale),
    getCollection('notes', (n) => !n.data.draft && n.data.lang === locale),
  ]);

  const pairKey = (id: string) => id.replace(/\.(en|zh)$/, '');

  const items = [
    ...posts.map((p) => ({
      title: p.data.title,
      description: p.data.description ?? '',
      pubDate: p.data.pubDate,
      categories: ['blog'],
      link: `${locale === 'zh' ? '/zh' : ''}/posts/${pairKey(p.id)}/`,
    })),
    ...notes.map((n) => ({
      title: n.data.title,
      description: n.data.description ?? '',
      pubDate: n.data.pubDate,
      categories: [categoryLabel(n.data.category, locale)],
      link: `${locale === 'zh' ? '/zh' : ''}/notes/${pairKey(n.id)}/`,
    })),
  ].sort((a, b) => b.pubDate.valueOf() - a.pubDate.valueOf());

  return rss({
    title: `${SITE.name} — ${SITE.handle}`,
    description:
      locale === 'zh'
        ? '后端与平台工程师，做去中心化、点对点的系统。'
        : 'Backend and platform engineer building cloudless, decentralized computing.',
    site: context.site!,
    items,
    customData: `<language>${locale === 'zh' ? 'zh-CN' : 'en'}</language>`,
  });
}
