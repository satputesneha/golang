package main

import "fmt"

func sliceInfo(slice []int) {
	fmt.Printf("len=%v,cap=%v,val=%v\n ", len(slice), cap(slice), slice)

}
func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7}
	sliceInfo(s)
	s = s[:0]
	sliceInfo(s)
	s = s[:4]
	sliceInfo(s)
	s = s[2:]
	sliceInfo(s)
	var slice []int
	sliceInfo(slice)
	fmt.Println(slice == nil)
	slice1 := make([]int, 100)
	sliceInfo(slice1)

}
