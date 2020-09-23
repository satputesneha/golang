package main

import "fmt"

func main() {
	a := StudyHours{lang: "golang", handsonDuration: 120}
	a.min = 60
	fmt.Printf("%v,%T", a, a)
	p := &a
	fmt.Printf("%v", (*p).min)

}

type StudyHours struct {
	lang            string
	min             int
	handsonDuration int
}
