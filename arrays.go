// In Go, an array is a numbered sequence of elements of a specific length. In typical Go code, slices are much more common; arrays are useful in some special scenarios.

package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a) // emp: [0 0 0 0 0]

	a[4] = 100 
	fmt.Println("set:",a) // set: [0 0 0 0 100]
	fmt.Println("get:",a[4]) // get: 100

	fmt.Println("len:",len(a)) // len: 5

	b := [5]int{1,2,3,4,5} 
	fmt.Println("dcl:",b) // dcl: [1 2 3 4 5]

	b = [...]int{1,2,3,4,5} // [...] means the compiler will count the number of elements in the array
	fmt.Println("dcl:",b) // dcl: [1 2 3 4 5]

	b = [...]int{100,3:400,500}
	fmt.Println("dcl:",b) // dcl: [100 0 0 400 500]

	var twoD [2][3]int
	for i := 0; i<2; i++ {
		for j := 0;j<3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ",twoD) // 2d:  [[0 1 2] [1 2 3]]

	twoD = [2][3]int{
		{1,2,3},
		{4,5,6},
	}
	fmt.Println("2d: ",twoD) // 2d:  [[1 2 3] [4 5 6]]
}


//Note : In Go, arrays are fixed-size sequences of elements of the same type.
// Array is used in spesial cases where the size of the array is known at compile time and does not change.