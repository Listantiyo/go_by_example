// Slices are an important data type in Go, giving a more powerful interface to sequences than arrays.

package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("uninitialized:", s, s == nil, len(s) == 0, cap(s) == 0)

	s = make([]string, 3)// make is a built-in function that allocates and initializes slices, maps, and channels.
	fmt.Println("emp:", s,"len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s,"d")
	s = append(s, "e", "f") // append is a built-in function that appends elements to a slice.
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c,s) // copy is a built-in function that copies elements from one slice to another.
	fmt.Println("cpy:", c)

	l := s[2:5] // slicing a slice
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2){
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLength := i + 1
		twoD[i] = make([]int, innerLength)
		for j := 0; j < innerLength; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD) 

	//Note: Slices can be composed into multi-dimensional data structures. The length of the inner slices can vary, unlike with multi-dimensional arrays.
}	