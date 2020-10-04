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
	fmt.Println(i)
	m = m + i
	fmt.Println(i)
	m = m * i
	fmt.Println(i)
}

func main() {
	for i := 0; i < 100; i++ {
		go calc(i)
	}
	time.Sleep(3 * time.Second)
}
