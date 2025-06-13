package main

import (
	"errors"
	"fmt"
)

// Umumnya struct untuk membuat custom error pakai prefix "Error"
type argError struct {
	arg int
	message string
}

// Ketika sebuah struct memiliki method Error() maka struct tersebut akan menjadi sebuah error
// penambahan method Error() akan otomatis membuat struct mengimplementasikan interface error
func (arg *argError) Error() string {
	return fmt.Sprintf("%d - %s", arg.arg, arg.message)
}

func f(arg int) (int, error) {
	if arg == 42 {
		// errors.New() akan membuat sebuah error baru
		// return -1, errors.New("can't work with 42")
		// &argError akan memiliki perilaku yang sama dengan errors.New()
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	_, err := f(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	}else{
		fmt.Println("err doesn't math argError")
	}
}