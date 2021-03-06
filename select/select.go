package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	q := make(chan int)
	go func(c, q chan int) {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		q <- 1
		for {
			fmt.Println(<-c)
			//time.Sleep(1 * time.Second)

		}
	}(c, q)
	fib(c, q)
	fib2(c)
}
func fib(c chan int, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y

		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func fib2(c chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y

		case <-time.After(100 * time.Microsecond):
			fmt.Println("quit")
			return
		}
	}
}
