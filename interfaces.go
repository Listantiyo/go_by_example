// Interfaces are named collections of method signatures.

package main

import (
	"fmt"
	"math"
)

type geometry interface { // interface to define methods for area and perimeter
	area() float64
	perimeter() float64
}

type rect struct { // struct to define rectangle
	width, height float64
}

type circle struct { // struct to define circle
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func detecCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius:", c.radius)
	}
}

func main() {
	r := rect{width:3,height:4}
	c := circle{radius:5}

	measure(r)
	measure(c)

	detecCircle(r)
	detecCircle(c)
}


// Tujuan Interfaces adalah untuk membuat kode lebih modular dan reusable.
// Interfaces juga dapat digunakan untuk membuat kode yang lebih mudah dipelihara dan diperbarui.

// Untuk mengimplementasikan interfaces, kita harus membuat struct yang memiliki method yang sama dengan interface yang kita buat.
// Kemudian kita bisa membuat objek dari struct tersebut dan menggunakannya sebagai parameter untuk function yang menerima interface sebagai parameter.

// Tidak seperti bahasa perograman lain yang menerapkan interface langsung pada struct, di Go interface hanya mendefinisikan method yang harus diimplementasikan oleh struct.
// Artinya, struct bisa mengimplementasikan lebih dari satu interface, asalkan struct memiliki method yang sama dengan interface yang diimplementasikan.

// Interface juga bisa digunakan untuk membuat function yang menerima parameter dengan tipe data yang berbeda-beda.
// Hal ini bisa dilakukan dengan membuat function yang menerima parameter dengan tipe data interface.