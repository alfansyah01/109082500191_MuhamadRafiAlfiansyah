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
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[penjelasan]

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
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[penjelasan]

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
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[penjelasan]