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