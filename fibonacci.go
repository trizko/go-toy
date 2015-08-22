// Package fibonacci finds the nth element of the fibonacci sequence
package fibonacci

// Fibonacci finds the nth element of the fibonacci sequence
func Fibonacci(n int64) int64 {
	mem := map[int64]int64{0: 0, 1: 1}
	if _, present := mem[n]; !present {
		for i := int64(2); i <= n; i++ {
			mem[i] = mem[i-1] + mem[i-2]
		}
	}

	return mem[n]
}
