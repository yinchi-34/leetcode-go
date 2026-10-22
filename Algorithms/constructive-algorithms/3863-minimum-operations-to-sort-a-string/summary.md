# 3863. Minimum Operations to Sort a String

English | [简体中文](./summary.zh-CN.md)

Use the [shared glossary](../../../docs/GLOSSARY.md) and follow the repository's [GitHub math conventions](../../../docs/GITHUB_MATH.md).

## 1. Problem Information

- **Problem**: [3863. Minimum Operations to Sort a String](https://leetcode.com/problems/minimum-operations-to-sort-a-string/)
- **Official Difficulty**: Medium
- **Primary Category**: Constructive Algorithms
- **Secondary Category**: Greedy, String
- **Tags**: `Constructive Algorithms`, `Greedy`, `String`, `Case Analysis`, `Upper and Lower Bounds`
- **Code**: [solution.go](./solution.go)
- **Tests**: [solution_test.go](./solution_test.go)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: 1860.00 ([source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode Official Difficulty: Medium

**Four-Dimensional Rating — Case Analysis on Extreme Positions**

- Evaluation Scope: $`1 \leq n \leq 10^5`$, lowercase letters only; time $`O(n)`$, extra space $`O(1)`$; Go implementation.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 3/5 | One must derive "at least one endpoint stays fixed" from "the whole string may not be chosen", turning a sorting problem into a question about where the global minimum and maximum sit |
| Knowledge Prerequisites | 1/5 | Only character comparison and one linear scan |
| Correctness Reasoning Difficulty | 4/5 | Needs several lemmas: the iff condition for one operation, the two- and three-operation constructions, and the lower bound that two operations are not enough |
| Implementation Difficulty | 1/5 | One scan for sortedness and extremes, then branch on the conditions; the only edge case is length 2 |

Ratings follow the repository's [RATING.md](../../../docs/RATING.md). Each dimension is rated independently and is not combined into a total.

## 2. Problem Understanding

Given a string `s` of lowercase letters. In one operation, choose a contiguous substring that is **not the whole string** and sort it in non-decreasing order. Return the minimum number of operations to make all of `s` non-decreasing, or `-1` if impossible.

**Example**

- Input: `s = "dog"`, Output: `1`. Sorting `"og"` gives `"dgo"`.
- Input: `s = "card"`, Output: `2`. Sort `"car"` to get `"acrd"`, then sort `"rd"` to get `"acdr"`.
- Input: `s = "gf"`, Output: `-1`. The only proper substrings are single characters, and sorting them changes nothing.

**Key Requirements**

1. $`1 \leq n \leq 10^5`$, and `s` contains only lowercase English letters.
2. Each operation's interval $`[l, r]`$ must satisfy $`[l, r] \neq [0, n-1]`$, though its length may be $`n-1`$.
3. An already sorted string needs `0`; an unsorted string of length 2 has no solution.

## 3. Algorithm Specification

**Algorithm Name:** Case Analysis on Extreme Positions

### 3.1 Input

A string $`s = \langle s_0, s_1, \ldots, s_{n-1} \rangle`$ such that:

- $`1 \leq n \leq 10^5`$;
- each $`s_i`$ is a lowercase letter, compared alphabetically.

Let $`m = \min_i s_i`$ and $`M = \max_i s_i`$.

### 3.2 Output

Let $`f(s)`$ be the minimum number of operations needed to make `s` non-decreasing, with $`f(s) = -1`$ if impossible. The algorithm returns $`f(s)`$.

For unsorted `s` with $`n \geq 3`$, define four booleans:

| Symbol | Meaning |
|---|---|
| $`L`$ | $`s_0 = m`$: the left end is already the global minimum |
| $`R`$ | $`s_{n-1} = M`$: the right end is already the global maximum |
| $`I_m`$ | some $`1 \leq i \leq n-2`$ has $`s_i = m`$: the minimum appears in the interior |
| $`I_M`$ | some $`1 \leq j \leq n-2`$ has $`s_j = M`$: the maximum appears in the interior |

This note proves:

```math
f(s) =
\begin{cases}
0, & s \text{ is sorted} \\
-1, & s \text{ is unsorted and } n = 2 \\
1, & s \text{ unsorted}, n \geq 3, L \lor R \\
2, & s \text{ unsorted}, n \geq 3, \lnot L \land \lnot R, I_m \lor I_M \\
3, & \text{otherwise}
\end{cases}
```

### 3.3 Pseudocode

Indices are 0-based.

```text
MIN-OPERATIONS(s)
    n ← length(s)
    m ← 'z', M ← 'a'
    sorted ← true
    for i ← 0 to n - 1
        M ← max(M, s[i])
        m ← min(m, s[i])
        if i > 0 and s[i-1] > s[i]
            sorted ← false

    if sorted
        return 0
    if n = 2
        return -1
    if s[0] = m or s[n-1] = M
        return 1
    for i ← 1 to n - 2
        if s[i] = m or s[i] = M
            return 2
    return 3
```

**State Definitions**

- `sorted`: whether `s` is already non-decreasing.
- `m`, `M`: the global minimum and maximum characters, starting from `'z'` and `'a'` and updated during the scan.

The interior scan only checks indices $`1`$ to $`n-2`$. By this point $`\lnot L \land \lnot R`$ is known, that is $`s_0 \neq m`$ and $`s_{n-1} \neq M`$, so "the minimum appears somewhere other than the last position" is equivalent to $`I_m`$, and "the maximum appears somewhere other than the first position" is equivalent to $`I_M`$. One interior scan is enough.

### 3.4 Correctness

**Proof Method:** For each case, give an upper-bound construction (it can be done) and a lower-bound proof (it cannot be done with fewer).

#### I. Key Properties Required for Correctness

1. **Endpoint invariance**: every legal operation leaves at least one endpoint unchanged.
2. **Endpoints of the target**: a sorted string starts with $`m`$ and ends with $`M`$.
3. **Power of a length $`n-1`$ sort**: one operation can sort the first $`n-1`$ or the last $`n-1`$ characters, so the middle characters never need separate handling.

#### II. Correctness Proof

**Lemma 0 (endpoint invariance)**

Let an operation sort the interval $`[l, r]`$. Since $`[l, r] \neq [0, n-1]`$, either $`l \geq 1`$ or $`r \leq n-2`$. In the first case $`s_0`$ is unchanged; in the second, $`s_{n-1}`$ is unchanged.

**Lemma 1 (iff condition for one operation)**

For unsorted `s`, $`f(s) = 1`$ if and only if $`L \lor R`$.

*Necessity*: if one operation makes `s` sorted, some endpoint did not move by Lemma 0. A sorted string starts with $`m`$ and ends with $`M`$, and an operation does not change the multiset of characters, so the unmoved endpoint already satisfied $`s_0 = m`$ or $`s_{n-1} = M`$.

*Sufficiency*: if $`L`$ holds, sort $`[1, n-1]`$. The result starts with $`m`$, which is no greater than any character, and the rest is sorted, so the whole string is sorted. The case $`R`$ is symmetric: sort $`[0, n-2]`$.

**Lemma 2 (length 2)**

When $`n = 2`$, the only proper substrings are single characters, and sorting them leaves the string unchanged. An unsorted `s` can never become sorted, so $`f(s) = -1`$.

When $`n = 1`$, the string is always sorted, and the `sorted` branch returns 0.

**Lemma 3 (exactly two)**

Let `s` be unsorted with $`n \geq 3`$, $`\lnot L \land \lnot R`$, and $`I_m \lor I_M`$. Then $`f(s) = 2`$.

*Lower bound*: by Lemma 1, one operation is not enough when $`\lnot L \land \lnot R`$.

*Upper bound*: if $`I_m`$ holds, the minimum appears at some $`1 \leq i \leq n-2`$. Sorting $`[0, n-2]`$ makes the first character $`m`$, so $`L`$ holds, and Lemma 1 finishes with one more operation. If $`I_M`$ holds, symmetrically sort $`[1, n-1]`$ first so that the last character becomes $`M`$, then use one more operation.

**Lemma 4 (exactly three)**

Let `s` be unsorted with $`n \geq 3`$, and let $`\lnot L`$, $`\lnot R`$, $`\lnot I_m`$, $`\lnot I_M`$ all hold.

Since $`m`$ must appear, $`\lnot L`$ rules out index 0, and $`\lnot I_m`$ rules out the interior, it appears **only** at $`n-1`$. Likewise $`M`$ appears **only** at $`0`$. Because `s` is unsorted, $`m < M`$.

*Lower bound*: we show that after any single operation, $`\lnot L \land \lnot R`$ still holds. Let the operation's interval be $`[l, r]`$:

- If $`l \geq 1`$: $`s_0 = M \neq m`$ is unchanged, so $`\lnot L`$. If $`r \leq n-2`$, $`s_{n-1} = m \neq M`$ is unchanged; if $`r = n-1`$, the new last character is $`\max(s_l, \ldots, s_{n-1})`$, which is strictly less than $`M`$ because $`M`$ appears only at index 0. Either way $`\lnot R`$.
- If $`l = 0`$: then $`r \leq n-2`$, so $`s_{n-1} = m \neq M`$ is unchanged and $`\lnot R`$. The new first character is $`\min(s_0, \ldots, s_r)`$, which is strictly greater than $`m`$ because $`m`$ appears only at index $`n-1`$. So $`\lnot L`$.

After one operation the string is still unsorted (its first character is not $`m`$) and $`\lnot L \land \lnot R`$ holds, so by Lemma 1 a second operation cannot finish. Hence $`f(s) \geq 3`$.

*Upper bound*: sort the first $`n-1`$, then the last $`n-1`$, then the first $`n-1`$ characters.

1. Sort $`[0, n-2]`$: $`M`$ lies in this range and is its unique maximum, so it moves to index $`n-2`$.
2. Sort $`[1, n-1]`$: since $`n \geq 3`$, index $`n-2 \geq 1`$ lies in this range, so $`M`$ moves to index $`n-1`$ and $`R`$ holds.
3. By the sufficiency part of Lemma 1, sorting $`[0, n-2]`$ makes the whole string sorted.

So $`f(s) = 3`$.

For example, `"edcba"`:

| Operation | Interval | Result |
|---|---|---|
| Start | — | `edcba` |
| 1st | `[0, 3]` | `bcdea` |
| 2nd | `[1, 4]` | `bacde` |
| 3rd | `[0, 3]` | `abcde` |

#### III. Conclusion

The five cases are mutually exclusive and cover every input: sortedness and $`n = 2`$ separate the first two; the remaining unsorted strings with $`n \geq 3`$ are split by $`L \lor R`$ and then $`I_m \lor I_M`$. Lemmas 1 through 4 give matching upper and lower bounds for each case, and the pseudocode checks the cases in the same order, so the algorithm returns $`f(s)`$ for every valid input.

In addition, while writing this note, a breadth-first search over every string of length at most 6 over the letters `a` to `d` matched the formula above in every case. This is only a supporting check; the lemmas are the proof.

### 3.5 Complexity Analysis

Let $`n`$ be the length of the string.

**Time Complexity: $`O(n)`$**

The first scan computes `sorted`, `m`, and `M` in $`O(n)`$; the scan for $`I_m \lor I_M`$ is $`O(n)`$ in the worst case. Total $`O(n)`$.

**Space Complexity: $`O(1)`$**

Only a constant number of variables are used, and the string is not copied.

## 4. Learning Notes

### Challenge 1: Why at most 3 operations even when the length exceeds 3?

After seeing that `"cba"` needs 3 operations, a natural question is whether longer strings need more.

**Understanding Process**

One operation can cover $`n-1`$ characters, so its reach grows with $`n`$. The three-step construction (prefix, suffix, prefix) depends only on moving $`M`$ to the right end, not on the order of the middle characters or the size of $`n`$, so the upper bound is always 3. The lower bound looks only at the two endpoints and is also independent of $`n`$.

**Transferable Insight:** When the reach of one operation grows with the problem size, the answer may be a constant independent of size; look for a constant upper bound first.

## 5. Key Takeaways

1. **Endpoint Invariance**: here, "the whole string may not be chosen" implies that every operation leaves at least one endpoint fixed, which gives the necessary condition for finishing in one step. It applies to sorting, reversing, or modifying problems with restricted intervals, provided the restriction really prevents an interval from covering both ends.
   - Example: [2337. Move Pieces to Obtain a String](https://leetcode.com/problems/move-pieces-to-obtain-a-string/). `L` only moves left, `R` only moves right, and pieces cannot pass each other, so the order of pieces with `_` removed never changes, each `L` ends at an index no greater than where it started, and each `R` at an index no smaller. As in this problem, derive an invariant from the operation's restriction first, then use it to decide feasibility.
2. **Upper and Lower Bounds**: a minimum-operations answer needs two proofs: a construction showing it can be done, and an invariant showing it cannot be done with fewer. Here, the three-step construction gives the upper bound and Lemma 4 shows two operations are not enough. This applies to every minimum-operations problem.
   - Example: [1529. Minimum Suffix Flips](https://leetcode.com/problems/minimum-suffix-flips/). Flipping a suffix changes whether exactly one pair of adjacent positions differs, so the answer is at least the number of adjacent differing positions in the target (with a `0` prepended); flipping greedily from left to right whenever they differ uses exactly that many, so the bounds match.
3. **Value-Preserving Abstraction**: compress a full arrangement into a few features and prove that the optimum depends only on them. It suits constructive problems whose answer takes only a few values; its limit is that it does not guarantee transitions can be expressed on the compressed states.
   - Example: [1033. Moving Stones Until Consecutive](https://leetcode.com/problems/moving-stones-until-consecutive/). With the three stones sorted, the minimum number of moves is only 0, 1, or 2 and depends only on the two gaps: 0 if both gaps are 1, 1 if some gap is at most 2, otherwise 2. The exact coordinates compress into a few features: whether a gap is 1 and whether it is at most 2.

These methods fit constructive and minimum-operations problems where each operation is powerful and the answer is a small constant. For such problems, first study what must stay fixed and what can change, rather than simulating operations.

### CLRS Connections

Section numbers follow CLRS, 3rd edition. Relation is either "Direct" or "Related idea"; if no section fits, write "No matching section" and keep only the concept name.

| Concept | CLRS Section | Relation | How It Appears Here |
|---|---|---|---|
| Lower-bound argument | §8.1 Lower bounds for sorting | Related idea | Lemma 4's lower bound: after any single operation both endpoints are still wrong, so every strategy needs at least three operations, not just one particular sequence |
| Upper-bound construction | No matching section | — | Upper bounds in Lemmas 3 and 4: explicit two- and three-operation sequences |
| Value-preserving abstraction | No matching section | — | The answer is determined by sortedness, $`L`$, $`R`$, $`I_m`$, and $`I_M`$ alone |

## 6. Review Log

| Date | Independent Completion | Understanding Gained | Areas for Improvement |
|---|---|---|---|
| 2026-10-21 | Wrote the code independently | Revised "like merge sort, three segments" into "three overlapping intervals, only one endpoint needs fixing"; asked about lengths above 3 and whether the state compression can be proved | Write Lemma 4's lower bound without looking |

### Next Review Checklist

- [ ] Can I implement the algorithm without looking at the solution?
- [ ] Can I accurately define the input and output?
- [ ] Can I write the pseudocode independently?
- [ ] Can I explain why the algorithm is correct?
- [ ] Can I derive its time and space complexity?
- [ ] Can I recognize the same algorithmic pattern in other problems?
- [ ] Can I explain why `"cba"` cannot be done in two operations?
- [ ] Can I explain why length 2 has no solution while length 1 gives 0?

## 7. Related Problems

- [969. Pancake Sorting](https://leetcode.com/problems/pancake-sorting/) — also sorts with a restricted interval operation (prefix reversal), built around moving one extreme to its final position at a time.

---

**One-Sentence Summary**

"The whole string may not be chosen" means every operation leaves an endpoint fixed, so the answer depends only on where the global minimum and maximum sit, and matching upper-bound constructions and lower-bound proofs pin it to 0, 1, 2, 3, or -1.
