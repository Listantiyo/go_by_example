// Go supports constants of character, string, boolean, and numeric values.

package main

import (
	"fmt"
	"math"
)

const s string = "constant"
func main() {
	fmt.Println(s)

	const n = 500000000 //500.000.000
	
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d)) // Explicit conversion from float64 to int64
	fmt.Println(n/d) // Implicit conversion from int to float64
	fmt.Println(math.Sin(n)) // math.Sin() function takes a float64 argument and returns a float64 value.
}

// Note : Constants are typed, but they can be implicitly converted to other types when used in expressions.