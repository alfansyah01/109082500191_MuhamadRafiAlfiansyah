# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">Muhamad Rafi Alfiansyah - 109082500191</p>

## Unguided 

### 1. SoalA
#### 12A.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}

```

### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[penjelasan]

### 2. SoalB
#### 12B.go

```go
package main
import "fmt"
func main(){
	var m,k,h,u string
	suc := true
	i := 1
	for i <= 5 {
		fmt.Printf("Pecobaan %d : ",i)
		fmt.Scan(&m,&k,&h,&u)
		if m != "merah" || k != "kuning" || h != "hijau" || u != "ungu"{
			suc = false
		}
		i++
	}
	fmt.Println("BERHASIL:",suc)
}

```

### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[penjelasan]

### 3. SoalC
#### 12c.go

```go
package main

import "fmt"

func main() {
	var berat int
	fmt.Print("Berat Parsel (gram): ")
	fmt.Scan(&berat)

	kg := berat / 1000
	gr := berat % 1000

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gr)
	biayakg := kg * 10000
	biayagr := 0

	if kg >= 10 {
		biayagr = 0
	} else if gr < 500 {
		biayagr = gr * 15
	} else if gr >= 500 {
		biayagr = gr * 5
	}
	
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayakg, biayagr)
	fmt.Print("Total biaya: Rp. ", biayakg+biayagr)
}

```

### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[penjelasan]
