---
name: site-content
description: 在 printfcoder 个人站点（Astro，本仓库）上撰写、修改、校对任何内容时使用——博客文章（src/content/posts，含「Agent 是如何工作的」系列）或页面长文案（src/content/copy）。涵盖中英成对规则、frontmatter、标签/系列注册、架构图、证据对账与作者文风硬约束。
---

# site-content — 本站内容写作

本 skill **不复制规则**。所有写作规范的唯一真源是：

**`docs/writing-guide.md`**（仓库根目录下）

README 的 "Writing content" 一节也指向同一文件。规则变更只改那一份文档，
不要在本 skill 或 README 里另立副本。

## 执行流程

1. **先读** `docs/writing-guide.md`，再动笔。任何与内容相关的任务都不例外，
   包括只改一句话——文风约束对一句话同样生效。
2. 按文档判断目标类型：博客 → `src/content/posts/`；固定页面散文 →
   `src/content/copy/`；短 UI 标签 → `src/lib/i18n.ts`。
3. 中英成对处理：改动必须同时落到 `.en.md` 与 `.zh.md`；
   新增标签/系列先改 `src/data/tags.json` / `src/data/series.json`。
   每篇 posts 必须声明 `authorship`（`human` 或 `human-ai`，闸门强制，
   定义见 guide 第 3 节）；AI 参与研究/起草/校改的一律标 `human-ai`，
   不得标 `human` 冒充纯手写。
4. 中文必须符合作者本人的文风（资深后端工程师口吻），英文用自然技术英语重写，
   不逐字机翻。逐条文风自检以文档第 6 节为准。
   **「Agent 是如何工作的」系列额外执行第 7 章**：范围护栏、本地研究素材
   （读仓库根 `.research-sources.json`，路径不进 git）、证据表
   `architecture/<slug>.evidence.md`、四层写作验收、图的"位置即语义"布局。
5. **验证**（交付前必跑，在仓库根目录）：
   - `pnpm check:i18n` —— 成对、lang、标签、系列闸门；
   - `pnpm exec astro build` —— 确认构建通过并检查渲染产物。
6. 未经作者明确许可，**不得 push 或开 PR**。

## 禁止事项

- 不得凭记忆套用 frontmatter 字段或目录约定，以 `docs/writing-guide.md` 当前内容为准。
- 不得只写一种语言然后声称"另一份稍后补"。
- 不得在内容里出现公司名、内部代号、SLA 数字，除非作者当次明确点名要求。
- 不得引入个人英雄主义句式、AI 腔金句和空词。
