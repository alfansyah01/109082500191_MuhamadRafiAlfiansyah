# <h1 align="center">Laporan Praktikum Modul 10 - PENCARIAN NILAI MAX/MIN </h1>
<p align="center">Muhamad Rafi Alfiansyah - 109082500191</p>

## Unguided 

### 1. [Soal]
#### Soal1Modul10.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[penjelasan]

### 2. [Soal]
#### Soal2Modul10.go

```go
// Muhamad Rafi Alfiansyah - 109082500191
package main

import "fmt"

func main() {
	var x, y int
	var ikan [1000]float64

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	var totalSemua float64
	jumlahWadah := 0

	for i := 0; i < x; i += y {
		var totalWadah float64

		for j := i; j < i+y && j < x; j++ {
			totalWadah += ikan[j]
		}

		fmt.Printf("%.2f ", totalWadah)

		totalSemua += totalWadah
		jumlahWadah++
	}

	rataRata := totalSemua / float64(jumlahWadah)

	fmt.Printf("\n%.2f\n", rataRata)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[penjelasan]

### 3. [Soal]
#### Soal3Modul10.go

```go
// Muhamad Rafi Alfiansyah - 109082500191
package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}

		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64

	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}

	return total / float64(n)
}

func main() {
	var arrBerat arrBalita
	var n int
	var bMin, bMax float64

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&arrBerat[i])
	}

	hitungMinMax(arrBerat, n, &bMin, &bMax)

	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerata(arrBerat, n))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[penjelasan]