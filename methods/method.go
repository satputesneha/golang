package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

func add(x, y int) int {
	return x + y

}
func (v rect) area() int {
	return v.width * v.height
}

func main() {
	fmt.Println(add(10, 20))
	r := rect{width: 3, height: 4}
	fmt.Println(r)
	fmt.Println(r.area())
}
