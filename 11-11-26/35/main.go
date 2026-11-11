package main

/*
35. 搜索插入位置

给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。

示例 1：
	输入: nums = [1,3,5,6], target = 5
	输出: 2

示例 2：
	输入: nums = [1,3,5,6], target = 2
	输出: 1

示例 3：
	输入: nums = [1,3,5,6], target = 7
	输出: 4

示例 4：
	输入: nums = [1,3,5,6], target = 0
	输出: 0

提示：
	1. 1 <= nums.length <= 10^4
	2. -10^4 <= nums[i] <= 10^4
	3. nums 为无重复元素的升序排列数组
	4. -10^4 <= target <= 10^4
*/

import "fmt"

// searchInsert 返回 target 的下标；不存在时返回它应插入的位置。
// 即第一个 nums[i] >= target 的 i，时间复杂度 O(log n)。
func searchInsert(nums []int, target int) int {
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

	return left
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
		{"示例1 命中", []int{1, 3, 5, 6}, 5, 2},
		{"示例2 插中间", []int{1, 3, 5, 6}, 2, 1},
		{"示例3 大于全部", []int{1, 3, 5, 6}, 7, 4},
		{"示例4 小于全部", []int{1, 3, 5, 6}, 0, 0},
		// 边界情况
		{"命中首位", []int{1, 3, 5, 6}, 1, 0},
		{"命中末位", []int{1, 3, 5, 6}, 6, 3},
		{"两元素间隔", []int{1, 3}, 2, 1},
		{"单元素-小于", []int{1}, 0, 0},
		{"单元素-命中", []int{1}, 1, 0},
		{"单元素-大于", []int{1}, 2, 1},
		{"空数组", []int{}, 1, 0},
	}

	pass := 0
	for _, c := range cases {
		got := searchInsert(c.nums, c.target)
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
