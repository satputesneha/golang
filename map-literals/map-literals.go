package main

import "fmt"

func main() {
	m := map[string]map[string]int{}
	fmt.Println(m)
	m["gopal"] = map[string]int{
		"sneha-age":    25,
		"shiva-age":    30,
		"srikanth-age": 28,
	}
	fmt.Println(m)
	value, exist := m["motba"]
	fmt.Println(value, exist)

}
