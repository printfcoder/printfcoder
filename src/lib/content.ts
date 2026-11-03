// Content entry ids are authored as "<pair>.<locale>" (e.g. "foo.en" and
// "foo.zh" form one translated pair). Routes, feeds, and listing views all
// link by the locale-stripped pair key, so the rule lives here once.
export const pairKey = (id: string): string => id.replace(/\.(en|zh)$/, '');
