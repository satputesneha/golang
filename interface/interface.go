package main

import "fmt"

func main() {
	s1 := samsung{length: 3, width: 2}
	fmt.Println(s1.size())
	n1 := nokia{length: 30, width: 20}
	fmt.Println(n1.size())
	var p phone
	p = &s1
	fmt.Println(p.size())
	p = &n1
	fmt.Println(p.size())
	a1 := apple{length: 50, width: 30}
	fmt.Println(a1.size())
	p = &a1

}

type phone interface {
	size() (int, int)
}
type samsung struct {
	length, width int
}

func (s *samsung) size() (int, int) {
	return s.length, s.width
}

type nokia struct {
	length, width int
}

func (s *nokia) size() (int, int) {
	return s.length, s.width
}

type apple struct {
	length, width int
}

func (a *apple) size() (int, int) {
	return a.length, a.width
}
