package main

import (
	"math"
)

/**
给定n个物品，第i个物品的重量为wgt[i-1]、价值为ual[i-1]，和一个容量为cap
 的背包。每个物品只能选择一次，问在限定背包容量下能放入物品的最大价值。
背包：knapscak
i当前物品 c 代表重量
1：如果放入当前物品，则c重量-wgt[i-1]
0： 如果不放入当前物品，则c重量不变
放入和不放入两种状态取价值最大的
状态转移方程：dp[i,c] = max(dp[i-1,c],dp[i-1, c-wgt[i-1]+ual[i-1])
*/

// 暴力求解
func knapscakDFS(wgt, ual []int, i, c int) int {
	//选完所有物品和背包无空间了
	if i == 0 || c == 0 {
		return 0
	}
	if wgt[i-1] > c {
		return knapscakDFS(wgt, ual, i-1, c)
	}

	no := knapscakDFS(wgt, ual, i-1, c)
	yes := knapscakDFS(wgt, ual, i-1, c-wgt[i-1]) + ual[i-1]
	// 返回两种方案中价值更大的那一个
	return int(math.Max(float64(no), float64(yes)))

}

// 暴力求解+dp备忘录
func knapscakDFSDP(wgt, ual []int, i, c int, mem [][]int) int {
	//选完所有物品和背包无空间了
	if i == 0 || c == 0 {
		return 0
	}
	if mem[i][c] != -1 {
		return mem[i][c]
	}

	if wgt[i-1] > c {
		return knapscakDFS(wgt, ual, i-1, c)
	}

	no := knapscakDFS(wgt, ual, i-1, c)
	yes := knapscakDFS(wgt, ual, i-1, c-wgt[i-1]) + ual[i-1]
	// 返回两种方案中价值更大的那一个
	mem[i][c] = int(math.Max(float64(no), float64(yes)))
	return mem[i][c]
}

// knapsackDP 动态规划
func knapsackDP(wgt, ual []int, cap int) int {
	n := len(wgt)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, cap+1)
	}

	// 状态转移
	for i := 1; i < n; i++ {
		for j := 1; j < cap; j++ {
			if wgt[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i-1][j-wgt[i-1]]+ual[i-1])))
			}
		}
	}
	return dp[n][cap]

}
