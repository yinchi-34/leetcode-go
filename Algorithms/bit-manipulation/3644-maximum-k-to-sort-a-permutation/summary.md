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

**Four-Dimensional Rating — Bitwise AND over the Support**

- Evaluation Scope: $`1 \le n \le 10^5`$, `nums` is a permutation of $`0 \ldots n-1`$; time $`\Theta(n)`$, auxiliary space $`O(1)`$; Go implementation.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 3 | One must notice that every element of the support takes part in at least one swap, which limits $`k`$ to a common submask of the support, and then think of sorting through a hub element |
| Knowledge Prerequisites | 2 | Needs the submask property of bitwise AND and basic facts about permutations; no specialized algorithm |
| Correctness Reasoning Difficulty | 3 | Needs separate proofs of the upper bound (necessity) and reachability (sufficiency); sufficiency requires constructing a swap sequence whose hub may be a fixed point |
| Implementation Difficulty | 1 | A single pass; the only detail is starting the AND at `-1` (all ones) and telling apart the empty-support case |

Ratings follow [RATING.md](../../../docs/RATING.md). Each dimension is evaluated independently and not combined into a total score.

## 2. Problem Understanding

We are given a permutation `nums` of $`0 \ldots n-1`$. After choosing a non-negative integer $`k`$, we may swap two positions only when `nums[i] & nums[j] == k`. Return the largest $`k`$ that lets us sort the array in increasing order; if the array is already sorted, return 0.

**Example**

- Input: `nums = [0,3,2,1]`
- Output: `1` ($`3 \mathbin{\&} 1 = 1`$, and swapping 3 with 1 sorts the array)

**Key Requirements**

1. Input constraints: $`1 \le n \le 10^5`$, and `nums` contains each of $`0 \ldots n-1`$ exactly once.
2. The swap condition is that the AND **equals** $`k`$ exactly, not that it contains $`k`$. Any number of swaps is allowed, and fixed points may take part in swaps.
3. Boundary: the answer is defined as 0 for a sorted array; otherwise $`k = 0`$ is always feasible (see 4.4), so an answer always exists.

## 3. Initial Approach

### First Idea: This Is a Bit Manipulation Problem

It was clear from the statement that bit manipulation was needed, but the step from the statement (the swap condition, whether the array can be sorted) to the requested answer (a specific $`k`$) read awkwardly.

### Difficulties Encountered

The code was written as "skip the numbers already in place and apply `k &= nums[i]` to the rest", and the tests passed, but two things were unclear:

- `k` only records "the current value"; why does ANDing all the way through give the answer? Is this greedy?
- Does the number 0 need special handling? How should the test cases be classified?

### Thought Transformation

From:

> Focusing on the swap condition: which numbers can swap with which, whether 0 is in place, and which "kinds of disorder" the tests must cover.

To:

> The problem is really about finding the support of the permutation and taking the AND of the support. The AND can be computed while scanning, without storing the support first.

After this shift, 0 is no longer a special case: if 0 is out of place, it is in the support and takes part in the AND under the same rule. The tests simplify as well: the code has only two paths, "empty support" and "non-empty support", and one test case covers each (see 4.6).

## 4. Algorithm Specification

**Algorithm Name:** Bitwise AND over the Support

Indices in this section start at 0.

### 4.1 Input

View the input as a permutation $`\pi`$ with $`\pi(i) = \mathrm{nums}[i]`$, a bijection from $`\{0, \ldots, n-1\}`$ to itself.

- An $`i`$ with $`\pi(i) = i`$ is a fixed point.
- The support is $`\mathrm{supp}(\pi) = \{i : \pi(i) \ne i\}`$.
- The array is sorted exactly when $`\pi = \mathrm{id}`$, that is, when the support is empty.

### 4.2 Output

Call $`k`$ **feasible** if some finite sequence of swaps, each between two values whose AND is exactly $`k`$, ends with $`\pi = \mathrm{id}`$.

The output $`\mathrm{res}`$ satisfies:

- If $`\mathrm{supp}(\pi) = \varnothing`$, then $`\mathrm{res} = 0`$;
- Otherwise $`\mathrm{res}`$ is feasible, and every feasible $`k`$ satisfies $`k \le \mathrm{res}`$.

We will prove that it equals the AND of all elements of the support:

```math
M = \mathop{\&}_{v \in \mathrm{supp}(\pi)} v
```

### 4.3 Pseudocode

```text
MAX-K-TO-SORT(A)
    res ← -1                      ▷ all ones in binary, the identity of AND
    for i ← 0 to n - 1
        if A[i] ≠ i               ▷ i is in the support
            res ← res AND A[i]
    if res == -1                  ▷ the support is empty
        return 0
    return res
```

**State Definitions**

- `res`: the AND of the support elements in the scanned prefix; `-1` while none has been seen.
- `i`: the position being checked.

### 4.4 Correctness

**Proof Method: Loop Invariant + Upper Bound (Necessity) + Construction (Reachability)**

Notation: $`x \subseteq y`$ means every 1 bit of $`x`$ is also 1 in $`y`$, that is, $`x \mathbin{\&} y = x`$.

#### I. Key Properties Required for Correctness

1. **Submask property (bitwise fact)**: if $`a \mathbin{\&} b = k`$, then $`k \subseteq a`$ and $`k \subseteq b`$. If $`k \subseteq v`$ for every $`v`$ in a set, then $`k \subseteq \mathop{\&} v`$. $`x \subseteq y`$ implies $`x \le y`$.
2. **The support is closed under $`\pi`$ (Lemma 1)**: $`\{\pi(i) : i \in \mathrm{supp}(\pi)\} = \mathrm{supp}(\pi)`$. So "misplaced values" and "misplaced positions" are the same set, and the AND of `A[i]` in the code is $`M`$.
3. **Upper bound (Lemma 2)**: when the support is non-empty, every feasible $`k`$ satisfies $`k \subseteq M`$.
4. **Reachability (Lemma 3)**: when the support is non-empty, $`M`$ is feasible.

**Lemma 1.** The set $`F`$ of fixed points maps to itself under $`\pi`$. Since $`\pi`$ is a bijection, it maps the complement of $`F`$, which is the support, onto the complement of $`F`$.

**Lemma 2.** Take any $`v \in \mathrm{supp}(\pi)`$. Value $`v`$ is not at position $`v`$ now but must end there, and only swaps move elements, so $`v`$ takes part in at least one swap. If the partner is $`u`$, then $`v \mathbin{\&} u = k`$, and property 1 gives $`k \subseteq v`$. This holds for every $`v \in \mathrm{supp}(\pi)`$, so $`k \subseteq M`$ and hence $`k \le M`$.

**Lemma 3 (construction).** Since $`M \subseteq \min \mathrm{supp}(\pi) \le n-1`$, the value $`M`$ occurs in the array. Let the participating set be $`T = \mathrm{supp}(\pi) \cup \{M\}`$, and call $`M`$ the **hub**.

- If $`M \in \mathrm{supp}(\pi)`$, then $`T = \mathrm{supp}(\pi)`$;
- Otherwise $`M`$ is a fixed point sitting at position $`M`$.

By Lemma 1, the positions occupied by the values of $`T`$ are exactly $`T`$. For every $`x \in T`$ we have $`M \subseteq x`$, so $`M \mathbin{\&} x = M`$: **swapping the hub with any element of $`T`$ is allowed.** Using only such swaps, repeat:

- **Step A**: the hub is at position $`p \ne M`$. Since $`p \in T`$, value $`p`$ is in $`T`$ and is not the hub; swap the hub with value $`p`$, which puts $`p`$ home.
- **Step B**: the hub is at position $`M`$, but some $`x \in T`$ is not at position $`x`$. Swap the hub with $`x`$; the hub is now away from position $`M`$, so the next step is Step A.
- If neither applies, every element of $`T`$ is home, and the elements outside $`T`$ are fixed points that were never moved, so $`\pi = \mathrm{id}`$.

Termination: let $`\Phi`$ be the number of non-hub elements of $`T`$ that are not home. Step A lowers $`\Phi`$ by 1. Step B leaves $`\Phi`$ unchanged but is always followed by Step A (Step B moves the hub to $`x`$'s old position $`q \ne M`$, and Step A never moves an element that is already home). So $`\Phi`$ drops by at least 1 every two steps, and the process ends after at most $`2\lvert T \rvert`$ swaps.

#### II. Correctness Proof

First we show that the code computes $`M`$.

**Loop Invariant**

At the start of iteration $`i`$:

**Invariant I: AND of the support elements in the prefix**

`res` equals the AND of all values in $`\{\pi(j) : 0 \le j < i,\ \pi(j) \ne j\}`$; when that set is empty, `res` $`= -1`$.

**Initialization**

For $`i = 0`$ the set is empty and `res` $`= -1`$, so the invariant holds.

**Maintenance**

- If $`i`$ is a fixed point: the set is unchanged and so is `res`.
- If $`i \in \mathrm{supp}(\pi)`$: the new set gains $`\pi(i)`$. If the old set was empty, `res` becomes $`-1 \mathbin{\&} \pi(i) = \pi(i)`$ (`-1` is all ones in two's complement); otherwise, by associativity of AND, `res` becomes the old value ANDed with $`\pi(i)`$. The invariant holds in both cases.

**Termination**

The loop ends at $`i = n`$; by Lemma 1 the set is then exactly $`\mathrm{supp}(\pi)`$.

- Empty support: `res` $`= -1`$ and the function returns 0, as required.
- Non-empty support: `res` $`= M`$. All values are non-negative, so their AND has sign bit 0 and cannot equal $`-1`$; the support is not mistaken for empty, and $`M`$ is returned.

#### III. Conclusion

When $`\pi = \mathrm{id}`$ the result is 0. Otherwise the result is $`M`$: by Lemma 3, $`M`$ is feasible; by Lemma 2, no feasible $`k`$ exceeds $`M`$. Hence $`M`$ is the largest feasible value, and the algorithm is correct for every valid input.

### 4.5 Complexity Analysis

Let $`n`$ be the array length.

**Time Complexity: $`\Theta(n)`$**

The loop runs exactly $`n`$ times, each with one comparison and at most one AND, all constant time.

**Space Complexity: $`O(1)`$**

Only `res`, `i`, and `n` are used; the input is not modified and the support is not stored. The swap procedure of Lemma 3 exists only for the proof and is not executed by the code.

### 4.6 Test Design

The tests follow the algorithm's mechanism. The code has only two paths: with an empty support it performs no AND and returns 0 directly; with a non-empty support it ANDs the support elements one by one. The number 0 needs no special handling: when it is in the support, it takes part in the AND under the same rule. So each path gets one test case.

| Test | Input | Expected | Path Covered |
|---|---|---|---|
| `identity_permutation` | `[0,1,2,3]` | 0 | $`\pi = \mathrm{id}`$, empty support |
| `nonempty_support` | `[0,1,2,3,4,6,7,5]` | 4 | Support $`\{5, 6, 7\}`$, $`6 \mathbin{\&} 7 \mathbin{\&} 5 = 4`$ |

The second test is chosen so that common mistakes all produce a different result:

| If the Code Were Wrong | Result |
|---|---|
| Initial value written as `k := 0` | 0 |
| Fixed points also included in the AND | 0 |
| Only the first two support elements used | $`6 \mathbin{\&} 7 = 6`$ |

In addition, the answer 4 is a fixed point, not an element of the support, which shows that the hub can be a fixed point (Lemma 3).

**Additional check**: outside the repository, an exhaustive search over all permutations with $`n \le 7`$ (searching every state reachable by swaps for each $`k`$) agreed with $`M`$ in every case.

## 5. Learning Notes

### Challenge 1: Is ANDing all the way through a greedy algorithm?

`k` only records "the current value" and tightens each time a support element appears, which looks greedy.

**Understanding Process**

Look at each binary bit separately. A bit of the answer can stay 1 exactly when every element of the support has a 1 in that bit; if any element has a 0 there, the bit must be cleared. Bits do not affect one another, and keeping one bit never prevents keeping another, so keeping every 1 that can be kept gives the largest candidate. `k &= nums[i]` makes this decision for all bits at once.

There is no step of "choosing a locally best option among several"; each support element just adds a constraint that must hold, and AND merges the constraints one by one. So it is more accurate to call this "AND accumulation" than greedy. Also, the AND only gives an upper bound; Lemma 3 is still needed to prove that the bound is reachable.

**Transferable Insight:** When each bit of the target can be decided independently and there is no trade-off between bits, "make every bit 1 if possible" is optimal, and no greedy exchange argument is needed. Whether something is greedy depends on whether it chooses among several options.

### Challenge 2: Why do fixed points not limit the answer, yet may be borrowed?

**Understanding Process**

The upper bound comes only from "support elements must move". A fixed point can stay put, so it imposes no constraint.

When proving reachability, however, a fixed point may be useful. For example, in `[0,1,2,5,4,6,3,7]` the support is $`\{3, 5, 6\}`$, with pairwise ANDs $`5 \mathbin{\&} 6 = 4`$, $`5 \mathbin{\&} 3 = 1`$, and $`6 \mathbin{\&} 3 = 2`$, none of which equals $`M = 0`$, so no two support elements can be swapped directly. What works is borrowing the fixed point 0: its AND with anything is 0, so swap it out to act as an "empty slot", carry the support elements home one by one, and 0 returns to its place at the end.

**Transferable Insight:** When proving that a maximum is reachable, do not restrict the available resources to the part of the input the problem seems to be about. Intermediate states may temporarily break what is already correct, as long as it is restored at the end.

### Challenge 3: How should the test cases be designed?

At first many test cases were listed by "kind of disorder": whether 0 is misplaced, whether the common bit is low or high, how many numbers are misplaced, and so on. Later they were partitioned into equivalence classes by the position of 0 and the size of the support.

**Understanding Process**

Once the problem is understood as "find the support, then take its AND", these differences are only different inputs to the same rule, not different handling in the code. The code has only two real paths, so two test cases are enough. The key is choosing the second test well, so that common mistakes are exposed (see 4.6).

**Transferable Insight:** Tests should follow your own understanding of the mechanism, rather than adding one case for every surface difference. First state which paths the code has, then pick for each path an input that exposes common mistakes; when a safety net is needed, cross-check against a brute-force solution.

## 6. Key Takeaways

This problem illustrates the following reusable concepts:

1. **Submask**: Here, $`a \mathbin{\&} b = k`$ means $`k`$ is a common submask of $`a`$ and $`b`$, so "every support element must take part in a swap" becomes "$`k`$ is a common submask of the support elements", and the largest common submask of a set of numbers is their AND. It applies to constraints of the form "the AND equals / contains some value". Limitation: being a submask implies being numerically no larger, but not the converse, so when comparing "largest" values first obtain the bound in the submask sense.
2. **Upper Bound plus Construction**: Here, Lemma 2 proves that no feasible $`k`$ exceeds $`M`$, and Lemma 3 constructs a swap sequence proving that $`M`$ is reachable; only together do they show $`M`$ is the maximum. It applies to "find the largest / smallest feasible parameter" problems. Limitation: proving only the bound misses the case where it is unreachable, and giving only a construction does not rule out larger values; both halves are required.
3. **Star Transposition**: Here, every swap goes through the hub $`M`$, following the rule "whatever position the hub is in, bring that position's element home", which sorts $`T`$ in at most $`2\lvert T \rvert`$ swaps. It applies to sorting problems in which only one special element can be swapped with the others, such as moving items through an empty slot. It requires that swapping the hub with every element of the set is allowed and that the positions occupied by the set are exactly its elements' target positions (guaranteed here by Lemma 1).
4. **Fold with Identity Element**: Here, `res` starts at `-1` and is ANDed with each value, and `-1` also tells whether the support is empty. It applies to any scan that combines a collection with an associative operation. When the identity doubles as an "empty set" sentinel, valid data must never produce it; here all values are non-negative, so their AND can never be `-1`.
5. **Loop Invariant**: Here, "`res` is the AND of the support elements in the prefix" ties the code to the mathematical definition of $`M`$, so the rest of the proof in 4.4 only needs to reason about $`M`$ itself. It applies to every accumulating scan.

### CLRS Connections

Section numbers follow CLRS, 3rd edition. Relation is either "Direct" or "Related idea"; if no section fits, write "No matching section" and keep only the concept name.

| Concept | CLRS Section | Relation | How It Appears Here |
|---|---|---|---|
| Loop invariant | §2.1 Insertion sort | Direct | 4.4 proves the code computes $`M`$ with the invariant "`res` is the AND of the support elements in the prefix" in the initialization, maintenance, and termination steps |
| Submask and bitwise AND | No matching section | Related idea | Lemma 2 derives $`k \subseteq M`$ from "support elements must take part in a swap" |
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
- [Codeforces 1828B. Permutation Swap](https://codeforces.com/problemset/problem/1828/B) — Also asks for the largest swap parameter that lets a permutation be sorted: only positions a multiple of $`k`$ apart may be swapped, and the answer is the GCD of all $`\lvert p_i - i \rvert`$, again with "support elements must move" giving the bound and a construction proving it is reachable.

---

**One-Sentence Summary**

The problem is really about taking the AND of the permutation's support: every support element must take part in a swap, so $`k`$ cannot exceed their AND $`M`$; and since $`M`$ can be swapped legally with every support element, it serves as a hub that sorts the array, so the answer is $`M`$.
