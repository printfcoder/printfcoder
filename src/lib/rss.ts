import rss from '@astrojs/rss';
import { getCollection } from 'astro:content';
import type { APIContext } from 'astro';
import { SITE, tagLabel, type Locale } from './i18n';
import { pairKey } from './content';

export async function buildFeed(locale: Locale, context: APIContext) {
  const posts = await getCollection('posts', (p) => !p.data.draft && p.data.lang === locale);

  const items = posts
    .sort((a, b) => b.data.pubDate.valueOf() - a.data.pubDate.valueOf())
    .map((p) => ({
      title: p.data.title,
      description: p.data.description ?? '',
      pubDate: p.data.pubDate,
      categories: p.data.tags.map((tag) => tagLabel(tag, locale)),
      link: `${locale === 'zh' ? '/zh' : ''}/posts/${pairKey(p.id)}/`,
    }));

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
