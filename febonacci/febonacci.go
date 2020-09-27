package main

import "fmt"

func createFeb() func() int {
	first := 1
	isFirstTime := false
	second := 1
	isSecondTime := false

	return func() int {
		if !isFirstTime {
			isFirstTime = true
			return 1
		} else if !isSecondTime {
			isSecondTime = true
			return 1

		} else {
			result := first + second
			first = second
			second = result
			return result
		}

	}
}

func main() {
	feb := createFeb()
	for i := 0; i < 10; i++ {
		fmt.Println(feb())
	}
}
