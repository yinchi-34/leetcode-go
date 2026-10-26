# 856. Score of Parentheses

English | [简体中文](./summary.zh-CN.md)

Use the [shared glossary](../../../docs/GLOSSARY.md) and follow the repository's [GitHub math conventions](../../../docs/GITHUB_MATH.md). The definition, implementation, and common patterns of the stack itself are in the [stack overview](../README.md); this note covers only how this problem uses a stack.

## 1. Problem Information

- **Problem**: [856. Score of Parentheses](https://leetcode.cn/problems/score-of-parentheses/)
- **Official Difficulty**: Medium
- **Primary Category**: Stack
- **Secondary Category**: Math
- **Tags**: `Stack`, `Math`, `Nested Structure`, `Structural Induction`, `Contribution Technique`
- **Code**: [solution.go](./solution.go)
- **Tests**: [solution_test.go](./solution_test.go)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: 1562.72 ([source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode Official Difficulty: Medium

**Four-Dimensional Rating — Solution 1: Layered Score Stack**

- Evaluation scope: $`2 \le n \le 50`$, `s` is a balanced parentheses string; time $`\Theta(n)`$, extra space $`O(n)`$; Go implementation.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 2 | Requires translating the recursive definition into "one partial sum per open layer", folding a layer's result into the enclosing layer when it closes |
| Knowledge Prerequisites | 2 | Requires the last-in first-out property of a stack |
| Correctness Reasoning Difficulty | 3 | Requires the unique decomposition of a prefix into balanced blocks separated by unmatched opening parentheses, and an invariant on the stack contents built on it |
| Implementation Difficulty | 2 | Needs a sentinel 0 at the bottom for the outermost layer, and the order "pop first, then add $`\max(2v, 1)`$ to the new top" |

**Four-Dimensional Rating — Solution 2: Depth Contribution Sum**

- Evaluation scope: as above; time $`\Theta(n)`$, extra space $`O(1)`$; Go implementation.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 3 | Requires noticing that only adjacent `()` pairs create score, and using distributivity of multiplication over addition to rewrite the total as a sum of $`2^{d}`$ contributions |
| Knowledge Prerequisites | 1 | Requires only a loop, a counter, and a left shift |
| Correctness Reasoning Difficulty | 3 | Requires a structural induction on the grammar for the contribution formula, then a loop invariant showing the code computes that sum |
| Implementation Difficulty | 2 | Must add after decrementing `depth`, and must look at the previous character to tell whether the current `)` closes a `()` |

Ratings follow [RATING.md](../../../docs/RATING.md). Each dimension is evaluated independently; no overall score is calculated.

## 2. Problem Understanding

Given a balanced parentheses string `s`, compute its score by three rules:

1. `()` scores 1.
2. `AB` scores $`A + B`$, where `A` and `B` are balanced strings.
3. `(A)` scores $`2A`$, where `A` is a balanced string.

**Example**

- Input: `s = "(()(()))"`
- Output: `6` (the inner `()(())` scores $`1 + 2 = 3`$; one more layer gives $`2 \times 3 = 6`$)

**Key Requirements**

1. Input constraints: $`2 \le n \le 50`$, `s` contains only `(` and `)`, and is balanced.
2. Rule 2 may split a string in more than one way (e.g. `()()()`), but addition is associative, so the score does not depend on the split.
3. Boundary: the deepest input is 25 nested layers with score $`2^{24}`$; in general the score is at most $`2^{n/2 - 1}`$ (by induction on the grammar), so Go's `int` cannot overflow.

## 3. Initial Approach

This section needs a personal solving record. None is available, so it is omitted.

## 4. Algorithm Specification

### 4.0 Shared Definitions

Let $`s = s_0 s_1 \cdots s_{n-1}`$; all pseudocode uses **0-based** indexing. Write $`\mathrm{score}(A)`$ for the score of a balanced string $`A`$, with the empty string $`\varepsilon`$ scoring 0.

**Lemma 0 (a nonempty balanced string scores at least 1)**: By induction on the grammar. `()` scores 1; `AB` scores $`\mathrm{score}(A) + \mathrm{score}(B) \ge 2`$; `(A)` scores $`2\,\mathrm{score}(A) \ge 2`$.

Hence, for any balanced $`B`$ (possibly empty), $`\mathrm{score}((B)) = \max(2\,\mathrm{score}(B), 1)`$: if $`B`$ is empty this is rule 1; otherwise Lemma 0 gives $`2\,\mathrm{score}(B) \ge 2 > 1`$, and this is rule 3.

**Prefix decomposition**: Let $`P`$ be a prefix of $`s`$ and $`d`$ the number of unmatched opening parentheses in $`P`$ (the current depth). Using those $`d`$ unmatched parentheses as separators, $`P`$ is uniquely written as

```math
P = B_0 \, ( \, B_1 \, ( \, B_2 \cdots ( \, B_d
```

where each $`B_k`$ is a balanced string (possibly empty). $`B_k`$ is the already-closed part of layer $`k`$.

### Solution 1: Layered Score Stack

#### 4.1 Input

A balanced parentheses string $`s`$ of length $`n`$.

#### 4.2 Output

$`\mathrm{score}(s)`$.

#### 4.3 Pseudocode

```text
SCORE-STACK(s)
    st = [0]                       // st[0] is the outermost layer
    for i = 0 to n - 1
        if s[i] == '('
            PUSH(st, 0)            // open a new layer
        else
            v = POP(st)            // the layer closes, v = score(B_d)
            st[top] = st[top] + max(2v, 1)
    return st[0]
```

**State Definitions**

- `st`: the stack. After processing prefix $`P`$, `st[k]` is $`\mathrm{score}(B_k)`$ for $`0 \le k \le d`$.
- `v`: the score of the closed part of the layer that just closed.

#### 4.4 Correctness

**Proof method: loop invariant**

##### 1. Key Conditions

1. **Given**: $`s`$ is balanced, so $`d \ge 1`$ whenever a `)` is read, and the stack is nonempty after POP; it never underflows.
2. **Proved lemma**: Lemma 0 and its corollary $`\mathrm{score}((B)) = \max(2\,\mathrm{score}(B), 1)`$.
3. **Rule 2**: concatenation adds scores.

##### 2. Proof

**Loop invariant**: At the start of iteration $`i`$, let the processed prefix $`P = s_0 \cdots s_{i-1}`$ decompose as $`B_0 ( B_1 \cdots ( B_d`$. Then the stack has length $`d + 1`$ and `st[k]` $`= \mathrm{score}(B_k)`$ for $`0 \le k \le d`$.

**Initialization**: At $`i = 0`$, $`P = \varepsilon`$, $`d = 0`$, $`B_0 = \varepsilon`$, and `st = [0]`. The invariant holds.

**Maintenance**:

- $`s_i =`$ `(`: the new prefix is $`B_0 ( \cdots ( B_d ( \varepsilon`$, depth becomes $`d + 1`$, and the new innermost closed part is empty. The algorithm pushes 0 and leaves the first $`d + 1`$ entries unchanged, so the invariant holds.
- $`s_i =`$ `)`: by condition 1, $`d \ge 1`$. This `)` matches the $`d`$-th unmatched parenthesis, the new prefix is $`B_0 ( \cdots ( B_{d-1} ( B_d )`$, depth becomes $`d - 1`$, and the closed part of layer $`d - 1`$ becomes $`B_{d-1} ( B_d )`$. By rule 2 and the corollary, its score is $`\mathrm{score}(B_{d-1}) + \max(2\,\mathrm{score}(B_d), 1)`$. The algorithm pops $`v = \mathrm{score}(B_d)`$ and adds $`\max(2v, 1)`$ to the new top `st[d-1]`, which gives exactly this value; the other entries are unchanged. The invariant holds.

**Termination**: The loop runs exactly $`n`$ times. At the end $`P = s`$; since $`s`$ is balanced, $`d = 0`$ and the decomposition is just $`B_0 = s`$, so `st[0]` $`= \mathrm{score}(s)`$.

##### 3. Conclusion

For every balanced parentheses string the stack never underflows and the return value equals $`\mathrm{score}(s)`$.

#### 4.5 Complexity Analysis

Let $`n = |s|`$ and let $`D`$ be the maximum nesting depth of $`s`$, $`1 \le D \le n/2`$.

**Time: $`\Theta(n)`$**. Each character does one PUSH, or one POP plus one addition. By section 2.2 of the [overview](../README.md#22-go-idiom), `append` is amortized $`O(1)`$, so the total is $`\Theta(n)`$.

**Space: $`O(D)`$, worst case $`O(n)`$**. The stack length is always $`d + 1 \le D + 1`$; it reaches $`n/2 + 1`$ when `s` is $`n/2`$ fully nested layers.

### Solution 2: Depth Contribution Sum

#### 4.1 Input and Output

Same as Solution 1.

#### 4.2 Key Lemma

Call $`s_j s_{j+1} =`$ `()` an **atomic pair**, and let $`d_j`$ be the number of unmatched opening parentheses before position $`j`$, i.e. the number of layers enclosing the pair.

**Lemma 1 (contribution formula)**: For every nonempty balanced string $`A`$,

```math
\mathrm{score}(A) = \sum_{\text{atomic pairs } j} 2^{d_j}
```

**Proof** (structural induction on the grammar):

- `()`: the only atomic pair is at depth 0, and the sum is $`2^0 = 1`$.
- `AB`: $`A`$ is nonempty and balanced, so it ends with `)`; $`B`$ begins with `(`. The junction is `)(`, which creates no new atomic pair. $`A`$ returns to depth 0, so every atomic pair of $`B`$ keeps its depth. The sum is the sum of the two parts, matching rule 2.
- `(A)`: the outer `(` is followed by the first character of $`A`$, which is `(`, and the last character of $`A`$, which is `)`, is followed by the outer `)`; neither forms an atomic pair. Every atomic pair of $`A`$ goes one level deeper, so its contribution doubles. By distributivity the sum doubles, matching rule 3.

Intuitively, this distributes the factor 2 in $`2A`$ over every term of $`A`$, all the way down to the atomic pairs.

#### 4.3 Pseudocode

```text
SCORE-DEPTH(s)
    res = 0
    depth = 0
    for i = 0 to n - 1
        if s[i] == '('
            depth = depth + 1
        else
            depth = depth - 1
            if s[i - 1] == '('     // s[i-1] s[i] is an atomic pair
                res = res + 2^depth
    return res
```

**State Definitions**

- `depth`: the number of unmatched opening parentheses in the processed prefix.
- `res`: the total contribution of the atomic pairs that lie entirely in the processed prefix.

When a `)` is read, $`i \ge 1`$ (a balanced string starts with `(`), so `s[i-1]` is valid.

#### 4.4 Correctness

**Proof method: Lemma 1 plus a loop invariant**

**Loop invariant**: At the start of iteration $`i`$, `depth` equals the number of unmatched opening parentheses in $`s_0 \cdots s_{i-1}`$, and `res` equals the total contribution of the atomic pairs lying entirely in that prefix.

**Initialization**: For the empty prefix both are 0. The invariant holds.

**Maintenance**:

- $`s_i =`$ `(`: depth increases by 1; no atomic pair ends at $`i`$, so `res` is unchanged. The invariant holds.
- $`s_i =`$ `)` and $`s_{i-1} =`$ `)`: depth decreases by 1; there is no new atomic pair. The invariant holds.
- $`s_i =`$ `)` and $`s_{i-1} =`$ `(`: $`(i-1, i)`$ is a new atomic pair. Before reading $`s_{i-1}`$ the depth was $`d_{i-1}`$; after `(` it was $`d_{i-1} + 1`$, and after the decrement it is $`d_{i-1}`$ again. So the added $`2^{\mathrm{depth}}`$ is exactly $`2^{d_{i-1}}`$. The invariant holds.

**Termination**: After $`n`$ iterations, `res` is the total contribution of all atomic pairs of $`s`$, which equals $`\mathrm{score}(s)`$ by Lemma 1.

#### 4.5 Complexity Analysis

**Time: $`\Theta(n)`$**, $`O(1)`$ per character.

**Space: $`O(1)`$**, only the two integers `res` and `depth`.

### How the Two Solutions Relate

Solution 2 is **not** Solution 1's stack compressed into a counter. The stack in Solution 1 holds a different score for each layer, so its contents are not determined by its length alone. Solution 2 computes the answer differently: Lemma 1 rewrites it in a form that depends only on depths, so only the depth needs to be tracked. See the two cases in section 4 of the [overview](../README.md#4-when-a-stack-is-not-needed).

## 5. Learning Notes

### Challenge 1: Why a closing layer contributes $`\max(2v, 1)`$

Rules 1 and 3 look like two separate rules, yet the stack implementation merges them into one expression. The key is Lemma 0: a nonempty balanced string scores at least 1, so $`v = 0`$ exactly when the layer is empty, giving 1; otherwise the layer gives $`2v`$.

**Transferable insight:** Defining the value of the empty case as the identity element of the operation (here 0 for addition) often merges the "empty" and "nonempty" branches into one formula.

### Challenge 2: Why the depth solution only looks at atomic pairs

Only rule 1 creates new score; rules 2 and 3 just add or double existing scores. Pushing every doubling inward by distributivity gives each atomic pair the coefficient $`2^{\text{depth}}`$.

**Transferable insight:** When a result is built recursively from "base units plus linear combination rules", it can be split into one contribution per base unit, whose coefficient is the product of all multipliers on the path from the root to that unit.

## 6. Key Takeaways

1. **Layered Partial Sums on a Stack**
   - Role here: each open layer has an accumulator on the stack; when the layer closes, its result is transformed and folded into the enclosing layer.
   - Applies to: evaluating recursively defined nested structures, such as parenthesized expressions or weighted sums of nested lists.
   - Conditions: the outer result depends only on an aggregate of the inner layer, not on its internal structure.
2. **Structural Induction**
   - Role here: proves Lemma 0 and Lemma 1 along the three grammar rules.
   - Applies to: objects generated by a grammar or recursive rules, when the property to prove is preserved by every rule.
   - Conditions: every object is generated by finitely many rule applications; when the concatenation rule is ambiguous, the conclusion must not depend on how the string is split.
3. **Contribution Technique**
   - Role here: rewrites the total score as the sum of $`2^{d}`$ over atomic pairs, so only the depth must be tracked.
   - Applies to: answers that can be written as a sum of contributions of base units, each computable locally.
   - Conditions: the combination rules must be linear in the contributions (here, distributivity of multiplication over addition); rules with nonlinear operations such as max or min generally break it.

### CLRS Connections

Section numbers follow CLRS, 3rd edition.

| Concept | CLRS Section | Relation | How It Appears Here |
|---|---|---|---|
| Stack | §10.1 Stacks and queues | Direct | Solution 1 opens a layer with PUSH and closes it with POP; a balanced input guarantees the stack is nonempty before POP, so it never underflows |
| Loop invariant | §2.1 Insertion sort | Direct | Solution 1 uses "`st[k]` $`= \mathrm{score}(B_k)`$" and Solution 2 uses "`res` is the contribution of the atomic pairs in the prefix", each proved by initialization, maintenance, and termination |
| Structural induction | No matching section | — | Proves Lemma 0 and Lemma 1 along the three grammar rules |
| Contribution technique | No matching section | — | Solution 2 splits the total into contributions $`2^{d}`$ of the atomic pairs |

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

- [1021. Remove Outermost Parentheses](../../../Algorithms/simulation/1021-remove-outermost-parentheses/summary.md) — also scans a parentheses string by depth; its stack would hold only `(`, so it collapses directly into a counter, whereas this problem's depth solution needs a separate contribution formula.
- [1541. Minimum Insertions to Balance a Parentheses String](../../../Algorithms/greedy-algorithms/1541-minimum-insertions-to-balance-a-parentheses-string/summary.md) — also replaces a stack with counting, but counts unmet demand for closing parentheses.
- [20. Valid Parentheses](https://leetcode.cn/problems/valid-parentheses/) — has several bracket types, so stack elements differ and cannot collapse into a counter; a useful contrast.

---

**One-Sentence Summary**

Evaluate recursively defined nested structures with a stack of per-layer partial sums; when the combination rules are linear, distributivity splits the answer into contributions of base units, so tracking the depth replaces the stack.
