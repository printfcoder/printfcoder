import type { CollectionEntry } from 'astro:content';

// Page copy is pre-rendered to HTML by the content pipeline. We inject it
// with set:html, so external anchors authored in Markdown need their new-tab
// attributes here (internal/anchor links are left alone).
const EXTERNAL_ANCHOR = /<a\s+href="(https?:\/\/[^"]+)"/g;

function withExternalLinkAttrs(html: string): string {
  return html.replace(
    EXTERNAL_ANCHOR,
    '<a href="$1" target="_blank" rel="noopener"',
  );
}

export function renderedCopyHtml(entry: CollectionEntry<'copy'> | undefined): string {
  if (!entry?.rendered) return '';
  return withExternalLinkAttrs(entry.rendered.html);
}
