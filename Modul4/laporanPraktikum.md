# <h1 align="center">Laporan Praktikum Modul 4 - Prosedur </h1>
<p align="center">Muhamad Rafi Alfiansyah - 109082500191</p>

## Unguided 

### 1. [Soal]
#### Modul4Soal1.go

```go
package main

import "fmt"

func factorial(n int, hasil *int64) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= int64(i)
	}
}

func permutation(n, r int, hasil *int64) {
	var fn, fn_r int64

	factorial(n, &fn)
	factorial(n-r, &fn_r)

	*hasil = fn / fn_r
}

func combination(n, r int, hasil *int64) {
	var fn, fr, fn_r int64

	factorial(n, &fn)
	factorial(r, &fr)
	factorial(n-r, &fn_r)

	*hasil = fn / (fr * fn_r)
}

func main() {
	var a, b, c, d int
	var p, k int64

	fmt.Scan(&a, &b, &c, &d)
	permutation(a, c, &p)
	combination(a, c, &k)
	fmt.Println(p, k)

	permutation(b, d, &p)
	combination(b, d, &k)
	fmt.Println(p, k)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[penjelasan]

### 2. [Soal]
#### Modul4Soal2.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[penjelasan]

### 3. [Soal]
#### Modul4Soal3.go

```go
ppackage main

import "fmt"

func cetakDeret(n int) {
	for {
		fmt.Print(n)

		if n == 1 {
			break
		}

		fmt.Print(" ")

		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	cetakDeret(n)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[penjelasan]

