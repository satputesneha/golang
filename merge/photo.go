package main

import "fmt"

func photo(big []int, m int, small []int, n int) {
	photo_last := len(big) - 1
	free_last := m - 1
	for free_last >= 0 {
		big[free_last], big[photo_last] = big[photo_last], big[free_last]
		free_last--
		photo_last--
	}
	fmt.Println(big)
	free_index := 0
	photo_index := len(big) - m
	small_index := 0
	for small_index < n {
		if photo_index == len(big) || big[photo_index] > small[small_index] {
			big[free_index] = small[small_index]
			free_index++
			small_index++
		} else {
			big[free_index] = big[photo_index]
			free_index++
			photo_index++
		}
	}
	fmt.Println(big)

}
func main() {
	photo([]int{2, 3, 4, 5, 0, 0, 0}, 4, []int{1, 2, 3}, 3)
}
