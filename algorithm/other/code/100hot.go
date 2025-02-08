package main

import (
	"fmt"
	"sort"
)

/**
49. 字母异位词分组
中等
相关标签
相关企业
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
Anagrams : 判断回文
*/

func groupAnagrams(strs []string) [][]string {
	buckets := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		str := strs[i]
		s := strSort(str)
		buckets[s] = append(buckets[s], str)
	}

	var res [][]string
	for _, group := range buckets {
		res = append(res, group)
	}
	return res

}

func strSort(str string) string {
	strr := []rune(str)
	sort.Slice(strr, func(i, j int) bool {
		return strr[i] < strr[j]
	})
	return string(strr)
}

/*
*给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
*/

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)

	// 1. 把所有数加入 Set
	for _, num := range nums {
		numSet[num] = true
	}

	maxLength := 0

	// 2. 遍历所有数字，查找连续序列
	for num := range numSet {
		// 只从序列的起点开始检查
		if !numSet[num-1] { // num-1 不在 Set 里，说明 num 是起点
			length := 1
			curr := num

			// 计算当前连续序列的长度
			for numSet[curr+1] { // 检查下一个数是否存在
				curr++
				length++
			}

			// 更新最大长度
			if length > maxLength {
				maxLength = length
			}
		}
	}

	return maxLength
}

/*
*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。
*/
func threeSumI(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == nums[i+1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right] + nums[i]
			if sum == 0 {
				res = append(res, []int{nums[left], nums[right], nums[i]})
				if left < right && nums[left] == nums[left+1] {
					left++
				}

				if left < right && nums[right] == nums[right+1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return res
}

/*
*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
*/
func trap(height []int) int {
	stack := []int{}
	water := 0
	for i := 0; i < len(height); i++ {
		for i > 0 && height[i] > height[stack[len(stack)-1]] {
			bottom := stack[len(stack)-1] // 洼地
			stack := stack[:len(stack)-1] // 弹出
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			right := i
			width := right - left - 1
			heightDiff := min(height[left], height[right]) - height[bottom]
			if heightDiff > 0 {
				water += heightDiff * width
			}
		}
	}
	return water
}

func trapI(height []int) int {
	n := len(height)
	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(height[i-1], height[i])
	}
	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(height[i+1], height[i])
	}
	res := 0
	for i := 0; i < n; i++ {
		res += min(rightMax[i], leftMax[i]) - height[i]
	}

	return res
}

/*
*
560. 和为 K 的子数组
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。

子数组是数组中元素的连续非空序列。
前缀和
当前和-k 存在，就代表从有一段数据和为k
*/
func subarraySum(nums []int, k int) int {
	prefixCount := map[int]int{0: 1}
	currentSum := 0
	count := 0
	for _, num := range nums {
		currentSum += num
		if _, ok := prefixCount[currentSum-k]; ok {
			count += prefixCount[currentSum-k]
		}
		prefixCount[currentSum]++
	}
	return count
}

/**
239. 滑动窗口最大值
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回 滑动窗口中的最大值 。
示例 1：
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
示例 2：

输入：nums = [1], k = 1
输出：[1]
*/

/*
*
解析：
思路：
滑动窗口：我们有一个大小为 k 的滑动窗口，每次窗口右移一个位置。每当窗口发生变化时，我们需要找出当前窗口内的最大值。

双端队列（Deque）：使用双端队列来维护当前窗口内可能的最大值的下标。队列的性质是从队头到队尾是递减的，即队头的元素是当前窗口的最大值。

在遍历数组时，如果队尾元素比当前元素小，则可以将队尾元素弹出，因为当前元素更大，且它可以成为未来窗口的最大值。
每当窗口的左边界移出时，要检查队头元素是否超出了当前窗口的范围，如果超出，则将队头元素弹出。
队头元素始终是当前窗口内的最大值。
具体步骤：
遍历数组中的每个元素，对于每个元素执行以下操作：
移除队头：如果队头元素已经不在窗口内，移除它。
移除队尾：如果队尾元素小于当前元素，移除队尾元素，因为当前元素在未来会成为一个更大的候选元素。
加入当前元素的索引：将当前元素的索引添加到队列中。
记录最大值：当索引 i 达到 k-1 以后，每次窗口更新时，将队头元素（即当前窗口的最大值）添加到结果中。
最终返回记录的所有最大值。
*/
func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	deque := []int{}
	for i := 0; i < len(nums); i++ {
		if i > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}

		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)
		if i >= k-1 {
			res = append(res, nums[deque[0]])
		}

	}

	return res

}

/**
56. 合并区间
中等
相关标签
相关企业
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
*/

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}
	for _, interval := range intervals {
		if len(res) == 0 || res[len(res)-1][1] < interval[0] {
			res = append(res, interval)
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], interval[1])
		}
	}
	return res

}

/*
*
189. 轮转数组
中等
相关标签
相关企业
提示
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]
*/

/*
*
方法一：使用额外的数组
最直接的方法是使用一个额外的数组来存储结果：

计算新的位置：对于一个元素 nums[i]，它的新位置应该是 (i + k) % n，其中 n 是数组的长度。
将数组中的元素依次放入新数组中。
复制新数组中的内容回原数组。
*/
func rotate(nums []int, k int) {
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < len(nums); i++ {
		index := (i + k) % n
		res[index] = nums[i]
	}
	for i := 0; i < len(res); i++ {
		nums[i] = res[i]
	}
}
func rotateII(nums []int, k int) {
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}
func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

/**
238. 除自身以外数组的乘积
中等
相关标签
相关企业
提示
给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。

请 不要使用除法，且在 O(n) 时间复杂度内完成此题。



示例 1:

输入: nums = [1,2,3,4]
输出: [24,12,8,6]
示例 2:

输入: nums = [-1,1,0,-3,3]
输出: [0,0,9,0,0]

*/

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	prefix := 1
	for i := 0; i < len(nums); i++ {
		res[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= suffix
		suffix *= nums[i]
	}

	return res

}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	i := longestConsecutive(nums)
	fmt.Println(i)
	subarraySums := []int{1, 2, 3}
	s := subarraySum(subarraySums, 3)
	fmt.Println(s)

	sliding := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	res := maxSlidingWindow(sliding, k)
	fmt.Println(res)

}
