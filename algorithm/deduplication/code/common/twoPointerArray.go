package main

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
