package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("exiting")
			return
		}
	}
}
