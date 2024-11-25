package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 6
	sum := twoSum(nums, target)
	fmt.Println("sum:", sum)
}

func twoSum(nums []int, target int) []int {
	sumMap := make(map[int]int, len(nums))
	for i, num := range nums {
		complement := target - num
		if idx, ok := sumMap[complement]; ok {
			return []int{idx, i}
		}
		sumMap[num] = i
	}
	return []int{}
}
