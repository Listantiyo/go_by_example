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