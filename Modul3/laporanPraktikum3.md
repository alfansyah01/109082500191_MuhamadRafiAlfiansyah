# <h1 align="center">Laporan Praktikum Modul 3 - FUNGSI </h1>
<p align="center">Muhamad Rafi Alfiansyah - 109082500191</p>

## Unguided 

### 1. [Soal]
#### Modul3Soal1.go

```go
package main

import "fmt"

func factorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutation(n,r int) int{
	return factorial(n) / factorial(n-r)
}

func combination(n,r int) int{
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main(){
	var a,b,c,d int
	fmt.Scan(&a,&b,&c,&d)
	fmt.Println(permutation(a,c), combination(a,c))
	fmt.Println(permutation(b,d), combination(b,d))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alfansyah01/109082500191_MuhamadRafiAlfiansyah/blob/main/Modul3/Output/SSModul3Soal1.png)
[penjelasan]
Buat menghitung permutasi dan kombinasi dari dua pasang bilangan, yaitu (a, c) dan (b, d) dengan syarat a ≥ c dan b ≥ d. Di dalam programnya ada fungsi factorial yang dipake buat menghitung nilai faktorial dengan perulangan. Terus ada fungsi permutation untuk menghitung banyaknya cara milih tanpa merhatiin urutan dengan rumus P(n, r) = n!/(n−r)!, dan fungsi combination buat menghitung banyaknya cara memilih tanpa memperhatikan urutan dengan rumus C(n, r) = n!/(r!(n−r)!). Di bagian utama, program menerima empat input, kemudian menampilkan hasil permutasi dan kombinasi dari masing-masing pasangan dalam dua baris output.

### 2. [Soal]
#### Modul3Soal2.go

```go
package main
import "fmt"

func f(x int)int {
	return x*x
}

func g(x int)int {
	return x - 2
}

func h(x int)int {
	return x + 1
}

func main(){
	var a,b,c int
	fmt.Scan(&a,&b,&c)
	hasil1 := f(g(h(a)))
	hasil2 := g(h(f(b)))
	hasil3 := h(f(g(c)))
	fmt.Println(hasil1)
	fmt.Println(hasil2)
	fmt.Println(hasil3)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alfansyah01/109082500191_MuhamadRafiAlfiansyah/blob/main/Modul3/Output/SSModul3Soal2.png)
[penjelasan]
Program ini dibuat untuk menghitung hasil dari fungsi komposisi berdasarkan tiga fungsi, yaitu f(x) = x², g(x) = x − 2, dan h(x) = x + 1. Di dalam programnya ada tiga fungsi, yaitu f, g, dan h, yang masing-masing merepresentasikan rumus tersebut. Lalu di bagian utama, program menerima tiga input bilangan a, b, dan c. Setelah itu dihitung hasil komposisi fungsi, yaitu f(g(h(a))) untuk baris pertama, g(h(f(b))) untuk baris kedua, dan h(f(g(c))) untuk baris ketiga. Hasil dari masing-masing perhitungan kemudian ditampilkan dalam tiga baris output.

### 3. [Soal]
#### Modul3Soal3.go

```go
package main
import (
	"fmt"
	"math"
)

func jarak(a,b,c,d float64)float64{
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}

func didalam(cx,cy,r,x,y float64)bool{
	return jarak(cx,cy,x,y) <= r
}

func main(){
	var cx1,cy1,r1 float64
	var cx2,cy2,r2 float64
	var x,y float64
	fmt.Scan(&cx1,&cx2,&r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)
	in1 := didalam(cx1, cy1, r1, x, y)
	in2 := didalam(cx2, cy2, r2, x, y)
	if in1 && in2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/alfansyah01/109082500191_MuhamadRafiAlfiansyah/blob/main/Modul3/Output/SSModul3Soal3.png)
[penjelasan]
Program ini dibuat buat nentuin posisi suatu titik terhadap dua lingkaran, jadi bisa kelihatan apakah titiknya ada di lingkaran 1, lingkaran 2, dua-duanya, atau malah di luar semua. Di dalamnya ada fungsi jarak(a, b, c, d) buat ngitung jarak antara dua titik (a, b) ke (c, d) pakai rumus akar kuadrat. Terus ada fungsi didalam(cx, cy, r, x, y) yang dipakai buat ngecek apakah titik (x, y) masuk ke lingkaran dengan pusat (cx, cy) dan radius r, caranya dengan ngebandingin hasil jarak ke pusat sama nilai r.

Program nerima input cx1, cy1, r1 buat lingkaran pertama, cx2, cy2, r2 buat lingkaran kedua, sama x, y sebagai titik yang mau dicek. Dari situ ditentuin apakah titiknya masuk ke lingkaran 1 (in1) atau lingkaran 2 (in2). Kalau dua-duanya masuk berarti titiknya ada di dalam kedua lingkaran, kalau cuma salah satu berarti di lingkaran itu aja, dan kalau gak masuk dua-duanya berarti titiknya ada di luar kedua lingkaran.
