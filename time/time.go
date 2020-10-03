package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for {
		select {
		case <-time.After(2 * time.Microsecond):
			fmt.Println("exiting")
			return
		default:
			i++
			fmt.Println(i)
		}
	}
}
