package main

import "fmt"

func main() {
	var a [10]int
	a[0] = 1
	a[5] = 3
	for i := 0; i < 10; i++ {
		a[i] = 2 * i
		fmt.Println(i, a[i])
	}
	p := &a

	fmt.Printf("%T ,%v, %v,%v", a, a, len(a), p[3])

}
