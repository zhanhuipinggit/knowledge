package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

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

/**
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。





示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。
示例 2：

输入：nums = [0,1,1]
输出：[]
解释：唯一可能的三元组和不为 0 。
示例 3：

输入：nums = [0,0,0]
输出：[[0,0,0]]
解释：唯一可能的三元组和为 0 。
*/

func threeSum(nums []int) [][]int {

	var res [][]int
	buckets := make(map[string]bool)
	backtrack(nums, &res, []int{}, 0, buckets)
	return res

}

func isValidPath(res [][]int, path []int) bool {
	if len(res) == 0 {
		return true
	}
	for _, r := range res {
		sort.Ints(r)
		sort.Ints(path)
		for i := 0; i < 3; i++ {
			if r[i] != path[i] {
				return true
			}
		}
	}

	return false
}

func backtrack(nums []int, res *[][]int, path []int, start int, buckets map[string]bool) {
	if len(path) == 3 {
		if path[0]+path[1]+path[2] == 0 {
			temp := make([]int, 3)
			sortI := make([]int, 3)
			copy(temp, path)
			copy(sortI, path)
			sort.Ints(sortI)
			s := arrayToString(sortI)
			if !buckets[s] {
				*res = append(*res, temp)
				buckets[s] = true
			}
		}
		return
	}

	for i := start; i < len(nums)-1; i++ {
		path = append(path, nums[i])
		backtrack(nums, res, path, i+1, buckets)
		path = path[:len(path)-1]
	}
}

func arrayToString(arr []int) string {

	var sb strings.Builder // 使用 strings.Builder 高效拼接字符串
	for _, num := range arr {
		sb.WriteString(strconv.Itoa(num)) // 将每个数字转换为字符串并拼接
	}
	return sb.String()
}

func threeSumIII(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				// 移动指针
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res

}

// ThreeSumII 三数之和常规解法
func threeSumII(nums []int) [][]int {
	result := [][]int{}
	// 排序数组，确保同样的数值在一起
	sort.Ints(nums)

	// 遍历数组
	for i := 0; i < len(nums)-2; i++ {
		// 跳过重复的值
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 双指针查找
		left, right := i+1, len(nums)-1
		for left < right {
			// 计算三数之和
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				// 找到一个三元组，添加到结果中
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// 跳过重复的元素
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// 移动指针
				left++
				right--
			} else if sum < 0 {
				// 如果和小于 0，左指针右移
				left++
			} else {
				// 如果和大于 0，右指针左移
				right--
			}
		}
	}

	return result
}

func main() {
	nums := []int{3, 2, 4}

	target := 6
	res := twoSum(nums, target)
	fmt.Println("两位数相加：", res)

	numsThree := []int{-1, 0, 1, 2, -1, -4}
	resT := threeSumIII(numsThree)
	fmt.Println("三位数相加：", resT)

}
