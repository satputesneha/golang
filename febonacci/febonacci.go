package main

import "fmt"

func createFeb() func() int {
	isFistTime := false
	isSecondTime := false
	first := 1
	second := 1
	return func() int {
		if !isFistTime {
			isFistTime = true
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
