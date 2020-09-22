package main

import "fmt"

func main() {
	i := 30
	j := &i
	fmt.Println(i, j, *j)
	*j = 31
	fmt.Println(i, j, *j)
	*j = *j / 2
	fmt.Println(i, j, *j)

}
