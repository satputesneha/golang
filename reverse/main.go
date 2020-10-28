package main

import "fmt"

func reverse(x int) int {
	result := 0
	for ; x != 0; x = x / 10 {
		result = result*10 + x%10
	}
	return result
}
func main() {

	fmt.Println(reverse(234))
}
