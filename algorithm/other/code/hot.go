package main

import (
	"fmt"
	"math"
	"os"
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

/**
24. 两两交换链表中的节点
中等
相关标签
相关企业
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）
*/

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	perv := dummy
	for perv.Next != nil && perv.Next.Next != nil {
		fist := perv.Next
		second := fist.Next
		nextPail := second.Next
		//交换
		second.Next = fist
		fist.Next = nextPail
		perv.Next = second
		perv = fist
	}

	return dummy.Next

}

/**
给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
*/

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast.Next == nil || fast.Next.Next == nil {
		return nil
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow

}

/**
给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。

k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
1,2,3,4,5 k=2
2,1,4,3,5

1,2,3,4,5 k=3
3,2,1,4,5

*/

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	// 计算链表长度
	length := 0
	temp := head
	for temp != nil {
		length++
		temp = temp.Next
	}

	// 创建一个哑节点，方便操作
	dummy := &ListNode{Next: head}
	prevGroupEnd := dummy

	for length >= k {
		// 反转 k 个节点
		prev := prevGroupEnd
		curr := prev.Next
		next := curr.Next
		for i := 1; i < k; i++ {
			curr.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
			next = curr.Next
		}

		// 移动 prevGroupEnd 到下一组的前一个位置
		prevGroupEnd = curr
		length -= k
	}

	return dummy.Next

}

/*
*

322. 零钱兑换
中等
相关标签
相关企业
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

你可以认为每种硬币的数量是无限的。
状态转移方程：d[n][m] = min(d[n-1][m], d[n-1][m+coins[
*/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 0; i < amount; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 0; i <= amount; i++ {
		for coin := range coins {
			if i >= coin && dp[i-coin] != math.MaxInt32 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]

}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

/*
*
200. 岛屿数量
中等
相关标签
相关企业
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。
*/

func isLands(grid [][]byte, row, col int) {

}

func numIslands(grid [][]byte) int {

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 判断是否是岛屿
		}
	}
	return 0
}

/**
207. 课程表
中等
相关标签
相关企业
提示
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。

在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。

例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。



示例 1：

输入：numCourses = 2, prerequisites = [[1,0]]
输出：true
解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
示例 2：

输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
输出：false
解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)
	for _, pre := range prerequisites {
		course, prerequisite := pre[0], pre[1]
		graph[prerequisite] = append(graph[prerequisite], course)
		inDegree[course]++
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	var count int
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		count++
		for _, neighbor := range graph[course] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	return count == numCourses

}

/*
*
51. N 皇后
困难
相关标签
相关企业
按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。

n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。

每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
*/
func solveNQueens(n int) [][]string {
	// 判断是否符合规格
	var isSafe func([][]int, int, int) bool
	isSafe = func(board [][]int, row, col int) bool {
		// 正上方检测
		for i := row - 1; i >= 0; i-- {
			if board[i][col] == 1 {
				return false
			}
		}

		// 左上方
		for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if board[i][j] == 1 {
				return false
			}
		}

		// 右上方
		for i, j := row, col; i >= 0 && j < n; i, j = i-1, j+1 {
			if board[i][j] == 1 {
				return false
			}
		}
		return true
	}
	var backtrack func(*[][]string, int, [][]int)
	backtrack = func(res *[][]string, row int, broad [][]int) {
		if row == n {
			subStrArr := make([]string, n)
			for i := 0; i < n; i++ {
				subStr := ""
				for j := 0; j < n; j++ {
					if broad[i][j] == 1 {
						subStr = subStr + "Q"
					} else {
						subStr = subStr + "."
					}
				}
				subStrArr[i] = subStr
			}
			*res = append(*res, subStrArr)
			return
		}

		for j := 0; j < n; j++ {
			if isSafe(broad, row, j) {
				broad[row][j] = 1
				backtrack(res, row+1, broad)
				// 回溯，移除皇后
				broad[row][j] = 0
			}
		}
	}

	var res [][]string
	board := make([][]int, n)
	for i := range board {
		board[i] = make([]int, n)
	}
	backtrack(&res, 0, board)

	return res

}

/**
79. 单词搜索
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
*/

func main() {

	n := 4
	result := solveNQueens(n)
	//fmt.Println(result)
	for _, solution := range result {
		for _, row := range solution {
			fmt.Println(row)
		}
		fmt.Println("---------------------")
	}
	return
	numCourses := 2
	prerequisites := [][]int{{1, 0}, {0, 1}}

	isCan := canFinish(numCourses, prerequisites)
	fmt.Println(isCan)
	os.Exit(-1)
	return

	//TestTr()
	//// 构建数结构
	//root := buildTreeFromArray([]int{3, 9, 20, -1, -1, 15, 7})
	//
	//preorderTraversal(root)

	return
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

	swapPairsL := buildLinkedList([]int{1, 2, 3, 4})
	swap := swapPairs(swapPairsL)
	printList(swap)

	// 返回环节点
	detectCycleL := &ListNode{}
	d1 := &ListNode{}
	d2 := &ListNode{}
	d3 := &ListNode{}
	detectCycleL.Next = d1
	detectCycleL.Val = 3
	d1.Next = d2
	d1.Val = 2
	d2.Next = d3
	d2.Val = 0
	d3.Next = d1
	d3.Val = -4
	//detectCycleL := buildLinkedList([]int{3, 2, 0, -4})
	detectCycle(detectCycleL)
	//printList(detect)

	reverseKGroupL := buildLinkedList([]int{1, 2, 3, 4, 5})

	reverseKGroup(reverseKGroupL, 2)

}
