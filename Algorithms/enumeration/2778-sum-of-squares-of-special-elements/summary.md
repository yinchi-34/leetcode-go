# 2778. Sum of Squares of Special Elements

English | [简体中文](./summary.zh-CN.md)

[Shared Glossary](../../../docs/GLOSSARY.md)

## 1. Problem Information

- **Problem**: [2778. Sum of Squares of Special Elements](https://leetcode.cn/problems/sum-of-squares-of-special-elements/)
- **Official Difficulty**: Easy
- **Primary Category**: Enumeration
- **Secondary Category**: Number Theory
- **Tags**: `Enumeration`, `Factorization`, `Square-Root Optimization`, `Paired Enumeration`, `Index Mapping`
- **Code**: [solution.go](./solution.go)
- **Tests**: [solution_test.go](./solution_test.go)

### Difficulty Rating

**External Ratings**

- zerotrac Rating: Pending verification against the [source data](https://github.com/zerotrac/leetcode_problem_rating/blob/main/ratings.txt).
- LeetCode Official Difficulty: Easy

**Four-Dimensional Rating - Paired Divisor Enumeration Up to $`\sqrt{n}`$**

Evaluation scope: The problem's input constraints; divisor enumeration with paired processing; target time complexity $`O(\sqrt{n})`$ and auxiliary space complexity $`O(1)`$; implementation in Go.

| Dimension | Score | Reason |
|---|---|---|
| Modeling Difficulty | 1/5 | The statement directly gives the condition that an index divides $`n`$; it only needs to be translated into divisor enumeration. |
| Knowledge Prerequisites | 2/5 | Requires knowing that divisors occur in pairs and that enumeration only needs to reach $`\sqrt{n}`$. |
| Correctness Reasoning Difficulty | 2/5 | Requires showing that pairing neither misses nor duplicates any divisor, especially when `i == n/i`. |
| Implementation Difficulty | 2/5 | One-based indexing requires subtracting 1, and a perfect square can cause duplicate accumulation. |

The four dimensions are evaluated independently and are not combined into an overall score.

## 2. Problem Understanding

Given an integer array `nums` of length $`n`$, with **one-based indexing**.

If an index `i` satisfies `n % i == 0`, then `nums[i]` is a **special element**.

Return the sum of the squares of all special elements.

Key requirements:

1. The statement uses one-based indexing, while Go slices use zero-based indexing.
2. The condition is that **`i` divides the array length $`n`$**, rather than the element's value.
3. The input is small ($`1 \leq n \leq 50`$), so brute force is sufficient.

### Core Problem

The task is to **find every divisor of $`n`$ and use each divisor as an index.**

For example, when `n = 4`:

- The divisors are `1, 2, 4`.
- The corresponding elements are `nums[1], nums[2], nums[4]` in one-based indexing.
- In Go, these are `nums[0], nums[1], nums[3]`.

## 3. Initial Approach

### 3.1 The Most Direct Idea: Enumerate From 1 to $`n`$

```go
for i := 1; i <= n; i++ {
    if n%i == 0 {
        res += nums[i-1] * nums[i-1]
    }
}
```

This is correct and takes $`O(n)`$ time, which is sufficient for $`n \leq 50`$.

### 3.2 Further Optimization: Divisors Occur in Pairs

If `i` is a divisor of $`n`$, then `n/i` is also a divisor of $`n`$.

One member of the pair is at most $`\sqrt{n}`$, and the other is at least $`\sqrt{n}`$.

Therefore, it is sufficient to enumerate `i` from 1 to $`\sqrt{n}`$:

- Whenever a divisor `i` is found, its paired divisor `n/i` is obtained immediately.
- The squared elements at both indices must be added to the result.

### 3.3 Special Case: Perfect Squares

When `i * i == n`, `i` and `n/i` are **the same divisor**.

The corresponding square must be added only once to avoid double counting.

This is the purpose of `if i*i < n` in the code.

### Thought Transformation

From:

> Check each index individually to see whether it divides $`n`$.

To:

> Divisors occur in pairs. Check only the smaller members, and obtain the larger members directly from the pairing relation.

**Key insight: When the target objects have a symmetric or paired structure, enumerate one half and obtain the other half through a mapping.**

## 4. Algorithm Specification

**Algorithm Name:** Paired Divisor Enumeration

### 4.1 Input

An integer array of length $`n`$:

```math
nums = \langle a_1, a_2, \ldots, a_n \rangle
```

**Preconditions:**

- **Length constraint**: $`1 \leq n \leq 50`$.
- **Element constraint**: $`1 \leq a_i \leq 50`$.
- **Indexing convention**: The statement uses one-based indexing; the Go implementation accesses `nums[i-1]`.

### 4.2 Output

An integer `res`:

```math
res = \sum_{\substack{1 \leq i \leq n \\ i \mid n}} a_i^2
```

### 4.3 Pseudocode

```text
SUM-OF-SQUARES(nums)
    n <- length(nums)
    res <- 0

    for i <- 1 while i * i <= n
        if n mod i == 0
            res <- res + nums[i]^2

            if i * i < n
                res <- res + nums[n / i]^2

    return res
```

**State Definitions**

- `n`: The array length, whose divisors are being enumerated.
- `i`: The current candidate for the smaller divisor, satisfying $`i \leq \sqrt{n}`$.
- `n/i`: The larger divisor paired with `i` when `i` divides `n`.
- `res`: The accumulated sum of the squares of special elements.

### 4.4 Correctness

**Proof Method: Bijection Between Paired Divisors and a Loop Invariant**

#### I. Key Properties

**Property A: Divisors Occur in Pairs**

If `d` is a divisor of $`n`$, then `n/d` is also a divisor of $`n`$, and

```math
d \cdot \frac{n}{d} = n
```

Therefore, at least one of `d` and `n/d` is at most $`\sqrt{n}`$.

**Property B: The Mapping $`d \mapsto n/d`$ Is a Bijection**

This mapping gives a one-to-one correspondence between the divisors at most $`\sqrt{n}`$ and those at least $`\sqrt{n}`$.

- When $`d < \sqrt{n}`$, $`n/d > \sqrt{n}`$, so the two divisors are distinct.
- When $`d = \sqrt{n}`$, $`n/d = d`$, so both members of the pair are the same divisor.

#### II. Loop Invariant

Suppose the iterations for $`i = 1, 2, \ldots, k`$ have been completed.

**Invariant:** `res` equals the sum of the squared elements at all indices `d` satisfying $`d \mid n`$ and ($`d \leq k`$ or $`n/d \leq k`$), with each divisor counted exactly once.

#### III. Proof

**1. Initialization**

When $`k = 0`$, no value of `i` has been processed and `res = 0`, so the loop invariant holds.

**2. Maintenance**

Consider the iteration for $`i = k+1`$:

- If `n % i != 0`, `i` is not a divisor. `res` stays unchanged, preserving the loop invariant.
- If `n % i == 0` and `i*i < n`, `i` and `n/i` are distinct divisors. Their corresponding squares are each added once.
- If `n % i == 0` and `i*i == n`, `i == n/i`. The corresponding square is added only once.

For either divisor case, the smaller member `i` has not been processed before, so the pair is new. Thus, each divisor covered so far is counted exactly once in all three cases.

**3. Termination**

The loop ends when `i*i > n`, so $`i > \sqrt{n}`$.

Every divisor `d` satisfies $`d \leq \sqrt{n}`$ or $`n/d \leq \sqrt{n}`$. Hence, every divisor has been covered.

By the loop invariant, `res` is exactly the sum of the squares of the elements at all divisor indices.

### 4.5 Complexity Analysis

**Time Complexity: $`O(\sqrt{n})`$**

The loop variable `i` satisfies $`i^2 \leq n`$, so there are at most $`\lfloor\sqrt{n}\rfloor`$ iterations. Each iteration performs a constant amount of work.

**Auxiliary Space Complexity: $`O(1)`$**

Only three variables are used: `n`, `res`, and `i`.

> Note: For $`n \leq 50`$, the practical performance difference between $`O(n)`$ and $`O(\sqrt{n})`$ is negligible. The value of the square-root approach here is practicing a reusable pattern.

## 5. Learning Notes

### Challenge 1: Why Is `i*i < n` Necessary?

Take `n = 4` as an example:

```text
i = 1: Divisor pair (1, 4); add the squares of nums[0] and nums[3].
i = 2: Divisor pair (2, 2); add the square of nums[1] only once.
```

Without `i*i < n`, the square of `nums[1]` would be added twice when `i = 2`.

**Transferable insight: In paired enumeration, first ask whether the two members can coincide. Any shared member requires separate handling.**

### Challenge 2: The Condition Concerns the Index Dividing $`n`$

The condition checks whether the **index `i`** divides the **array length $`n`$**. It is independent of the value of `nums[i]`.

**Transferable insight: Distinguish positions from values when reading an array problem, because the two are easily confused.**

## 6. Key Takeaways

### 6.1 A Pattern for Paired Divisor Enumeration Up to $`\sqrt{n}`$

```go
for i := 1; i*i <= n; i++ {
    if n%i == 0 {
        // Process divisor i.
        if i*i < n {
            // Process the paired divisor n/i.
        }
    }
}
```

This pattern can be adapted to:

- Counting divisors or summing divisors.
- Checking primality or testing for perfect numbers.
- Finding the $`k`$-th divisor.
- Counting the common divisors of two numbers.

### 6.2 The General Idea of Paired Enumeration

When objects have a paired structure:

1. Enumerate only the smaller members.
2. Obtain the other members through a mapping.
3. Handle any shared member separately.

### 6.3 Choosing Between Brute Force and Optimization

For small inputs, direct $`O(n)`$ enumeration is less prone to mistakes.

During practice, deliberately using the $`O(\sqrt{n})`$ approach helps establish a reusable pattern for problems with larger inputs.

## 7. Review Log

| Date | Independent Completion | Understanding Gained | Areas for Improvement |
|---|---|---|---|
| 2026-10-11 | Notes organized; level of independent completion to be recorded. | Understood divisor pairing and how `i*i < n` prevents duplicate accumulation. | Reproduce the pattern without referring to the code; complete solution.go and tests. |

### Next Review Checklist

- [ ] Can I write the Go implementation independently without looking at the code?
- [ ] Can I explain why enumeration only needs to reach $`\sqrt{n}`$?
- [ ] Can I explain which duplicate `i*i < n` prevents?
- [ ] Can I explain why the loop condition is `i*i <= n` rather than `i*i < n`?
- [ ] Can I correctly convert between one-based and zero-based indexing?

## 8. Related Problems

- [0507. Perfect Number](https://leetcode.cn/problems/perfect-number/) - Uses the same paired divisor enumeration up to $`\sqrt{n}`$ to sum divisors.
- [1492. The kth Factor of n](https://leetcode.cn/problems/the-kth-factor-of-n/) - A direct application of divisor enumeration that also requires handling divisor order.
- [1952. Three Divisors](https://leetcode.cn/problems/three-divisors/) - Counts divisors and similarly requires handling `i == n/i`.
- [2427. Number of Common Factors](https://leetcode.cn/problems/number-of-common-factors/) - A variation on divisor enumeration involving common divisors of two numbers.

---

**One-Sentence Summary**

Problem 2778 translates "an index divides the array length" into divisor enumeration: enumerate only up to $`\sqrt{n}`$ and obtain each paired divisor through `n/i`, using `i*i < n` to avoid double counting for perfect squares and converting between one-based and zero-based indexing.
