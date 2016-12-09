package fib

// Fib returns fibonacci number of n.
func Fib(n int64) int64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}
