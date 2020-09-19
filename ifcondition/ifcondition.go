package main

import "fmt"

func foodType(age int) string {
	if age < 1 {
		return "liquid"
	} else if age >= 1 && age <= 3 {
		return "semi solid"
	} else {
		return "solid"
	}

}
func main() {
	fmt.Println(foodType(1), foodType(3), foodType(0))
}
