package main

import "fmt"

func maxProfit(prices []int) int {
	profit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > profit {
				profit = prices[j] - prices[i]
			}
		}
	}

	return profit
}

func main() {
	fmt.Println(maxProfit([]int{7, 2, 5, 3, 6, 1, 4}))
}
