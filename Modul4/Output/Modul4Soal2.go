package main

import "fmt"

func hitungSkor(soal *int, skor *int) {
	*soal = 0
	*skor = 0

	var waktu int

	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)

		if waktu <= 300 {
			*soal += 1
			*skor += waktu
		}
	}
}

func main() {
	var nama string

	var pemenang string
	var maxSoal, minSkor int
	var soal, skor int

	first := true

	for {
		fmt.Scan(&nama)

		if nama == "Selesai" {
			break
		}

		hitungSkor(&soal, &skor)

		// tentukan pemenang
		if first || soal > maxSoal || (soal == maxSoal && skor < minSkor) {
			pemenang = nama
			maxSoal = soal
			minSkor = skor
			first = false
		}
	}

	fmt.Println(pemenang, maxSoal, minSkor)
}
