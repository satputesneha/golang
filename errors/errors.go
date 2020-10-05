package main

import (
	"fmt"
)

type MyError2 struct {
	What string
}

func (e *MyError2) Error() string {
	return fmt.Sprintf("you got error: %s", e.What)
}

func run() error {
	return &MyError2{
		"It didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
