package main

import "fmt"

func main() {
	var i interface {
	}
	i = 5
	fmt.Println(i)
	value, ok := i.(string)
	fmt.Println(value, ok)
	if ok {
		fmt.Println(i.(string))
	}
	fmt.Println(i.(int))
	fmt.Println("the end")

}
