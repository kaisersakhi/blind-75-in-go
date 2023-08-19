/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.


Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]

*/

package main

import "fmt"

func main() {
	var nums = []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 1))
}

func twoSum(nums []int, target int) []int {
	return twoSumWithHashMap(nums, target)
}

// Better approach, takes O(n)
func twoSumWithHashMap(nums []int, target int) []int {
	var res = []int{-1, -1}
	var hash = make(map[int]int)

	for i, _ := range nums {
		value, exists := hash[target-nums[i]]
		if exists {
			res[0] = value
			res[1] = i
			return res
		}
		hash[nums[i]] = i
	}
	return res
}

// Not a good approach to solve this problem, needs n^2 time
func twoSumNSquare(nums []int, target int) []int {
	for i, _ := range nums {
		for j := range nums {
			if i == j {
				continue
			}

			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
