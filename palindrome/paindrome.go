package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	palindrome := 0
	y := x

	for x != 0 {

		palindrome = palindrome*10 + x%10
		x = x / 10
		fmt.Println(x, palindrome)

	}

	return y == palindrome
}
func main() {
	fmt.Println(isPalindrome(123))
}
