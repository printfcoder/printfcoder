// Bilingual UI dictionary. Every user-facing chrome string lives here;
// content collections carry their own language (see content.config.ts).

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

// Learning category metadata. Slugs stay URL-safe and language-neutral.
export const CATEGORIES = [
  {
    slug: 'tech',
    label: { en: 'Tech', zh: '科技' },
    description: {
      en: 'Distributed systems, agents, and engineering craft.',
      zh: '分布式系统、agent 与工程实践。',
    },
  },
  {
    slug: 'history',
    label: { en: 'History', zh: '历史' },
    description: { en: 'Notes and readings from history.', zh: '历史阅读与札记。' },
  },
  {
    slug: 'philosophy',
    label: { en: 'Philosophy', zh: '哲学' },
    description: { en: 'Ideas, arguments, and slow thinking.', zh: '观念、论证与慢思考。' },
  },
] as const;

export type CategorySlug = (typeof CATEGORIES)[number]['slug'];

export function categoryLabel(slug: CategorySlug, locale: Locale): string {
  return CATEGORIES.find((c) => c.slug === slug)!.label[locale];
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

interface TimelineItem {
  phase: string;
  title: string;
  body: string[];
}

interface WorkItem {
  title: string;
  body: string;
  foot: string;
  href: string;
}

interface UIShape {
  nav: { hello: string; learn: string; blog: string };
  switchTo: string;
  footer: { rights: string };
  home: {
    eyebrow: string;
    role: string;
    intro: string;
    now: { eyebrow: string; title: string; body: string };
    track: { eyebrow: string; title: string };
    timeline: TimelineItem[];
    work: { eyebrow: string; title: string; items: WorkItem[] };
    learning: {
      eyebrow: string;
      prefix: string;
      suffix: string;
      sub: string;
    };
    contact: { github: string; email: string; rss: string };
  };
  learn: {
    eyebrow: string;
    title: string;
    subtitle: string;
    all: string;
  };
  blog: {
    eyebrow: string;
    title: string;
    subtitle: string;
  };
  empty: { mark: string; text: string };
  article: { langNote: string };
  notFound: { title: string; body: string; back: string };
}

export const UI: Record<Locale, UIShape> = {
  en: {
    nav: { hello: 'Hello', learn: 'Learning', blog: 'Blog' },
    switchTo: '中',
    footer: { rights: 'All rights reserved' },
    home: {
      eyebrow: "Hello, I'm",
      role: 'Backend & platform engineer',
      intro:
        'I build backend platforms and the infrastructure around them — and I now work on cloudless, decentralized computing: peer-to-peer systems where people own their data and agents run beside them.',
      now: {
        eyebrow: 'Now',
        title: "What I'm building",
        body: 'is a decentralized social network with built-in AI agents: peer-to-peer, self-hostable, with users retaining ownership of their social graph and data. My focus is the client–station federation architecture and agent-native developer workflows.',
      },
      track: {
        eyebrow: 'Track record',
        title: 'Thirteen years of backend and platforms',
      },
      timeline: [
        {
          phase: 'Early career',
          title: 'Insurance & finance platforms',
          body: [
            'Built multi-tenant insurance SaaS and service-oriented frameworks from zero: service discovery, circuit breaking, rate limiting, configuration governance, and end-to-end tracing.',
            'Learned that platforms are judged by reliability and auditability, not by how many features they ship.',
          ],
        },
        {
          phase: 'Platform engineering',
          title: 'Unified R&D foundation at a large global internet company',
          body: [
            'Owned a unified resource-access layer (RPC, DB, cache, message bus) with hot configuration reload, connection-pool versioning, and structured observability.',
            'Turned fragmented open-platform APIs into one gateway: hundreds of endpoints migrated via dual-run, grayscale rollout, and traffic comparison.',
            'Ran the stability program end to end — capacity, load testing, alert tiers, on-call SOPs — lifting core SLA from 99.8% to 99.96%.',
          ],
        },
        {
          phase: 'General intelligence architecture',
          title: 'Agentic coding infrastructure',
          body: [
            'Designed the MCP + RAG coding assistant: semantic + keyword hybrid retrieval with reciprocal-rank fusion, LLM orchestration, and IDE integration over a standard protocol.',
            'Built the surrounding agent stack — reusable agent skills, spec-driven code generation, and sub-agent tooling for DevOps and project workflows.',
          ],
        },
        {
          phase: 'Now',
          title: 'Cloudless, decentralized computing',
          body: [
            'Working on Peers Touch: a peer-to-peer, self-hostable social network with built-in AI agents, where people own their social data and connections.',
            'Maintaining open-source Go microservice education — the Go-Micro China & Stack-RPC tutorials.',
          ],
        },
      ],
      work: {
        eyebrow: 'Selected work',
        title: 'Things I maintain or have built',
        items: [
          {
            title: 'Peers Touch',
            body: 'A decentralized social networking service with built-in AI agents — peer-to-peer, self-hostable, user-owned social data.',
            foot: 'TYPESCRIPT · GO · FEDERATION',
            href: SITE.peersTouch,
          },
        ],
      },
      learning: {
        eyebrow: 'Learning',
        prefix: 'Readings across',
        suffix: '',
        sub: 'Short notes and study records, kept in their original language.',
      },
      contact: { github: 'github.com/printfcoder', email: SITE.email, rss: 'RSS feed' },
    },
    learn: {
      eyebrow: 'Learning',
      title: 'Engineering, history, philosophy.',
      subtitle: '工程、历史与哲学的阅读与实验记录。Notes stay in the language they were written in.',
      all: 'All',
    },
    blog: {
      eyebrow: 'Blog',
      title: 'Long-form writing.',
      subtitle:
        'Distributed systems, platforms, agents, and the trade-offs worth writing down. Articles stay in their original language.',
    },
    empty: {
      mark: '∅ nothing published',
      text: 'The first entries are on their way. Notes are written as paired Markdown files and appear here automatically.',
    },
    article: { langNote: '中文' },
    notFound: {
      title: 'Page not found',
      body: 'This path does not exist — it may have moved or never shipped.',
      back: '← Back to Hello',
    },
  },

  zh: {
    nav: { hello: '介绍', learn: '学习', blog: '博客' },
    switchTo: 'EN',
    footer: { rights: '保留所有权利' },
    home: {
      eyebrow: '你好，我是',
      role: '后端与平台工程师',
      intro:
        '我构建后端平台及其周边基础设施——现在专注于无云、去中心化计算：点对点系统，让用户拥有自己的数据，agent 在人身边运行。',
      now: {
        eyebrow: '现在',
        title: '我正在做的事',
        body: '是一个内置 AI agent 的去中心化社交网络：点对点、可自托管，用户拥有自己的社交图谱与数据。我的重点是客户端–工作站联邦架构与 agent 原生的研发工作流。',
      },
      track: {
        eyebrow: '经历',
        title: '十三年后端与平台工程',
      },
      timeline: [
        {
          phase: '早期',
          title: '保险与金融平台',
          body: [
            '从零构建多租户保险 SaaS 与服务化框架：服务发现、熔断限流、配置治理与全链路追踪。',
            '学到的一课：平台的价值由可靠性与可审计性衡量，而不是功能数量。',
          ],
        },
        {
          phase: '平台工程',
          title: '某大型国际互联网公司 · 统一研发底座',
          body: [
            '负责统一资源访问层（RPC、数据库、缓存、消息总线），支持配置热更新、连接池版本化与结构化可观测。',
            '将碎片化的开放平台接口收敛为统一网关：数百个接口通过双跑、灰度发布与流量比对完成迁移。',
            '端到端负责稳定性体系——容量、压测、告警分级、值班 SOP，核心 SLA 从 99.8% 提升至 99.96%。',
          ],
        },
        {
          phase: '泛智能架构',
          title: 'Agentic Coding 基建',
          body: [
            '设计 MCP + RAG 编码助手：语义与关键词混合检索、倒数排名融合（RRF）、LLM 编排，通过标准协议接入 IDE。',
            '建设配套 agent 技术栈：可复用的 Agent Skill、规格驱动代码生成（SDD）、面向 DevOps 与研发流程的子代理工具。',
          ],
        },
        {
          phase: '现在',
          title: '无云、去中心化计算',
          body: [
            '打造 Peers Touch：点对点、可自托管的社交网络，内置 AI agent，用户拥有自己的社交数据与连接。',
            '维护开源 Go 微服务教学：Go-Micro China 与 Stack-RPC 中文教程。',
          ],
        },
      ],
      work: {
        eyebrow: '精选作品',
        title: '我在维护与构建的东西',
        items: [
          {
            title: 'Peers Touch',
            body: '去中心化社交网络服务，内置 AI agent——点对点、可自托管，社交数据归用户所有。',
            foot: 'TYPESCRIPT · GO · 联邦架构',
            href: SITE.peersTouch,
          },
        ],
      },
      learning: {
        eyebrow: '学习',
        prefix: '横跨',
        suffix: '的阅读',
        sub: '短文与学习札记，保留写作时的原语言。',
      },
      contact: { github: 'github.com/printfcoder', email: SITE.email, rss: 'RSS 订阅' },
    },
    learn: {
      eyebrow: '学习',
      title: '工程、历史与哲学。',
      subtitle: '阅读与实验记录，文章以原语言呈现。',
      all: '全部',
    },
    blog: {
      eyebrow: '博客',
      title: '长文写作。',
      subtitle: '关于分布式系统、平台、agent 与值得记下的权衡，文章以原语言呈现。',
    },
    empty: {
      mark: '∅ 暂无内容',
      text: '首批文章正在路上。笔记以成对的 Markdown 文件编写，构建后自动出现在这里。',
    },
    article: { langNote: 'English',
    },
    notFound: {
      title: '页面未找到',
      body: '这个路径不存在——可能已移动，或从未发布。',
      back: '← 返回首页',
    },
  },
};
