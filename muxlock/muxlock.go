package main

import (
	"fmt"
	"sync"
)

var m int
var mut sync.Mutex

func calc(i int, c chan int) {
	mut.Lock()
	defer mut.Unlock()
	fmt.Printf("1.%v\n", i)
	m = m + i
	fmt.Printf("2.%v\n", i)
	m = m * i
	fmt.Printf("3.%v\n", i)
	c <- i
}

func main() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go calc(i, c)
	}
	for j := 0; j < 10; j++ {
		fmt.Printf("Finished %v\n", <-c)
	}
}
