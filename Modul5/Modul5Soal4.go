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
