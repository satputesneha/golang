package main

import "fmt"

func main() {
	var m map[int]string
	fmt.Printf("type: %T,value: %v", m, m)
	m = make(map[int]string)
	m[1] = "sneha"
	m[2] = "samvega"
	fmt.Println(m)
	r := rect{length: 5, height: 4}
	fmt.Println(r)
	var r1 map[string]rect
	r1 = make(map[string]rect)
	r1["smallrect"] = rect{2, 3}
	r1["bigrect"] = rect{5, 4}
	fmt.Println(r1)

}

type rect struct {
	length int
	height int
}
