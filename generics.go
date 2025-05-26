package main

import "fmt"

// Generic function to find the maximum of two values
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// Note: arti ~ adalah bahwa S adalah slice dari E, dan E adalah tipe yang dapat dibandingkan (comparable).
// dan slice bisa termasuk alias dari slice yang sudah ada, misalnya:
// type MySlice []int

/* Generic Singly Linked List (List of Nodes) */

// Ini adalah Object yang menyimpan Object linked list
type List[T any] struct {
	head, tail *element[T]
}

// Ini akan menjadi Object linked list
type element[T any] struct {
	next *element[T]
	val T
}

// Method Push untuk menambahkan elemen baru ke dalam linked list.
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {

	var s = []string{"foo","bar","baz"}

	// Menggunakan fungsi generik untuk mencari indeks elemen dalam slice
	// Memanfaatkan typeinference untuk menentukan tipe parameter generik
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))
	// Menentukan tipe parameter generik secara eksplisit
	_ = SlicesIndex[[]string, string](s, "zoo")

	// Membuat linked list dengan tipe data int

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}

// What : Generic
// Why : Membuat kode lebih reusable
// How : Menggunakan type parameter
// Note:
// 1. Generic function
// 2. Generic type
// 3. Generic method
// 4. Generic constraint
// 5. Generic interface
// 6. Generic type parameter
// 7. Generic type set