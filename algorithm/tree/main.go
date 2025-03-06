package main

import (
	"fmt"
	"math"
	"strconv"
)

// 快排
func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivot := nums[len(nums)/2]
	left, mid, right := []int{}, []int{}, []int{}
	for num := range nums {
		if num == pivot {
			mid = append(mid, num)
		} else if num > pivot {
			right = append(right, num)
		} else {
			left = append(left, num)
		}
	}

	left = quickSort(left)
	right = quickSort(right)
	return append(append(left, mid...), right...)
}

// 归并排序
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

// 合并两个已排序数组
func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// 递归遍历单链表，倒序打印链表元素
type ListNode struct {
	Val  int
	Next *ListNode
}

func traverse(head *ListNode) {
	if head == nil {
		return
	}
	traverse(head.Next)
	// 后序位置
	fmt.Println(head.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树重复数
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	memo := make(map[string]int)
	res := []*TreeNode{}
	serialize(root, &memo, &res)
	return res

}

func serialize(root *TreeNode, memo *map[string]int, res *[]*TreeNode) string {
	if root == nil {
		return "#"
	}

	left := serialize(root.Left, memo, res)
	right := serialize(root.Right, memo, res)
	subTree := left + "," + right + "," + strconv.Itoa(root.Val)
	count, exists := (*memo)[subTree]
	if exists && count == 1 {
		*res = append(*res, root)
	}

	(*memo)[subTree] = count + 1
	return subTree

}

// 二叉树的直径
func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	var maxDepth func(*TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := maxDepth(root.Left)
		right := maxDepth(root.Right)
		maxD := left + right
		maxDiameter = max(maxDiameter, maxD)
		return max(left, right) + 1
	}

	maxDepth(root)
	return maxDiameter
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	path := []int{}
	res := []string{}
	var dfs func(node *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		path = append(path, root.Val)
		if root.Left == nil && root.Right == nil {
			pString := ""
			for i := 0; i < len(path); i++ {
				if i == len(path)-1 {
					pString = pString + strconv.Itoa(path[i])
				} else {
					pString = pString + strconv.Itoa(path[i]) + "->"
				}
			}
			res = append(res, pString)
			return
		}
		dfs(root.Left)
		dfs(root.Right)
		path = path[:len(path)-1]
	}
	dfs(root)
	return res

}

func sumNumbers(root *TreeNode) int {
	res := 0
	path := ""
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		path += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			num, _ := strconv.Atoi(path)
			res += num
		}
		dfs(root.Left)
		dfs(root.Right)
		path = path[:len(path)-1]

	}

	dfs(root)
	return res
}

/**
199. 二叉树的右视图
中等
相关标签
相关企业
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
*/

func rightSideView(root *TreeNode) []int {
	queue := []*TreeNode{root}
	res := []int{}
	for len(queue) > 0 {
		last := queue[0].Val
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
		}
		res = append(res, last)
	}
	return res
}

/*
*
988. 从叶结点开始的最小字符串
中等
相关标签
相关企业
给定一颗根结点为 root 的二叉树，树中的每一个结点都有一个 [0, 25] 范围内的值，分别代表字母 'a' 到 'z'。

返回 按字典序最小 的字符串，该字符串从这棵树的一个叶结点开始，到根结点结束。

注：字符串中任何较短的前缀在 字典序上 都是 较小 的：

例如，在字典序上 "ab" 比 "aba" 要小。叶结点是指没有子结点的结点。
节点的叶节点是没有子节点的节点。
*/
func smallestFromLeaf(root *TreeNode) string {
	res := ""
	path := []byte{}
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		path = append(path, byte('a'+root.Val))
		if root.Left == nil && root.Right == nil {
			reverse(&path)
			s := string(path)
			if s != "" && res > s {
				res = s
			}
			reverse(&path)
			path = path[:len(path)-1]
			return
		}
		dfs(root.Right)
		dfs(root.Left)
		path = path[:len(path)-1]
	}
	dfs(root)
	return res
}

func reverse(res *[]byte) {
	for i, j := 0, len(*res)-1; i < j; i, j = i+1, j-1 {
		(*res)[i], (*res)[j] = (*res)[j], (*res)[i]
	}

}

// 零钱兑换
func slave(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
	}

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin && dp[i-coin] == math.MaxInt {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	return dp[amount]
}

// 链表排序
func mergeList(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	prev := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	return dummy.Next
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

func deleteDuplicatesII(head *ListNode) *ListNode {
	dummy2 := &ListNode{Val: 101}
	p := dummy2
	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			head = head.Next
		} else {
			p.Next = head
			p = p.Next
			head = head.Next
		}
	}

	return dummy2.Next

}

// 链表去重
func uniqueList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	prev := dummy
	curr := head
	for curr != nil {
		if curr.Next != nil && curr.Val == curr.Next.Val {
			for curr.Next != nil && curr.Val == curr.Next.Val {
				curr = curr.Next
			}
		}
		prev.Next = curr
		prev = prev.Next
		curr = curr.Next
	}
	return dummy.Next
}

func reverseList(l *ListNode) *ListNode {
	var newH *ListNode
	tailH := newH
	var prev *ListNode
	curr := l
	for curr != nil {
		node := &ListNode{Val: curr.Val}
		if newH == nil {
			newH = node
			tailH = node
		} else {
			tailH.Next = node
			tailH = tailH.Next
		}

		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next

	}
	println("ccc")
	return prev

}

func reorderListI(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// 1. 找到链表中点（快慢指针法）
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 2. 反转后半部分链表
	second := reverseLists(slow.Next)
	slow.Next = nil // 断开前后两部分

	// 3. 交错合并两个链表
	first := head
	mergeLists(first, second)
}

// 反转链表
func reverseLists(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

// 交错合并两个链表
func mergeLists(l1, l2 *ListNode) {
	for l1 != nil && l2 != nil {
		l1Next, l2Next := l1.Next, l2.Next
		l1.Next = l2
		if l1Next == nil {
			break
		}
		l2.Next = l1Next
		l1, l2 = l1Next, l2Next
	}
}

func main() {
	l1 := &ListNode{1, nil}
	l2 := &ListNode{10, nil}
	l3 := &ListNode{11, nil}
	l4 := &ListNode{12, nil}
	//l5 := &ListNode{20, nil}
	//l6 := &ListNode{20, nil}
	//l7 := &ListNode{30, nil}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	//l4.Next = l5
	//l5.Next = l6
	//l6.Next = l7

	reorderListI(l1)

	r1 := &ListNode{3, nil}
	r2 := &ListNode{2, nil}
	r3 := &ListNode{0, nil}
	r4 := &ListNode{-4, nil}
	r1.Next = r2
	r2.Next = r3
	r3.Next = r4
	r4.Next = r2

	mergeList(l1, r1)

	//s := mergeTwoList(l1, r1)
	//for s != nil {
	//	fmt.Println(s.Val)
	//	s = s.Next
	//}

	return
	//
	//head := &ListNode{nil, 1}
	//head2 := &ListNode{nil, 2}
	//head3 := &ListNode{nil, 3}
	//head.Next = head2
	//head2.Next = head3
	//
	//traverse(head)
	//
	//arr := []int{3, 6, 8, 10, 1, 2, 1}
	//sortedArr := mergeSort(arr)
	//fmt.Println("归并排序结果:", sortedArr)
}
