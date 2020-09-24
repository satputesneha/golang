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
	var b []int = a[1:]
	c := a[3:5]
	c[0] = 9
	fmt.Printf("b %T,%v\n", b, b)

	fmt.Printf("c %T,%v\n", c, c)

	fmt.Printf("%T ,%v, %v,%v", a, a, len(a), p[3])

}
