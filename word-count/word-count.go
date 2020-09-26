package main

import (
	"fmt"
	"strings"
)

func wordCount(s string) map[string]int {
	m := map[string]int{}
	a := strings.Fields(s)
	for _, value := range a {
		fmt.Println(value)
		_, exists := m[value]
		if exists {
			m[value]++
		} else {
			m[value] = 1
		}
	}
	return m
}

func main() {
	fmt.Println(wordCount("I am I am going to learn Go"))
}
