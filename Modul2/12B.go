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