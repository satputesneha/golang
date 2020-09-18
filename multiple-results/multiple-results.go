package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	a, b := swap("hello", "samvega")
	fmt.Println(a, b)
	fmt.Println(swap("hello", "world"))
}
