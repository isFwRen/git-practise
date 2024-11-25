package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	target := 6
	search := twoSearch(nums, target)
	search1 := search11(nums, target)
	fmt.Println(search)
	fmt.Println(search1)
	t := Teacher{}
	t.ShowB()
	p := 1
	incr(&p)
	fmt.Println(p)
}

func twoSearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

// 时间复杂度 O(logn)
func search11(nums []int, target int) int {
	// 初始化左右边界
	left := 0
	right := len(nums)

	// 循环逐步缩小区间范围
	for left < right {
		// 求区间中点
		mid := left + (right-left)>>1
		fmt.Println("11", mid)
		// 根据 nums[mid] 和 target 的大小关系
		// 调整区间范围
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	// 在输入数组内没有找到值等于 target 的元素
	return -1
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func incr(p *int) int {
	fmt.Println(p)
	*p++
	return *p
}
