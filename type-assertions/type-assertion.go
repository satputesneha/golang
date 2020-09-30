package main

import "fmt"

func main() {
	var i interface {
	}
	i = 5
	fmt.Println(i)
	fmt.Println(i.(string))
	fmt.Println(i.(int))

}
