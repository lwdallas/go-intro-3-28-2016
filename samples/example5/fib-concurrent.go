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
	answer := make(chan int)
	fib_worker := func(i int) {
		answer <- fib(i)
	}

	go fib_worker(42)
	go fib_worker(42)
	go fib_worker(42)
	go fib_worker(42)

	fmt.Println(<-answer, <-answer, <-answer, <-answer)

	elapsed := time.Since(start)
	fmt.Println(elapsed)

}
