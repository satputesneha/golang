package main

import (
	"fmt"
	"math/cmplx"
)

var (
	a              = false
	max uint64     = 1<<64 - 1
	z   complex128 = cmplx.Sqrt(-5 + 12i)
	c   complex64  = complex(10, 11)
)

func main() {
	b := "type:%T,value:%v\n"
	fmt.Printf(b, a, a)
	fmt.Printf(b, max, max)
	fmt.Printf(b, z, z)
	fmt.Printf(b, c, c)
}
