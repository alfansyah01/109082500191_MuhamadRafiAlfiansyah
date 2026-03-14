# <h1 align="center">Laporan Praktikum Modul 2 - REVIEW ALGORITMA DAN PEMROGRAMAN 1 </h1>
<p align="center">Muhamad Rafi Alfiansyah - 109082500191</p>

## unguided 

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

### Output unguided :

##### Output 
![Screenshot Output unguided A_1](https://github.com/alfansyah01/109082500191_MuhamadRafiAlfiansyah/blob/main/Modul2/Output/Output_SoalA.png)
### Penjelasan :
 Program di soal itu bertujuan untuk menerima tiga buah input string dari user. User diminta memasukkan tiga string yang kemudian disimpan ke dalam tiga variabel berbeda, yaitu variabel satu, dua, dan tiga. Setelah ketiga input dimasukkan, program akan tampilkan urutan awal dari ketiga string tersebut menggunakan perintah output sehingga user dapat melihat urutan string sesuai dengan yang pertama kali diinputkan.

 Setelah menampilkan urutan awal, program kemudian melakukan pertukaran posisi nilai pada ketiga variabel tersebut dengan menggunakan variabel sementara yang bernama temp. Pertama nilai dari variabel satu disimpan ke dalam temp, kemudian nilai dua dipindahkan ke satu, dan nilai tiga dipindahkan ke dua. Lalu nilai yang disimpan di temp dimasukkan ke dalam tiga. Dari proses ini urutan string mengalami pergeseran posisi, kemudian program menampilkan kembali hasilnya sebagai output akhir yang menunjukkan bahwa urutan string telah berubah dari urutan awal.

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
![Screenshot Output unguided B_1](https://github.com/alfansyah01/109082500191_MuhamadRafiAlfiansyah/blob/main/Modul2/Output/Output_SoalB.png)
### Penjelasan :
Program ini bertujuan untuk meminta user melakukan input sebanyak lima kali percobaan. Pada tiap percobaannya user diminta memasukkan empat buah string yang disimpan ke dalam variabel m, k, h, dan u (nama varibel berupa inisal dari warna merah, kuning, hijau dan ungu). Kemudian program mengecek apakah nilai yang dimasukkan sesuai dengan urutan warna yang sudah ditentukan, yaitu merah, kuning, hijau, dan ungu. Di awal program, variabel suc (nama variabel singkatan dari kata succeed yang dimaksudkan untuk menentukan keberhasilannya false atau true) diberi nilai true yang menandakan bahwa percobaan dianggap berhasil jika semua input sesuai dengan yang diminta.

Selama proses lima kali percobaan tersebut, program akan memeriksa tiap input menggunakan kondisi if. Jika ada satu saja nilai yang tidak sesuai dengan urutan warna, maka variabel suc akan berubah jadi false. Artinya percobaan dianggap gagal karena ada input yang berbeda dari yang ditentukan. Setelah lima percobaan selesai dilakukan, program akan menampilkan gasil akhir berupa tulisan BERHASIL diikitu nilai dari variabel suc, yang menunjukkan apakah seluruh input yang dimasukkan benar atau ada yang salah.

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
![Screenshot Output unguided 1_1](https://github.com/alfansyah01/109082500191_MuhamadRafiAlfiansyah/blob/main/Modul2/Output/Output_SoalC.png)
### Penjelasan :
Program ini dibuat untuk menghitung biaya pengiriman parsel berdasarkan berat yang dimasukkan oleh user. Pertama, program meminta pengguna memasukkan berat parsel dalam satuan gram melalui variabel berat. Setelah berat dimasukkan, program membagi berat tersebut menjadi dua bagian, yaitu kilogram (kg) dan sisa gram (gr). Nilai kilogram didapat dari hasil pembagian berat dengan 1000, sedangkan sisa gram didapat dari modulus. Setelah itu program menampilkan detail berat parsel dalam bentuk kilogram dan gram.

Selanjutnya program menghitung biaya pengirimian. Biaya untuk setiap kilogram dihitung dengan rumus kg x 10000 dan disimpan dalam variabel biayakg. Kemudian program menentukan biaya tambahan dari  sisa gram yang disimpan dalam variabel biayagr. Jika berat parsel 10 kg atau lebih, maka sisa gram digratiskan sehingga biayagr bernilai 0. Namun jika beratnya kurang dari 10 kg, maka biaya gram dihitung berdasarkan jumlahnya: jika sisa gram kurang dari 500 gram dikennakan biaya Rp15 per gram, sedangkan jika sisa gram 500 gram atau lebih dikenakan biaya Rp5 per gram. Setelah semua biaya dihitung, program menampilkan detail biaya per kilogram, biaya tambahan dari gram, serta total biaya pengiriman yang merupakan penjumlahan dari keduanya.
