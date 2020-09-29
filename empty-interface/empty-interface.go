package main

import "fmt"

func main() {
	var i interface {
	}
	intr(i)
	i = 2
	intr(i)
	k := 8
	intr(k)

}
func intr(i interface{}) {
	fmt.Printf("type:%T,value:%v", i, i)
}
