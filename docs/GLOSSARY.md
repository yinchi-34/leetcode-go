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
