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
	switch i.(type) {
	case int:
		fmt.Println(i.(int))
	case string:
		fmt.Println(i.(string))

	}
	fmt.Println("the end")

}
