package main

import (
	"fmt"
)

func main() {
	fmt.Println(onePass([]int{7, 1, 5, 3, 6, 4}))
}

// Pretty fair approach; takes O(n)
func onePass(prices []int) int {
	totalDays := len(prices)

	if totalDays < 1 {
		return 0
	}
	lastBuy := prices[0]
	maxProfit := 0

	for i := 0; i < totalDays; i++ {
		currentProfit := prices[i] - lastBuy
		if currentProfit > maxProfit { // max(currentProfit, maxProfit)
			maxProfit = currentProfit
		}
		if prices[i] < lastBuy { //min(prices[i], lastBuy)
			lastBuy = prices[i]
		}
	}
	return maxProfit
}

// Bute-Force, super slow on large inputs; takes O(n^2)
func bruteForce(prices []int) int {
	profit := 0

	for i := 0; i < len(prices); i++ {
		for j := 0; j < len(prices); j++ {
			sub := prices[j] - prices[i]
			if i < j && sub > profit {
				profit = sub
			}
		}
	}
	return profit
}
