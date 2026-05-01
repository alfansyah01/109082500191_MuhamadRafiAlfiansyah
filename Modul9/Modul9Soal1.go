package main
import (
	"fmt"
	"math"
)

type titik struct{
	x,y int
}

type lingkaran struct{
	center titik
	radius int
}

func jarak(p,q titik) float64{
	return math.Sqrt(math.Pow(float64(p.x-q.x),2) + math.Pow(float64(p.y-q.y),2))
}

func didalam(c lingkaran,p titik) bool{
	return jarak(c.center, p) <= float64(c.radius)
}

func main(){
	var c1, c2 lingkaran
	var p titik

	fmt.Scan(&c1.center.x, &c1.center.y, &c1.radius)
	fmt.Scan(&c2.center.x, &c2.center.y, &c2.radius)
	fmt.Scan(&p.x, &p.y)

	in1 := didalam(c1, p)
	in2 := didalam(c2, p)

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