# {{1021}}. {{Remove Outermost Parentheses}}

English | [简体中文](./summary.zh-CN.md)

## 1. Problem Information & Rating

- **Problem**: [1021. Remove Outermost Parentheses](https://leetcode.cn/problems/remove-outermost-parentheses/)
- **Difficulty**: {{Easy}}
- **Category**: Simulation（模拟）
- **Tags**: `Counting` `State Invariant`
- **Code**: [solution.go](./solution.go)
- **Tests**: [solution_test.go](./solution_test.go)

### 1.1 Difficulty Ratings

**External Ratings**

- **zerotrac Rating**: To be checked ([Source](https://github.com/zerotrac/leetcode_problem_rating))
- **Leedcode Official Difficulty**: {{Easy}}

**Four-Dimensional Rating - Counting**

**Evaluation Scope**: Valid parentheses strings; 

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 2/5 | Requires translating the identification of outermost parentheses into a nesting-depth calculation |
| Knowledge Prerequisites | 1/5 | Requires only loops, conditional statements, and a counter |
| Correctness Reasoning Difficulty | 2/5 | Requires case analysis by parenthesis type and a simple state invariant |
| Implementation Difficulty | 2/5 | Uses few branches, but the order of state updates and condition checks requires care |

## 2. Problem Understanding

### Problem Description

Given a valid parentheses string, remove the outermost pair of parentheses from each primitive component.

Key Requirements：

- The input is guaranteed to be a vaild parentheses string.
- Primitive: A nonempty valid parentheses string that cannot be split into the concateation of two nonempty valid parentheses strings.
- Only the outermost pair of each primitive is removed. All inner parentheses retain their original order.

## 3. Initial Approach 

### First Idea: Stack

When I first saw a parentheses problem, my inital idea was to use a stack.

A stack is suitable for handing matching parentheses and nested structures because it can store unmatched opening parentheses.

However, after re-analyzing the problem, I realized:

**This problem does not require identifying the matching partner of every parenthesis. It only requires determining whether each parenthesis belongs to the outermost layer.**

This individual elements stored in a stack provide more information than this problem needs.

The state can therefore be simplified to a single varible, `depth`, representing the current nesting depth.

### Thought Transformation

From:
> How can I record and match every parenthesis?

To:

> How can I determine whether the current parenthesis belongs to the outermost layer?

The first question focuses on matching individual elements. The second focuses on the current state.

This change reduces the auxiliary state from a stack to a single counter.

## 4. Algorithm Specification

### 4.1 Input

A valid parentheses string of length n:

$$
s = \langle s_1, s_2, \ldots, s_n \rangle
$$

**Constraints:**

- **Character constraint**: $s_i \in \{\texttt{'('}, \texttt{')'}\}$.
- **Length constraint**: $n \geq 2$, and $n$ is even.
- **Prefix validity**: The parentheses depth of every prefix satisfies $d(k) \geq 0$.
- **Global balance**: The final parentheses depth satisfies  $d(n) = 0$.

The parentheses depth is defined as:

$$
d(k) = \sum_{i=1}^{k}
\begin{cases}
+1, & s_i = \texttt{'('} \\
-1, & s_i = \texttt{')'}
\end{cases}
$$

Therefore, a valid parentheses string satisfies:

$$
\boxed{
\begin{aligned}
&d(k) \geq 0,\quad \forall k \in \{1,\ldots,n\} \\
&d(n) = 0
\end{aligned}
}
$$

### 4.2 Output

Output a parentheses string of length $m$:

$$
s' = \langle s'_1, s'_2, \ldots, s'_m \rangle
$$

where $0 \leq m < n$.

Suppose the input string can be uniquely decomposed into $k$ primitives:

$$
s = P_1 \Vert P_2 \Vert \cdots \Vert P_k
$$

**Constraints:**

$$
s' = \operatorname{inner}(P_1)
\Vert \operatorname{inner}(P_2)
\Vert \cdots
\Vert \operatorname{inner}(P_k)
$$

Here, $\Vert$ denotes string concatenation, and $\operatorname{inner}(P)$ denotes the string obtained by removing the first opening parenthesis and the last closing parenthesis from primitive $P$.

### 4.3 Pseudocode

```text
REMOVE-OUTER-PARENTHESES(s)
    res ← ""
    depth ← 0

    for each ch in s
        if ch == '('
            if depth > 0
                APPEND(res, ch)
            depth ← depth + 1
        else
            depth ← depth - 1
            if depth > 0
                APPEND(res, ch)

    return res
```

**state Definitions**:

At the beginning of each iteration:
- `depth`: The number of unmatched opening parentheses in the processed prefix.
- `res`: The inner parentheses retained from processed prefix, in their original order.

### 4.4 Correctness

### 4.4 Correctness

**Proof method: Loop Invariant**

**Premise**: The input $s$ is a valid parentheses string, satisfying prefix validity ($d(k) \geq 0$) and overall closure ($d(n) = 0$). $s$ decomposes uniquely into primitives $P_1 \Vert P_2 \Vert \cdots \Vert P_k$.

#### 1. Key Properties

For $1 \leq j \leq n$:

- **Property A (left boundary)**: $s_j$ is the first character of some primitive if and only if $s_j = \texttt{'('}$ and $d(j-1) = 0$.
- **Property B (right boundary)**: $s_j$ is the last character of some primitive if and only if $s_j = \texttt{')'}$ and $d(j) = 0$.

**Proof:**

- The positions where the depth is $0$ are exactly the empty prefix and the end of each primitive. A new primitive starts at the character right after such a position. If $s_j = \texttt{')'}$ there, then $d(j) = -1$, which violates prefix validity, so the first character must be `(`.
- A primitive cannot be split into two non-empty valid parentheses strings, so the depth inside it (before its last character) is always $\geq 1$, and only the last character brings the depth back to $0$. If $s_j = \texttt{'('}$ there, then $d(j-1) = -1$, which is again a contradiction, so the last character must be `)`.
- The two conditions require `(` and `)` respectively, so they are mutually exclusive: a character is never both a left and a right boundary.

#### 2. Loop Invariant

Suppose the first $i$ characters have been processed, $0 \leq i \leq n$. Before each iteration:

1. **Invariant I (depth correctness)**: `depth` $= d(i)$.
2. **Invariant II (result correctness)**: `res` equals the string obtained from $s_1 \cdots s_i$ by deleting the first and last character of every primitive, with the remaining characters concatenated in their original order.

#### 3. Proof of Correctness

**1. Initialization**

When $i = 0$:

- `depth` $= 0 = d(0)$.
- `res` $= ""$.

The empty prefix contains no characters, so Invariants I and II hold.

**2. Maintenance**

Assume Invariants I and II hold after the first $i$ characters have been processed. Let the current character be $s_{i+1}$.

- **If $s_{i+1} = \texttt{'('}$**:

  At the time of the check, `depth` $= d(i)$.

  By Property A:

  - If `depth` $= 0$, it is the outermost left parenthesis of a primitive and is not appended.
  - If `depth` $> 0$, it is an inner parenthesis and is appended to `res`.

  Then the update:

  $$
  depth \leftarrow depth + 1 = d(i+1)
  $$

- **If $s_{i+1} = \texttt{')'}$**:

  First execute:

  $$
  depth \leftarrow depth - 1 = d(i+1)
  $$

  By the validity of the input, `depth` $> 0$ before the update, so it remains non-negative afterwards.

  By Property B:

  - If `depth` $= 0$ after the update, it is the outermost right parenthesis of a primitive and is not appended.
  - If `depth` $> 0$ after the update, it is an inner parenthesis and is appended to `res`.

In both cases, `depth` is updated to $d(i+1)$ and `res` receives exactly the non-boundary characters, so Invariants I and II hold at $i+1$.

**3. Termination**

When $i = n$:

By Invariant I:

$$
depth = d(n) = 0
$$

By Invariant II:

$$
res = \operatorname{inner}(P_1)
\Vert \cdots
\Vert \operatorname{inner}(P_k)
$$

which is exactly the output required in Section 4.2.

**Therefore, the algorithm is correct.**

### 4.5 Complexity Analysis

Let the length of the input string be n.

**Time complexity: O(n)**

The algorithm scans each character once, and each iteration performs only a constant number of checks, state updates, and append operations.

Therefore, the total time complexity is O(n).

**Space complexity: O(n)**

- `depth` uses O(1) space.
- `res` stores up to O(n) characters in the worst case.

Therefore, the total space complexity is O(n).

If the output buffer is not counted, the extra auxiliary state takes O(1) space.


## 5. Learning Notes

### Difficulty 1: When is a stack not needed?

When first seeing a parentheses problem, it is natural to think of a stack.

But before choosing a data structure, first decide what information the problem actually needs to keep.

- If, when a parenthesis closes, you need to retrieve its matching parenthesis or the contents of its level, a stack is usually a good fit. For example, the stack solution of 0856 needs to retrieve the accumulated score of the current level at every right parenthesis.
- If you only need to know the current nesting level, a counter is enough. This problem only asks "is the current parenthesis at the outermost level?", and the answer depends only on whether `depth` is 0.

**Transferable thinking: do not reach for the data structure that matches a problem's typical features; first decide what information is actually needed.**

### Difficulty 2: Why do the check and the update happen in different orders?

This is the part of the problem that most needs to be understood.

The left and right parentheses actually follow **the same rule**: if the depth of the level a parenthesis belongs to is 0, it is outermost and must be removed. The only difference is *when* that depth is read:

| Character | Depth of its level | When it is read |
| --- | --- | --- |
| `(` | Depth before entering, $d(j-1)$ | Check first, then update |
| `)` | Depth after leaving, $d(j)$ | Update first, then check |

This corresponds one-to-one with Properties A and B in Section 4.4.

**Typical mistake**: checking before updating for `)` as well. Verify with `"()"`:

1. `(`: `depth = 0`, not appended; after the update, `depth = 1`.
2. `)`: at the check, `depth = 1 > 0`, so `)` is wrongly appended; then `depth = 0`.

The output is `")"`, while the correct answer is `""`.

To generalize further:

**First define precisely what the state variable means, then decide whether the condition needs the state before or after the update.**

### Difficulty 3: Why is there no need to split primitives explicitly?

The problem says to process "each primitive separately", yet the code has no segmentation logic at all.

The reason: the positions where `depth` returns to 0 are exactly the primitive boundaries. The segmentation information is already implicit in the counter, so there is no need to record it separately.

This is also the intuition behind Properties A and B in Section 4.4.

**Transferable thinking: first ask whether structural information can be read directly from existing state, then decide whether to maintain it explicitly.**

## 6. Key Takeaways

This problem yields three reusable methods:

1. **Information Reduction**: maintain only the minimum information needed to solve the problem, without storing every intermediate detail. Here `depth` replaces the stack, and primitive boundaries need not be recorded either; they are read directly from `depth`.
2. **State Invariant**: be clear about what a variable represents at a specific moment of every iteration. `depth` always equals the depth $d(i)$ of the processed prefix, which is the foundation of the proof in Section 4.4.
3. **Condition-Update Order**: this is a direct corollary of point 2. Once the meaning of the state variable is fixed, whether a condition reads the value before or after the update is determined by "the level of the object being judged", not by coding habit.

The third point in particular applies to other simulation problems involving counters, nested structures, or state changes.

## 7. Review Log

| Date | Solved independently | Change in understanding | To improve |
|---|---|---|---|
| 2026-10-09 | To be filled | Moved from a stack to depth counting; focused on understanding the update order | Practice stating state invariants further |
| 2026-10-10 | Done | Understood that left and right parentheses follow the same rule (remove if the depth of their level is 0) and differ only in when it is read; made clear that primitive boundaries are implicit in `depth` | Without looking at the notes, write out Properties A and B and their proofs independently |

### Checklist for the Next Review

- Can I implement it independently without looking at the code?
- Can I explain why no stack is needed?
- Can I explain why the check order differs for left and right parentheses? (It must be explained with "depth of the level", not by memorizing the rule.)
- Can I give the `"()"` counterexample for "checking before updating on a right parenthesis"?
- Can I prove the algorithm's correctness without relying on specific code? (At minimum, write out Properties A and B and the maintenance step.)
- Can I explain why primitives need not be split explicitly?
- Can I recognize the same pattern in other problems?

## 8. Related Problems

- [0856. Score of Parentheses](https://leetcode.cn/problems/score-of-parentheses/) — Also involves parentheses nesting depth, but the goal of the computation differs. This problem only needs to decide whether a parenthesis is outermost, while that one needs to retrieve the accumulated value of a level when a parenthesis closes, which makes it a good contrast for the choice between a stack and a counter.
- [1614. Maximum Nesting Depth of the Parentheses](https://leetcode.cn/problems/maximum-nesting-depth-of-the-parentheses/) — Maintains `depth` directly and takes the maximum; the simplest version of this problem's counting idea.
- [0020. Valid Parentheses](https://leetcode.cn/problems/valid-parentheses/) — Requires checking whether the parentheses match validly; a contrast showing that a stack is more appropriate when matching relationships are needed.

---

**One-Sentence Summary**

The core of this problem is not just replacing the stack with `depth`, but first defining clearly what `depth` means and then deciding the order of checking and updating accordingly: left and right parentheses follow the same rule and differ only in when the "depth of their level" is read.