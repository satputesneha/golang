package main

import (
	"fmt"
	"strings"
)

type mystring string

func main() {
	s := wordCount("sneha hoe are you")
	fmt.Println(s)
	m := mystring("samvega how r u")
	fmt.Println(m)

}
func wordCount(s string) []string {

	return strings.Fields(s)
}
func (m mystring) wordCount1() []string {

	return strings.Fields(string(m))
}
