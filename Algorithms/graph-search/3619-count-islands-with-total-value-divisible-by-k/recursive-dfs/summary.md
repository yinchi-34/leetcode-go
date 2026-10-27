# 3619. Count Islands With Total Value Divisible by K — Recursive DFS

English | [简体中文](./summary.zh-CN.md)

Use the [shared glossary](../../../../docs/GLOSSARY.md) and follow the repository's [GitHub math conventions](../../../../docs/GITHUB_MATH.md).

This problem has two solutions in two directories that share one set of test cases:

- **Recursive DFS** (this page): visited cells are set to 0 in place.
- [**Explicit-stack DFS**](../iterative-dfs/summary.md): marks cells in a `visited` array and replaces recursion with a slice used as a stack.

## 1. Problem Information

- **Problem**: [3619. Count Islands With Total Value Divisible by K](https://leetcode.cn/problems/count-islands-with-total-value-divisible-by-k/)
- **Official Difficulty**: Medium
- **Primary Category**: Graph Search
- **Tags**: `Graph Search`, `Depth-First Search`, `Connected Components`, `Recursion`, `Matrix`
- **Code**: [solution.go](./solution.go)
- **Tests**: [solution_test.go](./solution_test.go), with cases in [testcases.go](../testcases/testcases.go)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: 1461.19 ([source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode Official Difficulty: Medium

**Four-Dimensional Rating — Recursive Depth-First Search**

- Evaluation Scope: $`1 \le m, n \le 1000`$, $`mn \le 10^5`$, $`0 \le \mathrm{grid}[i][j] \le 10^6`$, $`1 \le k \le 10^6`$; time $`\Theta(mn)`$, auxiliary space $`O(mn)`$ (recursion stack); Go implementation.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 2 | One direct transformation: view the grid as a graph whose vertices are land cells and whose edges join horizontally or vertically adjacent cells; an island is a connected component and its value is the sum of its vertex weights |
| Knowledge Prerequisites | 2 | Needs DFS flood fill on a grid and recursion, but not the full DFS theory of discovery times or edge classification |
| Correctness Reasoning Difficulty | 3 | Needs strong induction over recursive calls to show that one call zeroes and sums exactly one whole island, plus the invariant "every island is either intact or fully zeroed" to show each island is counted once |
| Implementation Difficulty | 2 | The current cell must be set to 0 before recursing, and all four directions need bounds checks; in Go the recursive closure must first be declared with `var dfs func(int, int) int` |

Ratings follow [RATING.md](../../../../docs/RATING.md). Each dimension is evaluated independently and not combined into a total score.

## 2. Problem Understanding

We are given an $`m \times n`$ matrix `grid` of non-negative integers and a positive integer $`k`$. Positive cells are land and 0 is water. Land cells that are adjacent horizontally or vertically form an island, and the total value of an island is the sum of its cells. Return the number of islands whose total value is divisible by $`k`$.

**Example**

- Input: `grid = [[0,2,1,0,0],[0,5,0,0,5],[0,0,1,0,0],[0,1,4,7,0],[0,2,0,0,8]]`, `k = 5`
- Output: `2`

The four islands have total values $`2+1+5 = 8`$, $`5`$, $`1+1+4+7+2 = 15`$, and $`8`$; of these, 5 and 15 are divisible by 5.

**Key Requirements**

1. Input constraints: $`1 \le m, n \le 1000`$, $`1 \le mn \le 10^5`$, $`0 \le \mathrm{grid}[i][j] \le 10^6`$, $`1 \le k \le 10^6`$.
2. Only horizontal and vertical neighbors are adjacent; diagonals are not.
3. The test is on the sum of the whole island, so each island must be found completely before it can be judged.
4. A sum can reach $`mn \cdot 10^6 \le 10^{11}`$, which exceeds the 32-bit integer range.

## 4. Algorithm Specification

**Algorithm Name:** Recursive Depth-First Search

Indices in this section are 0-based.

### 4.1 Input

A grid $`\mathrm{grid}`$ of size $`m \times n`$ and a positive integer $`k`$.

View the grid as an undirected graph $`G = (V, E)`$:

- the vertex set $`V = \{(i, j) : \mathrm{grid}[i][j] > 0\}`$ is the set of all land cells;
- two land cells are joined by an edge when they are adjacent vertically or horizontally.

Each vertex $`v`$ has weight $`w(v) = \mathrm{grid}[v]`$ (its value in the input).

### 4.2 Output

Islands are the connected components of $`G`$. For a connected component $`C`$, its total value is

```math
S(C) = \sum_{v \in C} w(v)
```

The output $`\mathrm{res}`$ is the number of connected components $`C`$ with $`S(C) \bmod k = 0`$.

### 4.3 Pseudocode

```text
COUNT-ISLANDS(grid, k)
    res ← 0
    for i ← 0 to m - 1
        for j ← 0 to n - 1
            if grid[i][j] > 0 and DFS(grid, i, j) mod k == 0
                res ← res + 1
    return res

DFS(grid, i, j)                       ▷ precondition: grid[i][j] > 0
    s ← grid[i][j]
    grid[i][j] ← 0                    ▷ mark first, then visit neighbors
    for each (x, y) in {(i, j-1), (i, j+1), (i-1, j), (i+1, j)}
        if 0 ≤ x < m and 0 ≤ y < n and grid[x][y] > 0
            s ← s + DFS(grid, x, y)
    return s
```

In the code, `ans` corresponds to `res` in the pseudocode, and the local variable `res` inside `dfs` corresponds to `s`.

**State Definitions**

- `grid[x][y]`: greater than 0 means an unvisited land cell; 0 means water or a land cell already visited.
- `s`: the sum of the cell values accumulated so far by this `DFS` call.
- `res`: the number of islands already processed whose total value is divisible by $`k`$.

### 4.4 Correctness

**Proof method: strong induction over recursive calls + a loop invariant for the main loop**

Let $`U`$ be the set of cells that are still positive in `grid` at some moment. Then $`U \subseteq V`$, and cells in $`U`$ still hold their input values. $`G[U]`$ denotes the subgraph of $`G`$ induced by $`U`$.

#### I. Key Properties Required for Correctness

1. **All values are non-negative** (given): land cells are exactly the positive cells, so "set to 0" can also mean "visited" without being confused with real land.
2. **Mark before recursing**: `DFS` sets the current cell to 0 as soon as it is entered, so no later call can enter that cell again. Hence the recursion terminates and every cell is added at most once.
3. **Lemma 1 (one call zeroes one whole component)**: see below.
4. **Removing whole components leaves the others unchanged**: if some complete connected components of $`G[U]`$ are removed from $`U`$, every remaining connected component stays the same, because there are no edges between different components.

**Lemma 1.** Suppose that when `DFS(i, j)` starts, the set of positive cells is $`U`$ and $`(i, j) \in U`$. Let $`R`$ be the connected component of $`(i, j)`$ in $`G[U]`$. Then the call:

- sets exactly the cells of $`R`$ to 0 and leaves every other cell unchanged;
- returns $`\sum_{v \in R} w(v)`$.

*Proof.* Strong induction on $`\lvert R \rvert`$.

The call first adds $`w(i, j)`$ and sets $`(i, j)`$ to 0, so the set of positive cells becomes $`U_0 = U \setminus \{(i, j)\}`$. The set $`R \setminus \{(i, j)\}`$ is made up of some connected components $`D_1, \ldots, D_t`$ of $`G[U_0]`$, and each $`D_s`$ contains a neighbor of $`(i, j)`$. This holds because a path from any other cell of $`R`$ to $`(i, j)`$ does not pass through $`(i, j)`$ before its last step.

While the four neighbors are processed in turn, the following property holds: **the recursive calls made so far have zeroed some complete $`D_s`$, and their return values add up to the total weight of those $`D_s`$.** When neighbor $`u`$ is processed:

- If $`u`$ is out of bounds or water: no call is made, and the property holds.
- If $`u`$ is land but already 0: either $`u`$ had been zeroed before this call started and is not in $`R`$, or the $`D_s`$ containing $`u`$ has already been zeroed as a whole. In both cases no call is made, so nothing is added twice.
- If $`u`$ is still positive: by property 4, the connected component of $`u`$ in the subgraph induced by the current positive cells is exactly its $`D_s`$, and $`\lvert D_s \rvert < \lvert R \rvert`$. By the induction hypothesis, `DFS(u)` zeroes exactly $`D_s`$ and returns its total weight, so the property holds.

After all four neighbors, each $`D_s`$ contains some neighbor $`u`$, and when $`u`$ was processed, $`D_s`$ either had already been zeroed or was zeroed at that step. So every $`D_s`$ is zeroed and added exactly once. Together with $`(i, j)`$ itself, the call zeroes exactly $`R`$ and returns $`\sum_{v \in R} w(v)`$. Cells outside $`R`$ belong to no $`D_s`$ and are not modified. ∎

#### II. Correctness Proof

**Loop Invariant**

Scan the cells in row-major order. Before cell $`(i, j)`$ is examined:

**Invariant I: every island is either intact or fully zeroed**

For each connected component $`C`$ of $`G`$: if some cell of $`C`$ comes before $`(i, j)`$, then all cells of $`C`$ are 0; otherwise all cells of $`C`$ keep their input values.

**Invariant II: the count is correct**

`res` equals the number of zeroed connected components with $`S(C) \bmod k = 0`$.

**Initialization**

Before the first cell, no cell comes earlier, every cell keeps its input value, and `res` $`= 0`$, so both invariants hold.

**Maintenance**

Let the current cell be $`(i, j)`$ with value $`x`$.

- **$`x = 0`$**: the cell is water, or its component has already been fully zeroed. `DFS` is not called and `res` is unchanged. The component of $`(i, j)`$, if any, is already zeroed, which is what Invariant I requires for "cells before the next cell", so both invariants hold.
- **$`x > 0`$**: by Invariant I, the component $`C`$ of $`(i, j)`$ is intact, and $`(i, j)`$ is the first cell of $`C`$ in scan order. The current set of positive cells $`U`$ consists of complete components, so the component of $`(i, j)`$ in $`G[U]`$ is $`C`$. By Lemma 1, `DFS(i, j)` zeroes exactly $`C`$ and returns $`S(C)`$. If $`S(C) \bmod k = 0`$, `res` increases by 1. Other components are unaffected, so both invariants hold.

Short-circuit evaluation of `&&` ensures that `DFS` is called only when $`x > 0`$.

There is one more Go detail. In `for j, x := range row`, `row` is the slice header `grid[i]`, which shares its backing array with `grid`, and `x` reads `row[j]` at each iteration. So a cell zeroed by `DFS` is read as 0 when the scan reaches it.

**Termination**

Each `DFS` call sets one positive cell to 0, and there are at most $`mn`$ positive cells, so the total number of recursive calls is finite; the outer loop runs exactly $`mn`$ times.

After the last cell, every component has a cell "before the end", so by Invariant I every component is zeroed, and by Invariant II `res` is the number of components that satisfy the condition.

#### III. Conclusion

Each island is found and zeroed completely when the scan reaches its first cell, and it is judged exactly once; after that its cells are all 0 and it is never counted again. Therefore the algorithm returns the correct answer for every valid input.

### 4.5 Complexity Analysis

Let the grid have $`m`$ rows and $`n`$ columns, with $`mn \le 10^5`$.

**Time Complexity: $`\Theta(mn)`$**

The outer double loop runs $`mn`$ times. Each `DFS` call sets one positive cell to 0, so the number of calls equals the number of land cells, at most $`mn`$. Each call checks 4 neighbors in constant time. The total is $`\Theta(mn)`$.

**Space Complexity: $`O(mn)`$**

The algorithm marks cells in the input and allocates no extra array. Auxiliary space comes from the recursion stack, whose depth equals the length of the current chain of calls and is at most the size of one island.

This bound is reached in the worst case. For example, on a grid that is all land, with the code's direction order "left, right, up, down", the search from $`(0, 0)`$ goes right to the end of the row, down one cell, then left along the next row, snaking through the whole grid with recursion depth $`mn`$. The test case `full_grid_max_sum` is a $`100 \times 1000`$ all-land grid, so its recursion depth reaches $`10^5`$.

Goroutine stacks in Go grow on demand (the default limit is 1 GB on 64-bit platforms), so this depth runs normally. The same code may overflow the stack in languages with a small default stack; in that case use the [explicit-stack DFS](../iterative-dfs/summary.md).

**Integer Range**

An island sum is at most $`10^5 \times 10^6 = 10^{11}`$, which is larger than $`2^{31} - 1`$ and smaller than $`2^{63} - 1`$. Go's `int` is 64 bits on 64-bit platforms, so the sum can be accumulated directly. To support 32-bit platforms, reduce modulo $`k`$ while accumulating, since only the remainder of the sum modulo $`k`$ matters.

### 4.6 Test Design

Both solutions share the cases in [testcases.go](../testcases/testcases.go). Each case copies the grid with `CloneGrid` before the call, because this solution sets visited cells to 0.

| Case | Expected | What It Checks |
|---|---|---|
| `example_1` | 2 | Official example 1, four islands of different shapes |
| `example_2_diagonal_not_connected` | 6 | Diagonals are not connected; treating them as adjacent would merge the six 3s into one island with sum 18 and give 1 |
| `all_water` | 0 | No land, so `DFS` is never called |
| `single_cell_not_divisible` | 0 | A one-cell island whose sum is not divisible |
| `square_island_has_cycle` | 1 | A $`2 \times 2`$ all-land grid with a cycle, so one cell is reachable from two neighbors; checks that nothing is added twice |
| `spiral_island_needs_all_directions` | 1 | A spiral path from $`(0, 0)`$ that moves right, down, left, up, then right; dropping any one direction splits the island into pieces whose sums are all below 14 |
| `full_grid_max_sum` | 1 | $`100 \times 1000`$ cells of $`10^6`$, sum $`10^{11}`$; with 32-bit integers the truncated sum is not divisible by $`10^6`$ and the result would be 0; recursion depth reaches $`10^5`$ |

A mutation check was run for `spiral_island_needs_all_directions`: deleting any one of the four directions from `dirs` makes this case fail.

## 5. Learning Notes

### Challenge 1: Why is each cell visited only once?

**Understanding Process**

The key is "set to 0 first, then visit the neighbors". `DFS(i, j)` sets $`(i, j)`$ to 0 as soon as it is entered, so any later call that checks this cell skips it because `grid[x][y] > 0` is false. Each land cell therefore enters `DFS` at most once, which guarantees three things at the same time: the recursion terminates, each cell's value is added once, and the total time is $`\Theta(mn)`$.

If the cell were set to 0 only after its neighbors were visited, any two adjacent land cells would keep entering each other and the recursion would never end.

**Transferable Insight:** In graph search, when a vertex is marked as visited decides both correctness and complexity. Mark it at the moment you decide to visit it, not when the visit is finished.

### Challenge 2: Why does one `DFS` call find exactly one whole island?

**Understanding Process**

"Recursion reaches every cell it can" is the intuition, but we need to show it finds neither more nor fewer cells:

- No more: `DFS` only moves to adjacent land cells, so it cannot leave the island.
- No fewer: Lemma 1 uses strong induction to show that every not-yet-zeroed neighbor of the current cell is handled recursively, and each recursive call zeroes a whole sub-component.

Invariant I of the outer scan completes the argument: an island is always intact when it is found, it is fully 0 afterward, and it is never counted again.

**Transferable Insight:** To prove a recursive search correct, use strong induction on the size of what a call must handle, with "a child call handles a whole sub-component" as the induction hypothesis.

### Challenge 3: The cost of in-place marking

**Understanding Process**

Setting visited cells to 0 saves the `visited` array but changes the input. This problem never uses `grid` again, so that is fine; the tests, however, have to copy the grid before each call. The approach also relies on the problem's guarantee that 0 is never land. If land could have value 0, we would need a marker that cannot be confused with land, or a `visited` array.

**Transferable Insight:** In-place marking requires two conditions: the caller no longer needs the original input, and there is a marker value that cannot be confused with valid data.

## 6. Key Takeaways

This problem yields the following reusable methods:

1. **Implicit Graph**: Here the grid cells are vertices and horizontal or vertical adjacency gives edges; there is no need to build adjacency lists, because a direction array enumerates neighbors on the fly. It applies to grids, boards, state spaces, and similar problems, provided neighbors can be enumerated in constant time.
2. **Flood Fill**: Here the search starts from a land cell and marks the whole island connected to it. It applies to grid problems that count or process each connected region. Limitation: the recursive version can recurse as deep as the region is large, so with limited stack space rewrite it with an explicit stack or BFS.
3. **Aggregate per Connected Component**: Here the search adds up the weights of the component as it goes and judges the whole component once at the end. It applies to problems that need a per-component area, weight sum, maximum, whether it touches the border, and so on, provided the aggregate can be accumulated one vertex at a time (sum, count, maximum).
4. **In-place Marking**: Here visited land is set to 0 instead of using a `visited` array. The conditions are those in Challenge 3: the input may be modified, and there is a marker value that cannot be confused with valid data.
5. **Strong Induction on Recursive Calls**: Here the induction is on the size of the component a call must handle, proving that one call zeroes exactly one whole component. It applies to proving recursive searches and divide-and-conquer algorithms correct.

### CLRS Connections

Section numbers follow CLRS, 3rd edition. Relation is either "Direct" or "Related idea"; if no section fits, write "No matching section" and keep only the concept name.

| Concept | CLRS Section | Relation | How It Appears Here |
|---|---|---|---|
| Representations of graphs | §22.1 Representations of graphs | Related idea | 4.1 views the grid as a graph whose vertices are land cells and enumerates neighbors with a direction array instead of an explicit adjacency list |
| Depth-first search | §22.3 Depth-first search | Direct | `DFS` sets a cell to 0 as soon as it is entered, like coloring a vertex gray in CLRS; unvisited neighbors are visited recursively; the outer scan over all cells is like the DFS main procedure calling DFS-VISIT on every white vertex |
| Connected components | §21.1 Disjoint-set operations | Related idea | CLRS uses disjoint sets in this section to find the connected components of an undirected graph; this problem finds components with DFS instead and sums each component's weights along the way |

## 7. Review Log

| Date | Independent Completion | Understanding Gained | Areas for Improvement |
|---|---|---|---|
| To be recorded | To be recorded | To be recorded | To be recorded |

### Next Review Checklist

- [ ] Can I implement the algorithm without looking at the solution?
- [ ] Can I accurately define the input and output?
- [ ] Can I write the pseudocode independently?
- [ ] Can I explain why the algorithm is correct?
- [ ] Can I derive its time and space complexity?
- [ ] Can I recognize the same algorithmic pattern in other problems?

## 8. Related Problems

- [200. Number of Islands](https://leetcode.cn/problems/number-of-islands/) — The same grid connected-component model, counting components only; it is this problem without the aggregate and the test.
- [695. Max Area of Island](https://leetcode.cn/problems/max-area-of-island/) — Aggregates per connected component, with the cell count as the aggregate and the maximum as the answer.
- [2658. Maximum Number of Fish in a Grid](https://leetcode.cn/problems/maximum-number-of-fish-in-a-grid/) — Almost the same aggregate as this problem: positive cells are connected and each component's value sum is computed, but the answer is the maximum instead of a count.
- [1905. Count Sub Islands](https://leetcode.cn/problems/count-sub-islands/) — Also "search a whole component, then judge it once", with the condition that every cell of the component is land in another grid.

---

**One-Sentence Summary**

View the grid as a graph so that islands are connected components; DFS sets a cell to 0 as soon as it is entered, which ensures that each island is found completely when the scan reaches its first cell and is counted exactly once, while summing along the way gives the island's total value.
