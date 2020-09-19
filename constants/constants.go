package main

import "fmt"

const (
	a     = 3
	big   = 1 << 60
	small = big >> 59
)

func needInt(x int) int {
	return x*10 + 1

}
func needFloat(y float64) float64 {
	return y*0.1 + 1
}

func main() {

	fmt.Println(a, needInt(small), needFloat(small))
	fmt.Println(a, needInt(big), needFloat(big))
}
