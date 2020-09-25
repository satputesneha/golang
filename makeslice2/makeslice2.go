package main

import (
	"fmt"
	"strings"
)

func main() {

	s1 := []int{2, 3, 4}
	fmt.Println(s1)
	s2 := []int{5, 6, 7}
	fmt.Println(s2)
	s3 := []int{8, 9, 10}
	fmt.Println(s3)
	m := [][]int{s1, s2, s3}
	fmt.Println(m)
	n := make([][]int, len(m))
	//for i := 0; i < len(m); i++ {
	for i, value := range m {
		n[i] = make([]int, len(value))
		for j := 0; j < len(value); j++ {
			n[i][j] = m[i][j]

		}

	}
	n[1] = append(n[1], 99)
	n[1] = append(n[1], 100)
	fmt.Println(n[1])

	fmt.Println(n)
	r := strings.Join([]string{"s", "t", "u"}, "_")
	fmt.Println(r)
	fmt.Println(n[2])

}
