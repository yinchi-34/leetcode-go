# CLRS 反查索引

[English](CLRS_INDEX.md) | 简体中文

章节号以《算法导论》第三版为准。本索引汇总各题题解「CLRS 关联」表中写了章节号的条目，按章节排序，只列出已有题目关联的章节。写“无对应章节”的知识点不收录。

每节一张表，格式如下：

```markdown
## 第 16 章 贪心算法

### §16.2 贪心策略的要素

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [题号. 题名](../Algorithms/<分类>/<题目录>/summary.zh-CN.md) | 直接对应 / 相关思想 | 具体到本题的哪一步 |
```

## 第 2 章 算法基础

### §2.1 插入排序

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [118. 杨辉三角](../Algorithms/simulation/0118-pascals-triangle/summary.zh-CN.md) | 直接对应 | 用“前 $`i`$ 行已正确”的外层循环不变式，按初始化、保持、终止三步证明 |
| [3644. 排序排列](../Algorithms/bit-manipulation/3644-maximum-k-to-sort-a-permutation/summary.zh-CN.md) | 直接对应 | 用“`res` 是前缀中支撑集元素的按位与”的不变式，按初始化、保持、终止三步证明代码算出 $`M`$ |
| [856. 括号的分数](../data-structures/stack/0856-score-of-parentheses/summary.zh-CN.md) | 直接对应 | 解法一用“`st[k]` $`= \mathrm{score}(B_k)`$”、解法二用“`res` 是前缀内原子对的贡献和”，分别按初始化、保持、终止三步证明 |

## 第 3 章 函数的增长

### §3.2 标准记号与常用函数

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [118. 杨辉三角](../Algorithms/simulation/0118-pascals-triangle/summary.zh-CN.md) | 相关思想 | 斐波那契递推与本题递推同类，且杨辉三角浅对角线之和等于斐波那契数 |

## 第 10 章 基本数据结构

### §10.1 栈和队列

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [856. 括号的分数](../data-structures/stack/0856-score-of-parentheses/summary.zh-CN.md) | 直接对应 | 解法一用 PUSH 打开一层、POP 闭合一层；平衡串保证 POP 前栈非空，不会下溢 |

## 第 15 章 动态规划

### §15.3 动态规划原理

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [118. 杨辉三角](../Algorithms/simulation/0118-pascals-triangle/summary.zh-CN.md) | 相关思想 | 上一行每个数被下一行两个位置复用，按行号递增填表保证先算后用 |

## 附录 C 计数与概率

### C.1 计数

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [118. 杨辉三角](../Algorithms/simulation/0118-pascals-triangle/summary.zh-CN.md) | 相关思想 | 证明 $`P(i, j) = \binom{i}{j}`$，并据此得出 $`\binom{29}{14}`$ 的溢出上界 |
