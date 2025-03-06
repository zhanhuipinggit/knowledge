package main

import (
	"container/list"
	"fmt"
	"math"
	"sort"
	"sync"
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

/*
*
73. 矩阵置零
给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
*/
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	firstRowZero := false
	firstColZero := false
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstRowZero = true
		}
	}

	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			firstColZero = true
		}
	}

	// 用首行和首列来标记是否为0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	// 根据首行首列来标记
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 标记首行首列
	if firstColZero {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}

	// 标记首行首列
	if firstRowZero {
		for i := 0; i < n; i++ {
			matrix[i][0] = 0
		}
	}

}

/*
*
54. 螺旋矩阵
中等
相关标签
相关企业
提示
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
*/
func spiralOrder(matrix [][]int) []int {
	res := []int{}

	// 边界初始化
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

	for top <= bottom && left <= right {
		// 从左到右遍历上边界
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++ // 完成上边界遍历后，top 边界下移

		// 从上到下遍历右边界
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right-- // 完成右边界遍历后，right 边界左移

		if top <= bottom {
			// 从右到左遍历下边界
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom-- // 完成下边界遍历后，bottom 边界上移
		}

		if left <= right {
			// 从下到上遍历左边界
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++ // 完成左边界遍历后，left 边界右移
		}
	}

	return res
}

func spiralOrderII(matrix [][]int) []int {
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	res := []int{}
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if top <= bottom {
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}

	}
	return res
}

/*
*
给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
*/
func rotateI(matrix [][]int) {
	n := len(matrix)
	// 行变成列
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 水平翻转
	for i := 0; i < n; i++ {
		left, right := 0, n-1
		for left < right {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
			left++
			right--
		}
	}

}

/**
240. 搜索二维矩阵 II
中等
相关标签
相关企业
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列

需要注意解法，必须从右上角开始往下
*/

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return false
	}
	row, col := 0, n-1
	for row < m && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}
	return false

}

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

// isPalindrome 判断链表是否是回文链表
func isPalindrome(head *ListNode1) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 1. 使用快慢指针找到链表中点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 2. 反转后半部分链表
	reverse := func(head *ListNode1) *ListNode1 {
		var prev *ListNode1
		curr := head
		for curr != nil {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}
		return prev
	}

	secondHalf := reverse(slow)
	firstHalf := head

	// 3. 比较前半部分和后半部分
	p1, p2 := firstHalf, secondHalf
	for p2 != nil { // 只需比较后半部分
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	return true
}

func TestIsPalindrome() {
	head := &ListNode1{1, nil}
	head1 := &ListNode1{2, nil}
	head2 := &ListNode1{2, nil}
	head3 := &ListNode1{1, nil}
	head.Next = head1
	head1.Next = head2
	head2.Next = head3

	isPalindrome(head)
}

/*
*
2. 两数相加
中等
相关标签
相关企业
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
*/
func addTwoNumbers(l1 *ListNode1, l2 *ListNode1) *ListNode1 {
	dummy := &ListNode1{}
	curr := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		val1, val2 := 0, 0
		if l1 != nil {
			l1 = l1.Next
			val1 = l1.Val
		}

		if l2 != nil {
			l2 = l2.Next
			val2 = l2.Val
		}
		sum := val1 + val2 + carry
		carry = sum / 10
		node := &ListNode1{Val: sum % 10}
		curr.Next = node
		curr = curr.Next
	}
	return dummy.Next // 返回结果链表的实际头节点
}

func sortList(head *ListNode1) *ListNode1 {
	if head == nil {
		return head
	}
	countH := head
	maxNum := math.MinInt32
	for countH != nil {
		if maxNum < countH.Val {
			maxNum = countH.Val
		}
		countH = countH.Next
	}

	newSliceList := make([]int, maxNum+1)
	curr := head
	for curr != nil {
		val := curr.Val
		newSliceList[val]++
		curr = curr.Next
	}
	dummy := &ListNode1{}
	pre := dummy
	for i, v := range newSliceList {
		for v > 0 {
			node := &ListNode1{Val: i}
			pre.Next = node
			pre = pre.Next
			v--
		}
	}

	return dummy.Next

}

func TestSortList() {
	head := &ListNode1{4, nil}
	head1 := &ListNode1{2, nil}
	head2 := &ListNode1{1, nil}
	head3 := &ListNode1{3, nil}
	head.Next = head1
	head1.Next = head2
	head2.Next = head3

	p := sortList(head)
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}

func trapII(height []int) int {
	if len(height) < 3 {
		return 0
	}

	leftMax := make([]int, len(height))
	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(height[i-1], height[i])
	}

	rightMax := make([]int, len(height))
	rightMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(height[i], height[i+1])
	}
	res := 0
	for i := 0; i < len(height); i++ {
		res += min(leftMax[i], rightMax[i]) - height[i]
	}
	return res

}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxAreaS := 0
	for left < right {
		widthH := right - left
		minH := 0
		if height[left] < height[right] {
			minH = height[left]
			left++
		} else {
			minH = height[right]
			right--
		}

		currArea := widthH * minH
		if currArea > maxAreaS {
			maxAreaS = currArea
		}
	}

	return maxAreaS
}

func trapIII(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}

	left, right := 0, n-1
	maxLeft, maxRight := height[left], height[right]
	res := 0
	for left <= right {
		if height[left] <= height[right] {
			if height[left] >= maxLeft {
				maxLeft = height[left]
			} else {
				res += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				res += maxRight - height[right]
			}
			right--
		}
	}

	return res

}

// 单调栈
func trapIIII(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	stack := []int{}
	res := 0
	for i := 0; i < n-1; i++ {
		// 大于栈顶元素，说明可以储水
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			top := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}

			width := i - stack[len(stack)-1] - 1
			heightDiff := min(height[stack[len(stack)-1]], height[i]) - height[top] // 水的高度
			res += heightDiff * width
		}
		stack = append(stack, i)
	}

	return res

}

func orangesRotting(grid [][]int) int {
	queue := [][]int{} // 存储腐烂橘子的坐标
	freshOranges := 0
	row, col := len(grid), len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			} else if grid[i][j] == 1 {
				freshOranges++
			}
		}
	}

	if freshOranges == 0 {
		return 0
	}

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	minuteCount := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			x, y := queue[0][0], queue[0][1]
			queue = queue[1:]
			for _, dir := range directions {
				nx, ny := x+dir[0], y+dir[1]
				if nx > 0 && nx < row && ny > 0 && ny < col && grid[nx][ny] == 1 {
					grid[nx][ny] = 2
					freshOranges--
					queue = append(queue, []int{nx, ny})
				}
			}
		}
		minuteCount++
	}

	return minuteCount

}

func generateParenthesis(n int) []string {
	res := []string{}
	left, right := 0, 0
	var backtracking func(int, int, string)
	backtracking = func(left int, right int, s string) {
		if left == n && right == n {
			res = append(res, s)
			return
		}
		if left < n {
			backtracking(left+1, right, s+"(")
		}

		if right < n {
			backtracking(left, right+1, s+")")
		}

	}

	backtracking(left, right, "")
	return res
}

// 汉诺塔
func move(src, tar *list.List) {
	pan := src.Back()
	tar.PushBack(pan.Value)
	src.Remove(pan)
}

func hanota(i int, src, buf, tar *list.List) {
	if i == 0 {
		move(src, tar)
		return
	}

	hanota(i-1, src, tar, buf)
	move(src, tar)
	hanota(i-1, buf, src, tar)
}

/**
79. 单词搜索
已解答
中等
相关标签
相关企业
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。


*/

func exist(board [][]byte, words string) bool {
	row, col := len(board), len(board[0])
	var dfs func(int, int, int) bool
	dfs = func(index int, r, c int) bool {
		if len(words) == index {
			return true
		}

		if r < 0 && r > row && c < 0 && c > col && board[r][c] != words[index] {
			return false
		}
		temp := board[r][c]
		board[r][c] = '#'
		found := dfs(index+1, r+1, c) || dfs(index+1, r, c+1) || dfs(index+1, r-1, c) || dfs(index+1, r, c-1)
		board[r][c] = temp
		return found
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if dfs(0, i, j) {
				return true
			}
		}
	}
	return false
}

/*
*
121. 买卖股票的最佳时机
简单
相关标签
相关企业
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

示例 1：

输入：[7,1,5,3,6,4]
输出：5
解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。

	注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。

示例 2：

输入：prices = [7,6,4,3,1]
输出：0
解释：在这种情况下, 没有交易完成, 所以最大利润为 0。
*/
func maxProfit(prices []int) int {
	n := len(prices)
	maxProfitc := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if prices[j]-prices[i] > maxProfitc {
				maxProfitc = prices[j] - prices[i]
			}
		}
	}
	return maxProfitc
}

func maxProfitI(prices []int) int {
	minPrice := math.MaxInt
	maxPrice := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}

		currPrice := price - minPrice
		if currPrice > maxPrice {
			maxPrice = currPrice
		}
	}

	return maxPrice
}

/**
55. 跳跃游戏
中等
相关标签
相关企业
给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。
*/

func canJump(nums []int) bool {
	maxJump := 0
	for i := 0; i < len(nums); i++ {
		if i > maxJump {
			return false
		}

		maxJump = max(maxJump, i+nums[i])
		if maxJump >= len(nums)-1 {
			return true
		}
	}

	return false
}

func canJumpII(nums []int) bool {
	maxJump := 0
	for i := 0; i < len(nums); i++ {
		if i > maxJump {
			return false
		}
		maxJump = max(maxJump, i+nums[i])
		if maxJump > len(nums)-1 {
			return true
		}
	}
	return false
}

// 小偷
func slave(nums []int) int {
	count := make([]int, len(nums))
	count[0] = nums[0]
	count[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		count[i] = max(count[i-1], count[i-2]+nums[i])
	}
	return count[len(nums)-1]
}

// 括号组队
func compute(n int) []string {
	res := []string{}
	var dfs func(int, int, string)
	dfs = func(left, right int, path string) {
		if left == n && right == n {
			res = append(res, path)
			return
		}

		// 先尝试添加左括号
		if left < n {
			dfs(left+1, right, path+"(")
		}

		// 再尝试添加右括号（必须保证右括号数量小于左括号）
		if right < left {
			dfs(left, right+1, path+")")
		}
	}
	dfs(0, 0, "")
	return res
}

// MemoryPool 分配内存池
type MemoryPool struct {
	pool *sync.Pool
}

const (
	poolSize   = 100
	bufferSize = 4096
)

func NewMemoryPool() *MemoryPool {
	mp := &MemoryPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return make([]byte, bufferSize)
			},
		},
	}

	for i := 0; i < poolSize; i++ {
		mp.pool.Put(make([]byte, bufferSize))
	}

	return mp

}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy1, dummy2 := &ListNode{Val: 101}, &ListNode{Val: 101}
	p, q := dummy1, dummy2
	for head != nil {
		if (head.Next != nil && head.Next.Val == head.Val) || (head.Val == p.Val) {
			p.Next = head
			p = p.Next

		} else {
			q.Next = head
			q = q.Next
		}
		head = head.Next
		p.Next = nil
		q.Next = nil
	}

	return dummy2.Next

}

func main() {
	compute(2)

	prices := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
	return
	generateParenthesis(3)
	return
	TestSortList()
	return
	TestIsPalindrome()
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

	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	ress := spiralOrderII(matrix)
	fmt.Println("spiralOrder", ress)

}
