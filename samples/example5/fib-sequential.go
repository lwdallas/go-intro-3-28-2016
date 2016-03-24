package main

import (
	"fmt"
	"time"
)

func fib(i int) int {
	if i < 2 {
		return 1
	}
	return fib(i-1) + fib(i-2)
}

func main() {
	start := time.Now()

	fmt.Println(fib(42), fib(42), fib(42), fib(42))

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
