package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	s:=[][]uint8
	s=make([][]uint8,dy)
	return s



}

func main() {
	pic.Show(Pic)
}
