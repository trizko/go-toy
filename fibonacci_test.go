package fibonacci

import (
	"fmt"
	"testing"
	"time"
)

func TestFibonacci(t *testing.T) {
	cases := []struct {
		in   int64
		want int64
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{42, 267914296},
		{100, 3736710778780434371},
	}
	for _, c := range cases {
		start := time.Now()
		got := Fibonacci(c.in)
		elapsed := time.Since(start)

		if got != c.want {
			t.Errorf("Fibonacci(%v)\tFAIL | returns: %v\t | wanted : %v\n", c.in, got, c.want)
		} else if elapsed.Seconds() > 1. {
			t.Errorf("Fibonacci(%v)\tPASS but takes too long to run | returns: %v in %vs", c.in, got, elapsed.Seconds())
		} else {
			fmt.Printf("Fibonacci(%v)\tPASS | returns: %v\n", c.in, got)
		}
	}
}
