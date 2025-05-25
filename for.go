// for is Go’s only looping construct. Here are some basic types of for loops.

package main

import (
	"fmt"
)

func main() {

	//The most basic type, with a single condition.
	i := 1
	for i<=3 {
		fmt.Println(i)
		i = i + 1
	}
	//A classic initial/condition/after for loop.
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
	//Another way of accomplishing the basic “do this N times” iteration is range over an integer.
	for i := range 3{
		fmt.Println("range", i)
	}
	//for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
	for {
		fmt.Println("infinite loop")
		break // Break out of the infinite loop
	}
	//You can also continue to the next iteration of the loop.
	for n := range 6 {
		if n%2 == 0 {
			continue // Skip even numbers
		}

		fmt.Println(n)
	}	
}

// Note : The for loop is the only loop available in Go. It can be used to create a while loop or a do-while loop by using the appropriate syntax.
// The for loop can also be used to iterate over a range of values or a collection.