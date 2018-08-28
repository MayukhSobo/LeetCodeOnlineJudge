/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

package main

import (
	"fmt"
)

// func twoSum(nums []int, target int) []int {
// // Time: O(n^2), Space: O(1)
// 	for i := 0; i < len(nums); i++ {
// 		for j := 0; j < len(nums); j++ {
// 			if i != j && nums[i]+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return []int{}
// }

func twoSum(nums []int, target int) []int {
	table := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		table[nums[i]] = i
	}

	for index, num := range nums {
		residual := target - num
		_, ok := table[residual]
		if ok && table[residual] != index {
			return []int{index, table[residual]}
		}
	}
	return []int{}
}

func main() {
	a := []int{2, 7, 11, 15}
	fmt.Println(twoSum(a, 9))
}
