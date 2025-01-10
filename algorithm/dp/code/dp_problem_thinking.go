package main

/*
*
给定一个n*m

	的二维网格 grid ，网格中的每个单元格包含一个非负整数，表示该单元格的代价。

机器人以左上角单元格为起始点，每次只能向下或者向右移动一步，直至到达右下角单元格。请返回从左上角到右下角的最小路径和。
dp[i,j] = min(dp[i-1,j], dp(i, j-1)) + grid[i,j]
i >= 0 && i < n
j >= && j < m
*/
func minPathSumDp(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}

	dp[0][0] = grid[0][0]

	n, m := len(grid), len(grid[0])

	for i := 0; i < n; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for k := 0; k < m; k++ {
		dp[0][k] = dp[0][k-1] + grid[0][k]
	}

	min := func(j, k int) int {
		if j < k {
			return j
		}
		return k
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[n-1][m-1]

}

/*
*
空间优化
*/
func minPathSumDp1(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	dp := make([]int, m)
	dp[0] = grid[0][0]
	for j := 1; j < m; j++ {
		dp[j] = dp[j] + grid[0][j]
	}

	mins := func(i, j int) int {
		if i > j {
			return j
		}
		return i
	}

	for i := 1; i < n; i++ {
		dp[0] = dp[0] + grid[i][0]
		for j := 1; j < m; j++ {
			dp[j] = mins(dp[j-1], dp[j]) + grid[i][j]
		}
	}
	return dp[m-1]

}
