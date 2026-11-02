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
      role: '后端 / 平台架构',
      intro:
        '后端和平台做了十几年：服务框架、资源访问层、稳定性治理、AI Coding 基建都干过。现在做去中心化、点对点的系统——没有中心化服务器，数据归用户自己，agent 跑在用户自己的节点上。',
      now: {
        eyebrow: '现在',
        title: '手上在做的事',
        body: '是个去中心化的社交网络，内置 AI agent，点对点、可以自己部署，社交关系和数据都在用户手里。我在里面主要负责客户端和 Station 之间的联邦架构，另外一摊是怎么让 agent 真正参与研发流程。',
      },
      track: {
        eyebrow: '经历',
        title: '十三年，一直在做后端和平台',
      },
      timeline: [
        {
          phase: '前半段',
          title: '保险与金融',
          body: [
            '保险 SaaS 从 0 到 1，多租户、风控、审计都趟过；后来做金融的服务化框架，注册发现、熔断限流、配置治理、链路追踪一套落下来。',
            '这段想明白一件事：平台好不好不看功能多少，看出事的时候稳不稳、查不查得到。',
          ],
        },
        {
          phase: '后半段 · 平台',
          title: '某大型国际互联网公司 · 统一研发底座',
          body: [
            '负责统一资源访问层，RPC、数据库、缓存、消息总线收到一套调用模型里；配置热更新走新老客户端并存，流量按比例切，参数下发错了也能灰度回退。',
            '开放平台原来六个域名、接口各写各的，收成一个网关：两百多个接口迁移，鉴权错误码统一，全程双跑、灰度、监控对比，没敢直接切流。',
            '稳定性也归我：容量、压测、告警分级降噪、值班 SOP 做成闭环，核心链路 SLA 从 99.8% 提到 99.96%。',
          ],
        },
        {
          phase: '后半段 · 泛智能架构',
          title: 'AI Coding 基建',
          body: [
            '智能编码助手的服务端我是 Owner。内部知识库做向量化，但检索不盲信纯向量：语义和关键词一起查，再用 RRF 重排，解决专有名词召回不准；回答经 MCP 协议接进 IDE，研发查文档从挨个问人变成随手问。',
            '周边的 agent 工具链也一起做：可复用的 Agent Skill、按需求文档出框架代码的 SDD，还有 DevOps 和研发项目管理用的 Sub-Agent。',
          ],
        },
        {
          phase: '现在',
          title: '去中心化、点对点的系统',
          body: [
            '主线是 Peers Touch：点对点、可自托管的社交网络，内置 AI agent，数据和社交关系不押在任何一家公司的服务器上。',
            'Go-Micro China 和 Stack-RPC 的中文教程一直在更新，写了这么多年踩过的坑，能让后来的 Go 后端少绕几个弯。',
          ],
        },
      ],
      work: {
        eyebrow: '作品',
        title: '现在主要在维护的',
        items: [
          {
            title: 'Peers Touch',
            body: '去中心化的社交网络，内置 AI agent。点对点、能自托管，数据归用户，不是把中心化服务器换个皮。',
            foot: 'TYPESCRIPT · GO · 联邦架构',
            href: SITE.peersTouch,
          },
        ],
      },
      learning: {
        eyebrow: '学习',
        prefix: '技术是饭碗，也读点 ',
        suffix: '，都记在这儿',
        sub: '读过的书、查过的资料、想明白的问题。中文的用中文写，英文的用英文写，不硬翻。',
      },
      contact: { github: 'github.com/printfcoder', email: SITE.email, rss: 'RSS 订阅' },
    },
    learn: {
      eyebrow: '学习',
      title: '工程、历史、哲学。',
      subtitle: '技术是本分，历史是搞明白人怎么走到今天，哲学是怕自己只长手艺不长见识。读过的、想过的，记在这里。',
      all: '全部',
    },
    blog: {
      eyebrow: '博客',
      title: '长文。',
      subtitle: '分布式系统、平台、agent，还有一些值得记下来的权衡。想清楚一件就写一件。',
    },
    empty: {
      mark: '∅ 还没发内容',
      text: '东西还在攒，第一批笔记和文章写好就放上来。',
    },
    article: { langNote: 'English' },
    notFound: {
      title: '页面没找到',
      body: '这个地址不存在，可能东西挪过位置，也可能压根没发过。',
      back: '← 回首页',
    },
  },
};
