package main

import (
	"fmt"
)
func intSeq() func() int {
	i := 0
	// Fungsi closure yang mengembalikan fungsi lain
	return func() int {
		i++      // Increment i setiap kali fungsi dipanggil
		return i // Mengembalikan nilai i
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	newInts := intSeq() // Membuat fungsi baru dengan state baru
	fmt.Println(newInts()) // 1
}

// Note : Closure merupakan fungsi di dalam fungsi lain yang dapat mengakses variabel di dalam fungsi luar
// Closure juga dapat mengakses variabel di dalam fungsi luar meskipun fungsi luar sudah selesai die
// Closure mengembalikan fungsi lain yang dapat mengakses variabel di dalam fungsi luar

// Note : kenapa variable i tidak di reset ke 0 ketika fungsi intSeq() dipanggil lagi?
// Karena i adalah variabel lokal di dalam fungsi intSeq() yang hanya dapat diakses oleh fungsi closure yang dikembalikan oleh intSeq()
// Value i bisa bertambah karena lingkunan lexical yang dihasilkan oleh closure
// Lingkungan lexical ini menyimpan state dari variabel i yang diakses oleh fungsi closure
// Ketika fungsi intSeq() dipanggil lagi, fungsi closure yang dikembalikan adalah fungsi baru yang memiliki state baru