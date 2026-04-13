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
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[penjelasan]

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
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[penjelasan]

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
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[penjelasan]

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
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[penjelasan]

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
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[penjelasan]

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
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[penjelasan]

