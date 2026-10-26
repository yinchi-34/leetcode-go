# CLRS Index

English | [简体中文](CLRS_INDEX.zh-CN.md)

Section numbers follow CLRS, 3rd edition. This index collects the entries with section numbers from each problem's "CLRS Connections" table, sorted by section. Only sections linked to at least one problem are listed; concepts marked "No matching section" are not included.

Each section gets one table:

```markdown
## Chapter 16 Greedy Algorithms

### §16.2 Elements of the greedy strategy

| Problem | Relation | How It Appears Here |
|---|---|---|
| [Number. Title](../Algorithms/<category>/<problem>/summary.md) | Direct / Related idea | Which step of this problem |
```

## Chapter 2 Getting Started

### §2.1 Insertion sort

| Problem | Relation | How It Appears Here |
|---|---|---|
| [118. Pascal's Triangle](../Algorithms/simulation/0118-pascals-triangle/summary.md) | Direct | Proves correctness with the outer-loop invariant "the first $`i`$ rows are correct" in the initialization, maintenance, and termination steps |
| [3644. Maximum K to Sort a Permutation](../Algorithms/bit-manipulation/3644-maximum-k-to-sort-a-permutation/summary.md) | Direct | Proves the code computes $`M`$ with the invariant "`res` is the AND of the support elements in the prefix" in the initialization, maintenance, and termination steps |
| [856. Score of Parentheses](../data-structures/stack/0856-score-of-parentheses/summary.md) | Direct | Solution 1 uses "`st[k]` $`= \mathrm{score}(B_k)`$" and Solution 2 uses "`res` is the contribution of the atomic pairs in the prefix", each proved by initialization, maintenance, and termination |

## Chapter 3 Growth of Functions

### §3.2 Standard notations and common functions

| Problem | Relation | How It Appears Here |
|---|---|---|
| [118. Pascal's Triangle](../Algorithms/simulation/0118-pascals-triangle/summary.md) | Related idea | The Fibonacci recurrence is the same kind as this problem's, and the shallow-diagonal sums of Pascal's triangle are Fibonacci numbers |

## Chapter 10 Elementary Data Structures

### §10.1 Stacks and queues

| Problem | Relation | How It Appears Here |
|---|---|---|
| [856. Score of Parentheses](../data-structures/stack/0856-score-of-parentheses/summary.md) | Direct | Solution 1 opens a layer with PUSH and closes it with POP; a balanced input guarantees the stack is nonempty before POP, so it never underflows |

## Chapter 15 Dynamic Programming

### §15.3 Elements of dynamic programming

| Problem | Relation | How It Appears Here |
|---|---|---|
| [118. Pascal's Triangle](../Algorithms/simulation/0118-pascals-triangle/summary.md) | Related idea | Each entry of a row is reused by two positions in the next row; filling rows in increasing order computes values before use |

## Appendix C Counting and Probability

### C.1 Counting

| Problem | Relation | How It Appears Here |
|---|---|---|
| [118. Pascal's Triangle](../Algorithms/simulation/0118-pascals-triangle/summary.md) | Related idea | Proves $`P(i, j) = \binom{i}{j}`$ and uses it for the $`\binom{29}{14}`$ overflow bound |
