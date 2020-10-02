package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go add("ball", c)
	go add("football", c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	d := make(chan int, 3)
	go intchan(3, d)
	fmt.Println("3")
	go intchan(8, d)
	fmt.Println("8")
	go intchan(9, d)
	fmt.Println("9")

	fmt.Println(<-d)
	fmt.Println(<-d)
	fmt.Println(<-d)

	fmt.Println("Hello")
	go delayed()
}
func add(m string, c chan string) {
	c <- m
}

func intchan(m int, d chan int) {
	fmt.Println("intchan function")
	d <- m
}

func delayed() {
	time.Sleep(time.Second * 2)
	fmt.Println("delayed")

}
