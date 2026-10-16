# 2916. Subarrays Distinct Element Sum of Squares II

English | [简体中文](./summary.zh-CN.md)

Use the [shared glossary](../../../docs/GLOSSARY.md) and follow the repository's [GitHub math conventions](../../../docs/GITHUB_MATH.md).

## 1. Problem Information

- **Problem**: [2916. Subarrays Distinct Element Sum of Squares II](https://leetcode.cn/problems/subarrays-distinct-element-sum-of-squares-ii/)
- **Official Difficulty**: Hard
- **Primary Category**: Dynamic Programming
- **Secondary Category**: Segment Tree
- **Tags**: `Right-Endpoint Sweep`, `Last Occurrence`, `Range Add Range Sum`, `Lazy Segment Tree`, `Data-Structure-Optimized DP`
- **Code**: To be added (solution.go)
- **Tests**: To be added (solution_test.go)
- **Prerequisite**: [2913. Subarrays Distinct Element Sum of Squares I](../../enumeration/2913-subarrays-distinct-element-sum-of-squares-i/)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: 2816 ([source](https://github.com/zerotrac/leetcode_problem_rating))
- LeetCode Official Difficulty: Hard

**Four-Dimensional Rating - Right-Endpoint DP with a Lazy Segment Tree**

- Evaluation scope: $`1 \leq n \leq 10^5`$, $`1 \leq \mathrm{nums}[i] \leq 10^5`$; time O(n log n), auxiliary space O(n); implementation in Go.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 4 | Three conversions in a row: sweep the right endpoint, use the last occurrence to find the affected range of left endpoints, then expand the change of a square into a range sum. |
| Knowledge Prerequisites | 4 | Requires DP-style recurrences plus a lazy segment tree (or two Fenwick trees) that supports range add and range sum. |
| Correctness Reasoning Difficulty | 3 | Requires one lemma about which left endpoints gain one distinct element, then a loop invariant on what the tree and f represent. |
| Implementation Difficulty | 3 | Lazy push-down, querying before updating, and where to apply the modulus are all easy to get wrong. |

Ratings follow the repository's [RATING.md](../../../docs/RATING.md). The four dimensions are evaluated independently and are not combined into an overall score.

## 2. Problem Understanding

Given an integer array `nums`, the count of a non-empty subarray is the number of distinct values in it. Return the sum of the squared counts over all subarrays, modulo $`10^9 + 7`$.

The statement is identical to 2913; only the bound grows from $`n \leq 100`$ to $`n \leq 10^5`$.

**Example**

- Input: `nums = [1,2,1]`
- Output: `15`
- Explanation: the six subarrays have counts 1, 1, 1, 2, 2, 2, so the answer is $`3 \cdot 1 + 3 \cdot 4 = 15`$.

**Key Requirements**

1. There are $`n(n+1)/2`$ subarrays, about $`5 \times 10^9`$, so they cannot be enumerated one by one.
2. The target is the sum of **squared** counts, not the sum of counts.
3. The answer is taken modulo $`10^9 + 7`$, but the counts stored in the tree sum to at most about $`n^2 = 10^{10}`$, which fits in a 64-bit `int`.

## 3. Initial Approach

### First Idea: Reuse the 2913 Enumeration

In 2913, fix the left end `i`, extend the right end `j`, keep the distinct values in a set, and add `len(s) * len(s)` at every step. This takes O(n²).

### The Problem

With $`n = 10^5`$, O(n²) is about $`10^{10}`$ operations and times out. Instead of computing each subarray separately, the counts of many subarrays must be updated **in bulk**.

A natural question arises: problems like this are usually DP, so why do the editorials talk about segment trees? Do the two approaches compute different things?

### Shift in Thinking

From:

> Fix the left end, enumerate every right end, and compute each subarray's count.

To:

> Fix the right end r and treat all subarrays ending at r as one state; when the right end moves from r−1 to r, only a contiguous range of left ends gains one distinct element.

The key observation is that the new element `nums[r]` only increases the count of subarrays that do not already contain it, and their left ends form the range $`(p, r]`$, where p is the previous occurrence of `nums[r]`. Updating a whole range at once is exactly what a segment tree does well.

So both approaches compute the same answer, and the algorithm is still DP; the segment tree only makes each DP transition fast.

## 4. Algorithm Specification

**Algorithm Name:** Right-Endpoint DP with a Lazy Segment Tree

### 4.1 Input

$`A = \langle a_0, a_1, \ldots, a_{n-1} \rangle`$

subject to:

- $`1 \leq n \leq 10^5`$
- $`1 \leq a_i \leq 10^5`$

### 4.2 Output

Let $`c(l, r)`$ be the number of distinct values among $`a_l, \ldots, a_r`$. Output:

```math
\mathrm{ans} = \left( \sum_{0 \leq l \leq r < n} c(l, r)^2 \right) \bmod (10^9 + 7)
```

### 4.3 Pseudocode

```text
SUM-COUNTS(A)
    n ← length(A)
    T ← segment tree over n zeros, supporting range add and range sum
    last ← empty map              // value → index of its previous occurrence
    f ← 0
    ans ← 0

    for r ← 0 to n − 1
        p ← last[A[r]], or −1 if absent
        S ← T.SUM(p + 1, r)        // query first: position l holds c(l, r−1)
        T.ADD(p + 1, r, 1)         // then update: position l becomes c(l, r)
        f ← (f + 2·S + (r − p)) mod M
        ans ← (ans + f) mod M
        last[A[r]] ← r

    return ans
```

**State Definitions**

- `T[l]`: after processing right end r, equals $`c(l, r)`$; positions not yet used as a left end ($`l > r`$) hold 0.
- `f`: the sum of squared counts over subarrays ending at r, that is $`f(r) = \sum_{l=0}^{r} c(l, r)^2`$.
- `ans`: the sum of squared counts over all subarrays whose right end is at most r.
- `p`: the last index before r holding the value `A[r]`, or −1 if there is none.

### 4.4 Correctness

**Proof Method: A Lemma Plus a Loop Invariant**

#### 1. Conditions the Correctness Depends On

1. **A count grows by 0 or 1**: appending `A[r]` to a subarray either leaves its number of distinct values unchanged or increases it by one.
2. **The affected left ends form a range**: the counts that increase are exactly those with left end in $`(p, r]`$.
3. **The change of a square is a range sum**: $`(c+1)^2 - c^2 = 2c + 1`$, so the change of a range's sum of squares is twice its sum plus its length.

**Lemma**: for $`0 \leq l \leq r`$, with $`c(r, r-1) = 0`$,

```math
c(l, r) = c(l, r-1) + [\, l > p \,]
```

Proof: `A[r]` already appears among $`a_l, \ldots, a_{r-1}`$ if and only if that segment contains some $`q < r`$ with $`a_q = a_r`$, which holds if and only if $`l \leq p`$, because p is the largest such index. If it already appears the count is unchanged; otherwise it increases by one.

The lemma gives the transition of f:

```math
\begin{aligned}
f(r) &= \sum_{l=0}^{p} c(l, r-1)^2 + \sum_{l=p+1}^{r} \bigl( c(l, r-1) + 1 \bigr)^2 \\
     &= f(r-1) + 2 \sum_{l=p+1}^{r} c(l, r-1) + (r - p)
\end{aligned}
```

The convention $`c(r, r-1) = 0`$ makes the new subarray $`[r, r]`$ contribute exactly $`1^2 = 1`$, so it needs no special case.

#### 2. Proof

**Loop Invariant**

At the end of iteration r:

**Invariant I: Meaning of the Tree**

$`T[l] = c(l, r)`$ for $`l \leq r`$, and $`T[l] = 0`$ for $`l > r`$.

**Invariant II: Partial Results**

$`f = f(r) \bmod M`$ and $`\mathrm{ans} = \sum_{r'=0}^{r} f(r') \bmod M`$. The map `last` holds the last index of every value among $`a_0, \ldots, a_r`$.

**Initialization**

Before the loop (think of it as $`r = -1`$), T is all zeros, `f = ans = 0`, and `last` is empty. There are no subarrays, so the invariant holds.

**Maintenance**

Assume the invariant holds after iteration $`r-1`$. In iteration r:

- The query runs before any update, so by Invariant I, `S` equals $`\sum_{l=p+1}^{r} c(l, r-1)`$; position r holds 0, matching $`c(r, r-1) = 0`$.
- After adding 1 on $`[p+1, r]`$, the lemma shows every position $`l \leq r`$ now holds $`c(l, r)`$; positions $`l > r`$ were not touched and stay 0. Invariant I holds.
- By the transition, the new `f` equals $`f(r)`$; after adding `f`, `ans` covers every subarray whose right end is at most r. `last[A[r]]` becomes r. Invariant II holds.

**Termination**

The loop ends after n iterations with $`r = n - 1`$. Every subarray has exactly one right end, so `ans` equals the sum of squared counts over all subarrays, modulo M.

#### 3. Conclusion

The lemma turns "which subarrays change" into a single range, and expanding the square turns "how the sum of squares changes" into a sum over that range. The loop invariant keeps the tree equal to the counts of every left end for the current right end, so the algorithm is correct for every valid input.

### 4.5 Complexity Analysis

Let the input size be $`n`$.

**Time Complexity: O(n log n)**

The loop runs n times, each doing one range query and one range update in O(log n); reads and writes to `last` are amortized O(1).

**Space Complexity: O(n)**

The tree's sum and lazy arrays each use 4n slots, and `last` holds at most n keys. Because values are at most $`10^5`$, `last` can also be a fixed-size array.

## 5. Learning Notes

### Difficulty 1: Is This DP or a Segment Tree? Why Isn't DP Alone Enough?

The first question was: if f(r) can be written in terms of f(r−1), why not just do DP? Do the two approaches compute different results?

**How It Was Resolved**

The results are identical, and the algorithm really is DP. The obstacle is the term $`\sum_{l=p+1}^{r} c(l, r-1)`$ in the transition:

- it needs the sum of counts over **a range of left ends**, and since p differs at every step the range is arbitrary;
- after each step, the counts in that range must all increase by one.

So the state that must be kept is the whole array $`c(\cdot, r)`$, which cannot be compressed into O(1) scalars. Maintaining it in a plain array makes each range add and range sum O(n), which brings the total back to O(n²). A segment tree brings both operations down to O(log n).

Compare [2262. Total Appeal of A String](https://leetcode.cn/problems/total-appeal-of-a-string/): it asks for $`\sum c(l, r)`$, with no square. The same derivation gives an increment equal to the range length $`r - p`$, which does not depend on c at all, so a single scalar recurrence solves it in O(n). **Squaring is nonlinear**: the increment $`2c+1`$ of $`(c+1)^2`$ depends on c, so every c has to be maintained.

**Transferable Insight:** after writing a DP transition, check whether the information the increment needs can be summarized by a few scalars. If so, plain DP suffices; if not, a data structure has to maintain the state.

### Difficulty 2: Is a Segment Tree an Algorithm or a Data Structure?

**How It Was Resolved**

The two work at different levels:

- **DP decides what to compute**: it models the problem, defines the state f(r), and gives the transition; correctness comes from it.
- **The segment tree decides how to compute it quickly**: it only implements "add 1 on a range" and "sum a range". Replace it with a plain array and the answer is still correct, only slower.

This combination is usually called **data-structure-optimized DP**. Other examples include monotonic-queue-optimized DP and Fenwick-tree-optimized longest increasing subsequence. Conversely, segment trees also pair with sweep lines, greedy algorithms, and offline queries, so they are not specific to DP. That is why this problem is filed under Dynamic Programming, with the segment tree as a secondary category and tag.

**Transferable Insight:** when an editorial reads "DP + some data structure", strip the data structure away and write the naive O(n²) transition first; then identify which operation is the bottleneck and pick a structure that speeds up exactly that operation.

### Difficulty 3: Why Query Before Updating?

**How It Was Resolved**

The transition needs the sum of the **old** values $`c(l, r-1)`$. If the range add runs first, the query returns the sum of $`c(l, r)`$ instead, which is larger by the range length $`r - p`$, and the answer comes out too large.

Including position r in the query range is harmless: that position is exactly 0 before this iteration, matching $`c(r, r-1) = 0`$, and the contribution 1 of the new subarray $`[r, r]`$ comes from the $`r - p`$ term.

**Transferable Insight:** when a transition is written as "old state plus increment", read the old values it needs before writing the new state.

## 6. Key Takeaways

This problem yields the following reusable methods:

1. **Right-endpoint sweep**: when summing over all subarrays, fix the right end, treat all subarrays ending there as one state, and transition from r−1 to r.
2. **Last occurrence boundary**: a new element only affects subarrays whose left end lies in $`(p, r]`$. The same boundary appears in other distinct-element counting problems such as 828 and 2262.
3. **Data-structure-optimized DP**: when a transition needs range updates and range queries over the state, keep the state array in a segment tree or a pair of Fenwick trees.
4. **Expanding a nonlinear increment**: $`(c+1)^2 = c^2 + 2c + 1`$ reduces the change of a square to sums of first powers; higher powers need more maintained sums, for example cubes also need $`\sum c^2`$.

These methods fit problems that sum a statistic over all subarrays where extending the right end affects a contiguous range of left ends. If the affected left ends are not contiguous, or the increment cannot be expanded into maintainable range statistics, a different approach is needed.

## 7. Review Log

| Date | Independent Completion | Understanding Change | Needs Improvement |
|---|---|---|---|
| 2026-10-16 | Solved 2913 independently with O(n²) enumeration; 2916 not yet implemented | Understood that this problem is still DP and the segment tree only speeds up the transition; the nonlinearity of the square is why the tree is needed | Implement the lazy segment tree independently and add solution.go and tests |

### Next Review Checklist

- [ ] Can I implement it independently without looking at the code?
- [ ] Can I state the algorithm's input and output precisely?
- [ ] Can I write the pseudocode independently?
- [ ] Can I explain the correctness, especially why the affected left ends are $`(p, r]`$?
- [ ] Can I derive the time and space complexity?
- [ ] Can I explain why 2262 needs only plain DP while this problem needs a segment tree?

## 8. Related Problems

- [2913. Subarrays Distinct Element Sum of Squares I](https://leetcode.cn/problems/subarrays-distinct-element-sum-of-squares-i/) — the small-input version of the same problem, solvable by O(n²) enumeration.
- [2262. Total Appeal of A String](https://leetcode.cn/problems/total-appeal-of-a-string/) — sums distinct counts instead of their squares; the increment is constant, so plain DP suffices. A good contrast with this problem.
- [828. Count Unique Characters of All Substrings of a Given String](https://leetcode.cn/problems/count-unique-characters-of-all-substrings-of-a-given-string/) — also uses the previous (and next) occurrence to bound each character's contribution.

---

**One-Sentence Summary**

2916 is still a right-endpoint DP: a new element raises the count only for left ends in $`(p, r]`$, the expanded square needs the sum of those counts, and because that state cannot be compressed into scalars, a lazy segment tree handles the range add and range sum.
