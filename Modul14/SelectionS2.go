package main

import "fmt"

const NMAX int = 1000000

type arrInt [NMAX]int

func selectionSort(T *arrInt, n int) {
	var i, j, idxMin, temp int

	i = 1
	for i <= n-1 {
		idxMin = i - 1
		j = i

		for j < n {
			if T[idxMin] > T[j] {
				idxMin = j
			}
			j++
		}

		temp = T[idxMin]
		T[idxMin] = T[i-1]
		T[i-1] = temp

		i++
	}
}

func main() {
	var daerah, m, i, j int
	var rumah arrInt

	fmt.Scan(&daerah)

	for i = 0; i < daerah; i++ {

		fmt.Scan(&m)

		for j = 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		selectionSort(&rumah, m)

		for j = 0; j < m; j++ {
			if rumah[j]%2 != 0 {
				fmt.Print(rumah[j], " ")
			}
		}

		for j = m - 1; j >= 0; j-- {
			if rumah[j]%2 == 0 {
				fmt.Print(rumah[j], " ")
			}
		}

		fmt.Println()
	}
}