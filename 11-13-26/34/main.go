package main

/*
34. 在排序数组中查找元素的第一个和最后一个位置

给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。
请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

示例 1：
	输入：nums = [5,7,7,8,8,10], target = 8
	输出：[3,4]

示例 2：
	输入：nums = [5,7,7,8,8,10], target = 6
	输出：[-1,-1]

示例 3：
	输入：nums = [], target = 0
	输出：[-1,-1]

提示：
	1. 0 <= nums.length <= 10^5
	2. -10^9 <= nums[i] <= 10^9
	3. nums 是一个非递减数组
	4. -10^9 <= target <= 10^9
*/

import (
	"fmt"
	"math/rand"
	"slices"
)

// searchRange 返回 target 在 nums 中的 [第一个位置, 最后一个位置]，不存在返回 [-1,-1]。
// 用两次二分分别求「最后一个 < target」和「第一个 > target」，两者夹住的就是目标区间。
func searchRange(nums []int, target int) []int {
	LB := GetLB(nums, target)
	RB := GetRB(nums, target)

	if LB == -2 || RB == -2 {
		return []int{-1, -1}
	}

	if RB-LB > 1 {
		return []int{LB + 1, RB - 1}
	}

	return []int{-1, -1}
}

// GetLB 返回最后一个小于 target 的下标；-2 表示不存在（即 target 比所有元素都大）。
func GetLB(nums []int, target int) int {
	L := 0
	R := len(nums) - 1
	B := -2

	for L <= R {
		mid := L + ((R - L) >> 1)
		if nums[mid] >= target {
			R = mid - 1
			B = R
		} else {
			L = mid + 1
		}
	}

	return B
}

// GetRB 返回第一个大于 target 的下标；-2 表示不存在（即 target 比所有元素都小）。
func GetRB(nums []int, target int) int {
	L := 0
	R := len(nums) - 1
	B := -2

	for L <= R {
		mid := L + ((R - L) >> 1)
		if nums[mid] > target {
			R = mid - 1
		} else {
			L = mid + 1
			B = L
		}
	}

	return B
}

// brute 线性扫描的标准答案，用来对拍。
func brute(nums []int, target int) []int {
	first, last := -1, -1
	for i, v := range nums {
		if v == target {
			if first == -1 {
				first = i
			}
			last = i
		}
	}
	if first == -1 {
		return []int{-1, -1}
	}
	return []int{first, last}
}

func main() {
	type testCase struct {
		name   string
		nums   []int
		target int
		want   []int
	}

	cases := []testCase{
		// 题目示例
		{"示例1 重复两次", []int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{"示例2 不存在", []int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{"示例3 空数组", []int{}, 0, []int{-1, -1}},
		// 边界情况
		{"单元素-命中", []int{1}, 1, []int{0, 0}},
		{"单元素-未命中", []int{1}, 2, []int{-1, -1}},
		{"全部相同", []int{2, 2, 2, 2}, 2, []int{0, 3}},
		{"命中首位", []int{1, 1, 2, 3}, 1, []int{0, 1}},
		{"命中末位", []int{1, 2, 3, 3}, 3, []int{2, 3}},
		{"目标小于全部", []int{5, 6, 7}, 1, []int{-1, -1}},
		{"目标大于全部", []int{5, 6, 7}, 9, []int{-1, -1}},
		{"负数与重复", []int{-3, -2, -2, -2, 0}, -2, []int{1, 3}},
	}

	pass := 0
	for _, c := range cases {
		got := searchRange(c.nums, c.target)
		ok := slices.Equal(got, c.want)
		if ok {
			pass++
		}
		status := "✅"
		if !ok {
			status = "❌"
		}
		fmt.Printf("%s target=%-3d got=%-8s want=%-8s  %-12s nums=%v\n",
			status, c.target, fmt.Sprint(got), fmt.Sprint(c.want), c.name, c.nums)
	}

	// 随机对拍
	rng := rand.New(rand.NewSource(20261113))
	randFail := 0
	for range 20000 {
		n := rng.Intn(20)
		nums := make([]int, n)
		for j := range nums {
			nums[j] = rng.Intn(7) - 3 // 取值范围小，保证有大量重复
		}
		slices.Sort(nums)
		target := rng.Intn(9) - 4

		got := searchRange(nums, target)
		want := brute(nums, target)
		if !slices.Equal(got, want) {
			randFail++
			if randFail <= 5 {
				fmt.Printf("❌ 对拍失败 nums=%v target=%d got=%v want=%v\n", nums, target, got, want)
			}
		}
	}

	fmt.Printf("\n用例通过 %d/%d，随机对拍 20000 组，失败 %d 组\n", pass, len(cases), randFail)
}
