package collection

// Reduce reduces a collection of T into a single value of type R.
//
// Example:
//   sum := Reduce(nums, 0, func(acc, n int) int { return acc + n })
// // sum is the total of all numbers in nums
// concatenated := Reduce(strings, "", func(acc, s string) string { return acc + s })
// // concatenated is all strings in strings joined together
func Reduce[T any, R any](c Collection[T], initial R, fn func(R, T) R) R {
	acc := initial
	for _, v := range c.Items() {
		acc = fn(acc, v)
	}
	return acc
}
