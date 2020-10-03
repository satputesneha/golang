package main

import "fmt"

func main() {
	c := make(chan int)
	q := make(chan int)
	go func(c, q chan int) {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		q <- 1

	}(c, q)
	fib(c, q)
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
