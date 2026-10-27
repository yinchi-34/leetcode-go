# 3619. Count Islands With Total Value Divisible by K — Explicit-Stack DFS

English | [简体中文](./summary.zh-CN.md)

Use the [shared glossary](../../../../docs/GLOSSARY.md) and follow the repository's [GitHub math conventions](../../../../docs/GITHUB_MATH.md).

This problem has two solutions in two directories that share one set of test cases:

- [**Recursive DFS**](../recursive-dfs/summary.md): visited cells are set to 0 in place.
- **Explicit-stack DFS** (this page): marks cells in a `visited` array and replaces recursion with a slice used as a stack.

The problem statement, example, and constraints are the same as in the recursive version; this page focuses on how the two solutions differ.

## 1. Problem Information

- **Problem**: [3619. Count Islands With Total Value Divisible by K](https://leetcode.cn/problems/count-islands-with-total-value-divisible-by-k/)
- **Official Difficulty**: Medium
- **Primary Category**: Graph Search
- **Secondary Category**: Stack
- **Tags**: `Graph Search`, `Stack`, `Depth-First Search`, `Connected Components`, `Matrix`
- **Code**: [solution.go](./solution.go)
- **Tests**: [solution_test.go](./solution_test.go), with cases in [testcases.go](../testcases/testcases.go)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: 1461.19 ([source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode Official Difficulty: Medium

**Four-Dimensional Rating — Explicit-Stack Depth-First Search**

- Evaluation Scope: $`1 \le m, n \le 1000`$, $`mn \le 10^5`$, $`0 \le \mathrm{grid}[i][j] \le 10^6`$, $`1 \le k \le 10^6`$; time $`\Theta(mn)`$, auxiliary space $`\Theta(mn)`$; Go implementation.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 2 | Same as the recursive version: view the grid as a graph, so an island is a connected component and its value is the sum of its weights |
| Knowledge Prerequisites | 3 | Beyond graph search, one must replace recursion with a stack and understand the difference between marking on push and marking on pop, which combines a few common techniques |
| Correctness Reasoning Difficulty | 3 | Needs an invariant for the inner loop relating three kinds of cells (popped, on the stack, undiscovered), then connectivity to show that an empty stack means exactly one whole island has been processed |
| Implementation Difficulty | 3 | `visited`, the stack, and the running sum must be kept in step; marking must happen on push, or a cell is pushed and added more than once; bounds checks are also needed |

Ratings follow [RATING.md](../../../../docs/RATING.md). Each dimension is evaluated independently and not combined into a total score.

## 2. Problem Understanding

See [Section 2 of the recursive version](../recursive-dfs/summary.md#2-problem-understanding).

## 4. Algorithm Specification

**Algorithm Name:** Iterative Depth-First Search with an Explicit Stack

Indices in this section are 0-based.

### 4.1 Input

The same as the recursive version: a grid $`\mathrm{grid}`$ ($`m \times n`$) and a positive integer $`k`$. View the grid as an undirected graph $`G = (V, E)`$, where $`V`$ is the set of all positive cells, two land cells adjacent vertically or horizontally are joined by an edge, and each vertex has weight $`w(v) = \mathrm{grid}[v]`$.

### 4.2 Output

The total value of a connected component $`C`$ is

```math
S(C) = \sum_{v \in C} w(v)
```

The output $`\mathrm{res}`$ is the number of connected components $`C`$ with $`S(C) \bmod k = 0`$.

### 4.3 Pseudocode

```text
COUNT-ISLANDS(grid, k)
    visited ← m × n boolean array, all false
    res ← 0
    for r ← 0 to m - 1
        for c ← 0 to n - 1
            if grid[r][c] ≠ 0 and not visited[r][c]
                S ← empty stack
                PUSH(S, (r, c))
                visited[r][c] ← true          ▷ mark on push
                total ← 0
                while S is not empty
                    (x, y) ← POP(S)
                    total ← total + grid[x][y]
                    for each (nx, ny) in {(x+1, y), (x-1, y), (x, y-1), (x, y+1)}
                        if 0 ≤ nx < m and 0 ≤ ny < n
                                and grid[nx][ny] ≠ 0 and not visited[nx][ny]
                            PUSH(S, (nx, ny))
                            visited[nx][ny] ← true
                if total mod k == 0
                    res ← res + 1
    return res
```

**State Definitions**

- `visited[x][y]`: the cell has been discovered, that is, it has been pushed.
- `S`: cells that have been discovered but not yet processed (not yet added, neighbors not yet checked).
- `total`: the sum of the values of the popped cells of the current island.
- `res`: the number of islands already processed whose total value is divisible by $`k`$.

### 4.4 Correctness

**Proof method: a loop invariant for the inner loop + a loop invariant for the outer loop**

#### I. Key Properties Required for Correctness

1. **Mark on push**: a cell is pushed only while `visited` is false, it is set to true at the same moment, and it never becomes false again. So **each cell is pushed at most once**, and therefore popped and added at most once.
2. **Expand only through land**: only land cells are pushed, and each newly pushed cell is adjacent to a popped cell, so one search cannot leave the island of its starting cell.
3. **Connectivity**: any two cells of an island $`C`$ are joined by a path that stays inside $`C`$.

#### II. Correctness Proof

**Inner-loop invariant**

Fix one search with starting cell $`s`$ and island $`C`$. Split the cells of $`C`$ into three groups: the set $`T`$ of popped cells, the set $`P`$ of cells currently on the stack, and the cells not yet discovered. Each time the `while` condition is checked:

**Invariant I: the groups are disjoint and complete**

The cells pushed during this search are exactly $`T \cup P`$, $`T \cap P = \varnothing`$, and $`T \cup P \subseteq C`$.

**Invariant II: the sum is correct**

`total` $`= \sum_{v \in T} w(v)`$.

**Invariant III: neighbors of popped cells have been discovered**

Every land neighbor of every cell in $`T`$ has been pushed, that is, belongs to $`T \cup P`$.

**Initialization**: before the first check, $`T = \varnothing`$, $`P = \{s\}`$, and `total` $`= 0`$. By the outer-loop invariant (below), no cell of the island of $`s`$ has been discovered before, so all three invariants hold.

**Maintenance**: one iteration pops a cell $`v`$, moves it from $`P`$ to $`T`$, and adds $`w(v)`$, so Invariant II holds. It then checks the four neighbors of $`v`$: those out of bounds, water, or already discovered are skipped; the rest are undiscovered land cells of $`C`$, which are pushed, marked, and added to $`P`$. So all land neighbors of $`v`$ are now in $`T \cup P`$; Invariant III holds for the newly added $`v`$ and continues to hold for the earlier cells of $`T`$. Newly pushed cells are undiscovered cells of $`C`$, so Invariant I holds.

**Termination**: by property 1, each cell is pushed at most once, so the inner loop runs at most $`\lvert C \rvert`$ times before the stack is empty. At that point $`P = \varnothing`$.

Suppose $`T \ne C`$. Take a cell $`u`$ of $`C`$ not in $`T`$, and a path from $`s`$ to $`u`$ that stays inside $`C`$. Since $`s \in T`$, the first cell on the path that is not in $`T`$ is a land neighbor of some cell in $`T`$, so by Invariant III it lies in $`T \cup P = T`$, a contradiction. Hence $`T = C`$, and by Invariant II `total` $`= S(C)`$.

**Outer-loop invariant**

Scan the cells in row-major order. Before cell $`(r, c)`$ is examined:

- **Invariant A: every island is either fully discovered or fully undiscovered.** If island $`C`$ has a cell before $`(r, c)`$, then every cell of $`C`$ is `visited`; otherwise none is.
- **Invariant B: the count is correct.** `res` equals the number of fully discovered islands with $`S(C) \bmod k = 0`$.

At initialization no cell has been discovered and `res` $`= 0`$, so both hold.

Maintenance: if the current cell is water or already discovered, nothing happens and both invariants hold. Otherwise the current cell is undiscovered land, and by Invariant A its island $`C`$ is fully undiscovered. By the inner-loop analysis above, the search ends having discovered exactly the cells of $`C`$ with `total` $`= S(C)`$, and `res` is updated accordingly, so both invariants hold.

Termination: after all cells are scanned, every island is fully discovered, and `res` is the answer.

#### III. Conclusion

Each island is searched completely once, when the scan reaches its first cell, and at the end of that search `total` is exactly the island's total value; afterward all of its cells are marked and it is never searched again. Therefore the algorithm returns the correct answer for every valid input.

### 4.5 Complexity Analysis

Let the grid have $`m`$ rows and $`n`$ columns.

**Time Complexity: $`\Theta(mn)`$**

Initializing `visited` takes $`\Theta(mn)`$. The outer double loop runs $`mn`$ times. By property 1, each land cell is pushed and popped at most once, and each pop checks 4 neighbors in constant time. The total is $`\Theta(mn)`$.

**Space Complexity: $`\Theta(mn)`$**

- The `visited` array holds exactly $`mn`$ booleans and is allocated whatever the grid contains, so it takes $`\Theta(mn)`$.
- Cells on the stack are distinct (each is pushed at most once), so the stack never holds more cells than the current island, at most $`mn`$.
- The input `grid` is not modified.

`append` may grow the slice, and one growth step copies the elements, but amortized over all appends each append costs $`O(1)`$, so the total time is unaffected.

**Integer Range**

The same as the recursive version: an island sum is at most $`10^{11}`$ and needs a 64-bit integer. Go's `int` is 64 bits on 64-bit platforms.

### 4.6 Test Design

Both solutions share the cases in [testcases.go](../testcases/testcases.go); what each case checks is listed in [Section 4.6 of the recursive version](../recursive-dfs/summary.md#46-test-design).

The case that matters most for this solution is `square_island_has_cycle` (a $`2 \times 2`$ grid of 1s with `k = 4`). A mutation check was run: if marking is moved to pop time and pop does not check whether the cell was already visited, cell $`(1, 0)`$ is pushed once by $`(0, 0)`$ and once by $`(1, 1)`$, `total` becomes 5, and this case fails. The other cases still pass under this mutation, because their islands either have no cycle or have every cell value a multiple of $`k`$, so the repeated additions remain divisible.

## 5. Learning Notes

### Challenge 1: Why mark on push rather than on pop?

**Understanding Process**

`visited` means "discovered", not "finished". Marking on push guarantees that each cell is pushed at most once, so it is added at most once, and the stack is never larger than the island.

If marking happened on pop, a cell could be seen by several neighbors before it is popped, and each neighbor would push it again. Any island with a cycle (such as a $`2 \times 2`$ square) triggers this, with these effects:

- if pop does not check whether the cell was already visited, the cell's value is added several times and the answer is wrong;
- if pop checks and skips it, the answer is correct, but the number of pushes grows to the order of the number of edges and the stack gets larger.

**Transferable Insight:** In graph search with a stack or a queue, marking should happen together with adding a vertex to the pending set. BFS follows the same rule: mark on enqueue.

### Challenge 2: Is this really depth-first search?

**Understanding Process**

When this version pops a cell, it pushes and marks all of the cell's undiscovered neighbors at once. Its visiting order is not exactly that of recursive DFS. Recursive DFS enters one neighbor and explores everything reachable from it before looking at the next neighbor. Here, the later neighbors are marked right away, so even if another path reaches them first later on, they are not "discovered" a second time.

So this version does not guarantee the DFS discovery order of CLRS §22.3, and it cannot be used directly to compute discovery and finishing times or to classify edges. This problem only needs to find whole connected components, so any search order that visits every cell exactly once works, and the correctness proof in 4.4 does not depend on the order. If strict DFS order is needed, keep "cell + next direction to check" on the stack to simulate the call frames of the recursion.

**Transferable Insight:** First decide which property of the search the problem needs. If only connectivity matters, DFS, BFS, and an explicit stack are interchangeable; properties that depend on order, such as discovery times or a topological order, require simulating the recursion exactly.

### Challenge 3: What does the explicit stack buy?

**Understanding Process**

On an all-land grid, the recursive version recurses $`mn`$ levels deep. Go stacks grow automatically, so both versions pass in Go, but in languages with a small default stack the recursive version may overflow. The explicit stack moves the "call stack" into a slice on the heap, so the depth is limited only by memory. The price is extra code for the `visited` array and the stack.

Also, this solution uses `visited` instead of setting cells to 0, so it does not modify the input.

**Transferable Insight:** The general way to turn recursion into iteration is to keep the "unfinished work" on a stack. If we only care which vertices are reached, the stack can hold just vertices; if work must continue after a child call returns, the stack must also record how far each call has progressed.

## 6. Key Takeaways

This problem yields the following reusable methods:

1. **Recursion to Iteration with an Explicit Stack**: Here a slice serves as a stack for cells that have been discovered and are waiting to be processed, which avoids recursion depth up to $`mn`$. It applies to searches, tree traversals, and other problems whose recursion can get deep. Limitation: if work must continue after a child call returns (for finishing times or post-order traversal), the stack must also record progress.
2. **Mark on Discovery**: Here `visited` is set at the same time as the push, so each cell is pushed at most once. It applies to every stack- or queue-based graph search.
3. **Implicit Graph**: Here grid cells are vertices, horizontal or vertical adjacency gives edges, and neighbors are computed from coordinate offsets on the fly, provided they can be enumerated in constant time.
4. **Aggregate per Connected Component**: Here the search sums the component's weights and judges the whole component once the stack is empty, provided the aggregate can be accumulated one vertex at a time.
5. **Closure Argument for Search Completeness**: Here "the land neighbors of popped cells have been discovered" together with an empty stack shows that the set of popped cells is closed under adjacency, and connectivity then shows it is the whole island. It applies to proving that BFS, DFS, and their variants find whole connected components.

### CLRS Connections

Section numbers follow CLRS, 3rd edition. Relation is either "Direct" or "Related idea"; if no section fits, write "No matching section" and keep only the concept name.

| Concept | CLRS Section | Relation | How It Appears Here |
|---|---|---|---|
| Stack | §10.1 Stacks and queues | Direct | A Go slice implements PUSH and POP and holds cells that have been discovered and are waiting to be processed |
| Representations of graphs | §22.1 Representations of graphs | Related idea | Views the grid as a graph whose vertices are land cells and enumerates neighbors from coordinate offsets instead of an explicit adjacency list |
| Breadth-first search | §22.2 Breadth-first search | Related idea | Same structure as CLRS's BFS: mark a vertex as soon as it is discovered (gray in BFS), then add it to the pending set; with a stack instead of a queue, the search still finds the whole connected component |
| Depth-first search | §22.3 Depth-first search | Related idea | Replaces recursion with a stack, but pushes all undiscovered neighbors at once on each pop, so the visiting order is not guaranteed to match DFS-VISIT; see Challenge 2 |

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

- [200. Number of Islands](https://leetcode.cn/problems/number-of-islands/) — The same grid connected-component model, counting components only.
- [2658. Maximum Number of Fish in a Grid](https://leetcode.cn/problems/maximum-number-of-fish-in-a-grid/) — Almost the same aggregate as this problem: positive cells are connected and each component's value sum is computed, but the answer is the maximum instead of a count.
- [733. Flood Fill](https://leetcode.cn/problems/flood-fill/) — The most basic flood fill, good for practicing the explicit-stack version and marking on push.
- [856. Score of Parentheses](https://leetcode.cn/problems/score-of-parentheses/) — Also uses a stack for a nested structure where the inner part must be finished before returning to the outer part, a useful comparison for replacing recursion with a stack.

---

**One-Sentence Summary**

When a stack replaces recursion in an island search, set `visited` on push to mean "discovered": each cell is then pushed and added at most once, and when the stack is empty the popped cells are exactly the whole island.
