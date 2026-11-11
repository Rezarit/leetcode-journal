package main

/*
704. 二分查找

给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target，
写一个函数搜索 nums 中的 target，如果 target 存在返回下标，否则返回 -1。

你必须编写一个具有 O(log n) 时间复杂度的算法。

示例 1：
	输入：nums = [-1,0,3,5,9,12], target = 9
	输出：4
	解释：9 出现在 nums 中并且下标为 4

示例 2：
	输入：nums = [-1,0,3,5,9,12], target = 2
	输出：-1
	解释：2 不存在 nums 中因此返回 -1

提示：
	1. 你可以假设 nums 中的所有元素是不重复的。
	2. n 将在 [1, 10000] 之间。
	3. nums 的每个元素都将在 [-9999, 9999] 之间。
*/

import "fmt"

// search 在升序数组 nums 中二分查找 target，返回其下标，不存在返回 -1。
// 时间复杂度 O(log n)，空间复杂度 O(1)。
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := 0

	for left <= right {
		mid = (left + right) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	type testCase struct {
		name   string
		nums   []int
		target int
		want   int
	}

	cases := []testCase{
		// 题目示例
		{"示例1 命中中间", []int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{"示例2 不存在", []int{-1, 0, 3, 5, 9, 12}, 2, -1},
		// 边界情况
		{"单个元素-命中", []int{5}, 5, 0},
		{"单个元素-未命中", []int{5}, 3, -1},
		{"命中首位", []int{-1, 0, 3, 5, 9, 12}, -1, 0},
		{"命中末位", []int{-1, 0, 3, 5, 9, 12}, 12, 5},
		{"空数组", []int{}, 1, -1},
		{"目标小于全部", []int{1, 2, 3}, 0, -1},
		{"目标大于全部", []int{1, 2, 3}, 4, -1},
	}

	pass := 0
	for _, c := range cases {
		got := search(c.nums, c.target)
		ok := got == c.want
		if ok {
			pass++
		}
		status := "✅"
		if !ok {
			status = "❌"
		}
		fmt.Printf("%s target=%-3d got=%-2d want=%-2d  %-9s nums=%v\n",
			status, c.target, got, c.want, c.name, c.nums)
	}
	fmt.Printf("\n通过 %d/%d\n", pass, len(cases))
}
