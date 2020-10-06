package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

type MyError2 struct {
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func (e *MyError2) Error() string {
	return fmt.Sprintf("you got error: %s", e.What)
}

func run() error {
	return &MyError2{
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
