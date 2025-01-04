package main

import (
	"sort"
)

/**
给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。

考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：

更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
返回 k 。
*/
// removeDuplicates
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1

}

func RemoveDuplicatesList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head
	for fast != nil {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}

}

/**
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。元素的顺序可能发生改变。然后返回 nums 中与 val 不同的元素的数量。
*/

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow

}

func moveZeroes(nums []int) {
	p := removeElement(nums, 0)
	for ; p < len(nums); p++ {
		nums[p] = 0
	}
}

/*
*
二分查找、左右指针
*/
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

/**
给你一个下标从 1 开始的整数数组 numbers ，该数组已按 非递减顺序排列  ，请你从数组中找出满足相加之和等于目标数 target 的两个数。如果设这两个数分别是 numbers[index1] 和 numbers[index2] ，则 1 <= index1 < index2 <= numbers.length 。

以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。

你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。

你所设计的解决方案必须只使用常量级的额外空间。
*/

func twoSum(nums []int, target int) []int {
	left, right := 1, len(nums)
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++
		} else if sum > target {
			right--
		}
	}
	return []int{-1, -1}
}

// reverseString 反转数组
func reverseString(s []rune) {
	left, right := 0, len(s)-1
	for left < right {
		temp := s[left]
		s[left] = s[right]
		s[right] = temp
		left++
		right--
	}
}

/**
回文串判断
回文串就是正着读和反着读都一样的字符串。
比如说字符串 aba 和 abba 都是回文串，因为它们对称，反过来还是和本身一样；反之，字符串 abac 就不是回文串。
*/

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

/**
给你一个字符串 s，找到 s 中最长的 回文 子串。

示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"
提示：

1 <= s.length <= 1000
s 仅由数字和英文字母组成
*/

func palindrome(s string, l, r int) string {
	if l > 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}

func longestPalindrome(s string) string {
	var longestString string
	for i := 0; i < len(s)-1; i++ {
		s1 := palindrome(s, i, i)
		s2 := palindrome(s, i, i+1)
		if len(longestString) < len(s1) {
			longestString = s1
		}
		if len(longestString) < len(s2) {
			longestString = s2
		}
	}
	return longestString
}

// 最少会议室，每个人有开始时间和结束时间，求会议室最少数
type meetingTime struct {
	time  int
	start bool
}

func minMeetingRooms(nums [][]int) int {
	var events []*meetingTime
	for _, num := range nums {
		events = append(events, &meetingTime{time: num[0], start: true})
		events = append(events, &meetingTime{time: num[1], start: false})
	}

	// 排序
	// 按时间排序（如果时间相同，结束事件优先）
	sort.Slice(events, func(i, j int) bool {
		if events[i].time == events[j].time {
			return !events[i].start
		}
		return events[i].time < events[j].time
	})

	// 计算会议室
	var current int
	var sumMeeting int
	for _, event := range events {
		if event.start {
			current++
			if current > sumMeeting {
				sumMeeting = current
			}
		} else {
			current--
		}
	}
	return sumMeeting

}
