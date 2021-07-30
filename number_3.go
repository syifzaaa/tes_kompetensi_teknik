package main

import (
	"fmt"
	"strings"
	"strconv"
)


func main()  {
	// Deklarasi variabel date bertipe string
	var date string
	// Baca input dari user lalu masukan ke dalam variabel date
	fmt.Scanf()
	// Deklarasi variabel dateReplaced untuk menyimpan hasil replace ':' dengan ' '
	var dateReplaced = strings.ReplaceAll(date, ":", " ")
	// Deklarasi variabel dateSliced untuk menyimpan hasi split
	/*	Hasil berupa 
	 	dateSliced[0] untuk jam
		dateSliced[1] untuk menit
		dateSliced[2] untuk detik
		dateSliced[3] untuk format
	*/
	dateSliced = strings.Split(dateReplaced, " ")
	// Cek apakah format merupakan 'AM'
	if format == AM then {
		// Gabungkan jam dan menit lalu tidak perlu menambahkan 12 jam
		result := ...
		// Print result
		...
	} 
	// Selain itu
	else {
		// Tambahkan 12 jam pada jam
		dateSliced[0] = ...
		// Cek apakah hasil jumlah jam sama dengan 24
		if dateSliced[0] == 24 {
			// Jika iya ubah menjadi jam 00
			dateSliced[0] = ...
		}
		// Gabungkan jam dan menit
		result := ... 
		// Print result
		...
	}
}