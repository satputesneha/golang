package main

import "fmt"

func main() {
	//slice:= []string{
	//[]string{"one"}
	s := []int{0, 1, 2, 3, 4}
	fmt.Printf("%T,%v\n", s, s)
	s1 := []string{"one", "two", "three"}
	fmt.Println(s1)
	s2 := []string{"a", "b", "c"}
	fmt.Println(s2)
	m := [][]string{s1, s2, []string{"slice"}, s2}

	s2[1] = "sam"
	m[2][0] = "vega"
	fmt.Println(m)
	n1 := make([]string, 3)
	n1[0] = "one"
	n1[1] = "two"
	n1[2] = "three"
	fmt.Println(n1)
	n2 := make([]string, len(s2))
	n2[0] = s2[0]
	n2[1] = s2[1]
	n2[2] = s2[2]

	fmt.Println(n2)
	n := make([][]string, len(m))
	for i := 0; i < len(m); i++ {
		n[i] = make([]string, len(m[i]))
		for j := 0; j < len(m[i]); j++ {
			n[i][j] = m[i][j]
		}
	}
	fmt.Println(n)
}
