// Package math provides math helper functions.
package math

// IntAbs returns the absolute value of n.
func IntAbs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// IntMax returns the larger of a and b.
func IntMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Factorial returns n!.
func Factorial(n int) int {
	if n == 0 {
		return 1
	}

	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}

// IntPermutations returns a slice of integer slices with all permutations of
// the input slice s.
func IntPermutations(s []int) [][]int {
	results := make(chan []int)
	permutations := make([][]int, Factorial(len(s)))

	go generateIntPermutation(len(s), s, results)
	for i := 0; i < len(permutations); i++ {
		permutations[i] = <-results
	}
	return permutations
}

// generateIntPermutation implements a variant of Heap's algorithm as shown in
// Sedgewicks's slides at https://www.cs.princeton.edu/~rs/talks/perms.pdf. It
// is not ideal in that the number of swaps performed is larger than the
// factorial of len(s), but it is simpler than the suggested implementation on
// Wikipedia.
func generateIntPermutation(n int, s []int, out chan<- []int) {
	if n == 1 {
		perm := make([]int, len(s))
		copy(perm, s)
		out <- perm
		return
	}

	for c := 0; c < n; c++ {
		generateIntPermutation(n-1, s, out)
		if n%2 == 1 {
			s[0], s[n-1] = s[n-1], s[0]
		} else {
			s[c], s[n-1] = s[n-1], s[c]
		}
	}
}