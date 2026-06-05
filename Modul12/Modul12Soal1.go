package main

import "fmt"

func main() {
	var suara [21]int
	var x int
	var masuk, sah int

	for {
		fmt.Scan(&x)
		masuk++

		if x == 0 {
			break
		}

		if x >= 1 && x <= 20 {
			suara[x]++
			sah++
		}
	}

	fmt.Println("Suara masuk:", masuk)
	fmt.Println("Suara sah:", sah)

	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}