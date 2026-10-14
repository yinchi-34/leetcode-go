# 仓库 GitHub 数学书写规范

适用于本仓库的题解、学习笔记、模板和其他 Markdown 文档在 GitHub 网页上的显示。默认规则是：**行内公式统一使用美元符号加反引号；独立块公式使用顶层 `math` 围栏；嵌套容器内只写行内公式。**

这里的目标是提供已验证的默认写法与验收方法。GitHub Pages、本地编辑器和其他 Markdown 渲染器需要各自确认支持情况。

## 官方语法与仓库约定

GitHub 官方支持两种行内语法：普通美元符号，以及美元符号配合反引号；后一种有助于避免与 Markdown 语法冲突。官方也支持 `$$` 块和 `math` 围栏，`math` 围栏内部不需要再加 `$$`。[GitHub 官方数学表达式文档](https://docs.github.com/zh/get-started/writing-on-github/working-with-advanced-formatting/writing-mathematical-expressions)

本仓库选择其中一种默认写法，并限制块公式的位置，以规避本次实测遇到的渲染问题。这些位置限制是仓库约定，不是 GitHub 官方禁止其他写法。

| 位置 | 默认写法 | 说明 |
| --- | --- | --- |
| 普通正文中的简短公式 | 美元符号加反引号的行内公式 | 统一语法，减少 Markdown 干扰 |
| 顶层独立公式、多行推导 | `math` 围栏 | 与前后段落用空行分隔，围栏不缩进 |
| 有序或无序列表 | 行内公式 | 单独一段也使用行内语法；长推导移出列表 |
| 表格单元格 | 行内公式 | 避免裸竖线破坏分列 |
| 引用块、`details` 等 HTML 容器 | 行内公式 | 采用与列表相同的保守约定 |
| 程序代码、伪代码、故意展示的公式源码 | 相应语言的代码围栏 | 保留代码或源码的显示用途 |

已有、实际正常显示的顶层 `$$` 公式可以保留。新增公式或调整结构时按上述默认写法处理。

## 行内公式

原始 Markdown 写法：

```markdown
处理完前 $`i`$ 个字符后，深度满足 $`d(i) \geq 0`$。
```

预览：处理完前 $`i`$ 个字符后，深度满足 $`d(i) \geq 0`$。

## 列表里的独立更新式

原始 Markdown 写法：

```markdown
- **若当前字符为左括号**：

  判断时 `depth` 等于 $`d(i)`$。

  随后更新：

  $`\mathrm{depth} \leftarrow \mathrm{depth} + 1 = d(i+1)`$
```

预览：

- **若当前字符为左括号**：

  判断时 `depth` 等于 $`d(i)`$。

  随后更新：

  $`\mathrm{depth} \leftarrow \mathrm{depth} + 1 = d(i+1)`$

这条公式仍单独成段，但使用的是行内数学语法，不是居中的块公式。需要居中或多行排版时，按下一节将数学块放到顶层。

列表内容的缩进取决于列表标记和嵌套层级。上例中的 `- ` 对应两格内容缩进；两格不是所有列表或所有公式的通用规则。[GFM 列表规范](https://github.github.io/gfm/#list-items)

## 顶层块公式与多行推导

原始 Markdown 写法：

````markdown
**深度定义**

```math
d(k) = \sum_{j=1}^{k}
\begin{cases}
1, & s_j = \texttt{'('} \\
-1, & s_j = \texttt{')'}
\end{cases}
```
````

预览：

```math
d(k) = \sum_{j=1}^{k}
\begin{cases}
1, & s_j = \texttt{'('} \\
-1, & s_j = \texttt{')'}
\end{cases}
```

例如，处理左括号时，可以对齐多行等式：

```math
\begin{aligned}
d(i+1) &= d(i) + 1 \\
\mathrm{depth}_{\mathrm{new}} &= d(i+1)
\end{aligned}
```

若推导属于某个分支，把分支名称改为顶层段落标题或小标题，再放数学块；不要在列表里继续嵌套该块。

## 表格与特殊字符

表格单元格只放行内内容。公式中避免直接写裸 `|`，因为它参与 Markdown 表格分列；按数学含义使用 `\mid`、`\lvert` / `\rvert` 或 `\Vert`，不要机械地把所有竖线替换成同一种符号。[GFM 表格规范](https://github.github.io/gfm/#tables-extension-)

原始 Markdown 写法：

```markdown
| 含义 | 公式 |
| --- | --- |
| 长度约束 | $`n \geq 2`$ |
| 绝对值 | $`\lvert x \rvert \leq M`$ |
```

预览：

| 含义 | 公式 |
| --- | --- |
| 长度约束 | $`n \geq 2`$ |
| 绝对值 | $`\lvert x \rvert \leq M`$ |

需要在数学表达式里显示普通美元符号时，用 `\$`；普通美元符号与公式出现在同一行、但位于公式外时，可用 `<span>$</span>` 包起来。[官方美元符号说明](https://docs.github.com/zh/get-started/writing-on-github/working-with-advanced-formatting/writing-mathematical-expressions#writing-dollar-signs-in-line-with-and-within-mathematical-expressions)

## 记号与代码

正文中的程序变量用代码格式，例如 `depth`、`res`。数学式中的程序变量可用 `\mathrm{depth}`；普通数学变量可以保留斜体，例如 `d(i)`。排版约定与本次故障的根因是两回事。

| 用途 | LaTeX 示例 |
| --- | --- |
| 上下标 | `s_{i+1}`、`n^2` |
| 集合 | `\{1,\ldots,n\}` |
| 求和 | `\sum_{i=1}^{n} a_i` |
| 分数 | `\frac{a}{b}` |
| 说明文字 | `\text{otherwise}` |
| 程序变量 | `\mathrm{depth}` |
| 函数名的直立字形 | `\mathrm{inner}(P)` |

先区分容器识别问题与具体命令的兼容问题，再修改公式。保留数学语义；不要为了修复容器识别问题，盲目替换箭头、下标等命令。不要把完整 LaTeX 文档、任意宏包或未经验证的命令当作 GitHub 一定支持的输入。

## 命名算子的兼容性

MathJax 的 AMS 扩展支持 `\operatorname`，但这不能证明 GitHub 的实际渲染链路允许该命令。[MathJax 命令列表](https://docs.mathjax.org/en/stable/input/tex/macros/)

2026-10-10 的客户端对照中，`\operatorname{inner}(P)` 在普通正文行内、列表行内和顶层 `math` 围栏里都被 API 识别为数学组件，随后 GitHub 客户端均报错：`The following macros are not allowed: operatorname`。相同客户端中的 `\mathrm{inner}(P)` 行内与块公式对照均生成 MathML。这是命令被拒绝，与列表中的数学块退化为代码块是两个独立问题。

GitHub 自身的问题记录也包含不依赖列表的 `\operatorname` 失败案例。该问题在 2026 年 4 月关闭的原因是前端数学排版不属于 `github/markup` 仓库的职责，并不表示故障已经修复。[命名算子问题记录](https://github.com/github/markup/issues/1688)，[维护者说明](https://github.com/github/markup/issues/1688#issuecomment-4311348967)

因此，本仓库对于 `inner(P)` 这类只需直立函数名的表达式，默认使用：

```markdown
$`\mathrm{inner}(P)`$
```

`\mathrm` 只提供直立字形，不等同于 `\operatorname` 的算子类型、间距或星号版本的上下限行为。若这些排版语义必不可少，需要针对完整表达式验证替代写法；不要机械地全仓库替换，也不要仅换美元符号或缩进后就认定兼容问题解决。

## 避免本次复现的问题

2026-10-10 的最小实验中，纯文字列表里的 `math` 围栏能够渲染；在同一列表加入行内公式后，该围栏退化为普通代码块。即使行内式与数学块分属不同列表项，也能触发问题。这不意味着整篇文档不能同时使用行内公式和顶层数学块。

为避免重复触发，本仓库在嵌套容器中统一使用行内公式。`details` 中混用两种公式也有作者直接复现的社区反馈。[同类问题报告](https://github.com/orgs/community/discussions/57950)

完整实验与分工见 [本次排查记录](investigations/github-math-rendering-investigation.md)。

## 验收方法

修复数学渲染问题，或改变公式所在容器时，按以下顺序检查：

1. **源文件**：核对公式语义、围栏闭合、空行、缩进与列表或表格结构；区分真正公式和故意展示的公式源码。
2. **GitHub Markdown API**：向 `https://api.github.com/markdown` 提交实际上下文，使用 `mode: "gfm"`。确认公式被识别为数学组件，数量和所属容器正确。
3. **真实 GitHub 页面**：在固定提交地址打开实际文件，确认可见公式、数量、排版和错误提示。MathML、MathJax 等成功排版的 DOM 可作为辅助证据。
4. **记录证据**：保留测试输入、API 结果、页面地址和截图。

API 只返回 `<pre lang="math">`，表示仍是代码块，不能据此声称数学渲染成功。`math-renderer` 能证明识别阶段通过，但仍不能替代最终页面显示检查。页面保留隐藏 TeX 源码可能是正常行为，不要仅靠搜索反斜线判定失败。
