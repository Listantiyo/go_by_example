// Functions are central in Go. We’ll learn about functions with a few different examples.

package main

import (
	"fmt"
)

func plus(a int, b int)int{
	return a + b
}

func plusPlus(a, b, c int)int{ // a, b, c are all int
	return a + b + c
}

func main() {
	result := plus(1,2)
	fmt.Println("1 + 2 =", result)

	result = plusPlus(1,2,3)
	fmt.Println("1 + 2 + 3 =", result)
}