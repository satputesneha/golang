package main

import "fmt"

func main() {
	r := rect{width: 3, height: 10}
	fmt.Println(r)
	fmt.Printf("the rect width:%v  height:%v", r.width, r.height)
	r1 := rect{width: 30, height: 40}
	fmt.Println(r1)
	fmt.Println(r)
	fmt.Println(r1)
	fmt.Println(r)
}

type rect struct {
	width, height int
}

func (r rect) String() string {
	return fmt.Sprintf("the rect width:%v  height:%v", r.width, r.height)

}
