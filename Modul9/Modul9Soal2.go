package main
import (
	"fmt"
	"math"
)
func main() {
	var n int
	fmt.Print("Jumlah data: ")
	fmt.Scan(&n)
	var arr [100]int
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println("Semua data:")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println("\nIndeks ganjil:")
	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println("\nIndeks genap:")
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	var x int
	fmt.Print("\nMasukkan x: ")
	fmt.Scan(&x)
	fmt.Println("Kelipatan x:")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	var idx int
	fmt.Print("\nIndex yang dihapus: ")
	fmt.Scan(&idx)
	for i := idx; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	n--
	fmt.Println("Setelah dihapus:")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
	}
	rata := float64(sum) / float64(n)
	fmt.Println("\nRata-rata:", rata)
	var total float64
	for i := 0; i < n; i++ {
		total += math.Pow(float64(arr[i])-rata, 2)
	}
	std := math.Sqrt(total / float64(n))
	fmt.Println("Standar deviasi:", std)
	var cari, freq int
	fmt.Print("Cari angka: ")
	fmt.Scan(&cari)
	for i := 0; i < n; i++ {
		if arr[i] == cari {
			freq++
		}
	}
	fmt.Println("Frekuensi:", freq)
}