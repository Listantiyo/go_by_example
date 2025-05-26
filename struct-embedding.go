package main

import "fmt"

// creates a base struct with an int field
type base struct {
	num int
}
// create a method for the base struct
func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
	// Sprintf adalah fungsi dari package fmt yang digunakan untuk format string
	// dan mengembalikan string yang diformat.
}

type container struct {
	base // embed the base struct
	str string // add a new field
}

func main() {
	co := container{ // creating struct literal
		base: base{num: 42}, // initializing the embedded struct
		str:  "hello", // initializing the new field
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str) // accessing the embedded struct field and the new field
	// Printf adalah fungsi dari package fmt yang digunakan untuk format string
	// dan mencetak string yang diformat ke output standar (biasanya layar).

	fmt.Println("alse num:", co.base.num) // accessing the embedded struct field

	fmt.Println("describe:", co.describe()) // method describe() can be called directly
	// karena method describe() di-embed ke dalam struct container

	type describer interface {
		describe() string // method describe() is defined in the interface
	}

	// create a variable of type describer
	var d describer = co // assign the container to the describer interface
	fmt.Println("describer:", d.describe()) // call the describe() method from the interface
}

// What : embedding struct adalah sebuah fitur dalam bahasa pemrograman Go yang memungkinkan sebuah struct
// untuk mengandung atau menempelkan (embed) field atau method dari struct lain.

// Why : embedding struct digunakan untuk mengimplementasikan konsep pewarisan (inheritance) dalam bahasa pemrograman Go.
