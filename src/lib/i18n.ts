// Bilingual UI dictionary. Every user-facing chrome string lives here;
// content collections carry their own language (see content.config.ts).

import tagRegistry from '../data/tags.json';

export const LOCALES = ['en', 'zh'] as const;
export type Locale = (typeof LOCALES)[number];
export const DEFAULT_LOCALE: Locale = 'en';

export const SITE = {
  name: 'Shu Xian',
  handle: 'printfcoder',
  email: 'i@shux.me',
  location: { en: 'Shenzhen, China', zh: '中国 · 深圳' } as Record<Locale, string>,
  github: 'https://github.com/printfcoder',
  peersTouch: 'https://github.com/peers-labs/peers-touch',
} as const;

// Curated tags surfaced as blog filter chips. Everything is a post; these
// fixed tags replace the former top-level 学习/科技/历史/哲学 sections.
export const BLOG_TAGS = ['learning', 'tech', 'history', 'philosophy'] as const;

// Bilingual display names for all registered tags.
export function tagLabel(slug: string, locale: Locale): string {
  const hit = (tagRegistry as Record<string, { en: string; zh: string }>)[slug];
  return hit ? hit[locale] : slug;
}

// Locale-aware path helpers. English lives at the root, Chinese under /zh.
export function lpath(locale: Locale, path = '/'): string {
  const normalized = path.startsWith('/') ? path : `/${path}`;
  if (locale === 'zh') return normalized === '/' ? '/zh' : `/zh${normalized}`;
  return normalized;
}

// Path of the same page in the other language. Article pages pass an explicit
// counterpart when the translation pair exists.
export function switchPath(locale: Locale, pathname: string, counterpart?: string): string {
  if (counterpart) return counterpart;
  if (locale === 'zh') {
    const stripped = pathname.replace(/^\/zh(?=\/|$)/, '');
    return stripped || '/';
  }
  return pathname === '/' ? '/zh' : `/zh${pathname}`;
}

interface WorkItem {
  title: string;
  body: string;
  foot: string;
  href: string;
  logo?: string;
}

interface UIShape {
  nav: { hello: string; blog: string };
  switchTo: string;
  footer: { rights: string };
  home: {
    eyebrow: string;
    role: string;
    metaDescription: string;
    now: { eyebrow: string; title: string };
    work: { eyebrow: string; title: string; items: WorkItem[] };
    writing: {
      eyebrow: string;
      title: string;
      sub: string;
      cta: string;
    };
    contact: { github: string; email: string; rss: string };
  };
  blog: {
    eyebrow: string;
    title: string;
    subtitle: string;
    all: string;
    filterLabel: string;
  };
  article: {
    seriesNav: string;
    upcoming: string;
    sections: string;
    expandAll: string;
    collapseAll: string;
    authorship: {
      human: string;
      'human-ai': string;
    };
  };
  empty: { mark: string; text: string };
  notFound: { title: string; body: string; back: string };
}

export const UI: Record<Locale, UIShape> = {
  en: {
    nav: { hello: 'Hello', blog: 'Blog' },
    switchTo: '中',
    footer: { rights: 'All rights reserved' },
    home: {
      eyebrow: "Hello, I'm",
      role: 'Backend / platform engineer · full-stack experience',
      metaDescription:
        'Backend and platform engineer building decentralized social networks and agent systems.',
      now: {
        eyebrow: 'Now',
        title: "What I'm building",
      },
      work: {
        eyebrow: 'Selected work',
        title: 'Things I maintain or have built',
        items: [
          {
            title: 'Peers Touch',
            body: 'A decentralized social networking service with built-in AI agents — peer-to-peer, self-hostable, user-owned social data.',
            foot: 'TYPESCRIPT · GO · FEDERATION',
            href: SITE.peersTouch,
            logo: '/peers-icon.png',
          },
        ],
      },
      writing: {
        eyebrow: 'Writing',
        title: 'Writing & study notes',
        sub: 'Engineering, history, philosophy — essays and reading notes, each kept in the language it was written in.',
        cta: 'All posts',
      },
      contact: { github: 'github.com/printfcoder', email: SITE.email, rss: 'RSS feed' },
    },
    blog: {
      eyebrow: 'Blog',
      title: 'Writing & study notes.',
      subtitle:
        'Distributed systems, platforms, agents — plus history and philosophy. Posts stay in the language they were written in.',
      all: 'All',
      filterLabel: 'Filter by tag',
    },
    article: {
      seriesNav: 'In this series',
      upcoming: 'Upcoming',
      sections: 'Show sections',
      expandAll: 'See all {count} parts',
      collapseAll: 'Collapse',
      authorship: {
        human: 'Written by hand',
        'human-ai': 'Human + AI',
      },
    },
    empty: {
      mark: '∅ nothing published',
      text: 'The first entries are on their way. Posts are written as paired Markdown files and appear here automatically.',
    },
    notFound: {
      title: 'Page not found',
      body: 'This path does not exist — it may have moved or never shipped.',
      back: '← Back to Hello',
    },
  },

  zh: {
    nav: { hello: '介绍', blog: '博客' },
    switchTo: 'EN',
    footer: { rights: '保留所有权利' },
    home: {
      eyebrow: '你好，我是',
      role: '后端 / 平台架构 · 全栈经验',
      metaDescription: '后端 / 平台工程师，做去中心化社交网络与 Agent 系统。',
      now: {
        eyebrow: '现在',
        title: '手上在做的事',
      },
      work: {
        eyebrow: '作品',
        title: '现在主要在维护的',
        items: [
          {
            title: 'Peers Touch',
            body: '去中心化的社交网络，内置 AI agent。点对点、能自托管，数据归用户，不是把中心化服务器换个皮。',
            foot: 'TYPESCRIPT · GO · 联邦架构',
            href: SITE.peersTouch,
            logo: '/peers-icon.png',
          },
        ],
      },
      writing: {
        eyebrow: '写作',
        title: '文章和学习札记',
        sub: '技术是饭碗，也读点历史和哲学。读过的、踩过的、想明白的，都在博客里，用标签分。',
        cta: '全部文章',
      },
      contact: { github: 'github.com/printfcoder', email: SITE.email, rss: 'RSS 订阅' },
    },
    blog: {
      eyebrow: '博客',
      title: '文章与札记。',
      subtitle: '分布式系统、平台、agent，也有历史和哲学。想清楚一件就写一件，中文的用中文写，英文的用英文写。',
      all: '全部',
      filterLabel: '按标签筛选',
    },
    article: {
      seriesNav: '本系列目录',
      upcoming: '待写',
      sections: '展开小节',
      expandAll: '展开全部 {count} 篇',
      collapseAll: '收起',
      authorship: {
        human: '人手撰写',
        'human-ai': '人机共著',
      },
    },
    empty: {
      mark: '∅ 还没发内容',
      text: '东西还在攒，第一批文章写好就放上来。',
    },
    notFound: {
      title: '页面没找到',
      body: '这个地址不存在，可能东西挪过位置，也可能压根没发过。',
      back: '← 回首页',
    },
  },
};
