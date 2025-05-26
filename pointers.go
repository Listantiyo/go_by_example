// Go supports pointers, allowing you to pass references to values and records within your program.

package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0 // dereference the pointer to set the value to 0
}

func main() {
	i := 1
	fmt.Println("initial:", i) // prints 1
	zeroval(i)
	fmt.Println("zeroval:", i)
	zeroptr(&i) // pass the address of i to the function
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i) // prints the address of i
}

// What is a pointer?
// A pointer is a variable that stores the memory address of another variable.
// The & operator returns the memory address of a variable.
// The * operator returns the value stored at a memory address.

// The * operator is called the dereference operator or the indirection operator.
// The & operator is called the address operator.