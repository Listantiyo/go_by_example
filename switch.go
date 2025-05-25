package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	t := time.Now()
	switch { // Note : switch without expression equal as true
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) { // Note : interface{} is a type that can hold any value
		switch t := i.(type){ // i.( type ) is a type assertion that checks the type of i
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t) // Printf is a formatted print function that can print any type
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hello")
}