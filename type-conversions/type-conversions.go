package main

import (
	"fmt"
	"math"
)

func main() {
	x, y := 3, 4
	fmt.Printf("type:%T,value:%v\n", x*x+y*y, x*x+y*y)
	f := math.Sqrt(float64(x*x + y*y))
	fmt.Printf("type:%T,value:%v\n", f, f)
	z := uint(f)
	fmt.Printf("type:%T,value:%v\n", z, z)

}
