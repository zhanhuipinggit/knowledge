package main

import "fmt"

/*
*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。

你可以按任意顺序返回答案。

示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]
*/
func twoSum(nums []int, target int) []int {
	buckets := make(map[int]int)
	for i := 0; i <= len(nums)-1; i++ {
		index := target - nums[i]
		if j, ok := buckets[index]; ok {
			return []int{j, i}
		}
		buckets[nums[i]] = i
	}
	return nil
}

func threeSum(nums []int) [][]int {

	var res [][]int
	backtrack(nums, &res, []int{}, 0)

	return res

}

func backtrack(nums []int, res *[][]int, path []int, start int) {
	if len(path) == 3 {
		if path[0]+path[1]+path[2] == 0 {
			temp := make([]int, 3)
			copy(temp, path)
			*res = append(*res, temp)
		}
		return
	}

	for i := start; i < len(nums)-1; i++ {
		path = append(path, nums[i])
		backtrack(nums, res, path, i+1)
		path = path[:len(path)-1]
	}
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	res := twoSum(nums, target)
	fmt.Println("两位数相加：", res)

	numsThree := []int{-1, 0, 1, 2, -1, -4}
	resT := threeSum(numsThree)
	fmt.Println("三位数相加：", resT)

}
