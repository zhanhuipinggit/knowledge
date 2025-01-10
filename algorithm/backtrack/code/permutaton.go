package main

import (
	"fmt"
)

// 全排列
func backtrackI(nums []int, res *[][]int, path []int, used map[int]bool) {
	// 如果路径长度等于数组长度，说明找到一个完整的排列
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path) // 复制当前路径
		*res = append(*res, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		// 如果当前元素已经被使用，跳过
		if used[i] {
			continue
		}

		// 做选择
		path = append(path, nums[i])
		used[i] = true

		// 递归进入下一层
		backtrackI(nums, res, path, used)

		// 撤销选择
		path = path[:len(path)-1]
		used[i] = false
	}
}

func permute(nums []int) [][]int {
	var res [][]int
	used := make(map[int]bool)
	backtrackI(nums, &res, []int{}, used)
	return res
}

/*
*
给定一个正整数数组 nums 和一个目标正整数 target ，请找出所有可能的组合，使得组合中的元素和等于 target 。
给定数组无重复元素，每个元素可以被选取多次。请以列表形式返回这些组合，列表中不应包含重复组合。
*/
func backtrackII(nums []int, start int, res *[][]int, path []int, target int) {
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		if target-nums[i] < 0 {
			continue
		}
		backtrackII(nums, i, res, path, target-nums[i])
		path = path[:len(path)-1]
	}
}

func permuteII(nums []int, target int) [][]int {
	var res [][]int
	backtrackII(nums, 0, &res, []int{}, target)
	return res
}

/**
给定一个正整数数组 nums 和一个目标正整数 target ，请找出所有可能的组合，使得组合中的元素和等于 target 。
给定数组可能包含重复元素，每个元素只可被选择一次。请以列表形式返回这些组合，列表中不应包含重复组合。
*/

func backtrackIII(nums []int, res *[][]int, path []int, start int, target int) {
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	for i := start; i < len(nums); i++ {
		if target-nums[i] < 0 {
			continue
		}

		path = append(path, nums[i])
		backtrackIII(nums, res, path, i+1, target-nums[i])
		path = path[:len(path)-1]

	}

}

func permuteIII(nums []int, target int) [][]int {
	var res [][]int
	backtrackIII(nums, &res, []int{}, 0, target)
	return res
}

func main() {
	nums := []int{1, 2, 3}
	//permutations := permute(nums)
	//for _, perm := range permutations {
	//	fmt.Println(perm)
	//}

	permutationsII := permuteIII(nums, 3)
	for _, perm := range permutationsII {
		fmt.Println(perm)
	}
}
