// Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages).

package main

import (
	"fmt"
	"maps"
)

func main() {

	// To create a map, use the make function
	m := make(map[string]int) // Create a map with string keys and int values

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3) // if the key does not exist, it returns the zero value for the value type (0 for int, "" for string, etc.)
	
	fmt.Println("len:", len(m)) // len returns the number of key/value pairs in the map

	delete(m, "k2") // delete a key/value pair from the map
	fmt.Println("map:", m)

	clear(m) // clear the map
	fmt.Println("map:", m)

	_, prs := m["k2"] // _ is used to ignore the value returned by the map lookup & prs is a boolean that indicates whether the key exists in the map
	fmt.Println("prs:", prs) // check if a key exists in the map

	n := map[string]int{"foo": 1, "bar": 2} // create and initialize a map in one line
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2} // create and initialize a map in one line
	if maps.Equal(n, n2){
		fmt.Println("n == n2")
	}
}

// Map in go is like a object in javascript with key value pairs