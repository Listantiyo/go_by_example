package main

import (
	"fmt"
	"iter"
	"slices"
)

// List is a generic list
// T is the type of the list
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next  *element[T]
	val T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	} 
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	for e := range lst.All() {
		fmt.Println(e)
	}

	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for n := range genFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}

// Range over iterators artinya anda dapat menggunakan range untuk iterasi pada iterator
// iterator adalah objek yang memiliki method next() yang mengembalikan nilai dan boolean
// interator bukanlah data, melainkan alat untuk mengiterasi suatu data dengan suatu kerangka iterasi yang telah ditentukan.
// ciri iterator adalah memiliki method next() yang mengembalikan nilai dan boolean
// next() adalah method yang digunakan untuk mengiterasi data
// next() mengembalikan nilai dan boolean
// nilai adalah data yang akan diiterasi
// boolean adalah kondisi yang digunakan untuk menghentikan iterasi
// jika boolean bernilai true maka iterasi akan dihentikan
// jika boolean bernilai false maka iterasi akan dilanjutkan
