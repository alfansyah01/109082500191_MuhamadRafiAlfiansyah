// Muhamad Rafi Alfiansyah - 109082500191
package main

import "fmt"

func main() {
	var N int
	var berat [1000]float64

	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&berat[i])
	}

	minimum := berat[0]
	maksimum := berat[0]

	for i := 1; i < N; i++ {
		if berat[i] < minimum {
			minimum = berat[i]
		}

		if berat[i] > maksimum {
			maksimum = berat[i]
		}
	}

	fmt.Println("Berat terkecil =", minimum)
	fmt.Println("Berat terbesar =", maksimum)
}