package main

import "fmt"

func main() {
	s := func(x, y int) int {
		return (x + y)

	}
	fmt.Println(s(3, 4))
	fmt.Println(s(3, 4))
	fmt.Println(sub(s))
}
func sub(f func(a, b int) int) int {
	return f(4, 5)

}
