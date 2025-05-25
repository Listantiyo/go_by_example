//Branching with if and else in Go is straight-forward.

package main

import (
	"fmt"
)

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	}else{
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	}else if num < 10{
		fmt.Println(num, "is 1 digit")
	}else{
		fmt.Println(num, "has multiple digits")
	}
}

// Note : ternary operator is not available in Go. The if statement can be used to achieve similar functionality.
// The if statement can be used to achieve similar functionality.