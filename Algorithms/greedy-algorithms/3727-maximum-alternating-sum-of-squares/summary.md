# 3727. Maximum Alternating Sum of Squares

English | [简体中文](./summary.zh-CN.md)

[Global Glossary](../../../docs/GLOSSARY.md)

## 1. Problem Information

- **Problem**: [3727. Maximum Alternating Sum of Squares](https://leetcode.com/problems/maximum-alternating-sum-of-squares/)
- **Official Difficulty**: Medium
- **Primary Category**: Greedy
- **Secondary Category**: Sorting
- **Technique Tags**: `Greedy`, `Sorting`
- **Code**: [solution.go](./solution.go)
- **Tests**: [solution_test.go](./solution_test.go)

### Difficulty Rating

**External Ratings**

- zerotrac rating: 1454
- LeetCode official difficulty: Medium

**Four-Dimension Rating — Greedy + Sorting**

Scope: the given input constraints, using the most basic greedy + sorting solution (square, sort ascending, add the last ⌈n/2⌉ and subtract the first ⌊n/2⌋), time complexity O(n log n), evaluated on the Go implementation.

| Dimension | Score | Rationale |
|---|---|---|
| Modeling | 2/5 | Must realize that "free rearrangement" erases positional information, reducing the problem to "split the elements into a plus group and a minus group" |
| Prerequisite Knowledge | 1/5 | Only sorting and the basic greedy idea are needed |
| Correctness Argument | 2/5 | An upper-bound argument suffices, but the number of plus positions must be stated precisely as ⌈n/2⌉ |
| Implementation | 1/5 | Very short code; watch out for sorting by squared value, choosing the split point ⌊n/2⌋, and using `int64` for the result |

The four dimensions are rated independently and are not combined into a total score.

## 2. Problem Understanding

Given an integer array `nums`, the elements may be **rearranged in any order** to form an array `arr`. The alternating score is defined as:

$$
score = arr_0^2 - arr_1^2 + arr_2^2 - arr_3^2 + \cdots
$$

Find the maximum possible alternating score after rearranging.

Examples:

- `[1,2,3]`: rearrange to `[2,1,3]`, score `4 - 1 + 9 = 12`.
- `[1,-1,2,-2,3,-3]`: the maximum score is `16`.

Key constraints:

1. Any rearrangement is allowed; the original order is irrelevant.
2. Even indices take a plus sign, odd indices take a minus sign.
3. Elements may be negative: `-4·10^4 ≤ nums[i] ≤ 4·10^4`.
4. `1 ≤ n ≤ 10^5`.

### Core Problem

This problem is not about finding some clever permutation.

Because any rearrangement is allowed, **the only real decision is: which elements go to plus positions, and which go to minus positions.**

How elements are ordered within the plus positions, or within the minus positions, does not change the score at all.

Therefore:

**This is a grouping problem, not a permutation problem.**

## 3. Initial Approach

### 3.1 First Version: Take Absolute Values and Split into Two Groups Explicitly

The first idea was straightforward: the larger the absolute value, the larger the square, and the more it belongs in a plus position.

```text
MAX-ALTERNATING-SUM-V1(nums)
    n ← length(nums)
    for i ← 0 to n - 1
        d[i] ← |nums[i]|

    SORT-DESC(d)                         // descending

    if n mod 2 == 0
        plus   ← d[0 .. n/2 - 1]         // first n/2 elements
        minus  ← d[n/2 .. n - 1]
    else
        plus   ← d[0 .. ⌊n/2⌋]           // first ⌊n/2⌋ + 1 elements
        minus  ← d[⌊n/2⌋ + 1 .. n - 1]

    return Σ_{x ∈ plus} x² − Σ_{y ∈ minus} y²
```

The idea of this version is correct:

- Sorting by absolute value gives the same order as sorting by squared value ($`|x|`$ and $`x^2`$ increase and decrease together).
- After sorting in descending order, the first ⌈n/2⌉ go to the plus group and the rest to the minus group.

But the implementation has some redundancy:

- It needs two branches based on the parity of n to compute the split position.
- It builds two extra lists, `plus` and `minus`.
- It takes absolute values first and squares at the end, splitting the transformation into two steps.

### 3.2 Shift in Thinking: Unify Both Parities with One Split Point

Look at the two branches again:

- n even: n/2 in the plus group, n/2 in the minus group.
- n odd: (n+1)/2 in the plus group, (n−1)/2 in the minus group.

In both cases, **the minus group always has size ⌊n/2⌋**, and the plus group is simply the remaining ⌈n/2⌉.

So instead, sort the squared values in **ascending** order and let:

$$
k = \lfloor n/2 \rfloor
$$

Then `sq[:k]` is the minus group and `sq[k:]` is the plus group; both parities are covered by the same split point:

```text
n = 3: k = 1
  sq[:1] → minus (1 element)
  sq[1:] → plus  (2 elements = ⌈3/2⌉)

n = 4: k = 2
  sq[:2] → minus (2 elements)
  sq[2:] → plus  (2 elements = ⌈4/2⌉)
```

The simplified implementation:

```text
MAX-ALTERNATING-SUM-V2(nums)
    n ← length(nums)
    for i ← 0 to n - 1
        a[i] ← nums[i]²

    SORT(a)                              // ascending
    k ← ⌊n/2⌋                            // split point = size of minus group

    return Σ_{i=k}^{n-1} a[i] − Σ_{i=0}^{k-1} a[i]
```

### Summary of the Shift in Thinking

From:

> Depending on the parity of n, compute how many elements go to the plus group and to the minus group.

To:

> After sorting, the two groups are a prefix and a suffix; only one split point needs to be determined.

The former focuses on "how many in each group"; the latter focuses on "where the split is".

**Core insight: a grouping problem after sorting often reduces to finding a single split point; choosing the right definition for that point (the size of the smaller group, ⌊n/2⌋) eliminates the parity branches.**

## 4. Algorithm Specification

**Algorithm name:** Greedy

### 4.1 Input

An integer array of length n:

$$
nums = \langle x_0, x_1, \ldots, x_{n-1} \rangle
$$

**Preconditions:**

- **Length constraint**: $`1 \leq n \leq 10^5`$.
- **Value constraint**: $`-4 \times 10^4 \leq x_i \leq 4 \times 10^4`$.
- **Order constraint**: none; any rearrangement is allowed.

### 4.2 Output

An integer `k`, the maximum alternating score after rearranging.

Let $`S_n`$ be the set of all permutations of $`\{0, 1, \ldots, n-1\}`$. The objective is:

$$
\text{OPT} = \max_{\sigma \in S_n} \sum_{j=0}^{n-1} (-1)^j \, x_{\sigma(j)}^2
$$

Equivalent grouping form: let $`P`$ be the index set of the plus group with $`|P| = p = \lceil n/2 \rceil`$, $`a_i = x_i^2`$, and $`T = \sum_i a_i`$. Then:

$$
\text{OPT} = \max_{|P| = p} \left( \sum_{i \in P} a_i - \sum_{i \notin P} a_i \right)
  = \max_{|P| = p} \left( 2\sum_{i \in P} a_i - T \right)
$$

Since $`T`$ is a constant, **maximizing the score is equivalent to maximizing the sum of the plus group.**

### 4.3 Pseudocode

```text
MAX-ALTERNATING-SUM(nums)
    n ← length(nums)

    for i ← 0 to n - 1
        a[i] ← nums[i] * nums[i]

    SORT(a)                      // ascending
    k ← ⌊n/2⌋                    // size of minus group

    res ← 0
    for i ← 0 to n - 1
        if i ≥ k
            res ← res + a[i]     // largest ⌈n/2⌉: plus positions
        else
            res ← res - a[i]     // smallest ⌊n/2⌋: minus positions

    return res
```

**State Definitions**

- `a[i]`: the squared value of the i-th element (the Go implementation overwrites `nums[i]` in place).
- `k`: the number of minus positions, which is also the starting index of the plus group after sorting.
- `res`: the accumulated alternating score so far, stored as `int64`.

**Greedy Strategy**

Elements with larger squared values belong in plus positions; elements with smaller squared values belong in minus positions.

### 4.4 Construction

The algorithm only computes the score, but a permutation achieving it can be given explicitly:

1. After sorting ascending, take the plus group as $`a_{k}, \ldots, a_{n-1}`$ and the minus group as $`a_0, \ldots, a_{k-1}`$.
2. Place the plus-group elements at indices 0, 2, 4, …; place the minus-group elements at indices 1, 3, 5, ….

The plus group has exactly p elements and there are exactly p even indices; likewise for the minus group. So the construction is always feasible.

For example, `[1,2,3]`: the squares are `[1,4,9]`, k = 1, plus group `{4, 9}`, minus group `{1}`, constructing `[2,1,3]` with score 12.

### 4.5 Correctness

**Proof method: loop invariant (correct evaluation) + upper-bound argument (global optimality)**

> **Is the upper-bound argument itself greedy?** No. Greedy is the algorithm's selection strategy: "put larger squares in plus positions". The upper-bound argument is a method for proving that this strategy is optimal: it shows that no arrangement can score more than the value produced by greedy. It corresponds to the lower-bound argument in part four, "Why is it globally optimal?", of problem 1541; 1541 is a minimization, so it proves a lower bound, while this problem is a maximization, so it proves an upper bound. The exchange argument is another common way to prove greedy optimality; see the end of part four.

#### I. Key Properties

**Property A: the score depends only on the grouping**

The score is:

$$
\sum_{j \text{ even}} a_{\sigma(j)} - \sum_{j \text{ odd}} a_{\sigma(j)}
$$

Swapping two elements at even indices, or two elements at odd indices, leaves both sums unchanged. Hence the score is determined solely by the plus-group set $`P`$.

Conversely, for any set with $`|P| = p`$, a corresponding permutation can be built as in 4.4. So the permutation problem and the grouping problem are **fully equivalent**.

**Property B: there are exactly ⌈n/2⌉ plus positions**

Indices start at 0 and even indices take a plus sign:

$$
p = \lceil n/2 \rceil = n - \lfloor n/2 \rfloor,\qquad q = k = \lfloor n/2 \rfloor
$$

So every valid grouping satisfies $`|P| = p`$, and the minus group has size exactly $`k`$.

**Property C: maximizing the difference is equivalent to maximizing the plus-group sum**

$$
\text{score}(P) = 2\sum_{i \in P} a_i - T
$$

$`T`$ does not depend on the grouping, so the score increases monotonically with the plus-group sum.

#### II. Loop Invariants

Let the squared values be sorted ascending as $`a_0 \le a_1 \le \cdots \le a_{n-1}`$, and $`k = \lfloor n/2 \rfloor`$.

After processing the first $`i`$ elements:

**Invariant I (accumulated value)**

$$
res = \sum_{j=0}^{i-1} c_j \, a_j,\qquad
c_j =
\begin{cases}
-1, & j < k \\
+1, & j \ge k
\end{cases}
$$

**Invariant II (valid grouping)**

Among the processed elements, those negated are exactly the smallest $`\min(i, k)`$ of them, and there are at most $`k`$ of them; those added are at most $`p`$.

#### III. Proof

**1. Initialization**

When $`i = 0`$, $`res = 0`$ and no element has been processed; both invariants hold.

**2. Maintenance**

When processing element $`i`$:

- If $`i < k`$: execute `res -= a[i]`, corresponding to $`c_i = -1`$. The number of negated elements before this step is $`i < k`$, and after it is $`i + 1 \le k`$, within the minus-group capacity.
- If $`i \ge k`$: execute `res += a[i]`, corresponding to $`c_i = +1`$. The number of added elements is $`i - k + 1 \le n - k = p`$, within the plus-group capacity.

Because the array is sorted ascending, the first $`k`$ elements are always the smallest. Both invariants continue to hold.

**3. Termination**

When $`i = n`$:

$$
res = \sum_{j=k}^{n-1} a_j - \sum_{j=0}^{k-1} a_j
$$

The plus group $`P_g = \{k, \ldots, n-1\}`$ has exactly $`n - k = p`$ elements, so it is a valid grouping. By Property A and the construction in 4.4, some permutation achieves exactly this score, so:

$$
\text{OPT} \ge res
$$

So far this only shows the algorithm's result is **achievable**; it remains to show it **cannot be exceeded**.

#### IV. Why Is It Globally Optimal?

**Upper-bound argument**

For any grouping with $`|P| = p`$, the plus-group sum cannot exceed "the sum of the p largest squared values":

$$
\sum_{i \in P} a_i \le \sum_{j=k}^{n-1} a_j
$$

By Property C:

$$
\text{score}(P) = 2\sum_{i \in P} a_i - T
\;\le\; 2\sum_{j=k}^{n-1} a_j - T
= res
$$

So no permutation scores more than $`res`$:

$$
\text{OPT} \le res
$$

Combined with $`\text{OPT} \ge res`$ from part III:

$$
\text{OPT} = res
$$

**Therefore, the algorithm obtains the global maximum alternating score.**

**Alternative view: exchange argument**

Take any optimal grouping $`P^*`$. If it is not "the p largest", there exist $`i \in P^*`$ and $`j \notin P^*`$ with $`a_i < a_j`$. Swapping their groups changes the score by:

$$
\Delta = 2(a_j - a_i) > 0
$$

contradicting the optimality of $`P^*`$. So the plus group of an optimal grouping must be the p largest (when squared values tie, swapping does not change the score, so the conclusion still holds).

**Why must we sort by squared value?**

The upper-bound argument compares $`a_i = x_i^2`$. When the array contains negatives, $`x_i < x_j`$ does not imply $`x_i^2 < x_j^2`$, so the sort key must be the squared value or the absolute value. The Go implementation runs `nums[i] *= nums[i]` before `sort.Ints(nums)`, so the sort key is naturally the squared value.

**How the proof maps to intuition**: Property A corresponds to "after rearrangement only the grouping remains", Property B to "the split point is ⌊n/2⌋", and the upper-bound argument to "large values go to plus positions". Together they show there is no better scheme than "sort, then cut once".

### 4.6 Complexity Analysis

Let n be the length of the input array.

**Time complexity: O(n log n)**

- Squaring: O(n).
- Sorting: O(n log n).
- Summation: O(n).

$$
T(n) = O(n \log n)
$$

**Extra space complexity: O(log n)**

- The Go implementation overwrites `nums` with squared values in place, allocating no extra array.
- `sort.Ints` uses pdqsort with O(log n) recursion depth.

$$
S(n) = O(\log n)
$$

In the Python version, `sorted` creates a new list, so the extra space is O(n).

**Value range:**

- The largest single square is $`(4 \times 10^4)^2 = 1.6 \times 10^9`$, close to the 32-bit signed limit $`2^{31} - 1 \approx 2.147 \times 10^9`$; adding two of them overflows int32.
- The absolute value of the sum is bounded by $`10^5 \times 1.6 \times 10^9 = 1.6 \times 10^{14}`$, which fits in int64.
- In LeetCode's Go environment `int` is 64-bit, so `nums[i] *= nums[i]` does not overflow; accumulating into `int64` matches the function signature.

### 4.7 Optional Optimization: Linear Time

Sorting is not strictly necessary; we only need the p largest squared values:

1. **Quickselect**: find the k-th smallest squared value as the split point, expected O(n).
2. **Counting sort**: since $`|x_i| \leq 4 \times 10^4`$, count by absolute value in O(n + V) time and O(V) space, where V = 4·10⁴.

At this problem's scale, O(n log n) sorting is already sufficient; the optimization mainly illustrates that "only the split point is needed, not a total order".

## 5. Learning Notes

### Difficulty 1: Why is "any rearrangement allowed" the key?

Because free rearrangement makes the original order completely meaningless.

Once this is recognized, there is no need to think about "how to construct a permutation"; only "how to group" matters.

Compare with [1911. Maximum Alternating Subsequence Sum](https://leetcode.com/problems/maximum-alternating-subsequence-sum/): a subsequence must preserve the original order, so dynamic programming is needed; here the order is free, so greedy suffices.

**Transferable thinking: when a problem allows arbitrary rearrangement or arbitrary swaps, first check whether it reduces to a selection problem on a multiset.**

### Difficulty 2: Why sort by squared value rather than original value?

Because the score depends only on squared values, and negative numbers make the order of original values differ from the order of squared values.

For example, `[-5, 1, 2]`:

```text
Sorted by original value: [-5, 1, 2]
  n = 3, k = 1, minus group: -5 → subtract 25
  plus group: 1, 2 → add 1 + 4
  score: 5 - 25 = -20     ✗

Sorted by squared value: [1, 4, 25]
  minus group: 1
  plus group: 4, 25
  score: 29 - 1 = 28      ✓
```

Sorting by absolute value (as in the first version) is also correct, because $`|x|`$ and $`x^2`$ have the same ordering.

**Transferable thinking: greedy must compare the quantity that actually drives the objective. Transform the input into its "contribution" first, then apply greedy.**

### Difficulty 3: Why is the split point `n // 2`?

The key is whether there are more plus positions than minus positions, or the same number.

- Indices start at 0, and even indices take a plus sign.
- There are ⌈n/2⌉ even indices and ⌊n/2⌋ odd indices.

After sorting ascending, the first ⌊n/2⌋ go to the minus group, i.e. `sq[:k]`; the rest go to the plus group, i.e. `sq[k:]`.

The first version had two branches by parity, which were essentially computing ⌈n/2⌉ separately. Using the size of the smaller group, ⌊n/2⌋, as the split point unifies both cases automatically.

Writing `k = (n+1)//2` by mistake would subtract an element that should be added when n is odd.

**Transferable thinking: when dealing with "half", first decide which side gets the extra element when the count is odd, then choose a split-point definition that needs no branching.**

## 6. Key Takeaways

### 6.1 Free Rearrangement ⇒ Positions Reduce to Counts

When elements can be arranged arbitrarily and the objective depends only on "which class each position belongs to":

> Positional information can be compressed into "how many positions of each class there are".

This problem compresses n positions into "p plus positions and k minus positions".

Applicable to:

- Optimization problems that assign elements to several classes of slots.
- Weighted rearrangement-sum problems (e.g. rearrangement-inequality problems).

### 6.2 Linearizing the Objective: Maximizing a Difference ⇔ Maximizing One Side

$$
\sum_{P} a - \sum_{\bar P} a = 2\sum_{P} a - T
$$

When the total is fixed, "maximize the difference between two groups" is equivalent to "maximize the sum of one group".

This trick is common in grouping, partitioning, and sign-assignment problems.

### 6.3 Two Ways to Prove Greedy Optimality: Upper-Bound Argument and Exchange Argument

Greedy is a selection strategy; its optimality must be proven separately. Two common methods:

**Upper-bound (lower-bound) argument**

1. Show the greedy result is achievable (construct a valid solution).
2. Show no valid solution is better than greedy (upper bound for maximization, lower bound for minimization).
3. The two coincide, so greedy is optimal.

**Exchange argument**

1. Assume an optimal solution differs from the greedy one.
2. Find a pair of elements that are "out of order".
3. Show that swapping them does not make the result worse.
4. Swap repeatedly until it becomes the greedy solution.

Both methods are straightforward here: the upper bound comes from "the sum of the p largest", and the exchange gain is $`2(a_j - a_i)`$. Problem 1541 uses a lower-bound argument.

### 6.4 Transform First, Then Apply Greedy

When elements appear in the objective as $`f(x)`$, greedy should compare $`f(x)`$ (or a quantity with the same ordering), not $`x`$.

Here $`f(x) = x^2`$ is not monotonic in $`x`$ when negatives are present, but it is monotonic in $`|x|`$, so sorting by squared value or by absolute value both work.

### 6.5 Relation to 1911

| Aspect | 1911. Maximum Alternating Subsequence Sum | 3727. Maximum Alternating Sum of Squares |
|---|---|---|
| Allowed operation | Choose a subsequence, preserving order | Rearrange all elements freely |
| Element contribution | Original value | Squared value |
| Do positions matter? | Yes; index parity depends on how many elements were chosen | No; only the counts of plus and minus positions matter |
| Core model | State-machine DP (next sign plus or minus) | Grouping problem |
| Core algorithm | Dynamic programming / greedy (sum of rising segments) | Greedy + sorting |
| Correctness focus | Completeness of state transitions | Upper-bound argument / exchange argument |

Both problems emphasize:

**First determine whether an ordering constraint exists; it decides whether a state machine is needed or the problem can be grouped directly.**

## 7. Review Log

| Date | Independent Completion | Change in Understanding | To Improve |
|---|---|---|---|
| 2026-10-13 | Completed the Python version: first version used absolute values + parity branches, then simplified to sorted squares + split point; Go implementation completed | Realized free rearrangement reduces the problem to grouping; using ⌊n/2⌋ as the split point eliminates the parity branches | Add tests; write the upper-bound argument independently; try an O(n) version with quickselect or counting sort |

### Next Review Checklist

- [ ] Can I write the Go implementation without looking at the code?
- [ ] Can I explain why "any rearrangement allowed" turns the problem from permutation into grouping?
- [ ] Can I use `[-5, 1, 2]` to show why sorting by original value fails?
- [ ] Can I state that there are ⌈n/2⌉ plus positions, and explain why the split point is `n // 2`?
- [ ] Can I derive "maximize the difference ⇔ maximize the plus-group sum"?
- [ ] Can I distinguish "the greedy strategy" from "a method for proving greedy optimal"?
- [ ] Can I independently write the loop invariant's initialization, maintenance, and termination, plus the upper-bound argument?
- [ ] Can I explain why the result needs `int64`?
- [ ] Can I give an O(n) version with quickselect or counting sort?
- [ ] Can I compare with 1911 to explain how the presence of an ordering constraint affects modeling?

### Suggested Test Cases

| Input | Expected Output | What It Checks |
|---|---|---|
| `[1,2,3]` | 12 | Official example, odd n |
| `[1,-1,2,-2,3,-3]` | 16 | Official example, symmetric signs |
| `[5]` | 25 | Single element, only a plus position |
| `[-4,4]` | 0 | Equal squares cancel out |
| `[0,0,0]` | 0 | All zeros |
| `[-5,1,2]` | 28 | Negative number has the largest square; checks sorting by square |
| `[3,-5,2,0]` | 30 | Even n, mixed signs and zero |
| `[-40000,40000,40000]` | 1600000000 | Single square near the int32 limit |
| Length 10⁵, all `40000` | 0 | Large scale, even length, intermediate sums exceed int32 |

## 8. Related Problems

- [1911. Maximum Alternating Subsequence Sum](https://leetcode.com/problems/maximum-alternating-subsequence-sum/) — Also maximizes an alternating sum, but preserves order and needs state-machine DP; good for contrasting "with vs. without an ordering constraint".
- [3695. Maximize Alternating Sum Using Swaps](https://leetcode.com/problems/maximize-alternating-sum-using-swaps/) — Swaps are allowed only between given index pairs; use union-find to split into connected components, then apply this problem's "large values to plus positions" greedy within each component.
- [3701. Compute Alternating Sum](https://leetcode.com/problems/compute-alternating-sum/) — The basic version without rearrangement; direct simulation.
- [0561. Array Partition](https://leetcode.com/problems/array-partition/) — Another "sort, then group by position" greedy whose optimality can be proven by an exchange argument.

---

**One-Sentence Summary**

The key to 3727 is recognizing that "free rearrangement" leaves only the counting information "⌈n/2⌉ plus positions and ⌊n/2⌋ minus positions", reducing a permutation problem to a grouping problem; greedily assign the ⌈n/2⌉ largest squares a plus sign, which after sorting means a single cut at ⌊n/2⌋, and prove global optimality by "achievable + upper bound".