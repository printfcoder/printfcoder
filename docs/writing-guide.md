# 内容写作指南（单一真源）

本站所有"怎么写内容"的规则只维护在这一份文档。README 与 AI skill
都只做入口，指向这里。规则有变更，改这一个文件。

技术栈：Astro 7 + Tailwind v4，纯静态，部署 GitHub Pages（me.shux.me）。
英文在根路径，中文在 `/zh`。

---

## 1. 两类内容

| 类型 | 目录 | 用途 |
|------|------|------|
| 博客文章 | `src/content/posts/` | 技术、历史、哲学等文章，有标签、可归系列、进 RSS |
| 页面长文案 | `src/content/copy/` | 首页简介、Now 段等固定散文，不进博客列表 |

短 UI 标签（导航、按钮、空状态）不在 Markdown 里，在
`src/lib/i18n.ts` 的 `UI` 字典，一般不动。

## 2. 铁律：中英成对

- 每篇已发布内容必须是两个文件：`<slug>.en.md` + `<slug>.zh.md`。
- 文件里的 frontmatter `lang:` 必须与后缀一致。
- 改了一份，另一份必须同步改，不允许长期单边漂移。
- 翻译没写完：把两份都先标 `draft: true`（posts 支持；copy 不允许半成品上线）。
- 文章"用什么语言写就用什么语言写"，但两份文件都必须存在——不是机翻腔，
  是同一内容的两种母语表达。
- 校验：`pnpm check:i18n`，CI 会拦；本地提交前必跑。

## 3. 博客 frontmatter

```yaml
---
title: 文章标题
description: 一句话摘要，用于列表与 SEO
pubDate: 2026-11-02
updated: 2026-11-10        # 可选
tags: [learning, tech]     # 必须先在 src/data/tags.json 注册
series: how-agents-work    # 可选；必须在 src/data/series.json 注册
series_order: 1            # 有 series 时必填，正整数，中英一致
lang: zh
authorship: human-ai       # 必填：human（人手撰写）或 human-ai（人机共著，作者审定）
draft: false
---
```

### 署名（authorship，必填）

每篇发布文章必须声明怎么写成的，不允许含糊：

- `human`：作者独立撰写，无 AI 参与。
- `human-ai`：人机共著——AI 参与研究、起草或校改，作者逐条核实并对结论负责。
  不设"纯 AI"档，那种内容不发布。页面头部会显示对应徽章（"人手撰写"/"人机共著"）。

正文用标准 Markdown。

### 标签（tags）

- 标签是**扁平**的，没有层级；分类已废除，不要再建"学习/博客"两套。
- 固定主标签四个：`learning` `tech` `history` `philosophy`；
  开放主题标签如 `agent` `harness`。
- 新标签流程：先在 `src/data/tags.json` 加 `{ "slug": { "en": ..., "zh": ... } }`，
  再在文章里用。闸门会拒绝未注册标签。
- 标签用于无序筛选；有序阅读用系列，标签替代不了系列。

### 系列（series）

- 系列是有序阅读序列。在 `src/data/series.json` 登记系列和每一篇 planned part
  （slug、中英标题、大纲）。
- 属于某系列的文章，其 slug 必须出现在该系列的 `parts` 目录里，否则闸门失败。
- `series_order` 两份语言文件必须一致。

## 4. 页面长文案（copy）

`src/content/copy/<key>.en.md` 与 `<key>.zh.md`，frontmatter 只有 `lang`：

```yaml
---
lang: zh
---

正文直接写在这里，可换行、可分段。
```

- 首页当前引用了两个 key：`home-intro`（简介）、`home-now`（Now 段）。
- 加新块：按 `<key>.en.md`/`<key>.zh.md` 约定放文件，页面代码用
  `renderedCopyHtml()` 读取（见 `src/views/HomeView.astro`、`src/lib/copy.ts`）。
- 正文里的外部链接用 Markdown 写法 `[文字](https://...)`，构建时自动加
  `target="_blank" rel="noopener"`，不用手写 HTML。

## 5. Markdown 与架构图

- 外链自动新标签页；普通正文、强调、列表、引用、代码块正常用。
- 架构图不是内嵌图片。作者写 `architecture/*.json`（archify JSON IR），
  `pnpm diagrams` 渲染到被 gitignore 的 `public/architecture-assets/`
  （`pnpm build` 前也会自动跑）。文章里用占位 figure：

```html
<figure class="diagram-embed" data-slug="harness-scope" data-aspect="2.00">
<figcaption>图 1 · 图说</figcaption>
</figure>
```

## 6. 文风硬约束（写给人和 AI 共同遵守）

这是本站内容的验收口径，AI 起草时必须逐条自检：

1. **不出现**：公司名、内部系统/项目代号、SLA 等具体数字，除非作者明确要求公开。
2. **不要个人英雄主义**：不写"归我""我主导""我负责一切"这类揽功句式。
   讲事情本身、决策和机制，主语可以省略或用团队/系统。
3. **拒绝 AI 腔**：
   - 不写对仗升华的金句（"先设计失败，再设计功能"这类）；
   - 不用"体系化、端到端、闭环、赋能、持续下降"等空词，
     尤其是不敢给数字的"持续下降"；
   - 每段不要都是"动作 + 方法论升华"的同一模板；
   - 句式长短要有变化。
4. **落到具体**：链路、模块、取舍、真实事故，胜过形容词。
   能讲"一个故障、一个决策、一条边界"就不讲抽象成就。
5. **中文是作者本人的文风**：克制的书面技术语，句子完整、判断直接，
   但不口语化。严禁翻译腔、口水话、营销腔，也严禁为了"去 AI 味"滑向
   大白话唠嗑——作者与 AI 的聊天语气不是文章口吻。
   - 不用自造的"看似口语其实没人这么说"的词（如"最咬人""围着它转"），
     拿不准就用工程师日常书面说法（最难处理、最复杂、容易误判）。
   - 颜色等概念用日常叫法（浅绿，不写"青"）。
   - 判据：读出来像一篇技术文章，不像聊天记录，也不像翻译稿。
6. **英文不是中文的逐字翻译**：用自然的英语技术写作重写，同样禁止空话和口号。
7. 简历话术不上首页；已停更的项目/教程不宣传。
8. 中英文事实必须一致：同一件事的说法、范围、措辞立场不能两份矛盾。

## 7. Agent 系列专章（how-agents-work）

通用规则见前六节；本章只放这个系列额外的规矩。组织思想：《如何读一本书》
的四层阅读倒过来就是写作验收——读者怎么检验，作者就怎么写。

### 7.1 范围护栏

只讲"把模型变成能干活的系统所需的工程"。不碰训练、权重、推理内核、
GPU 调度；不写提示词偏方；不为厂商站队。

### 7.2 研究素材与证据（回应"方块是不是凑的"）

研究样本是作者本地 clone 的开源仓库（codex / opencode），**不提交、路径不进 git**。
本地仓库位置声明在仓库根 `.research-sources.json`（gitignore，各自机器自己建）：

```json
{
  "codex": "/abs/path/to/codex",
  "opencode": "/abs/path/to/opencode"
}
```

- 新事实必须能指到该仓库里的真实文件；指不到就写"待验证"，不许凭印象。
- 每张图配同名证据表 `architecture/<slug>.evidence.md`：方块名 → crate/文件路径
  → 锚定的 commit SHA → 归并说明（一个方块压了哪些 crate）。该文件可提交。
- 正文关键论断就近标 `crate/path/file.rs`，不许只写"codex 里是这么做的"。
- 灰/橙/红/浅绿着色是作者判断、不是源码事实，证据表里写归类理由，允许推翻。
- 引用锚定 revision；设计被后续版本推翻时写"演化"，不把旧设计当现状。

### 7.3 四层写作验收

1. **用词准确（对应基础阅读）**：术语一义、首次定义；模型是无状态推理函数，
   "记忆"是 harness 外部注入，不得说"模型记住了"；不发明新名词。
2. **两分钟抓骨架（对应检视阅读）**：标题断言式、零术语前置、制造认知冲突；
   开篇两段给结论；结构服务"一次 turn 生命周期"主线，不平铺知识点。
3. **统一有序可评判（对应分析阅读）**：一篇一个主旨；机制篇（01 起）按
   核心问题→机制→codex 代码切片→对照样本→边界组织，总起/收口篇不套此模板；
   关键论断给代码证据；区分"本质设计"与"codex 偶然选择"。
4. **跨样本不站队（对应主题阅读）**：先对齐中立词汇再比分歧得失，不比 star；
   第二样本无对应物时只引一个并说明，不许硬凑注水。

### 7.4 架构图布局

- **位置即语义**：左→右 = 人→模型数据流；上推理、下动作；最粗闭合线是
  结果回灌（agent 与请求/响应系统的本质区别）；不参与单次 turn 的底盘沉底。
  四色立柱"货架"布局已废弃。
- 颜色只分类，布局才讲运行逻辑，二者不互相替代。
- 图注写清"据什么整理 + revision"；`data-aspect` 与图一致；改图过 `pnpm diagrams`。

## 8. 本地验证

```bash
pnpm install --registry https://registry.npmjs.org   # 内部镜像缺个别包时
pnpm dev                # 本地预览 http://127.0.0.1:4321
pnpm check:i18n         # 中英成对 / 标签 / 系列 闸门，提交前必跑
pnpm exec astro build   # 只构建站点（跳过架构图渲染）
pnpm build              # 架构图渲染 + 完整构建
```

发布走 GitHub Actions（PR 合并到 main 后自动部署）。未经本人测试确认，
不要 push 或开 PR。

## 8. 安装写作 skill 到 IDE

skill 真源在 `tools/skill/`，IDE 目录只放软链，不复制：

```bash
pnpm skills          # 等价于 bash tools/skill/install.sh
```

脚本把 `tools/skill/*` 链进项目本地的 `.trae/skills/` 和 `.cursor/skills/`
（Trae CN、Cursor 均已验证可发现）。加新 skill 后重跑一次即可。
`.trae/`、`.cursor/` 已在 .gitignore，属于本机 IDE 私有目录；提交的是
`tools/skill/` 里的真源。
