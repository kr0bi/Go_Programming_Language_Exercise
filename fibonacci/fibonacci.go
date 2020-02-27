package main

import (
	"fmt"
	"math"
)

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func fib_formula(x float64) float64 {
	return math.Round(math.Pow((math.Sqrt(5)+1)/2, x) / math.Sqrt(5))
}

func main() {
	for i := 1.0; i < 1000; i++ {
		fmt.Printf("fib(x): %f\n", fib_formula(i))
	}
}
