package main

import (
	"container/list"
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	res = append(res, root.Val)
	fmt.Println(root.Val)
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return res
}

// 中序遍历
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

// 后序遍历
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, inorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return res
}

// 层序遍历
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		level := []int{}
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			res = append(res, level)
		}

	}

	return res

}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return max(left, right) + 1
}

func maxDepthII(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushBack(root)
	depth := 0
	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		depth++
	}
	return depth
}

/*
*
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：

节点的左
子树
只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
*/
func isValidBST(root *TreeNode) bool {
	return helper(root, nil, nil)
}

func helper(node *TreeNode, max, min *int) bool {
	if node == nil {
		return true
	}

	if min != nil && node.Val <= *min {
		return false
	}
	if max != nil && node.Val >= *max {
		return false
	}

	return helper(node.Left, min, &node.Val) && helper(node.Right, &node.Val, max)
}

/*
*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
*/
func mainc() {
	TestTr()
}

func TestTr() {
	root1 := &TreeNode{Val: 1}
	root2 := &TreeNode{Val: 2}
	root3 := &TreeNode{Val: 3}
	root4 := &TreeNode{Val: 4}
	root5 := &TreeNode{Val: 5}
	root6 := &TreeNode{Val: 6}
	root7 := &TreeNode{Val: 7}
	root8 := &TreeNode{Val: 8}
	root1.Left = root2
	root1.Right = root5
	root2.Left = root3
	root2.Right = root4
	root5.Left = root6
	root3.Left = root7
	root3.Right = root8

	s := lowestCommonAncestor(root1, root7, root4)
	fmt.Println(s)
}

func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

/**
      3
     / \
    5   1
   / \      \
  6   2     9
 /
8
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil || root == p || root == q {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)

	if l != nil && r != nil {
		return root
	}

	if l != nil {
		return l
	}

	return r

}

// 🚀 根据数组构建二叉树
func buildTreeFromArray(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	// 根节点
	root := &TreeNode{Val: arr[0]}
	queue := []*TreeNode{root}

	i := 1 // 遍历数组的索引
	for i < len(arr) {
		curr := queue[0] // 取出队列头部
		queue = queue[1:]

		// 添加左子节点
		if i < len(arr) && arr[i] != -1 {
			curr.Left = &TreeNode{Val: arr[i]}
			queue = append(queue, curr.Left)
		}
		i++

		// 添加右子节点
		if i < len(arr) && arr[i] != -1 {
			curr.Right = &TreeNode{Val: arr[i]}
			queue = append(queue, curr.Right)
		}
		i++
	}
	return root
}
