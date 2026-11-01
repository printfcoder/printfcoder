import rss from '@astrojs/rss';
import { getCollection } from 'astro:content';
import type { APIContext } from 'astro';
import { SITE, withBase } from '../lib/site';

export async function GET(context: APIContext) {
  const posts = (await getCollection('writing', (p) => !p.data.draft)).sort(
    (a, b) => b.data.pubDate.valueOf() - a.data.pubDate.valueOf(),
  );

  return rss({
    title: `${SITE.name} — ${SITE.handle}`,
    description: SITE.description,
    site: new URL(withBase('/'), context.site!).toString(),
    items: posts.map((post) => ({
      title: post.data.title,
      description: post.data.description ?? '',
      pubDate: post.data.pubDate,
      link: withBase(`/writing/${post.id}/`),
    })),
    customData: `<language>en</language>`,
  });
}
