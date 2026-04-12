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
![Screenshot Output Unguided 1_1](https://github.com/alfansyah01/109082500191_MuhamadRafiAlfiansyah/blob/main/Modul4/Output/Soal%201.png)
[penjelasan]
Program untuk menghitung permutasi dan kombinasi dari dua pasang angka yang diinput yaitu (a, c) dan (b, d). Perhitungannya dipisah jadi beberapa bagian biar lebih rapi, yaitu lewat prosedur factorial, permutation, dan combination. Prosedur factorial dipakai buat nyari nilai faktorial dengan cara mengalikan angka dari 1 sampai n. Hasil dari faktorial ini nanti dipakai lagi di prosedur permutation dan combination untuk dapetin hasil akhirnya. Intinya itu permutasi menghitung banyaknya susunan dengan memperhatikan urutan sedangkan kombinasi itu cuma memilih tanpa peduli urutan.

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
![Screenshot Output Unguided 1_1](https://github.com/alfansyah01/109082500191_MuhamadRafiAlfiansyah/blob/main/Modul4/Output/Soal%202.png)
[penjelasan]
Program ini intinya dipakai buat nentuin siapa pemenang dari beberapa peserta lomba. Setiap peserta ngerjain 8 soal, terus tiap soal ada waktu pengerjaannya. Kalau waktunya 301 berarti soal itu dianggap nggak selesai, jadi nggak dihitung. Nah, biar rapi, dibuat prosedur hitungSkor yang tugasnya baca 8 waktu tadi, terus ngitung berapa soal yang berhasil diselesaiin dan total waktunya. Jadi di bagian utama program cuma fokus baca nama peserta sama nentuin pemenangnya, sedangkan perhitungan skor diserahin ke prosedur itu.

Cara nentuin pemenangnya yang soal berhasilnya paling banyak dia menang. Kalau ternyata jumlah soalnya sama, baru dibandingin total waktunya, yang lebih kecil berarti lebih cepat jadi dia yang menang. Setiap ada peserta baru, datanya dibandingin sama pemenang sementara, jadi nanti terus di-update sampai ketemu yang paling bagus. Pas input "Selesai", program langsung nampilin hasil akhirnya, yaitu nama pemenang, jumlah soal yang dia selesain, sama total waktunya.

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
![Screenshot Output Unguided 1_1](https://github.com/alfansyah01/109082500191_MuhamadRafiAlfiansyah/blob/main/Modul4/Output/Soal%203.png)
[penjelasan]
Program ini buat nampilin deret angka dari satu angka awal yang kita masukin. Cara kerjanya tuh gini, kalau angkanya genap ya dibagi 2, kalau ganjil dikali 3 terus ditambah 1. Nah hasilnya itu dipakai lagi terus diulang-ulang sampai akhirnya dapet angka 1. Biar rapi, bagian yang ngurus deretnya dipisah ke prosedur cetakDeret, jadi di main cuma tinggal masukin angka terus manggil prosedurnya aja.

Di dalam cetakDeret program bakal nge-print angka satu per satu sambil terus ngubah nilainya sesuai aturan tadi. Setiap angka dipisahin pake spasi biar sesuai sama yang diminta. Loop-nya jalan terus sampai ketemu 1, baru berhenti. Jadi intinya program ini cuma muter angka dengan aturan tertentu sampai habis, terus nampilin semua langkahnya.