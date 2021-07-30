package main

import "fmt"

func main() {
	// Deklarasi variabel number bertipe integer
 	var number int
	// Baca input dari user lalu masukan ke dalam variabel number
	fmt.Scanf(number)
	// Cek apakah number habis dibagi 3 dan habis dibagi 5, 
	if number % 3 == 0 then {
		// Jika iya tampilkan "Hello World"
		fmt.Print("Hello World")
	}
	// Cek apakah number habis dibagi 3,
	if number % 3 != 0 then {
		// Jika iya tampilkan "Hello"
		fmt.Print("Hello")
	}
	// Cek apakah number habis dibagi 5,
	if number % 5 == 0 then {
		// Jika iya tampilkan "World"
		fmt.Print("World")
	}
}