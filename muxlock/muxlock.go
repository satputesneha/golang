package main

import (
	"fmt"
	"sync"
	"time"
)

var m int
var mut sync.Mutex

func calc(i int) {
	mut.Lock()
	defer mut.Unlock()
	fmt.Printf("1.%v\n", i)
	m = m + i
	fmt.Printf("2.%v\n", i)
	m = m * i
	fmt.Printf("3.%v\n", i)
}

func main() {
	for i := 0; i < 10; i++ {
		go calc(i)
	}
	time.Sleep(3 * time.Second)
}
