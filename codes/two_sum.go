package main

import "fmt"

func sumOfTwoNums(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return ([]int{i, j})
			}
		}
	}

	return ([]int{-1, -1})
}

func main() {
	fmt.Println(sumOfTwoNums([]int{3, 5, 1, 11, 19, 4}, 9))
}
