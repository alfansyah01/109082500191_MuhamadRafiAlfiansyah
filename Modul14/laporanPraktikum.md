# <h1 align="center">Laporan Praktikum Modul 12 - SEARCHING </h1>
<p align="center">Muhamad Rafi Alfiansyah - 109082500191</p>

## Unguided 

### 1. Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu. Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort. Masukan dimulai dengan sebuah integer 𝒏 (0 < n < 1000), banyaknya daerah kerabat Hercules tinggal. Isi 𝒏 baris berikutnya selalu dimulai dengan sebuah integer 𝒎 (0 < m < 1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan rangkaian bilangan bulat positif, nomor rumah para kerabat. Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar di masing-masing daerah.
#### soal1.go

```go
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
			fmt.Print(rumah[j], " ")
		}
		fmt.Println()
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alfansyah01/109082500191_MuhamadRafiAlfiansyah/blob/main/Modul14/Output/SS1.png)
[penjelasan]
Program ini digunakan untuk mengurutkan nomor rumah kerabat Hercules di setiap daerah secara menaik (ascending) menggunakan algoritma Selection Sort. Algoritma bekerja dengan mencari nilai terkecil dari bagian array yang belum terurut, kemudian menukarnya dengan elemen pada posisi paling depan. Proses ini diulang hingga seluruh data tersusun dari nomor rumah terkecil hingga terbesar sesuai dengan ketentuan soal.

### 2. Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil. Format Masukan masih persis sama seperti sebelumnya. Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah.
#### soal1.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alfansyah01/109082500191_MuhamadRafiAlfiansyah/blob/main/Modul14/Output/SS2.png)
[penjelasan]
Program ini memanfaatkan Selection Sort untuk mengurutkan seluruh nomor rumah secara menaik terlebih dahulu. Setelah data terurut, program menampilkan semua nomor rumah ganjil dalam urutan membesar, kemudian menampilkan nomor rumah genap dalam urutan mengecil. Dengan cara ini, rumah di sisi kiri jalan (ganjil) dikunjungi lebih dahulu, lalu dilanjutkan dengan rumah di sisi kanan jalan (genap) sesuai aturan pada soal.

### 3. Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya. Masukan terdiri dari sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya bilangan non negatif saja yang disimpan ke dalam array. Keluaran terdiri dari dua baris. Baris pertama adalah isi dari array setelah dilakukan pengurutan, sedangkan baris kedua adalah status jarak setiap bilangan yang ada di dalam array. "Data berjarak x" atau "data berjarak tidak tetap".
#### soal1.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alfansyah01/109082500191_MuhamadRafiAlfiansyah/blob/main/Modul14/Output/IS1.png)
[penjelasan]
Program ini membaca sekumpulan bilangan bulat non-negatif hingga ditemukan bilangan negatif sebagai penanda akhir input. Data yang telah dibaca kemudian diurutkan menggunakan algoritma Insertion Sort, yaitu dengan menyisipkan setiap elemen ke posisi yang sesuai pada bagian array yang sudah terurut. Setelah proses pengurutan selesai, program memeriksa apakah selisih antar elemen yang berurutan selalu sama. Jika sama, program menampilkan nilai jaraknya, sedangkan jika berbeda maka ditampilkan informasi bahwa data berjarak tidak tetap.

### 4. Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya. Masukan terdiri dari sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya bilangan non negatif saja yang disimpan ke dalam array. Keluaran terdiri dari dua baris. Baris pertama adalah isi dari array setelah dilakukan pengurutan, sedangkan baris kedua adalah status jarak setiap bilangan yang ada di dalam array. "Data berjarak x" atau "data berjarak tidak tetap".
#### soal1.go

```go
package main

import "fmt"

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	var i int

	fmt.Scan(n)

	for i = 0; i < *n; i++ {
		fmt.Scan(
			&pustaka[i].id,
			&pustaka[i].judul,
			&pustaka[i].penulis,
			&pustaka[i].penerbit,
			&pustaka[i].eksemplar,
			&pustaka[i].tahun,
			&pustaka[i].rating,
		)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	var i, idxMax int

	idxMax = 0

	for i = 1; i < n; i++ {
		if pustaka[i].rating > pustaka[idxMax].rating {
			idxMax = i
		}
	}

	fmt.Println("Buku Terfavorit")
	fmt.Println("Judul    :", pustaka[idxMax].judul)
	fmt.Println("Penulis  :", pustaka[idxMax].penulis)
	fmt.Println("Penerbit :", pustaka[idxMax].penerbit)
	fmt.Println("Tahun    :", pustaka[idxMax].tahun)
	fmt.Println()
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var i, j int
	var temp Buku

	i = 1

	for i <= n-1 {
		j = i
		temp = pustaka[j]

		for j > 0 && temp.rating > pustaka[j-1].rating {
			pustaka[j] = pustaka[j-1]
			j--
		}

		pustaka[j] = temp
		i++
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	var i, batas int

	if n < 5 {
		batas = n
	} else {
		batas = 5
	}

	fmt.Println("5 Buku Rating Tertinggi")

	for i = 0; i < batas; i++ {
		fmt.Println(pustaka[i].judul)
	}
	fmt.Println()
}

func CariBuku(pustaka DaftarBuku, n, r int) {
	var kiri, kanan, tengah int
	var ketemu bool

	kiri = 0
	kanan = n - 1
	ketemu = false

	for kiri <= kanan && !ketemu {

		tengah = (kiri + kanan) / 2

		if pustaka[tengah].rating == r {
			ketemu = true
		} else if r > pustaka[tengah].rating {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if ketemu {
		fmt.Println("Data Buku Ditemukan")
		fmt.Println("Judul     :", pustaka[tengah].judul)
		fmt.Println("Penulis   :", pustaka[tengah].penulis)
		fmt.Println("Penerbit  :", pustaka[tengah].penerbit)
		fmt.Println("Tahun     :", pustaka[tengah].tahun)
		fmt.Println("Eksemplar :", pustaka[tengah].eksemplar)
		fmt.Println("Rating    :", pustaka[tengah].rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var pustaka DaftarBuku
	var n, ratingCari int

	DaftarkanBuku(&pustaka, &n)

	CetakTerfavorit(pustaka, n)

	UrutBuku(&pustaka, n)

	Cetak5Terbaru(pustaka, n)

	fmt.Scan(&ratingCari)

	CariBuku(pustaka, n, ratingCari)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alfansyah01/109082500191_MuhamadRafiAlfiansyah/blob/main/Modul14/Output/IS1.png)
![Screenshot Output Unguided 1_1](https://github.com/alfansyah01/109082500191_MuhamadRafiAlfiansyah/blob/main/Modul14/Output/IS%232.png)
[penjelasan]
Program ini mengelola data buku perpustakaan menggunakan struktur data struct. Pertama, program mencari buku dengan rating tertinggi sebagai buku terfavorit. Selanjutnya data buku diurutkan secara menurun (descending) berdasarkan rating menggunakan algoritma Insertion Sort, sehingga buku dengan rating tertinggi berada di urutan awal. Setelah itu program menampilkan lima buku dengan rating tertinggi dan melakukan pencarian buku berdasarkan rating yang dimasukkan pengguna menggunakan metode Binary Search, sehingga proses pencarian menjadi lebih cepat dan efisien pada data yang sudah terurut.