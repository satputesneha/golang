package main

import (
	"fmt"
)

func main() {
	fmt.Println(sqrt(50000))

}
func sqrt(x float64) float64 {
	z := float64(1)
	var prevz float64
	for prevz != z {
		prevz = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z, prevz)
	}
	return z

}
