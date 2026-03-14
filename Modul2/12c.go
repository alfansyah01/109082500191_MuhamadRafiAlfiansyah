package main

import "fmt"

func main() {
	var berat int
	fmt.Print("Berat Parsel (gram): ")
	fmt.Scan(&berat)

	kg := berat / 1000
	gr := berat % 1000

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gr)
	biayakg := kg * 10000
	biayagr := 0

	if kg >= 10 {
		biayagr = 0
	} else if gr < 500 {
		biayagr = gr * 15
	} else if gr >= 500 {
		biayagr = gr * 5
	}
	
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayakg, biayagr)
	fmt.Print("Total biaya: Rp. ", biayakg+biayagr)
}