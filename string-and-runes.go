// A Go string is a read-only slice of bytes. The language and the standard library treat strings specially - as containers of text encoded in UTF-8. In other languages, strings are made of “characters”. In Go, the concept of a character is called a rune - it’s an integer that represents a Unicode code point.

package main

import (
	"fmt"
	"unicode/utf8" // package to work with UTF-8 encoded strings
)

func main() {
	const s = "สวัสดี" // "Hello" in Thai

	fmt.Println("Len:", len(s)) // prints 24 bytes

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i]) // prints the byte value of each character in hex format
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s)) // prints 6 runes

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, idx)
		// Note :%#U prints the rune in Unicode format
		// and the byte position in the string
		// range loops, treathing the string as a sequence of runes
		// instead of bytes
		// The range loop iterates over the string, returning the index and the rune value
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w { // i is the index, w is the width of the rune in bytes
		runeValue, widht := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = widht // width of the rune in bytes

		examineRune(runeValue) // call the function to examine the rune
	} 
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("Found tee")
	}else if r == 'ส' {
		fmt.Println("Found so sua")
	}
}

// What is a string anf rune?
// A string is a sequence of characters represented by a sequence of bytes
// A rune is a single character in a string