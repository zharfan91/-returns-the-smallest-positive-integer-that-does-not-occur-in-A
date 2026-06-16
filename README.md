# Smallest Missing Positive Integer

## Problem Details

Given an array `A` containing `N` integers, write a function that returns the smallest positive integer, greater than `0`, that does not occur in `A`.

Examples:

```text
A = [1, 3, 6, 4, 1, 2]
Result = 5
```

```text
A = [1, 2, 3]
Result = 4
```

```text
A = [-1, -3]
Result = 1
```

Assumptions:

- `N` is an integer within the range `1..100000`
- Each element of the array is an integer within the range `-1000000..1000000`

## Solution In Golang

The implementation is available in `main.go`.

```go
func Solution(A []int) int {
	seen := make([]bool, len(A)+2)

	for _, value := range A {
		if value > 0 && value < len(seen) {
			seen[value] = true
		}
	}

	for candidate := 1; candidate < len(seen); candidate++ {
		if !seen[candidate] {
			return candidate
		}
	}

	return len(seen)
}
```

## Code Explanation

This solution uses a boolean array named `seen` to mark positive numbers that appear in `A`.

The size of `seen` is `len(A) + 2` because the largest possible answer that needs to be checked is `N + 1`. For example, if `A = [1, 2, 3]`, then every number from `1` to `N` exists, so the correct answer is `4`.

How the code works:

1. Create the `seen` array.
2. Read each value in `A`.
3. If the value is positive and still within the range that needs to be checked, mark `seen[value] = true`.
4. Search for the first number starting from `1` that has not been marked.
5. The first unmarked number is the smallest positive integer that does not occur in the array.

## Complexity

- Time: `O(N)`, because the array is traversed linearly.
- Memory: `O(N)`, because the solution uses an additional boolean array.

## Running The Tests

Run the following command:

```bash
go test ./...
```
