package main

import "fmt"

func main() {

	c := make(chan int)
	go fib(5, c)
	for i := range c {
		fmt.Println(i)
	}

}
func fib(n int, c chan int) {
	x, y := 0, 1
	for {

		c <- x
		x, y = y, x+y

	}
	close(c)
}
