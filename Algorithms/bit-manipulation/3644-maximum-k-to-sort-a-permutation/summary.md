# 3644. Maximum K to Sort a Permutation

English | [简体中文](./summary.zh-CN.md)

Use the [shared glossary](../../../docs/GLOSSARY.md) and follow the repository's [GitHub math conventions](../../../docs/GITHUB_MATH.md).

## 1. Problem Information

- **Problem**: [3644. Maximum K to Sort a Permutation](https://leetcode.cn/problems/maximum-k-to-sort-a-permutation/)
- **Official Difficulty**: Medium
- **Primary Category**: Bit Manipulation
- **Secondary Category**: Constructive Algorithms
- **Tags**: `Bit Manipulation`, `Constructive Algorithms`, `Permutation`, `Bitwise AND`, `Upper Bound plus Construction`
- **Code**: [solution.go](./solution.go)
- **Tests**: [solution_test.go](./solution_test.go)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: 1775.45 ([source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode Official Difficulty: Medium

**Four-Dimensional Rating — Bitwise AND of Misplaced Elements**

- Evaluation Scope: $`1 \le n \le 10^5`$, `nums` is a permutation of $`0 \ldots n-1`$; time $`\Theta(n)`$, auxiliary space $`O(1)`$; Go implementation.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 3 | One must notice that every misplaced element takes part in at least one swap, which limits $`k`$ to a common submask of all misplaced elements, and then think of sorting through a hub element |
| Knowledge Prerequisites | 2 | Needs the submask property of bitwise AND and basic facts about permutations; no specialized algorithm |
| Correctness Reasoning Difficulty | 3 | Needs separate proofs of the upper bound (necessity) and reachability (sufficiency); sufficiency requires constructing a swap sequence whose hub may be an element that is already in place |
| Implementation Difficulty | 1 | A single pass; the only detail is starting the AND at `-1` (all ones) and telling apart the "no misplaced element" case |

Ratings follow [RATING.md](../../../docs/RATING.md). Each dimension is evaluated independently and not combined into a total score.

## 2. Problem Understanding

We are given a permutation `nums` of $`0 \ldots n-1`$. After choosing a non-negative integer $`k`$, we may swap two positions only when `nums[i] & nums[j] == k`. Return the largest $`k`$ that lets us sort the array in increasing order; if the array is already sorted, return 0.

**Example**

- Input: `nums = [0,3,2,1]`
- Output: `1` ($`3 \mathbin{\&} 1 = 1`$, and swapping 3 with 1 sorts the array)

**Key Requirements**

1. Input constraints: $`1 \le n \le 10^5`$, and `nums` contains each of $`0 \ldots n-1`$ exactly once.
2. The swap condition is that the AND **equals** $`k`$ exactly, not that it contains $`k`$. Any number of swaps is allowed, including swaps of elements already in place.
3. Boundary: the answer is defined as 0 for a sorted array; whenever some element is misplaced, $`k = 0`$ is always feasible (see 4.4), so an answer always exists.

## 4. Algorithm Specification

**Algorithm Name:** Bitwise AND of Misplaced Elements

Indices in this section start at 0.

### 4.1 Input

A permutation $`A = \langle a_0, a_1, \ldots, a_{n-1} \rangle`$ with $`\{a_0, \ldots, a_{n-1}\} = \{0, \ldots, n-1\}`$.

### 4.2 Output

Call $`k`$ **feasible** if some finite sequence of swaps, each between two values whose AND is exactly $`k`$, ends with $`a_i = i`$ for every $`i`$.

Let the set of misplaced values be $`S = \{a_i : a_i \ne i\}`$. The output $`\mathrm{res}`$ satisfies:

- If $`S = \varnothing`$, then $`\mathrm{res} = 0`$;
- Otherwise $`\mathrm{res}`$ is feasible, and every feasible $`k`$ satisfies $`k \le \mathrm{res}`$.

We will prove that it equals

```math
M = \mathop{\&}_{v \in S} v
```

### 4.3 Pseudocode

```text
MAX-K-TO-SORT(A)
    res ← -1                      ▷ all ones in binary, the identity of AND
    for i ← 0 to n - 1
        if A[i] ≠ i
            res ← res AND A[i]
    if res == -1                  ▷ no misplaced element, already sorted
        return 0
    return res
```

**State Definitions**

- `res`: the AND of all misplaced values in the scanned prefix; `-1` while none has been seen.
- `i`: the position being checked.

### 4.4 Correctness

**Proof Method: Loop Invariant + Upper Bound (Necessity) + Construction (Reachability)**

Notation: $`x \subseteq y`$ means every 1 bit of $`x`$ is also 1 in $`y`$, that is, $`x \mathbin{\&} y = x`$.

#### I. Key Properties Required for Correctness

1. **Submask property (bitwise fact)**: if $`a \mathbin{\&} b = k`$, then $`k \subseteq a`$ and $`k \subseteq b`$. If $`k \subseteq v`$ for every $`v`$ in a set, then $`k \subseteq \mathop{\&} v`$. $`x \subseteq y`$ implies $`x \le y`$.
2. **Misplaced values and misplaced positions are the same set (Lemma 1)**: $`S = \{i : a_i \ne i\}`$.
3. **Upper bound (Lemma 2)**: when $`S \ne \varnothing`$, every feasible $`k`$ satisfies $`k \subseteq M`$.
4. **Reachability (Lemma 3)**: when $`S \ne \varnothing`$, $`M`$ is feasible.

**Lemma 1.** The set of in-place values $`F = \{a_i : a_i = i\}`$ clearly equals the set of in-place positions $`\{i : a_i = i\}`$. $`S`$ is all values minus $`F`$, and the misplaced positions are all positions minus $`F`$; since both all values and all positions are $`\{0, \ldots, n-1\}`$, the two sets are equal.

**Lemma 2.** Take any $`v \in S`$. Value $`v`$ is not at position $`v`$ now but must end there, and only swaps move elements, so $`v`$ takes part in at least one swap. If the partner is $`u`$, then $`v \mathbin{\&} u = k`$, and property 1 gives $`k \subseteq v`$. This holds for every $`v \in S`$, so $`k \subseteq M`$ and hence $`k \le M`$.

**Lemma 3 (construction).** Since $`M \subseteq \min S \le n-1`$, the value $`M`$ occurs in the array. Let the participating set be $`T = S \cup \{M\}`$, and call $`M`$ the **hub**.

- If $`M \in S`$, then $`T = S`$;
- If $`M \notin S`$, then $`M`$ is an in-place element sitting at position $`M`$.

By Lemma 1, the positions occupied by the values of $`T`$ are exactly $`T`$. For every $`x \in T`$ we have $`M \subseteq x`$, so $`M \mathbin{\&} x = M`$: **swapping the hub with any element of $`T`$ is allowed.** Using only such swaps, repeat:

- **Step A**: the hub is at position $`p \ne M`$. Since $`p \in T`$, value $`p`$ is in $`T`$ and is not the hub; swap the hub with value $`p`$, which puts $`p`$ home.
- **Step B**: the hub is at position $`M`$, but some $`x \in T`$ is not at position $`x`$. Swap the hub with $`x`$; the hub is now away from home, so the next step is Step A.
- If neither applies, every element of $`T`$ is home, and elements outside $`T`$ were never moved and were already home, so the array is sorted.

Termination: let $`\Phi`$ be the number of non-hub elements of $`T`$ that are not home. Step A lowers $`\Phi`$ by 1. Step B leaves $`\Phi`$ unchanged but is always followed by Step A (Step B moves the hub to $`x`$'s old position $`q \ne M`$, and Step A never moves an element that is already home). So $`\Phi`$ drops by at least 1 every two steps, and the process ends after at most $`2\lvert T \rvert`$ swaps.

#### II. Correctness Proof

First we show that the code computes $`M`$.

**Loop Invariant**

At the start of iteration $`i`$:

**Invariant I: AND of the misplaced values in the prefix**

`res` equals the AND of all values in $`\{a_j : 0 \le j < i,\ a_j \ne j\}`$; when that set is empty, `res` $`= -1`$.

**Initialization**

For $`i = 0`$ the set is empty and `res` $`= -1`$, so the invariant holds.

**Maintenance**

- If $`a_i = i`$: the set is unchanged and so is `res`.
- If $`a_i \ne i`$: the new set gains $`a_i`$. If the old set was empty, `res` becomes $`-1 \mathbin{\&} a_i = a_i`$ (`-1` is all ones in two's complement); otherwise, by associativity of AND, `res` becomes the old value ANDed with $`a_i`$. The invariant holds in both cases.

**Termination**

The loop ends at $`i = n`$, when the set is exactly $`S`$.

- $`S = \varnothing`$: `res` $`= -1`$ and the function returns 0, as required.
- $`S \ne \varnothing`$: `res` $`= M`$. All $`a_i \ge 0`$, so their AND has sign bit 0 and cannot equal $`-1`$; the array is not mistaken for sorted, and $`M`$ is returned.

#### III. Conclusion

A sorted array returns 0. Otherwise the result is $`M`$: by Lemma 3, $`M`$ is feasible; by Lemma 2, no feasible $`k`$ exceeds $`M`$. Hence $`M`$ is the largest feasible value, and the algorithm is correct for every valid input.

In addition, an exhaustive Python search outside the repository (searching the swap state graph for each $`k`$) over all permutations with $`n \le 7`$ agreed with $`M`$ in every case. This is only a sanity check and does not replace the proof above.

### 4.5 Complexity Analysis

Let $`n`$ be the array length.

**Time Complexity: $`\Theta(n)`$**

The loop runs exactly $`n`$ times, each with one comparison and at most one AND, all constant time.

**Space Complexity: $`O(1)`$**

Only `res`, `i`, and `n` are used; the input is not modified and no extra array is allocated. The swap procedure of Lemma 3 exists only for the proof and is not executed by the code.

## 5. Learning Notes

### Challenge 1: The hub $`M`$ may already be in its correct place

It is tempting to think that in-place elements are irrelevant, and to look for swaps only among misplaced elements. But no pair of misplaced elements may have an AND exactly equal to $`M`$.

**Understanding Process**

Consider the test `fixed_zero_as_hub`: in `[0,1,2,5,4,6,3,7]` the misplaced values are $`\{5, 6, 3\}`$, with pairwise ANDs $`5 \mathbin{\&} 6 = 4`$, $`5 \mathbin{\&} 3 = 1`$, and $`6 \mathbin{\&} 3 = 2`$, none of which equals $`M = 0`$. What works is borrowing the in-place 0: its AND with anything is 0, so swap it out to act as an "empty slot", carry the misplaced elements home one by one, and 0 returns to its place at the end.

This also explains why in-place elements **do not affect the answer**: the upper bound comes only from "misplaced elements must move", and in-place elements can stay put, so they impose no constraint; sufficiency, meanwhile, lets them leave temporarily and come back.

**Transferable Insight:** When proving that a maximum is reachable, do not restrict the available resources to the part of the input the problem seems to be about. Intermediate states may temporarily break what is already correct, as long as it is restored at the end.

### Challenge 2: Why start at `-1`, and why can it detect "already sorted"?

**Understanding Process**

The identity of AND is all ones: ANDing all ones with any $`x`$ gives $`x`$. Go's `int` is signed two's complement, and `-1` is exactly all ones, so `res := -1` means "the AND of the empty set".

Detecting a sorted array relies on the value range: all values are non-negative, and the AND of non-negative numbers has sign bit 0, so it can never equal `-1`. Therefore `res == -1` holds exactly when no AND was performed, that is, when no element is misplaced. If values could be negative, this sentinel would no longer be reliable and a separate boolean flag would be needed.

**Transferable Insight:** When folding a collection with an operation (sum, product, AND, OR, minimum), start from that operation's identity. Before reusing the identity as an "empty set" sentinel, make sure valid data can never produce it.

## 6. Key Takeaways

This problem illustrates the following reusable concepts:

1. **Submask**: Here, $`a \mathbin{\&} b = k`$ means $`k`$ is a common submask of $`a`$ and $`b`$, so "every misplaced element must be swapped" becomes "$`k`$ is a common submask of all misplaced elements", and the largest common submask of a set of numbers is their AND. It applies to constraints of the form "the AND equals / contains some value". Limitation: being a submask implies being numerically no larger, but not the converse, so when comparing "largest" values first obtain the bound in the submask sense.
2. **Upper Bound plus Construction**: Here, Lemma 2 proves that no feasible $`k`$ exceeds $`M`$, and Lemma 3 constructs a swap sequence proving that $`M`$ is reachable; only together do they show $`M`$ is the maximum. It applies to "find the largest / smallest feasible parameter" problems. Limitation: proving only the bound misses the case where it is unreachable, and giving only a construction does not rule out larger values; both halves are required.
3. **Star Transposition**: Here, every swap goes through the hub $`M`$, following the rule "whatever position the hub is in, bring that position's element home", which sorts $`T`$ in at most $`2\lvert T \rvert`$ swaps. It applies to sorting problems in which only one special element can be swapped with the others, such as moving items through an empty slot. It requires that swapping the hub with every element of the set is allowed and that the positions occupied by the set are exactly its elements' target positions (guaranteed here by Lemma 1).
4. **Fold with Identity Element**: Here, `res` starts at `-1` and is ANDed with each value. It applies to any scan that combines a collection with an associative operation. When the identity doubles as an "empty set" sentinel, valid data must never produce it.
5. **Loop Invariant**: Here, "`res` is the AND of the misplaced values in the prefix" ties the code to the mathematical definition of $`M`$, so the rest of the correctness proof in 4.4 only needs to reason about $`M`$ itself. It applies to every accumulating scan.

### CLRS Connections

Section numbers follow CLRS, 3rd edition. Relation is either "Direct" or "Related idea"; if no section fits, write "No matching section" and keep only the concept name.

| Concept | CLRS Section | Relation | How It Appears Here |
|---|---|---|---|
| Loop invariant | §2.1 Insertion sort | Direct | 4.4 proves the code computes $`M`$ with the invariant "`res` is the AND of the misplaced values in the prefix" in the initialization, maintenance, and termination steps |
| Submask and bitwise AND | No matching section | Related idea | Lemma 2 derives $`k \subseteq M`$ from "misplaced elements must take part in a swap" |
| Sorting by star transpositions | No matching section | Related idea | Lemma 3 builds a swap sequence with $`M`$ as the hub and proves termination with $`\Phi`$ |

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

- [1202. Smallest String With Swaps](https://leetcode.cn/problems/smallest-string-with-swaps/) — Also asks which arrangements restricted swaps can reach: the swap relation forms a graph, and any arrangement within a connected component is reachable, which is the same reachability argument as "the hub can swap with every element of $`T`$, so $`T`$ can be rearranged arbitrarily" here.
- [2948. Make Lexicographically Smallest Array by Swapping Elements](https://leetcode.cn/problems/make-lexicographically-smallest-array-by-swapping-elements/) — The swap condition depends on element values (difference at most `limit`), so one first analyzes which values can be exchanged and then decides where each group can go, the same modeling as analyzing swaps through the bit relation of values here.
- [Codeforces 1828B. Permutation Swap](https://codeforces.com/problemset/problem/1828/B) — Also asks for the largest swap parameter that lets a permutation be sorted: only positions a multiple of $`k`$ apart may be swapped, and the answer is the GCD of all $`\lvert p_i - i \rvert`$, again with "misplaced elements must move" giving the bound and a construction proving it is reachable.

---

**One-Sentence Summary**

Every misplaced element must take part in a swap, so $`k`$ is a common submask of them and cannot exceed their AND $`M`$; and since $`M`$ can be swapped legally with every misplaced element, it serves as a hub that sorts the array, so the answer is the AND of the misplaced elements.
