package main

import "fmt"

// Factorial function using recursion
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	var fib func(int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}

// What is recursion?
// Recursion is a function that calls itself. It is used to solve problems that can be broken down into smaller subproblems. The base case is the condition under which the recursion stops. In the factorial example, the base case is when n is 0. In the Fibonacci example, the base case is when n is less than 2.

// Note : recursion is a function that calls itself. It is used to solve problems that can be broken down into smaller subproblems. The base case is the condition under which the recursion stops. In the factorial example, the base case is when n is 0. In the Fibonacci example, the base case is when n is less than 2.

