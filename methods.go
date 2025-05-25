package main

import "fmt"

type rect struct {
	width, height int
}

func (r * rect) area() int { // method using pointer receiver
	// This method modifies the receiver
	return r.width * r.height
}

func (r rect) perim() int  { // method using value receiver
	// This method does not modify the receiver
	return 2*r.width + 2*r.height 
}

// Perbedaan method dengan pointer receiver dan value receiver
// Pointer receiver: receiver yang di pass by reference, sehingga jika ada perubahan pada receiver, maka akan berpengaruh pada nilai asli dari receiver tersebut
// Value receiver: receiver yang di pass by value, sehingga jika ada perubahan pada receiver, maka tidak akan berpengaruh pada nilai asli dari receiver tersebut
// Pointer receiver lebih efisien untuk struct yang besar, karena tidak perlu copy struct tersebut
// Value receiver lebih efisien untuk struct yang kecil, karena tidak perlu mengubah pointer receiver tersebut

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r // go has a built-in conversion to get the pointer to the struct
	fmt.Println("area:", rp.area()) // pointer receiver
	fmt.Println("perim:", rp.perim()) // value receiver
}