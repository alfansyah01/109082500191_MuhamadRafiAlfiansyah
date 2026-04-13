# <h1 align="center">Laporan Praktikum Modul 5 - REKURSIF </h1>
<p align="center">Muhamad Rafi Alfiansyah - 109082500191</p>

## Unguided 

### 1. [Soal]
#### Modul5Soal1.go

```go
package main
import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i <= n; i++ {
		fmt.Print(fibonacci(i), " ")
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alfansyah01/109082500191_MuhamadRafiAlfiansyah/blob/main/Modul5/Output/Soal1.png)
[penjelasan]
Program ini digunakan untuk menampilkan deret Fibonacci menggunakan rekursif.
Fungsi fibonacci(n) akan memanggil dirinya sendiri untuk menghitung nilai Fibonacci ke-n dengan rumus fibonacci(n-1) + fibonacci(n-2). Kondisi berhentinya ada pada saat n = 0 dan n = 1. Di main, program membaca input n lalu mencetak deret Fibonacci dari 0 sampai n.

### 2. [Soal]
#### Modul5Soal2.go

```go
package main
import "fmt"

func pola(i,n int, b string){
	if i>n{
		return
	}
	b += "*"
	fmt.Println(b)
	pola(i+1,n,b)
}

func main(){
	var n int
	fmt.Scan(&n)
	pola(1,n,"")
	}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alfansyah01/109082500191_MuhamadRafiAlfiansyah/blob/main/Modul5/Output/Soal2.png)
[penjelasan]
Program ini digunakan untuk mencetak pola bintang bertambah menggunakan rekursif.
Fungsi pola(i, n, b) akan menambahkan satu karakter * ke string b, lalu mencetaknya. Setelah itu fungsi memanggil dirinya lagi dengan i+1 sampai nilai i lebih besar dari n. Hasilnya akan membentuk pola bintang yang semakin panjang di setiap baris.

### 3. [Soal]
#### Modul5Soal3.go

```go
package main
import "fmt"

func faktor(i,n int){
	if i>n{
		return
	}
	if n%i==0{
		fmt.Print(i," ")
	}
	faktor(i+1,n)

}
func main(){
	var n int
	fmt.Scan(&n)
	faktor(1,n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alfansyah01/109082500191_MuhamadRafiAlfiansyah/blob/main/Modul5/Output/Soal3.png)
[penjelasan]
Program ini digunakan untuk mencari faktor dari sebuah bilangan dengan rekursif.
Fungsi faktor(i, n) akan mengecek apakah n habis dibagi i. Jika iya maka i dicetak sebagai faktor. Setelah itu fungsi memanggil dirinya lagi dengan i+1 sampai i lebih besar dari n. Dengan cara ini semua faktor dari bilangan tersebut bisa ditampilkan.

### 4. [Soal]
#### Modul5Soal4.go

```go
package main
import "fmt"

func barisan(i int){
	if i == 0 {
		return
	}

	fmt.Print(i," ")
	barisan(i-1)
	if i!=1{
		fmt.Print(i," ")
	}
}

func main(){
	var n int
	fmt.Scan(&n)
	barisan(n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alfansyah01/109082500191_MuhamadRafiAlfiansyah/blob/main/Modul5/Output/Soal4.png)
[penjelasan]
Program ini menampilkan barisan angka turun lalu naik lagi menggunakan rekursif.
Fungsi barisan(i) pertama akan mencetak nilai i, lalu memanggil dirinya dengan i-1 sampai mencapai 0. Setelah proses rekursif selesai, angka akan dicetak lagi kecuali angka 1, sehingga membentuk pola seperti turun lalu kembali naik.

### 5. [Soal]
#### Modul5Soal5.go

```go
package main
import "fmt"

func ganjil(i,n int){
	if i>n{
		return
	}
	if i%2!=0{
		fmt.Print(i," ")
	}
	ganjil(i+1,n)
}

func main(){
	var n int
	fmt.Scan(&n)
	ganjil(1,n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alfansyah01/109082500191_MuhamadRafiAlfiansyah/blob/main/Modul5/Output/Soal5.png)
[penjelasan]
Program ini digunakan untuk mencetak bilangan ganjil dari 1 sampai n menggunakan rekursif.
Fungsi ganjil(i, n) akan mengecek apakah i adalah bilangan ganjil (i % 2 != 0). Jika iya maka angka tersebut dicetak. Setelah itu fungsi memanggil dirinya lagi dengan i+1 sampai i lebih besar dari n.

### 6. [Soal]
#### Modul5Soal6.go

```go
package main
import "fmt"

func pangkat(x, y int) int {
	if y == 0 { 
		return 1
	}

	return x * pangkat(x, y-1)
}

func main(){
	var x, y int
	fmt.Scan(&x, &y)

	fmt.Print(pangkat(x, y))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alfansyah01/109082500191_MuhamadRafiAlfiansyah/blob/main/Modul5/Output/Soal6.png)
[penjelasan]
Program ini digunakan untuk menghitung pangkat sebuah bilangan menggunakan rekursif.
Fungsi pangkat(x, y) akan mengalikan x dengan hasil pemanggilan fungsi pangkat(x, y-1) sampai y bernilai 0. Jika y = 0, fungsi langsung mengembalikan nilai 1 sebagai kondisi berhenti. Dengan cara ini program bisa menghitung hasil x^y.
