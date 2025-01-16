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

/**
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组
是数组中的一个连续部分。



示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
示例 2：

输入：nums = [1]
输出：1
示例 3：

输入：nums = [5,4,-1,7,8]
输出：23
分界：大于0小于0
大于0+任何数，都大于本身
小于0+任何数，都小于本身
*/

func maxSubArray(nums []int) int {

	// 辅助函数，用于求两个整数中的较大值
	maxSum := nums[0]
	currentSum := nums[0]
	for i := 1; i < len(nums); i++ {
		currentSum = maxs(nums[i], currentSum+nums[i])
		maxSum = maxs(currentSum, maxSum)
	}
	return maxSum
}

// 辅助函数，用于求两个整数中的较大值
func maxs(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。



示例 1：

输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
示例 2：

输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1
示例 3：

输入：nums = [1], target = 0
输出：-1
分析：主要核心是要知道：
中间点：mid = left + (right-left)/2
left := 0
right := len(nums)-1
如果中间点 大于 left点，则证明左侧是有序的
如果中间点小于left点，说明右边是有序的
再有序的空间里面查找target的值，则能达到log(n)
*/

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		// 右侧有序
		if nums[mid] < nums[right] {
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[mid] > target && nums[left] <= target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1

}

/*
*
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

// 根据数组构建链表
func buildLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	// 初始化头节点
	head := &ListNode{Val: arr[0]}
	current := head

	// 遍历数组并构建链表
	for i := 1; i < len(arr); i++ {
		current.Next = &ListNode{Val: arr[i]}
		current = current.Next
	}

	return head
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// 链表形成环
	p := head
	length := 1
	for p.Next != nil {
		length++
		p = p.Next
	}
	p.Next = head

	// 寻找头节点
	// 头节点是倒数第几个节点
	// 比如k== 2 那么头节点，就是倒数第二个node
	start := k % length
	if start == 0 {
		return head
	}
	newTail := head
	for i := 1; i < length-start; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next
	newTail.Next = nil
	return newHead
	//
	//for i := 0; i < k; i++ {
	//
	//}

}

/*
*
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
reverseBetween
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 将链表分成三个链表
	dummy := &ListNode{Next: head}
	pre := dummy
	for i := 1; i < left; i++ {
		pre = pre.Next
	}

	// 反转链表
	// 开始反转
	curr := pre.Next
	var next *ListNode
	for i := 0; i < right-left; i++ {
		next = curr.Next
		curr.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}

	return dummy.Next

}

// 链表两种反转
func listReverseI(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		temp := current.Next // 保留下一个节点
		current.Next = prev  // 断节点路径，指向前节点
		prev = current       // prev往前移动
		current = temp       // 当前节点后移动
	}
	return prev
}

// 打印链表
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	nums := []int{3, 2, 4}

	target := 6
	res := twoSum(nums, target)
	fmt.Println("两位数相加：", res)

	numsThree := []int{-1, 0, 1, 2, -1, -4}
	resT := threeSumIII(numsThree)
	fmt.Println("三位数相加：", resT)

	// 最大子串
	numsMaxSub := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	resM := maxSubArray(numsMaxSub)
	fmt.Println("最长子串：", resM)

	// 最大子串
	numsSearch := []int{5, 1, 3}
	resS := search(numsSearch, 3)
	fmt.Println("查找旋转排序数组：", resS)

	// 构建链表
	head := buildLinkedList([]int{1, 2, 3, 4, 5})
	//heads := rotateRight(head, 2)
	//fmt.Println("旋转链表，将链表每个节点向右移动 k 个位置", heads.Val)
	//
	//reverseBetween(head, 2, 4)
	heads := listReverseI(head)
	printList(heads)

}
