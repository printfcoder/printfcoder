---
name: site-publish
description: 把 printfcoder 个人站点（本仓库）已发布的文章同步到外部平台——当前为掘金（全文 Markdown 草稿）与 V2EX（一篇导读索引帖）。负责生成发布包、把架构图替换成已部署的静态 PNG、保留 canonical 首发声明与系列/GitHub 回链，并通过用户自己的 Chrome 登录态创建掘金草稿。绝不替用户点击发布。
---

# site-publish — 文章外发到掘金 / V2EX

写作规范仍以 `docs/writing-guide.md` 为准。本 skill 只负责"发布"这一段。

## 不可违背的规则

1. **只建草稿，永不代发**。掘金停在草稿页，V2EX 只产出文案；最终发布/发帖由用户本人点击并确认。
2. **不碰凭证**。使用用户自己 Chrome 的登录态（chrome 浏览器能力）；不读取、不存储、不要求用户提供 Cookie 或密码。登录失效时交还用户登录，不代填。
3. **首发在本站**。文章必须先在 me.shux.me 上线（PR 合并、Actions 部署完成、图 URL 可访问），
   然后才外发。掘金按"非首发/转载"处理，文首 canonical 声明由发布包自带，不要勾选掘金原创。
4. 发布包是生成物，不手改。要改内容就改 `src/content/posts/` 源文，重新生成。

## 前置条件（每次发布前核对）

```bash
pnpm check:i18n                    # 中英成对闸门
pnpm diagrams                      # 渲染 HTML reader + 导出 PNG（需要本机 Chrome）
pnpm exec astro build
```

- 确认 `public/architecture-assets/<slug>.zh.png` 存在；PNG 随仓库提交。
- 确认线上图可访问：`https://me.shux.me/architecture-assets/<slug>.zh.png` 返回 200。
  站点未部署就发掘金，正文图片会全部裂开。

## 生成发布包

```bash
node tools/publish/build-bundle.mjs <slug> --platform juejin
node tools/publish/build-bundle.mjs <slug> --platform v2ex
```

产物在 `.publish/`（gitignored）：剥掉 frontmatter、figure 占位替换为线上 PNG、
掘金包带 canonical 头 + 系列目录 + GitHub 尾链；V2EX 包是短导读帖（不是全文搬运）。

## 掘金流程（浏览器自动化，只到草稿）

1. 用 chrome 能力连接用户本机 Chrome，打开 https://juejin.cn/editor 。
   未登录则停止，请用户手动登录后继续。
2. 切到 Markdown 编辑器，把 `.publish/juejin-<slug>.zh.md` 全文粘入正文。
3. 标题取文章 frontmatter `title`；标签从正文 tags 中选掘金存在的分类
   （如 人工智能、后端），拿不准就留空让用户选。
4. 等待自动存草稿，截图核对：标题、两张图是否正常加载、文首出处链接、文末回链。
5. **停手**。向用户报告草稿链接和核对结果，由用户自己检查并点"发布"。

## V2EX 流程（人工发帖）

1. 读 `.publish/v2ex-<slug>.zh.md`，作为给用户的待发文案展示。
2. 建议节点 `/go/create`（程序员），最终节点由用户根据当下节点规则决定；
   注意 Anti-Flood 与新账号限制，不连续发多个主题。
3. 不自动提交。用户确认后自行发帖，或明确要求再协助粘贴。

## 出问题时

- 掘金编辑器改版导致粘贴流程失效：停下来报告差异，不要硬凑接口，不要改用 Cookie API。
- 图片 404：先让站点部署完成，不允许先发后补图。
