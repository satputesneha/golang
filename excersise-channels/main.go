package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	recursive(t, ch)
	close(ch)
}

func recursive(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	recursive(t.Left, ch)
	ch <- t.Value
	recursive(t.Right, ch)

}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		i, iok := <-ch1
		j, jok := <-ch2
		if iok != jok {
			return false

		}
		if i != j {
			return false

		}
		if iok == false {
			return true
		}

	}

	return false
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))

}
