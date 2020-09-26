package main

import "fmt"

func main() {

	s := map[string]int{"India": 30, "USA": 40}

	m := map[string]map[string]int{
		"sneha-age":    {"Australia": 22, "switzerland": 25},
		"shiva-age":    {"India": 3},
		"srikanth-age": s,
	}
	fmt.Println(m)
	value, exist := m["motba"]
	fmt.Println(value, exist)

}
