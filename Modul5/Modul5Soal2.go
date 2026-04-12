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
