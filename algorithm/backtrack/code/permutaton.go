package main

import (
	"fmt"
)

// 全排列
func backtrackI(nums []int, res *[][]int, path []int, used map[int]bool) {
	// 如果路径长度等于数组长度，说明找到一个完整的排列
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path) // 复制当前路径
		*res = append(*res, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		// 如果当前元素已经被使用，跳过
		if used[i] {
			continue
		}

		// 做选择
		path = append(path, nums[i])
		used[i] = true

		// 递归进入下一层
		backtrackI(nums, res, path, used)

		// 撤销选择
		path = path[:len(path)-1]
		used[i] = false
	}
}

func permute(nums []int) [][]int {
	var res [][]int
	used := make(map[int]bool)
	backtrackI(nums, &res, []int{}, used)
	return res
}

/*
*
给定一个正整数数组 nums 和一个目标正整数 target ，请找出所有可能的组合，使得组合中的元素和等于 target 。
给定数组无重复元素，每个元素可以被选取多次。请以列表形式返回这些组合，列表中不应包含重复组合。
*/
func backtrackII(nums []int, start int, res *[][]int, path []int, target int) {
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		if target-nums[i] < 0 {
			continue
		}
		backtrackII(nums, i, res, path, target-nums[i])
		path = path[:len(path)-1]
	}
}

func permuteII(nums []int, target int) [][]int {
	var res [][]int
	backtrackII(nums, 0, &res, []int{}, target)
	return res
}

/**
给定一个正整数数组 nums 和一个目标正整数 target ，请找出所有可能的组合，使得组合中的元素和等于 target 。
给定数组可能包含重复元素，每个元素只可被选择一次。请以列表形式返回这些组合，列表中不应包含重复组合。
*/

func backtrackIII(nums []int, res *[][]int, path []int, start int, target int) {
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	for i := start; i < len(nums); i++ {
		if target-nums[i] < 0 {
			continue
		}

		path = append(path, nums[i])
		backtrackIII(nums, res, path, i+1, target-nums[i])
		path = path[:len(path)-1]

	}

}

func permuteIII(nums []int, target int) [][]int {
	var res [][]int
	backtrackIII(nums, &res, []int{}, 0, target)
	return res
}

/*
*
根据国际象棋的规则，皇后可以攻击与同处一行、一列或一条斜线上的棋子。给定n个皇后和一个n*n大小的棋盘，寻找使得所有皇后之间无法相互攻击的摆放方案。
*/
func isSafe(broad [][]int, n, row, col int) bool {
	//// 检查当前列是否有皇后
	//for i := 0; i < col; i++ {
	//	// 当前位置有，则不能放置
	//	if broad[row][i] == 1 {
	//		return false
	//	}
	//}

	// 检查当前列是否有皇后
	for i := 0; i < row; i++ { // 修正：遍历当前列的所有行
		if broad[i][col] == 1 {
			return false
		}
	}

	// 检查左上对角线是否能放置
	for i, j := row, col; j >= 0 && i >= 0; j, i = j-1, i-1 {
		if broad[i][j] == 1 {
			return false
		}
	}

	// 检查右上角是否能放置
	for i, j := row, col; j < n && i >= 0; i, j = i-1, j+1 {
		if broad[i][j] == 1 {
			return false
		}
	}

	return true

}

// n皇后问题
// isSafe 检查当前位置 (row, col) 是否可以放置皇后
//func isSafe(board [][]int, row, col, n int) bool {
//	// 检查当前行是否有皇后
//	for i := 0; i < col; i++ {
//		if board[row][i] == 1 {
//			return false
//		}
//	}
//
//	// 检查左上对角线是否有皇后
//	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
//		if board[i][j] == 1 {
//			return false
//		}
//	}
//
//	// 检查左下对角线是否有皇后
//	for i, j := row, col; i < n && j >= 0; i, j = i+1, j-1 {
//		if board[i][j] == 1 {
//			return false
//		}
//	}
//
//	return true
//}

func nQueues(res *[][]string, row, n int, broad [][]int) {
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

	//for i := 0; i < n; i++ {
	//	if isSafe(broad, n, i, col) {
	//		broad[i][col] = 1
	//		nQueues(res, col+1, n, broad)
	//		// 回溯，移除皇后
	//		broad[i][col] = 0
	//	}
	//}

	for j := 0; j < n; j++ {
		if isSafe(broad, n, row, j) {
			broad[row][j] = 1
			nQueues(res, row+1, n, broad)
			// 回溯，移除皇后
			broad[row][j] = 0
		}
	}

}

// solveNQueens 解决 N 皇后问题
func solveNQueens(n int) [][]string {
	var result [][]string
	board := make([][]int, n)
	for i := range board {
		board[i] = make([]int, n)
	}

	nQueues(&result, 0, n, board)
	return result
}

func main() {
	//nums := []int{1, 2, 3}
	//permutations := permute(nums)
	//for _, perm := range permutations {
	//	fmt.Println(perm)
	//}

	//permutationsII := permuteIII(nums, 3)
	//for _, perm := range permutationsII {
	//	fmt.Println(perm)
	//}

	n := 4
	result := solveNQueens(n)
	//fmt.Println(result)
	for _, solution := range result {
		for _, row := range solution {
			fmt.Println(row)
		}
		fmt.Println("---------------------")
	}
}
