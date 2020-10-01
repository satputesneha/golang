package main

import "fmt"

func main() {
	c := make(chan string)
	go add("ball", c)
	go add("football", c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}
func add(m string, c chan string) {
	c <- m
}
