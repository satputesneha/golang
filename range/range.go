package main

import "fmt"

func main() {
	a := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	for i := 0; i < len(a); i++ {
		value := a[i]
		fmt.Println(i, value)
	}
	for i, value := range a {
		fmt.Println(i, value)

	}
}
