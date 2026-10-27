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

## 第 3 章 函数的增长

### §3.2 标准记号与常用函数

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [118. 杨辉三角](../Algorithms/simulation/0118-pascals-triangle/summary.zh-CN.md) | 相关思想 | 斐波那契递推与本题递推同类，且杨辉三角浅对角线之和等于斐波那契数 |

## 第 10 章 基本数据结构

### §10.1 栈和队列

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [3619. 总价值可以被 K 整除的岛屿数目（显式栈 DFS）](../Algorithms/graph-search/3619-count-islands-with-total-value-divisible-by-k/iterative-dfs/summary.zh-CN.md) | 直接对应 | 用 Go 切片实现 PUSH 和 POP，保存已发现、待处理的格子 |

## 第 15 章 动态规划

### §15.3 动态规划原理

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [118. 杨辉三角](../Algorithms/simulation/0118-pascals-triangle/summary.zh-CN.md) | 相关思想 | 上一行每个数被下一行两个位置复用，按行号递增填表保证先算后用 |

## 第 21 章 用于不相交集合的数据结构

### §21.1 不相交集合的操作

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [3619. 总价值可以被 K 整除的岛屿数目（递归 DFS）](../Algorithms/graph-search/3619-count-islands-with-total-value-divisible-by-k/recursive-dfs/summary.zh-CN.md) | 相关思想 | CLRS 在这一节用不相交集合求无向图的连通分量；本题改用 DFS 求分量，同时累加分量的权值和 |

## 第 22 章 基本的图算法

### §22.1 图的表示

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [3619. 总价值可以被 K 整除的岛屿数目（递归 DFS）](../Algorithms/graph-search/3619-count-islands-with-total-value-divisible-by-k/recursive-dfs/summary.zh-CN.md) | 相关思想 | 4.1 把网格看成以陆地格为顶点的图，用方向数组即时枚举邻居，代替显式的邻接表 |
| [3619. 总价值可以被 K 整除的岛屿数目（显式栈 DFS）](../Algorithms/graph-search/3619-count-islands-with-total-value-divisible-by-k/iterative-dfs/summary.zh-CN.md) | 相关思想 | 把网格看成以陆地格为顶点的图，用坐标偏移即时枚举邻居，代替显式的邻接表 |

### §22.2 广度优先搜索

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [3619. 总价值可以被 K 整除的岛屿数目（显式栈 DFS）](../Algorithms/graph-search/3619-count-islands-with-total-value-divisible-by-k/iterative-dfs/summary.zh-CN.md) | 相关思想 | 结构与 CLRS 的 BFS 相同：发现时立刻标记（BFS 中涂成灰色），再放入待处理集合；把队列换成栈后，仍然能找出整个连通分量 |

### §22.3 深度优先搜索

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [3619. 总价值可以被 K 整除的岛屿数目（递归 DFS）](../Algorithms/graph-search/3619-count-islands-with-total-value-divisible-by-k/recursive-dfs/summary.zh-CN.md) | 直接对应 | `DFS` 进入时立刻置 0，相当于 CLRS 中把顶点涂成灰色；未访问的邻居递归访问；外层按顺序扫描所有顶点，相当于 DFS 主过程对每个白色顶点调用 DFS-VISIT |
| [3619. 总价值可以被 K 整除的岛屿数目（显式栈 DFS）](../Algorithms/graph-search/3619-count-islands-with-total-value-divisible-by-k/iterative-dfs/summary.zh-CN.md) | 相关思想 | 用栈代替递归，但弹出时一次压入全部未发现邻居，访问顺序不保证与 DFS-VISIT 相同，见难点二 |

## 附录 C 计数与概率

### C.1 计数

| 题目 | 关联程度 | 本题中的体现 |
|---|---|---|
| [118. 杨辉三角](../Algorithms/simulation/0118-pascals-triangle/summary.zh-CN.md) | 相关思想 | 证明 $`P(i, j) = \binom{i}{j}`$，并据此得出 $`\binom{29}{14}`$ 的溢出上界 |
