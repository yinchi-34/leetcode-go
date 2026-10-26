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
| 栈 | stack | 支持 push、pop 的动态集合，pop 总是删除最近插入且尚未删除的元素。 |
| 后进先出 | last-in first-out (LIFO) | 最后插入的元素最先被删除的存取顺序。 |
| 下溢 | underflow | 对空栈执行 pop。 |
| 上溢 | overflow | 数组实现的栈中，元素数超过数组容量。 |
| 摊还代价 | amortized cost | 一列操作的总代价除以操作次数所得的每次操作的平均代价，不涉及概率。 |
| 单调栈 | monotonic stack | 栈中元素对应的值从栈底到栈顶保持单调的栈，用于查找每个元素左侧或右侧第一个更大（或更小）的元素。 |
| 结构归纳 | structural induction | 对由文法或递归规则生成的对象，证明性质对基本对象成立且被每条生成规则保持的归纳法。 |
| 贡献法 | contribution technique | 把答案写成各基本单元贡献之和，分别计算每个单元的贡献再求和的方法。 |
| 原子对 | atomic pair | 括号串中相邻的一对 `()`。 |
