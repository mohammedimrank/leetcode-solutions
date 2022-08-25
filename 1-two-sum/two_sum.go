package lc

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

func TwoSum(nums []int, target int) []int {

	result := make(map[int]int)

	for index, element := range nums {

		if value, ok := result[element]; ok {
			return []int{index, value}
		}
		result[target-element] = index

	}
	return nil
}
