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

## Chapter 3 Growth of Functions

### §3.2 Standard notations and common functions

| Problem | Relation | How It Appears Here |
|---|---|---|
| [118. Pascal's Triangle](../Algorithms/simulation/0118-pascals-triangle/summary.md) | Related idea | The Fibonacci recurrence is the same kind as this problem's, and the shallow-diagonal sums of Pascal's triangle are Fibonacci numbers |

## Chapter 10 Elementary Data Structures

### §10.1 Stacks and queues

| Problem | Relation | How It Appears Here |
|---|---|---|
| [3619. Count Islands With Total Value Divisible by K (explicit-stack DFS)](../Algorithms/graph-search/3619-count-islands-with-total-value-divisible-by-k/iterative-dfs/summary.md) | Direct | A Go slice implements PUSH and POP and holds cells that have been discovered and are waiting to be processed |

## Chapter 15 Dynamic Programming

### §15.3 Elements of dynamic programming

| Problem | Relation | How It Appears Here |
|---|---|---|
| [118. Pascal's Triangle](../Algorithms/simulation/0118-pascals-triangle/summary.md) | Related idea | Each entry of a row is reused by two positions in the next row; filling rows in increasing order computes values before use |

## Chapter 21 Data Structures for Disjoint Sets

### §21.1 Disjoint-set operations

| Problem | Relation | How It Appears Here |
|---|---|---|
| [3619. Count Islands With Total Value Divisible by K (recursive DFS)](../Algorithms/graph-search/3619-count-islands-with-total-value-divisible-by-k/recursive-dfs/summary.md) | Related idea | CLRS uses disjoint sets in this section to find the connected components of an undirected graph; this problem finds components with DFS instead and sums each component's weights along the way |

## Chapter 22 Elementary Graph Algorithms

### §22.1 Representations of graphs

| Problem | Relation | How It Appears Here |
|---|---|---|
| [3619. Count Islands With Total Value Divisible by K (recursive DFS)](../Algorithms/graph-search/3619-count-islands-with-total-value-divisible-by-k/recursive-dfs/summary.md) | Related idea | 4.1 views the grid as a graph whose vertices are land cells and enumerates neighbors with a direction array instead of an explicit adjacency list |
| [3619. Count Islands With Total Value Divisible by K (explicit-stack DFS)](../Algorithms/graph-search/3619-count-islands-with-total-value-divisible-by-k/iterative-dfs/summary.md) | Related idea | Views the grid as a graph whose vertices are land cells and enumerates neighbors from coordinate offsets instead of an explicit adjacency list |

### §22.2 Breadth-first search

| Problem | Relation | How It Appears Here |
|---|---|---|
| [3619. Count Islands With Total Value Divisible by K (explicit-stack DFS)](../Algorithms/graph-search/3619-count-islands-with-total-value-divisible-by-k/iterative-dfs/summary.md) | Related idea | Same structure as CLRS's BFS: mark a vertex as soon as it is discovered (gray in BFS), then add it to the pending set; with a stack instead of a queue, the search still finds the whole connected component |

### §22.3 Depth-first search

| Problem | Relation | How It Appears Here |
|---|---|---|
| [3619. Count Islands With Total Value Divisible by K (recursive DFS)](../Algorithms/graph-search/3619-count-islands-with-total-value-divisible-by-k/recursive-dfs/summary.md) | Direct | `DFS` sets a cell to 0 as soon as it is entered, like coloring a vertex gray in CLRS; unvisited neighbors are visited recursively; the outer scan over all cells is like the DFS main procedure calling DFS-VISIT on every white vertex |
| [3619. Count Islands With Total Value Divisible by K (explicit-stack DFS)](../Algorithms/graph-search/3619-count-islands-with-total-value-divisible-by-k/iterative-dfs/summary.md) | Related idea | Replaces recursion with a stack, but pushes all undiscovered neighbors at once on each pop, so the visiting order is not guaranteed to match DFS-VISIT; see Challenge 2 |

## Appendix C Counting and Probability

### C.1 Counting

| Problem | Relation | How It Appears Here |
|---|---|---|
| [118. Pascal's Triangle](../Algorithms/simulation/0118-pascals-triangle/summary.md) | Related idea | Proves $`P(i, j) = \binom{i}{j}`$ and uses it for the $`\binom{29}{14}`$ overflow bound |
