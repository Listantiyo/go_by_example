package main

import (
	"fmt"
	"time"
)

// Fungsi f ini akan mencetak angka dari 0 sampai 2, disertai dengan nama pemanggilnya.
// Tujuannya sederhana: menunjukkan siapa yang menjalankan, dan di urutan ke berapa.
func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Ini pemanggilan fungsi biasa, langsung.
	// Artinya: f akan dijalankan sekarang juga, dan program akan menunggu sampai selesai.
	f("direct")

	// Di sini kita memanggil fungsi f, tapi menggunakan kata kunci 'go'.
	// Artinya, kita menyuruh Go untuk menjalankan f sebagai goroutine.
	// Goroutine itu seperti "pekerja latar belakang" — dia akan jalan sendiri tanpa menunggu.
	// Karena main tidak menunggu goroutine ini selesai, kita perlu hati-hati agar dia tidak keburu "dilenyapkan" oleh program.
	go f("goroutine")

	// Ini juga goroutine, tapi bentuknya sedikit berbeda.
	// Kita buat fungsi secara langsung (anonymous function) dan langsung menjalankannya di background.
	// Tujuannya hanya mencetak satu kalimat: "going"
	go func(msg string) {
		fmt.Println(msg)
	}("going") // <-- argumen "going" dikirim ke fungsi di atas

	// Ini penting! Karena goroutine berjalan sendiri-sendiri,
	// kita kasih waktu sebentar (1 detik) agar mereka sempat menyelesaikan tugasnya.
	// Tanpa ini, program bisa keburu selesai dan goroutine tidak sempat mencetak apapun.
	time.Sleep(time.Second)

	// Setelah menunggu, kita tutup program dengan mencetak "done"
	fmt.Println("done")
}
