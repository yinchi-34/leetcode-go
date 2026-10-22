# 118. Pascal's Triangle

English | [简体中文](./summary.zh-CN.md)

Use the [shared glossary](../../../docs/GLOSSARY.md) and follow the repository's [GitHub math conventions](../../../docs/GITHUB_MATH.md).

## 1. Problem Information

- **Problem**: [118. Pascal's Triangle](https://leetcode.cn/problems/pascals-triangle/)
- **Official Difficulty**: Easy
- **Primary Category**: Simulation
- **Secondary Category**: Dynamic Programming
- **Tags**: `Simulation`, `Dynamic Programming`, `Recurrence`, `Binomial Coefficient`
- **Code**: [solution.go](./solution.go)
- **Tests**: [solution_test.go](./solution_test.go)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: No data available ([source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode Official Difficulty: Easy

**Four-Dimensional Rating — Row-by-Row Recurrence**

- Evaluation Scope: $`1 \le \mathrm{numRows} \le 30`$; time $`\Theta(n^2)`$, auxiliary space $`O(1)`$ excluding the output; Go implementation.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 1 | The statement already gives the rule "each number is the sum of the two numbers directly above it"; filling the table row by row follows the definition |
| Knowledge Prerequisites | 1 | Only 2D slices, nested loops, and a conditional are needed |
| Correctness Reasoning Difficulty | 2 | Needs one row-by-row loop invariant plus a check that reads from the previous row stay in bounds |
| Implementation Difficulty | 2 | Row `i` has length `i+1` with 1 at both ends, and `ret[i-1]` may be accessed only when `i ≥ 1` |

Ratings follow [RATING.md](../../../docs/RATING.md). Each dimension is evaluated independently and not combined into a total score.

## 2. Problem Understanding

Given an integer `numRows`, return the first `numRows` rows of Pascal's triangle. Each row starts and ends with 1, and every other number is the sum of the two numbers above it to the left and to the right.

**Example**

- Input: `numRows = 5`
- Output: `[[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]`

**Key Requirements**

1. Input constraint: $`1 \le \mathrm{numRows} \le 30`$.
2. The output is a 2D array whose row $`i`$ (0-based) has exactly $`i+1`$ numbers.
3. Edge cases: `numRows = 1` returns only `[[1]]`; rows 0 and 1 have no interior elements and consist only of the boundary 1s.

## 4. Algorithm Specification

**Algorithm Name:** Row-by-Row Recurrence

Row and column indices in this section are 0-based.

### 4.1 Input

A positive integer $`n = \mathrm{numRows}`$ with $`1 \le n \le 30`$.

### 4.2 Output

Return $`n`$ rows $`R_0, R_1, \ldots, R_{n-1}`$, where $`R_i = \langle R_i[0], \ldots, R_i[i] \rangle`$ has length $`i+1`$.

The problem's rule defines Pascal's triangle as the table $`P`$ satisfying:

```math
P(i, j) =
\begin{cases}
1, & j = 0 \text{ or } j = i \\
P(i-1, j-1) + P(i-1, j), & 0 < j < i
\end{cases}
```

The output must satisfy $`R_i[j] = P(i, j)`$ for all $`0 \le i < n`$ and $`0 \le j \le i`$.

### 4.3 Pseudocode

```text
GENERATE(n)
    ret ← array of length n, each element is a row
    for i ← 0 to n - 1
        ret[i] ← new array of length i + 1
        for j ← 0 to i
            if j == 0 or j == i
                ret[i][j] ← 1
            else
                ret[i][j] ← ret[i-1][j] + ret[i-1][j-1]
    return ret
```

**State Definitions**

- `ret[i]`: row `i` of Pascal's triangle. When the outer loop works on row `i`, `ret[0..i-1]` are complete and never modified again.
- `i`: the row being generated, which is also the index of its last element.
- `j`: the column being filled in the current row.

### 4.4 Correctness

**Proof Method: Outer Loop Invariant + Entry-by-Entry Check**

#### 1. Key Conditions for Correctness

1. **Recursive definition (given by the problem)**: $`P(i, j)`$ is uniquely determined by the recurrence in 4.2. Row $`i`$ depends only on row $`i-1`$, so computing rows in increasing order guarantees every needed value is already available.
2. **Valid indices (lemma to prove)**: in the `else` branch, both `ret[i-1][j-1]` and `ret[i-1][j]` lie within row $`i-1`$.
3. **Rows share no storage (implementation fact)**: each row is allocated separately with `make`, so writing row $`i`$ cannot change a finished row.

**Lemma (valid indices)**: the `else` branch requires $`j \ne 0`$ and $`j \ne i`$; together with $`0 \le j \le i`$ this gives $`1 \le j \le i-1`$. That forces $`i \ge 2`$, so $`i-1 \ge 1`$ and row $`i-1`$ exists. Row $`i-1`$ has valid indices $`0 \ldots i-1`$, and since $`0 \le j-1`$ and $`j \le i-1`$, both indices are valid. In particular, for $`i = 0`$ and $`i = 1`$ the inner loop only takes the `if` branch and never touches a nonexistent `ret[-1]`.

#### 2. Correctness Proof

**Loop Invariant**

At the start of each outer-loop iteration (about to process row $`i`$):

**Invariant I: finished rows are correct**

For every $`0 \le r < i`$, `ret[r]` has length $`r+1`$ and `ret[r][j]` $`= P(r, j)`$ for $`0 \le j \le r`$.

**Initialization**

When $`i = 0`$ there are no finished rows, so Invariant I holds vacuously.

**Maintenance**

Assume Invariant I holds before row $`i`$. The algorithm allocates a new row of length $`i+1`$ and writes each $`j \in \{0, \ldots, i\}`$ exactly once:

- **Case 1: $`j = 0`$ or $`j = i`$.** It writes 1, matching the definition of $`P(i, j)`$.
- **Case 2: $`0 < j < i`$.** By the lemma both indices are valid; by Invariant I, `ret[i-1][j-1]` $`= P(i-1, j-1)`$ and `ret[i-1][j]` $`= P(i-1, j)`$. Their sum is exactly $`P(i, j)`$.

By key condition 3, writing row $`i`$ does not change `ret[0..i-1]`. So after row $`i`$, Invariant I holds for every $`r < i+1`$.

**Termination**

The outer loop runs exactly $`n`$ times and the inner loop exactly $`i+1`$ times on row $`i`$, so the algorithm terminates. At termination $`i = n`$, and by Invariant I, `ret[0..n-1]` are exactly the first $`n`$ rows of Pascal's triangle.

#### 3. Conclusion

For every $`1 \le n \le 30`$, each returned row has the right length and every entry equals the value defined by the recurrence, so the algorithm is correct.

**Note: no overflow.** One can show $`P(i, j) = \binom{i}{j}`$ (see Challenge 2 in Section 5), so the largest value in the output is $`\binom{29}{14} = 77{,}558{,}760`$, below $`2^{31}-1`$. Even a 32-bit `int` would not overflow. This only bounds the values; the correctness proof above does not depend on it.

### 4.5 Complexity Analysis

Let $`n = \mathrm{numRows}`$.

**Time Complexity: $`\Theta(n^2)`$**

Row $`i`$ writes $`i+1`$ elements, each in constant time, so the total number of writes is:

```math
\sum_{i=0}^{n-1} (i+1) = \frac{n(n+1)}{2}
```

The time is therefore $`\Theta(n^2)`$; for $`n = 30`$ that is 465 numbers. The output itself contains $`n(n+1)/2`$ numbers and any algorithm must write each of them, so $`\Omega(n^2)`$ is an unavoidable lower bound for this problem, and this algorithm meets it asymptotically.

**Space Complexity: $`\Theta(n^2)`$ (including the output)**

The returned 2D array holds $`n(n+1)/2`$ integers plus $`n`$ slice headers.

Excluding the output, the only extra state is the loop variables `i` and `j`, so auxiliary space is $`O(1)`$.

## 5. Learning Notes

### Challenge 1: Why is `ret[i-1]` never accessed when `i = 0` or `i = 1`?

`ret[i-1]` looks like it would go out of bounds at `i = 0`, but it is never executed there.

**Understanding Process**

Whether the previous row is read is decided by the condition `j == 0 || j == i`. Row 0 has only $`j = 0`$, and row 1 has only $`j = 0`$ and $`j = 1 = i`$; all of them are boundary positions and take the `if` branch. Interior positions with $`0 < j < i`$ appear only when $`i \ge 2`$, and then the previous row exists. In other words, the rule "both ends are 1" not only supplies boundary values but also protects the indices.

**Transferable Insight:** When writing a recurrence, check that the condition that triggers the recursive branch implies that every referenced state already exists and has a valid index. If it does, no extra special case is needed; if it does not, add a base case.

### Challenge 2: How does Pascal's triangle relate to binomial coefficients?

The problem gives only the addition rule, yet every entry equals the binomial coefficient $`\binom{i}{j}`$.

**Understanding Process**

It suffices to show that $`\binom{i}{j}`$ has the same boundary values and the same recurrence as $`P`$; induction on $`i`$ then shows they agree everywhere.

- Boundary: $`\binom{i}{0} = \binom{i}{i} = 1`$.
- Recurrence (Pascal's rule): choose $`j`$ elements from $`\{1, \ldots, i\}`$ and split the choices by whether element $`i`$ is chosen; the two classes are disjoint and cover everything. If $`i`$ is chosen, the other $`j-1`$ come from the first $`i-1`$ elements, giving $`\binom{i-1}{j-1}`$ ways; if not, all $`j`$ come from the first $`i-1`$, giving $`\binom{i-1}{j}`$ ways.

```math
\binom{i}{j} = \binom{i-1}{j-1} + \binom{i-1}{j}, \qquad 0 < j < i
```

Same boundary, same recurrence, so induction on rows gives $`P(i, j) = \binom{i}{j}`$. This yields the overflow bound in 4.4, and it also shows that row $`i`$ can be computed without the previous row, term by term via $`\binom{i}{j} = \binom{i}{j-1} \cdot \frac{i-j+1}{j}`$.

**Transferable Insight:** When a table is defined only by an addition rule, ask whether it counts some set of choices. A counting interpretation gives a closed form, a value bound, and alternative ways to compute it.

## 6. Key Takeaways

This problem illustrates the following reusable concepts:

1. **Bottom-Up Tabulation**: here, row $`i`$ depends only on row $`i-1`$, so computing rows in increasing order guarantees dependencies are ready before use. It applies to recurrences and dynamic programming whose states have a clear topological order; the evaluation order must be a topological order of the dependencies.
2. **Loop Invariant**: here, the invariant "the first $`i`$ rows are correct" reduces the correctness of the whole table to a per-row check. It applies to simulations that build the result row by row, column by column, or segment by segment; the invariant must state when in the loop it holds (here, at the start of the outer loop).
3. **Pascal's Rule**: here, it links the addition rule to $`\binom{i}{j}`$ and bounds the values. It applies to path counting, subset counting, and binomial expansion coefficients. When using the closed form, watch for overflow in intermediate products and the order of division.
4. **Rolling Array**: this problem returns every row, so the whole table must be kept. If only the last row is needed (as in Problem 119), each row depends only on the previous one, so a single 1D array suffices, updated from right to left so that `row[j-1]` is read before it is overwritten. This holds only when only the final layer is needed.

### CLRS Connections

Section numbers follow CLRS, 3rd edition. Relation is either "Direct" or "Related idea"; if no section fits, write "No matching section" and keep only the concept name.

| Concept | CLRS Section | Relation | How It Appears Here |
|---|---|---|---|
| Loop invariant | §2.1 Insertion sort | Direct | 4.4 proves correctness with the outer-loop invariant "the first $`i`$ rows are correct" in the initialization, maintenance, and termination steps |
| Overlapping subproblems and bottom-up computation | §15.3 Elements of dynamic programming | Related idea | Each entry of row $`i-1`$ is reused by two positions in the next row, and filling rows in increasing order computes values before they are used; this is not an optimization problem, so only the tabulation idea carries over |
| Binomial coefficients | Appendix C.1 Counting | Related idea | Challenge 2 in Section 5 proves $`P(i, j) = \binom{i}{j}`$ and uses it to get the $`\binom{29}{14}`$ overflow bound |

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

- [119. Pascal's Triangle II](https://leetcode.cn/problems/pascals-triangle-ii/) — Same recurrence, but only row `rowIndex` is required; a rolling array cuts auxiliary space to $`O(k)`$ and must be updated right to left. It contrasts keeping the whole table with keeping a single layer.
- [2221. Find Triangular Sum of an Array](https://leetcode.cn/problems/find-triangular-sum-of-an-array/) — Repeatedly adds adjacent numbers (mod 10), the same structure as Pascal's triangle; the final value is a weighted sum with the binomial coefficients of row $`n-1`$ as weights, which Challenge 2 explains.
- [62. Unique Paths](https://leetcode.cn/problems/unique-paths/) — The transition "cell = above + left" is isomorphic to Pascal's rule, and the answer is $`\binom{m+n-2}{m-1}`$; it can be solved by tabulation or by computing the binomial coefficient directly.

---

**One-Sentence Summary**

Pascal's triangle is built by filling rows with the recurrence; correctness rests on the invariant "the first $`i`$ rows are correct," the boundary 1s also keep every index valid, and each entry is the binomial coefficient $`\binom{i}{j}`$.
