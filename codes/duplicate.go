package main

import "fmt"

func containsDuplicate(nums []int) bool {
	cache := make(map[int]int)
	for i, n := range nums {
		if _, ok := cache[n]; ok {
			return false
		}
		cache[n] = i
	}
	return true
}

func main() {
	fmt.Println(containsDuplicate([]int{7, 2, 5, 3, 6, 1, 4}))
}
