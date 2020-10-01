package main

import "fmt"

func main() {
	c := make(chan string)
	go add(c)
	fmt.Println(<-c)

}
func add(c chan string) {
	c <- "ball"
}
