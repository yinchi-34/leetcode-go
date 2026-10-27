# 全局术语表（Shared Glossary）

| 中文 | 英文 | 定义 |
| --- | --- | --- |
| 需求计数 | demand counting | 以未满足需求的数量表示状态的计数方法。 |
| 半完成匹配 | half-completed match | 由两步组成的匹配中，仅第一步完成的状态。 |
| 循环不变式 | loop invariant | 在每次循环的指定边界均成立的命题。 |
| 下界函数 | lower-bound function | 在定义域内逐点不超过目标函数的函数。 |
| 闭合 | closure | 下界函数 $`L`$ 对任意代价为 $`c`$ 的合法插入转移 $`x \to y`$ 满足 $`L(y) \le L(x) + c`$ 的性质。 |
| 二项式系数 | binomial coefficient | 从 $`n`$ 个不同元素中选出 $`k`$ 个的方案数，记作 $`\binom{n}{k}`$。 |
| 帕斯卡法则 | Pascal's rule | 对 $`0 < k < n`$，$`\binom{n}{k} = \binom{n-1}{k-1} + \binom{n-1}{k}`$。 |
| 自底向上填表 | bottom-up tabulation | 按依赖关系的拓扑序计算并存储子问题的值，使每个值在被使用前已经算出。 |
| 滚动数组 | rolling array | 当每层状态只依赖上一层时，只保留一层存储并原地更新的空间优化方法。 |
| 斐波那契数列 | Fibonacci sequence | 满足 $`F(1) = F(2) = 1`$、$`F(n) = F(n-1) + F(n-2)`$ 的整数序列。 |
| 子掩码 | submask | 若 $`x \mathbin{\&} y = x`$，即 $`x`$ 的每个为 1 的二进制位在 $`y`$ 中也为 1，则称 $`x`$ 是 $`y`$ 的子掩码。 |
| 星形对换 | star transposition | 都包含同一个固定元素（枢纽）的对换；只用这类对换就能实现任意排列。 |
| 单位元 | identity element | 对二元运算 $`\circ`$，满足 $`e \circ x = x \circ e = x`$ 的元素 $`e`$，例如按位与的全 1。 |
| 不动点 | fixed point | 置换 $`\pi`$ 中满足 $`\pi(i) = i`$ 的元素 $`i`$。 |
| 支撑集 | support | 置换 $`\pi`$ 中所有非不动点组成的集合，$`\mathrm{supp}(\pi) = \{i : \pi(i) \ne i\}`$。 |
| 连通分量 | connected component | 无向图中极大的、任意两顶点之间都有路径相连的顶点集合。 |
| 导出子图 | induced subgraph | 由顶点子集 $`U`$ 及原图中两端都在 $`U`$ 内的全部边组成的子图，记作 $`G[U]`$。 |
| 隐式图 | implicit graph | 不显式存储邻接表或邻接矩阵、而是在搜索时按规则即时计算邻居的图，例如以网格格子为顶点、相邻格子为边的图。 |
| 洪水填充 | flood fill | 从一个起点出发，用图搜索找出并标记与它连通的整个区域的方法。 |
