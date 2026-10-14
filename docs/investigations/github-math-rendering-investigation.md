# GitHub 数学渲染排查记录与 agent 分工

日期：2026-10-10（Asia/Shanghai）。对象：1021 题的中英文笔记。

最终修复位于 [5cfeeb0](https://github.com/yinchi-34/leetcode-go/commit/5cfeeb01b65ba450513c891310ad9357afa492fe)。两条深度更新式保留在列表里，改为独立段落的行内公式；中英文页面均经过真实浏览器检查。

仓库后续写法见 [GitHub 数学书写规范](../GITHUB_MATH.md)。以下记录描述这次实际做过的工作。

## 查询与定位

主 agent 使用 `rg` 在仓库中定位 `\mathrm{depth}`，读取两篇笔记的完整列表上下文，并检查适用的 `AGENTS.md`、Git 状态和提交历史。

首先搜索 GitHub 官方数学文档，查询词为：

```text
site:docs.github.com writing mathematical expressions markdown code block math dollar signs
```

实际打开 [官方数学表达式文档](https://docs.github.com/zh/get-started/writing-on-github/working-with-advanced-formatting/writing-mathematical-expressions)，确认行内公式、`$$` 和 `math` 围栏都是支持的语法。随后读取 [GFM 列表规范](https://github.github.io/gfm/#list-items) 和 [围栏规范](https://github.github.io/gfm/#fenced-code-blocks)，核实缩进规则。

第一次修复统一了围栏、空行和缩进，并用 Python `urllib.request` 请求 GitHub Markdown API。API 输出仍是 `<pre lang="math">`，主 agent 当时误将它作为渲染成功的证据。子 agent 的源文件审查也没有覆盖最终网页显示；检查 Markdown 语法不足以发现该问题。

用户提供灰底代码框截图后，主 agent 重新检查实际网页：这两条公式包在带复制功能的普通代码块中，其他公式已经是数学组件。因此故障发生在数学块识别阶段，而不是 `\mathrm` 或 `\leftarrow` 的数学排版阶段。

## 第二次查询与最小对照

主 agent 搜索：

```text
site:github.com/orgs/community math code block nested list not rendered github
```

找到作者直接复现的 [details 中混用公式问题](https://github.com/orgs/community/discussions/57950)，以及 [列表内多行公式问题](https://github.com/orgs/community/discussions/122438)。这些是辅助线索；本仓库的结论依靠自己的对照实验确认。

主 agent 用相同公式逐项改变上下文，并请求：

```http
POST https://api.github.com/markdown
Content-Type: application/json

{"text":"- Branch $i$\n\n  ```math\n  x + 1\n  ```","mode":"gfm"}
```

`render_root_cause` agent 补充了不同列表项、不同公式顺序和嵌套容器的对照。

| 实验输入 | API 输出或后续结果 |
| --- | --- |
| 顶层 `math` 围栏 | 显示数学组件 |
| 顶层正文含行内式，另有顶层 `math` | 显示数学组件 |
| 纯文字列表加 `math` 围栏 | 显示数学组件 |
| 同一列表出现行内式，再加入 `math` | 普通 `<pre lang="math">` |
| 行内式与 `math` 放在同一列表的不同条目 | 仍为普通代码块 |
| 将原行内式改成美元符号加反引号，但保留数学块 | 数学块仍为普通代码 |
| 混用行内式的列表内，多行 `$$` | 原始公式文本，未转换为数学组件 |
| 列表内同一行 `$$表达式$$` | 数学组件，但浏览器最终按行内方式显示 |
| 列表内独立段落，使用美元符号加反引号 | 行内数学组件，浏览器显示正常 |

结论是：当前 GitHub 渲染器在该列表结构中混用行内公式与块公式时，会漏掉块公式转换。两格与四格缩进并不是这次故障的根因，也不能据此认为 GitHub 不支持 `math` 围栏。

## 浏览器验证与最终修复

主 agent 在临时目录安装 `playwright-core`，使用本机 Chrome 的无界面模式打开实际 GitHub 页面；没有向代码仓库添加浏览器依赖。

候选比较时，先把 API 生成的两种候选 HTML 注入已加载的 GitHub 页面，检查客户端实际数学排版。这一步只是本地候选验证，不能替代发布后的真实页面检查。

最终选择：

```markdown
  $`\mathrm{depth} \leftarrow \mathrm{depth} + 1 = d(i+1)`$
```

主 agent 修改两篇笔记，并将修复纳入原来的单个汇总提交。推送后重新打开固定提交地址的真实页面，不注入候选 HTML：

- [英文页面](https://github.com/yinchi-34/leetcode-go/blob/5cfeeb01b65ba450513c891310ad9357afa492fe/Algorithms/simulation/1021-remove-outermost-parentheses/summary.md)
- [中文页面](https://github.com/yinchi-34/leetcode-go/blob/5cfeeb01b65ba450513c891310ad9357afa492fe/Algorithms/simulation/1021-remove-outermost-parentheses/summary.zh-CN.md)

两篇页面分别确认：63 个公式完成数学排版；两条深度更新式都生成可见 MathML；保留在正确的列表项内；没有遗留的 `pre[lang="math"]`，没有浏览器页面错误。主 agent 保存了两篇页面的截图，并查看英文截图确认布局。

本次整理规范时，又检查了规范中的八个预览公式：API 识别为六个行内式与两个块公式；在已加载的 GitHub 页面中注入规范 HTML 后，八个示例全部生成 MathML，包括两个表格单元格内的公式，没有数学错误元素。这是本地规范示例验证，不表示规范文档已发布到 GitHub。

## 命名算子的独立兼容问题（追加排查）

用户随后指出 `\operatorname` 曾经渲染失败。规范初稿把该命令列为推荐写法，仅凭 MathJax 支持和此前的容器故障作判断，证据不足。八个预览公式也没有使用该命令，因此它们通过验证不能证明命名算子可用。

主 agent 读取 [MathJax 命令列表](https://docs.mathjax.org/en/stable/input/tex/macros/)，确认它属于 AMS 扩展；同时委派 `operator_audit` 只读核查旧提交和 GitHub 一手故障报告。

该 agent 找到 [github/markup 的命名算子问题](https://github.com/github/markup/issues/1688)：其中有不依赖列表的失败案例和 2026 年的持续反馈。维护者在 [关闭说明](https://github.com/github/markup/issues/1688#issuecomment-4311348967) 中解释，前端数学排版不在该仓库职责内；关闭不代表修复。旧提交确实曾把 `\operatorname` 改成 `\mathrm`，但提交记录没有当时的浏览器诊断，不能反推出之前每次失败的原因。

主 agent 使用 GitHub Markdown API 创建五个最小输入，再用 Playwright 加载当天保存的 GitHub 页面 HTML 外壳及实际客户端资源，在页面中注入这些 API 结果。另注入两个只使用 `\mathrm` 的数学组件作为客户端对照。结果如下：

| 测试上下文 | 识别或客户端结果 |
| --- | --- |
| 正文行内 `\operatorname{inner}(P)` | API 识别为数学组件；客户端拒绝 `operatorname` |
| 列表行内同一式 | API 识别为数学组件；客户端拒绝 `operatorname` |
| 顶层 `math` 围栏内同一式 | API 识别为数学组件；客户端拒绝 `operatorname` |
| 列表行内式加内嵌 `math`，分别使用 `\operatorname` 和 `\mathrm` | 两种围栏都变成普通 `pre[lang="math"]`；客户端只渲染列表中的行内式 |
| `\mathrm{inner}(P)` 的行内与块组件对照 | 两者均生成 MathML，没有数学错误提示 |

前三种的可见错误完全一致：`The following macros are not allowed: operatorname`。没有浏览器页面异常。这验证了本次所加载的 GitHub 客户端拒绝 `operatorname`，与此前的列表识别故障独立；不能通过改变缩进或定界符消除宏拒绝。

这次是本地注入的客户端实验，没有把失败样例提交到仓库，也不是新规范已经发布后的页面验证。输入、结果与截图保存在本地临时目录和 `.git/codex-backups/operatorname-context-verification.json`、`.git/codex-backups/operatorname-context-verification.png`；主 agent 另外查看了简短对照截图。

规范已修正为默认使用 `\mathrm{inner}(P)` 的直立字形，并明确它不等同于命名算子的间距与上下限语义。该建议基于上述对照及已发布题解中 `\mathrm` 的实际渲染，不是对未来 GitHub 行为的永久保证。

## 各 agent 的实际分工

独立的只读调查并行进行；修改文件、修改 Git 历史和推送操作由主 agent 执行。各个子 agent 没有编辑题解，也没有操作暂存区。

| Agent | 实际工作 | 提供的证据与边界 |
| --- | --- | --- |
| 主 agent `/root` | 搜索官方资料、定位公式、读取实际 HTML、执行 API 对照、用浏览器验证、修复并推送 | 负责实际运行与最终判断；首次 API 验证不足由主 agent 纠正 |
| `math_audit` | 初次只读扫描 Markdown，审查两篇笔记的列表缩进、围栏、空行和数学语义 | 确认源文件结构；没有验证 GitHub 最终排版，因此未发现后续截图中的故障 |
| `render_root_cause` | 第二次只读调查官方资料、社区一手报告，并做补充 API 对照 | 确认同一列表的不同条目也能触发干扰；仅给行内式加反引号不能救回原来的数学块 |
| `squash_review` | 审查要合并的七个提交、最终文件范围和提交描述 | 确认范围是两篇 1021 笔记；属于 Git 历史检查，不是渲染验证 |
| `merge_check` | 审查失败 `pull` 后的历史、恢复方案和显式 lease 推送结果 | 对照旧七个提交与新提交内容；合并后人工修改和未跟踪文件的检查引用主 agent 执行结果 |
| `guide_review` | 本次整理规范时只读核对官方依据和规则边界 | 区分官方语法与仓库约定，补充表格竖线、美元符号和验收方法；没有重跑此前全部实验 |
| `operator_audit` | 追加只读检查旧提交、MathJax 支持情况和 GitHub 命名算子一手故障报告 | 发现独立的宏兼容问题；未修改文件或执行上述浏览器实验，客户端结果由主 agent 验证 |

主 agent 通过 `spawn_agent` 派发有明确问题的任务，通过 `send_message` 共享新证据，再读取子 agent 的报告合并结论。这种独立任务并行、由主 agent 汇总的方式也符合 [OpenAI 官方多 agent 文档](https://developers.openai.com/api/docs/guides/responses-multi-agent#overview) 的说明；上表描述的是本次实际分工，不是对所有任务固定采用的配置。

## 可复用的排查顺序

1. 官方文档确认支持语法。
2. 源文件确认实际上下文。
3. 最小对照隔离触发条件。
4. API 检查数学组件是否被识别。
5. 真实网页检查最终排版。
6. 记录输入、输出、页面地址与截图，再整理为仓库规则。

语法审查、API 结构检查和浏览器显示检查是三种不同证据。多个 agent 同时认可源文件格式，也不能替代最后的页面验证。
