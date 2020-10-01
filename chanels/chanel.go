package main

import "fmt"

func main() {
	c := make(chan string)
	go add("ball", c)
	go add("football", c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	d := make(chan int, 3)
	intchan(3, d)
	fmt.Println("3")
	intchan(8, d)
	fmt.Println("8")
	intchan(9, d)
	fmt.Println("9")

	fmt.Println(<-d)
	fmt.Println(<-d)
	fmt.Println(<-d)

	fmt.Println("Hello")

}
func add(m string, c chan string) {
	c <- m
}

func intchan(m int, d chan int) {
	d <- m
}
