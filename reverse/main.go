package main

import "fmt"

func reverse(x int) int {

	result := 0
	for ; x != 0; x = x / 10 {
		fmt.Println(x, result)
		result = result*10 + x%10
	}
	return result
}
func sum(x int) int {

	result := 0
	for ; x != 0; x = x / 10 {
		fmt.Println(x, result)
		result = result + x%10
	}
	return result
}
func ones(x int) bool {

	result := 0
	for ; x != 0; x = x / 10 {
		fmt.Println(x, result)
		result = x % 10
		if result == 0 || result == 1 {
			continue
		} else {
			return false
		}
	}
	return true

}

func main() {

	fmt.Println(reverse(234))
	fmt.Println(sum(234))
	fmt.Println(ones(1010))
}
