# 1541. Minimum Insertions to Balance a Parentheses String

English | [简体中文](./summary.zh-CN.md)

## 1. Problem Information

- **Problem**: [1541. Minimum Insertions to Balance a Parentheses String](https://leetcode.com/problems/minimum-insertions-to-balance-a-parentheses-string/)
- **Official difficulty**: Medium
- **Primary category**: Greedy
- **Secondary category**: Simulation
- **Technique tags**: `greedy`, `state simulation`, `counting`, `state invariant`, `locally optimal repair`
- **Code**: [solution.go](./solution.go) (to be added)
- **Tests**: [solution_test.go](./solution_test.go) (to be added)

### Difficulty Ratings

**External ratings**

- zerotrac rating: 1759.02 (to be verified against the [data source](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt))
- LeetCode official difficulty: Medium

**Four-dimension rating — right-parenthesis demand counting solution**

Scope: the input constraints of this problem, using greedy + state simulation, targeting O(n) time and O(1) extra space, with the Go implementation as the subject of evaluation.

| Dimension | Score | Rationale |
|---|---|---|
| Modeling difficulty | 3/5 | Parenthesis matching must be reformulated as right-parenthesis demand, and parity is used to express an unfinished consecutive match |
| Prerequisite knowledge | 2/5 | Requires understanding counting states, local greedy repair, and basic state transitions |
| Correctness argument | 3/5 | Beyond proving the state invariant, one must also argue that each local insertion is necessary and that the result is globally optimal |
| Implementation difficulty | 3/5 | `res` and `right` must be updated in coordination; the two repair branches and the update order are easy to get wrong |

The four dimensions are rated independently and are not combined into a total score.

## 2. Problem Understanding

Given a string `s` consisting only of `(` and `)`, we may insert left or right parentheses at any position. Find the minimum number of insertions needed to make it a balanced parentheses string.

The difference from ordinary parenthesis matching is:

**Every left parenthesis `(` must be matched by two consecutive right parentheses `))`.**

For example:

- `())`: valid, no insertion needed.
- `(()))`: one right parenthesis is missing.
- `(`: needs two right parentheses.
- `)`: needs one left parenthesis and one right parenthesis.

Key constraints:

1. Each `(` must correspond to one group of consecutive `))`.
2. A left parenthesis must appear before the two right parentheses that match it.
3. Characters may be inserted, but the original characters may not be deleted or reordered.
4. The number of insertions must be minimized.
5. The input string is not guaranteed to be balanced already.

### Core Issue

This problem is not simply a comparison of the counts of left and right parentheses.

For example:

`()(` has two left parentheses and one right parenthesis, but the total counts alone cannot decide where to insert.

The first `)` has already started matching the preceding `(`, so the second `)` must be supplied before the following `(` can be handled.

Therefore:

**In addition to the matching counts, we must also track the matching state of consecutive right parentheses.**

## 3. Initial Approach

### 3.1 First Idea: Count Left and Right Parentheses

The most direct idea is:

- `left`: the number of unmatched left parentheses.
- `right`: the number of unmatched right parentheses.
- Finally, supply parentheses according to the difference between the two.

I also considered using values like `0.5` to represent a half-completed right-parenthesis match.

However, this approach has a problem:

**Recording only the counts of parentheses cannot fully express the constraint that two `)` must be consecutive.**

For example, `()(`:

After processing the first `)`, one more right parenthesis is still needed, but it must follow immediately after the `)` that has already appeared, and it cannot jump over the later `(`.

If we only supply it uniformly at the end of the string, the consecutive-match relation may be broken.

### 3.2 Shift in Thinking: From Parenthesis Counts to Matching Demand

After re-analysis, I found that:

Each `(` can be directly converted into two `)` that will be needed in the future.

So instead of maintaining both left and right parenthesis counts, we maintain:

- `right`: the right-parenthesis demand that still has to be satisfied.
- `res`: the number of parentheses that have already been determined to be necessary insertions.

Encountering `(` means two new `)` are demanded in the future.

Encountering `)` means one `)` demand has been satisfied.

When a situation violates the matching constraint, perform the minimal repair immediately.

### 3.3 Further Thought: Why Do We Need to Check Parity?

`right` is not just a count; it also implicitly encodes matching progress.

- `right` is even: every unfinished left parenthesis still needs a complete group of `))`.
- `right` is odd: the most recent unfinished left parenthesis has already matched its first `)` and is missing the second `)`; the other unfinished left parentheses each still need two `)`.

Therefore:

`right % 2 == 1`

means there is currently a group of consecutive right parentheses that has not been completed.

If a new `(` appears at this point, we must first insert one `)` to complete the previous group.

### Summary of the Shift in Thinking

From:

> How many left and right parentheses are currently unmatched?

To:

> Given the prefix already processed, which characters are needed later? Which demands must be satisfied right now?

The former focuses on count differences; the latter focuses on state and constraints.

**Core insight: when a matching rule involves order or consecutiveness, the counter must express enough matching progress, and cannot merely record how many elements there are.**

## 4. Algorithm Specification

**Algorithm name:** Greedy + Simulation

### 4.1 Input

The input is a parenthesis string of length n:

$$
s = \langle s_1, s_2, \ldots, s_n \rangle
$$

**Preconditions:**

- **Character constraint**: $`s_i \in \{\texttt{'('}, \texttt{')'}\}`$.
- **Length constraint**: $`1 \leq n \leq 10^5`$.
- **Validity constraint**: the input is not required to be balanced beforehand.


### 4.2 Output

Output a non-negative integer `k`, the minimum number of insertions needed to make the original string balanced.

Let $`\mathcal{B}`$ denote the set of all balanced strings satisfying this problem's matching rule, and let $`s \preceq t`$ mean that $`s`$ is a subsequence of $`t`$.

The objective is:

$$
k = \min_{\substack{t \in \mathcal{B} \\ s \preceq t}}
\left(|t| - |s|\right)
$$

The subsequence relation expresses that we may only insert characters; we may not delete or rearrange the original string.

### 4.3 Pseudocode

```text
MIN-INSERTIONS(s)
    res ← 0
    right ← 0

    for each ch in s
        if ch == '('
            if right mod 2 == 1
                res ← res + 1
                right ← right - 1

            right ← right + 2

        else
            right ← right - 1

            if right < 0
                res ← res + 1
                right ← right + 2

    return res + right
```

**State definitions**

- `res`: the number of parentheses already determined to be necessary insertions during the scan; they may be `(` or `)`.
- `right`: the number of right parentheses still to be supplied in order to complete the matching of the prefix constructed so far.
- `ch`: the original input character currently being processed.

**Greedy strategy**

Perform the necessary minimal repair only when the current constraint is about to be violated; leave all other unsatisfied demands to later original characters as far as possible.

### 4.4 State Transitions

#### Case 1: The current character is `(`

A new left parenthesis requires two consecutive right parentheses to appear in the future.

In the normal case:

$$
right \leftarrow right + 2
$$

But if, before entering this branch:

`right % 2 == 1`

it means a group `))` has so far completed only its first `)`.

In this case we must first supply the second `)`:

$$
\begin{aligned}
res &\leftarrow res + 1 \\
right &\leftarrow right - 1
\end{aligned}
$$

Only then can the current `(` be processed:

$$
right \leftarrow right + 2
$$

**Why must the repair come first?**

Because the two earlier right parentheses must appear consecutively.

If we do not complete them, the current `(` would be inserted between the two `)` of the old matching group, which breaks validity.

#### Case 2: The current character is `)`

The current `)` satisfies one right-parenthesis demand:

$$
right \leftarrow right - 1
$$

If:

`right >= 0`

the `)` can satisfy an existing demand, and no insertion is needed.

If:

`right < 0`

there is no left parenthesis waiting for the current `)`.

So we need to insert a `(`:

$$
res \leftarrow res + 1
$$

The newly inserted `(` creates two right-parenthesis demands, and the current `)` has already satisfied one of them.

Since `right -= 1` was already executed, we only need:

$$
right \leftarrow right + 2
$$

For example, starting from `right = 0` and encountering `)`:

- `right = -1`.
- Insert `(`, `res += 1`.
- `right += 2`, giving `right = 1` in the end.

At this point only one `)` is missing, not two.

#### Case 3: End of the scan

After all input characters have been processed, no later original character can satisfy the remaining demand.

Therefore, all `right` right parentheses must be inserted at the end of the string.

Final result:

$$
\boxed{ans = res + right}
$$

### 4.5 Correctness

**Proof method: loop invariant + necessity of the greedy choices**

#### Part 1: Key Properties

**Property A: the right-parenthesis demand is never negative**

After the necessary repair in each iteration:

$$
right \geq 0
$$

Reasons:

- On `(`, two new demands are added.
- On `)`, one demand is removed.
- If the demand becomes negative, insert one `(` immediately to restore non-negativity.

**Property B: parity expresses consecutive-match progress**

After each iteration:

- When `right` is even, there is no group still waiting after having completed only its first `)`.
- When `right` is odd, exactly the most recent unfinished left parenthesis has completed only its first `)`.

Inductive reasoning:

1. Each left parenthesis contributes exactly 2 demands, so `right` equals "the sum of the remaining demands of all unfinished left parentheses".
2. Each `)` always consumes the demand of the most recent unfinished left parenthesis first, so apart from that one, the remaining demand of every other unfinished left parenthesis is the full 2.
3. Hence when `right` is even, the remaining demand of the most recent left parenthesis is either 2 or already fully completed; when `right` is odd, its remaining demand is exactly 1, i.e. it is half-completed.

Since the algorithm repairs the odd state before processing a new `(`, it never allows a new left parenthesis to enter the middle of an unfinished `))`.

**Property C: there is no need to fill all demands early**

As long as no order or consecutiveness constraint is violated, the unsatisfied `right` demand may be kept.

Future original `)` can continue to satisfy that demand, which avoids unnecessary early insertions.

#### Part 2: Loop Invariants

Suppose the algorithm has processed the first $`i`$ characters of the original string.

At the end of each iteration:

**Invariant I (state validity)**

`right >= 0`, and after `res` insertions, the processed input prefix can become a valid balanced string by appending exactly `right` right parentheses at the end; at the same time, `right` exactly represents the remaining right-parenthesis demand of that constructed prefix, and its parity conforms to Property B.

**Invariant II (deferred repair)**

The algorithm increases `res` only when the matching rule requires an immediate insertion; demands that can be satisfied by later original right parentheses are never inserted early.

#### Part 3: Proof of Correctness

**1. Initialization**

At the start:

$$
res = 0,\qquad right = 0
$$

The prefix processed is empty:

- No parentheses need to be inserted.
- There is no unfinished match.
- No constraint is violated.

Therefore, both invariants hold.

**2. Maintenance**

Assume the invariants hold after processing the first $`i`$ characters.

Consider the next character $`s_{i+1}`$.

**Case A: the current character is `(`**

If `right` is even:

- There is no half-completed `))`.
- The current `(` may legally enter a new nesting level.
- It suffices to add two future right-parenthesis demands.

If `right` is odd:

- An earlier `)` is waiting for a consecutive second `)`.
- The current `(` cannot appear in the middle of this `))` group.
- Therefore at least one `)` must be inserted so that the old matching group is complete.

The algorithm inserts exactly one `)`, which satisfies the necessary condition without any extra insertion.

Then two new demands are created, and the invariants continue to hold.

**Case B: the current character is `)`**

Execute `right -= 1`.

If the result is non-negative:

- The current right parenthesis can satisfy one existing demand.
- No extra character needs to be inserted.

If the result is negative:

- There is no left parenthesis that can match the current `)`.
- Without deleting or reordering the original characters, at least one `(` must be supplied.

The algorithm inserts exactly one `(`, and the current `)` satisfies one of the new demands.

Therefore this repair is both necessary and minimal in count.

In both cases, the correct remaining demand is maintained after the repair.

**3. Termination**

When all $`n`$ characters have been processed:

- `res` records the insertions that were necessary during the scan.
- `right` records all right-parenthesis demands that are still unsatisfied.

Since no original characters remain, every remaining demand must be satisfied by inserted `)`.

Appending `right` right parentheses completes all unclosed matches.

Therefore, the algorithm returns:

$$
ans = res + right
$$

and a valid balanced string can be constructed.

#### Part 4: Why Is This Globally Optimal?

Proving that the constructed result is valid is not enough to prove that the number of insertions is minimal. Here we give a rigorous lower-bound argument: for each prefix, characterize exactly "the minimum number of insertions needed to reach a given demand state", and prove that the algorithm's result equals the lower bound.

**1. Automaton Characterization of Balanced Strings**

A string is balanced if and only if, when scanned according to the following rules, it never gets "stuck" and ends with `need = 0` (`need` being the remaining right-parenthesis demand):

- On `(`: require `need` to be even (otherwise it would be inserted into the middle of an unfinished `))`), then `need += 2`.
- On `)`: require `need ≥ 1`, then `need -= 1`.

Hence, for a prefix $`s_{1..i}`$ of the original string `s`, every valid completion corresponds to a path in this automaton. Inserting a character means taking one extra step along the path:

- Insert `)`: `need → need - 1` (requires `need ≥ 1`), cost 1.
- Insert `(`: `need → need + 2` (requires `need` to be even), cost 1.

Reading an original character costs nothing. Let $`g_i(r)`$ be the minimum number of insertions required to reach `need = r` after processing the first $`i`$ original characters, with further insertions allowed. Then the optimal answer is:

$$
\text{OPT} = \min_{r \ge 0}\bigl(g_n(r) + r\bigr)
$$

Here the `+ r` means that `r` right parentheses must be appended at the end.

**2. Lemma: A Lower-Bound Function for $`g_i`$**

Let $`(res, right)`$ be the algorithm's state after processing the first $`i`$ characters. Define the function:

$$
G(r) =
\begin{cases}
res + (right - r), & r \le right \\[2pt]
res + \dfrac{r - right}{2}, & r > right,\ r \equiv right \pmod 2 \\[6pt]
res + \dfrac{r - right + 3}{2}, & r > right,\ r \not\equiv right \pmod 2
\end{cases}
$$

**Lemma**: for all $`r \ge 0`$, $`g_i(r) \ge G_i(r)`$.

Intuition: reaching a demand below `right` can only be done by inserting `)` one at a time; reaching a higher demand requires inserting `(`, which adds 2 each time; and when the parities differ, one extra insertion is needed to adjust the parity.

**3. Proof of the Lemma (by induction on $`i`$)**

First, a criterion. Let $`h`$ be the cost function after reading the next original character but before making any new insertion. The new cost function $`g_{i+1}`$ is the closure of $`h`$ under insertion steps. Therefore, as long as a function $`G'`$ satisfies both:

- **(a) Closure**: $`G'(r-1) \le G'(r) + 1`$ (for $`r \ge 1`$), and $`G'(r+2) \le G'(r) + 1`$ (for even $`r`$);
- **(b) Not exceeding $`h`$**: $`G' \le h`$ wherever $`h`$ is defined,

we have $`g_{i+1} \ge G'`$, because the cost of any path can only grow along the insertion steps by at least that much.

**Closure (holds for any $`(res, right)`$)**: check the three branches of $`G`$ piece by piece.

- $`r \le right`$: $`G(r-1) = G(r) + 1`$.
- $`r + 2 \le right`$: $`G(r+2) = G(r) - 2`$.
- $`r = right - 1`$ (so $`r`$ is even): $`G(r) = res + 1`$, $`G(r+2) = res + 2`$.
- $`r = right`$ (even): $`G(r+2) = G(r) + 1`$.
- $`r > right`$: $`r \to r + 2`$ preserves parity, and in both branches the value increases by exactly 1; $`r \to r - 1`$ changes parity, and the value increases by at most 1.

**Initialization**: $`i = 0`$, $`(res, right) = (0, 0)`$. $`g_0`$ is the shortest path using insertions starting from `need = 0`, while $`G_0(0) = 0`$ and $`G_0`$ is closed, so $`g_0 \ge G_0`$.

**Maintenance**: assume $`g_i \ge G_i`$ and consider the next character.

*Case 1: the next character is `(`.* Only states with even `r` can read it, giving $`h(r+2) = g_i(r) \ge G_i(r)`$. We need to show $`G_{i+1}(r+2) \le G_i(r)`$ (for even $`r`$).

- `right` is even: the algorithm sets $`res' = res,\ right' = right + 2`$. All three branches depend only on the difference $`r - right`$ and on parity, so the whole function is shifted by 2, hence $`G_{i+1}(r+2) = G_i(r)`$.
- `right` is odd: the algorithm sets $`res' = res + 1,\ right' = right + 1`$.
  - If $`r < right`$: both sides equal $`res + right - r`$.
  - If $`r > right`$: $`r`$ is even and `right` is odd, so $`G_i(r) = res + \tfrac{r - right + 3}{2}`$. Meanwhile $`r + 2`$ and $`right'`$ are both even, so $`G_{i+1}(r+2) = res + 1 + \tfrac{r - right + 1}{2}`$, and the two are equal.

*Case 2: the next character is `)`.* Only states with $`r \ge 1`$ can read it, giving $`h(r-1) = g_i(r) \ge G_i(r)`$. We need to show $`G_{i+1}(r-1) \le G_i(r)`$.

- $`right \ge 1`$: the algorithm sets $`res' = res,\ right' = right - 1`$, a uniform shift by 1, so $`G_{i+1}(r-1) = G_i(r)`$.
- $`right = 0`$: the algorithm sets $`res' = res + 1,\ right' = 1`$. Here $`G_i(r)`$ equals $`res + r/2`$ for even $`r`$ and $`res + (r+3)/2`$ for odd $`r`$. Substituting term by term:
  - $`r = 1`$: $`G_{i+1}(0) = res + 2 = G_i(1)`$.
  - $`r = 2`$: $`G_{i+1}(1) = res + 1 = G_i(2)`$.
  - $`r \ge 3`$: the parity relationship between $`r - 1`$ and $`right' = 1`$ corresponds exactly to the two formulas above, so they are equal again.

Hence in both cases $`G_{i+1}`$ satisfies criteria (a) and (b), so $`g_{i+1} \ge G_{i+1}`$, and the lemma holds. $`\blacksquare`$

**4. The Lower Bound Equals the Algorithm's Result**

By the lemma, for any $`r \ge 0`$:

- If $`r \le right`$: $`g_n(r) + r \ge res + (right - r) + r = res + right`$.
- If $`r > right`$: $`g_n(r) \ge res`$, so $`g_n(r) + r \ge res + r > res + right`$.

Therefore:

$$
\text{OPT} = \min_r\bigl(g_n(r) + r\bigr) \ge res + right
$$

On the other hand, Part 3 has shown that the algorithm constructs a valid string using exactly $`res + right`$ insertions, so $`\text{OPT} \le res + right`$.

**Therefore, the algorithm attains the globally minimum number of insertions:**

$$
\text{OPT} = res + right
$$

**This proof matches the earlier intuition of "three kinds of necessary insertions"**: the branch in Case 1 where `right` is odd corresponds to "complete the half-finished `))`", the branch in Case 2 where `right = 0` corresponds to "insert `(` for an unmatched `)`", and the final `right` at termination corresponds to "fill in at the end". The lemma guarantees that no other insertion order can avoid these three costs.

### 4.6 Complexity Analysis

Let the length of the input string be n.

**Time complexity: O(n)**

The algorithm scans the input string once from left to right.

Each iteration performs at most a constant number of:

- Conditional checks.
- Integer additions and subtractions.
- Modulo operations.
- State repairs.

Therefore:

$$
T(n) = O(n)
$$

**Extra space complexity: O(1)**

The algorithm maintains only:

- `res`
- `right`
- the current character being traversed

No auxiliary data structure that grows with the input size is needed.

Therefore:

$$
S(n) = O(1)
$$

## 5. Learning Notes

### Difficulty 1: Why can't we just count left and right parentheses?

Because this problem includes a consecutiveness constraint.

For ordinary parenthesis matching, the number of unmatched left parentheses usually carries enough information.

But in this problem:

- Two `)` must be consecutive.
- Having matched the first `)` and not having matched any `)` are different states.

If we only count, this information may be lost.

**Transferable thinking: a state should express not only how much remains, but also which stage we are currently in.**

### Difficulty 2: Why must we repair first when `right` is odd?

The point is not oddness itself, but the matching state that oddness encodes.

For example:

```text
Input: ()(

Process the first '(':
right = 2

Process ')':
right = 1

Encounter the second '(':
must first insert one ')'
res = 1
right = 0

Then process the new '(':
right = 2

End of scan:
res + right = 3
```

If we do not repair immediately, the old `))` would be separated by the new `(`.

Therefore, `right % 2 == 1` is a compressed expression of the consecutive-match constraint.

**Transferable thinking: parity of an integer, or some other computable property, can be used to encode extra state, without necessarily needing an explicit boolean variable.**

But such compression must be backed by a proof of correctness and must not be used on experience alone.

### Difficulty 3: Why do we execute `right += 2` after `right < 0`?

Consider the current state:

`right = 0`

and encountering one `)`:

```text
right -= 1
# right = -1
```

This indicates a right parenthesis with no left parenthesis available to match it.

So we insert one `(`:

```text
res += 1
right += 2
```

Here `right += 2` corresponds to the two demands produced by the newly inserted `(`.

But the current `)` has already consumed one of them.

So in the end:

`right = 1`

only one `)` remains to be matched.

**Transferable thinking: derive state transitions from the meaning of the state before and after an operation, rather than mechanically increasing or decreasing variables by intuition.**

### Difficulty 4: Why not fill in all right parentheses immediately?

Because the future original input may contain enough `)`.

For example:

`())`

After processing `(`, `right = 2`.

If we inserted two `)` early at this point, we would perform unnecessary operations.

The correct strategy is to keep the demand and wait for the two subsequent `)` to complete the match naturally.

Repair immediately only when the constraint is about to be violated.

This reflects an important idea in greedy strategies:

**Defer non-essential decisions, and handle only the conflicts that are already unavoidable.**

### Difficulty 5: How to distinguish correctness of the algorithm from optimality?

This is something to which 1541 requires more attention than 1021.

- **Correctness**: the string constructed by the algorithm ultimately satisfies the balance rule.
- **Optimality**: no valid construction uses fewer insertions.

For example, merely showing that `right` can eventually return to zero only proves that the string can be completed.

We must also show that:

- Every triggered repair is necessary.
- Every repair uses the fewest possible insertions.
- Deferring the satisfaction of the remaining demand does not lose a better solution.

Therefore, this problem requires combining the loop invariant with the necessity of the greedy choices.

## 6. Key Takeaways

This problem yields several transferable methods.

### 6.1 From Remaining Elements to Remaining Demand

Sometimes directly maintaining the counts of elements that have appeared is not the most suitable state.

We can switch to:

> How much demand has been generated so far? How much still needs to be satisfied later?

This problem uses `right` to maintain the unsatisfied right-parenthesis demand.

This approach transfers to:

- Resource allocation and inventory reservation.
- Sequential protocol parsing.
- Stateful input validation.
- Event-driven task processing.

### 6.2 Using State Compression to Express Matching Progress

The parity of `right` expresses whether there is a half-completed `))`.

This shows that a single state variable can sometimes carry both:

- The remaining quantity.
- The local matching stage.

But the encoding rule and the conditions under which it is preserved must be made explicit.

### 6.3 Local Minimal Repair

Greedy does not simply choose the option with the fewest operations right now.

The real strategy in this problem is:

1. Determine whether the current constraint is about to be violated.
2. If a repair is required, determine the unavoidable minimum cost.
3. Perform that repair.
4. For demands that can still be satisfied by later input, do not intervene early.

**Local repair works because the necessity of each repair can be proven.**

### 6.4 State Invariants Before Code

When handling state-simulation problems, first answer:

- What does each variable represent?
- What condition must it satisfy at the start or end of each iteration?
- Which inputs break the condition?
- How is the condition restored?
- Why is the repair both sufficient and necessary?

The core invariant of this problem is:

`right >= 0`

and the value and parity of `right` accurately express the remaining consecutive-match demand.

Non-negativity alone is not enough; the demand semantics must be preserved as well.

### 6.5 Connection to 1021

| Comparison | 1021 | 1541 |
|---|---|---|
| Input condition | Already valid | May be invalid |
| Core state | `depth` | `right` |
| Meaning of state | Number of unmatched left parentheses | Right-parenthesis demand yet to be satisfied |
| Goal | Remove the outermost parentheses of each primitive | Insert the fewest characters to make the string valid |
| Key check | Whether we are at the outermost level | Whether a necessary repair is triggered |
| Core idea | State counting | Greedy + state simulation |
| Correctness focus | Boundary checks and loop invariants | Invariants, necessity of repairs, and optimality |

Both problems emphasize:

**Define the meaning of the state first, then decide the checks and update order according to the invariant.**

1541 further requires proving that the insertions are optimal while keeping the state valid.

## 7. Review Log

| Date | Independent completion | Change in understanding | To improve |
|---|---|---|---|
| 2026-10-09 | Solution process written up; degree of independent completion to be filled in | Realized that merely counting left and right parentheses cannot express the consecutiveness constraint; began to understand demand counting | Understand the exact semantics of `res` and `right` |
| 2026-10-10 | Python version written up; Go implementation and tests still to be recorded | Clarified that parity encodes a half-completed match, and understood local minimal repair when a conflict occurs | Rewrite the Go implementation without looking at the code; prove greedy optimality independently |

### Checklist for the Next Review

- [ ] Can I write the Go implementation independently without looking at the code?
- [ ] Can I explain why we cannot just count left and right parentheses?
- [ ] Can I accurately describe the meanings of `res` and `right`?
- [ ] Can I explain why the parity of `right` can express matching progress?
- [ ] Can I use `()(` to explain why the repair must come before processing a new `(`?
- [ ] Can I explain why we need to insert `(` when `right < 0`?
- [ ] Can I explain why we execute `right += 2` rather than `right += 1`?
- [ ] Can I independently write the initialization, maintenance, and termination proofs of the loop invariant?
- [ ] Can I distinguish the proof that the result is valid from the proof that the number of insertions is optimal?
- [ ] Can I summarize what 1021 and 1541 have in common in state modeling?

### Suggested Test Cases

| Input | Expected output | What it verifies |
|---|---|---|
| `())` | 0 | Already balanced |
| `(()))` | 1 | A single right-parenthesis demand remains at the end |
| `))())(` | 3 | Multiple insertions and filling in at the end |
| `(` | 2 | A single left parenthesis |
| `)` | 2 | A single isolated right parenthesis |
| `()(` | 3 | Odd demand meeting a new left parenthesis |
| `(((` | 6 | All left parentheses |
| `()))` | 2 | Extra right parenthesis after everything is already closed |

## 8. Related Problems

- [1021. Remove Outermost Parentheses](https://leetcode.com/problems/remove-outermost-parentheses/) — uses a counter to describe the nesting state of parentheses; good for comparing state definitions and update order.
- [0921. Minimum Add to Make Parentheses Valid](https://leetcode.com/problems/minimum-add-to-make-parentheses-valid/) — also asks for the fewest insertions, but each `(` matches only one `)`; good for comparing how the state changes once the consecutiveness constraint is added.
- [0020. Valid Parentheses](https://leetcode.com/problems/valid-parentheses/) — requires maintaining bracket types and matching relations; good for comparing the applicability boundaries of stacks and state counting.
- [1249. Minimum Remove to Make Valid Parentheses](https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/) — also involves repairing the validity of parentheses, but the operation changes from insertion to deletion; good for comparing strategies under different operation constraints.

---

**One-sentence summary**

The core of 1541 is not just greedily adding parentheses, but converting the matching process into a right-parenthesis demand state, using parity to identify unfinished consecutive matches, and performing the necessary minimal repair only when a constraint is about to be violated; remaining demands are left to later input as far as possible, achieving the globally minimum number of insertions while maintaining the state invariant.