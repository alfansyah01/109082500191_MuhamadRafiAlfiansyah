package main

import "fmt"

const NMAX int = 1000

type arrInt [NMAX]int

func insertionSort(T *arrInt, n int) {
	var i, j, temp int

	i = 1
	for i <= n-1 {
		j = i
		temp = T[j]

		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j--
		}

		T[j] = temp
		i++
	}
}

func main() {
	var A arrInt
	var x, n, i int
	var tetap bool
	var selisih int

	n = 0

	fmt.Scan(&x)

	for x >= 0 {
		A[n] = x
		n++

		fmt.Scan(&x)
	}

	insertionSort(&A, n)

	for i = 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()

	if n <= 1 {
		fmt.Println("Data berjarak 0")
	} else {
		selisih = A[1] - A[0]
		tetap = true

		for i = 2; i < n && tetap; i++ {
			if A[i]-A[i-1] != selisih {
				tetap = false
			}
		}

		if tetap {
			fmt.Println("Data berjarak", selisih)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}