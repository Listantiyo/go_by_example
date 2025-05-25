package main

import "fmt"

// Structs di go adalah tipe data yang digunakan untuk mengelompokkan data yang berbeda menjadi satu kesatuan. Structs mirip dengan class di OOP, tetapi lebih sederhana dan tidak memiliki metode.
// Structs dapat memiliki field dengan tipe data yang berbeda, dan kita dapat membuat instance dari struct tersebut untuk menyimpan data yang sesuai dengan field yang ada di dalamnya.

type person struct {
	name string
	age int
}

func newPerson(name string) *person {
	p := person{name:name} // variable p bertipe person
	p.age = 42
	return &p // return pointer to struct
}

func main() {
	
	fmt.Println(person{"Bob", 26}) // creating struct with values literal
	
	fmt.Println(person{name:"Alice", age: 30}) // creating struct with named fields
	
	fmt.Println(person{name:"Fred"}) // creating struct with named fields, age will be 0

	fmt.Println(&person{name:"Ann", age: 40}) // using & to get pointer to struct

	fmt.Println(newPerson("Jon")) // using function to create struct with pointer, this idiomatic in go

	s := person{name:"Sean", age:50}
	fmt.Println(s.name) // accessing field of struct

	sp := &s // pointer to struct
	fmt.Println(sp.age) // accessing field of struct using pointer

	sp.age = 51 // changing field of struct using pointer
	fmt.Println(sp.age) // accessing field of struct using pointer

	// anonymous struct, this is not idiomatic in go, only use this if you need to create a struct for a short time and don't want to create a new type
	dog := struct{  
		name string
		isGood bool
	}{
		name:"Rex",
		isGood:true,
	}
	fmt.Println(dog) // creating struct with anonymous struct, this is not idiomatic in go
}