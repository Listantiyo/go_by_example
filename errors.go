package main

import (
	"errors"
	"fmt"
)

func f(arg int)(int, error){
	if arg == 42{
		// creating new error using constructor
		return -1, errors.New("can`t work with 42")
	}
	// nil value means no error
	return arg + 3, nil
}

// A sentinel error
// It is a common pattern to create named constants for all possible sentinel errors in a package
// and use them as the value of all sentinel errors in that package.
// Sentinel error adalah variabel error yang dideklarasikan di awal (biasanya global), dan dipakai sebagai penanda kondisi error tertentu yang dapat dikenali dan dibandingkan langsung.
var ErrOutOfTea = fmt.Errorf("no more tea avaliable")
var ErrPower = fmt.Errorf("can`t boil water")

func makeTea(arg int)error{
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		// wrapping error
		// We can wrap errors with higher-level errors to add context. The simplest way to do this is with the %w verb in fmt.Errorf. Wrapped errors create a logical chain (A wraps B, which wraps C, etc.) that can be queried with functions like errors.Is and errors.As.
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _,i := range []int{7, 42} {
		if r,e := f(i); e != nil {
			fmt.Println("f failed:", e)
		}else{
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err,ErrOutOfTea){
				fmt.Println("We should buy new tea!")
			}else if errors.Is(err, ErrPower){
				fmt.Println("Now it`s dark.")
			}else{
				fmt.Println("unknown error %s/n", err)
			}
			continue
		}
		fmt.Println("Tea is ready!")
	}
}