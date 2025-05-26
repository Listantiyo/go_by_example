// Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.

package main

import (
	"fmt"
)

// Variadic function that takes a variable number of integers and returns their sum
func sum(numbers ...int) {
	fmt.Print(numbers, " ")
	total := 0

	for _, number := range numbers {
		total += number
	}

	fmt.Println("Sum:", total)
}

func main() {
	
	sum(1,2)
	sum(1,2,3)

	nums := []int{1,2,3,4,5}
	sum(nums...) // Unpacking the slice into the variadic function
}

// What : Variadic functions in Go allow you to pass a variable number of arguments to a function.
// How : You define a variadic function by appending three dots (...) after the last parameter.
// Why : Variadic functions are useful when you don't know the number of arguments in advance.