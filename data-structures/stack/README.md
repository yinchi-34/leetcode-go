# Stack

English | [简体中文](./README.zh-CN.md)

Use the [shared glossary](../../docs/GLOSSARY.md) and follow the repository's [GitHub math conventions](../../docs/GITHUB_MATH.md).

This is the overview note for the stack: its definition, implementation, costs, and the usage patterns that recur across problems. Each application problem's note covers only how that problem uses a stack; knowledge about the stack itself lives here.

## 1. Definition

A stack is a **last-in first-out** (LIFO) dynamic set: deletion always removes the most recently inserted element that has not yet been removed. It supports four operations:

| Operation | Meaning | Name in CLRS |
|---|---|---|
| push(x) | Put `x` on top | PUSH |
| pop() | Remove and return the top element | POP |
| top() | Return the top element without removing it | No separate procedure; read `S[S.top]` |
| empty() | Test whether the stack is empty | STACK-EMPTY |

Popping an empty stack is an **underflow**; in an array implementation, exceeding the capacity is an **overflow**.

## 2. Implementation and Cost

### 2.1 The CLRS array implementation (§10.1)

CLRS implements a stack with an array `S[1..n]` and an index `S.top`: `S[1]` is the bottom, `S[S.top]` is the top, and `S.top = 0` means the stack is empty. PUSH increments `S.top` and then writes; POP reads `S[S.top]` and then decrements. Each operation takes $`O(1)`$.

### 2.2 Go idiom

Go has no built-in stack type. The idiom is a slice whose end is the top:

```go
st := []int{}             // empty stack
st = append(st, x)        // push
v := st[len(st)-1]        // top
st = st[:len(st)-1]       // pop (read with top first, then truncate)
empty := len(st) == 0     // empty
```

Costs:

- top, pop, and empty are $`O(1)`$.
- push is $`O(1)`$ when capacity suffices; otherwise `append` allocates a larger backing array and copies, costing $`O(k)`$ once ($`k`$ is the current length). The Go runtime grows capacity geometrically (the exact factor is an implementation detail not fixed by the language spec), so $`m`$ pushes cost $`O(m)`$ in total, i.e. **amortized** $`O(1)`$. This is the same argument as the dynamic-table analysis in CLRS §17.4.
- If the maximum depth can be estimated, `make([]int, 0, cap)` allocates once and avoids growth.
- pop only truncates the length and does not free the backing array. When elements are pointers, set `st[len(st)-1]` to the zero value before truncating so the garbage collector can reclaim it.

### 2.3 Space

The extra space of a stack is its **maximum length during the run**, not its final length. When analyzing complexity, first identify what determines that maximum (for example, the maximum nesting depth of a parentheses string), then state the worst case.

## 3. Common Usage Patterns

Each pattern lists "what the stack holds" and "when it applies". New problems are added under the matching pattern.

### 3.1 Nested structures: finish the inner layer, then return to the outer one

- **What the stack holds**: information for each layer that is still open; the top is the innermost layer.
- **When it applies**: the input is nested (parentheses, tags, a recursively defined grammar), and an outer result can only be settled after the inner one is computed.
- **Why a stack**: the last layer opened is always the first to close, which is exactly last-in first-out.
- **Proof method**: write a loop invariant "the $`i`$-th element of the stack is the partial result of layer $`i`$" and prove it is maintained in the push and pop cases.

### 3.2 Expression evaluation

- **What the stack holds**: an operand stack, plus an operator stack when needed.
- **When it applies**: expressions with parentheses or precedence; postfix expressions need only an operand stack.

### 3.3 Monotonic stack

- **What the stack holds**: indices whose values are monotonic from bottom to top.
- **When it applies**: for each element, find the first larger (or smaller) element to its left or right.
- **Cost**: each index is pushed and popped at most once, so the total time is $`O(n)`$. This is the same argument as the MULTIPOP example of aggregate analysis in CLRS §17.1: the total number of pops cannot exceed the total number of pushes.

### 3.4 An explicit stack in place of recursion

- **What the stack holds**: the state that a recursive call would have saved on the call stack.
- **When it applies**: recursion may be too deep and overflow, or the traversal must be paused and resumed. CLRS Exercise 10.4-3, which asks for a nonrecursive binary-tree traversal using a stack, is this pattern.

## 4. When a Stack Is Not Needed

Problems solved with a stack often have an $`O(1)`$ extra-space solution. There are two different ways to remove the stack, and they should be kept apart:

1. **The stack collapses into a counter**: if the pushed elements are indistinguishable (for example, only `(` is pushed), the stack's contents are determined by its length, so an integer recording the length suffices. push becomes increment, pop becomes decrement, and empty becomes a test for 0. This is an equivalent rewrite of the same algorithm, and the correctness proof is unchanged.
2. **A different computation**: if the stack elements differ (for example, each layer stores its own score), the stack cannot be compressed into a counter directly. The answer must be re-derived in a form that no longer depends on each layer's value. This is a different algorithm and needs its own proof.

## 5. Application Problems

| Problem | Location | Pattern | Notes |
|---|---|---|---|
| [856. Score of Parentheses](./0856-score-of-parentheses/summary.md) | This directory | 3.1 nested structures; case 2 of section 4 | The stack holds the score of each layer's closed part; the depth solution rewrites the answer as the sum of $`2^{d}`$ over `()` pairs |
| [1021. Remove Outermost Parentheses](../../Algorithms/simulation/1021-remove-outermost-parentheses/summary.md) | Simulation | Case 1 of section 4 | The stack would hold only `(`, so it collapses into a depth counter |
| [1541. Minimum Insertions to Balance a Parentheses String](../../Algorithms/greedy-algorithms/1541-minimum-insertions-to-balance-a-parentheses-string/summary.md) | Greedy | Case 1 of section 4 | Unmatched opening parentheses need not be distinguished, so demand counting replaces the stack |

## 6. CLRS Connections

Section numbers follow CLRS, 3rd edition.

| Concept | Section | Relation | How It Appears Here |
|---|---|---|---|
| Stack definition and array implementation | §10.1 Stacks and queues | Direct | The four operations in section 1 and the `S.top` implementation in section 2.1 |
| Amortized analysis of dynamic tables | §17.4 Dynamic tables | Direct | Amortized $`O(1)`$ growth of `append` in section 2.2 |
| Aggregate analysis | §17.1 Aggregate analysis | Direct | In section 3.3, the total number of monotonic-stack pops does not exceed the total number of pushes |
