import type { APIContext } from 'astro';
import { buildFeed } from '../../lib/rss';

export function GET(context: APIContext) {
  return buildFeed('zh', context);
}
