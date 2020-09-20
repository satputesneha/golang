package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Tuesday?")
	today := time.Now().Weekday()
	switch time.Tuesday {
	case today + 0:
		fmt.Println("today")
	case today + 1:
		fmt.Println("tommorow")
	case today + 2:
		fmt.Println("in 2 days")
	default:
		fmt.Println("too far")

	}

}
