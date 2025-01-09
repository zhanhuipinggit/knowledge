package main

import (
	"fmt"
	"math"
)

// 动态规划解斐波那契数列
func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	a, b := 1, 1
	var c int
	for i := 3; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}

/* 0-1 背包：空间优化后的动态规划 */
func knapsackDPComp(wgt, val []int, cap int) int {
	n := len(wgt)
	// 初始化 dp 表
	dp := make([]int, cap+1)
	// 状态转移
	for i := 1; i <= n; i++ {
		// 倒序遍历
		for c := cap; c >= 1; c-- {
			if wgt[i-1] <= c {
				// 不选和选物品 i 这两种方案的较大值
				dp[c] = int(math.Max(float64(dp[c]), float64(dp[c-wgt[i-1]]+val[i-1])))
			}
		}
	}
	return dp[cap]
}

/**
给定一个楼梯，你每步可以上
 1阶或者2阶，每一阶楼梯上都贴有一个非负整数，表示你在该台阶所需要付出的代价。给定一个非负整数数组cost
 ，其中cost[i]表示在第i个台阶需要付出的代价，为地面（起始点）。请计算最少需要付出多少代价才能到达顶部？
cost climbing stairs
*/
// 状态转移方程：d[i] = min(d[i-1],d[i-2])+cost[i]
func minCostClimbingStairsDp(cost []int) int {
	n := len(cost) - 1
	if n == 1 || n == 0 {
		return cost[n]
	}

	min := func(i, j int) int {
		if i > j {
			return j
		}
		return i
	}
	var dp []int
	dp[1] = cost[1]
	dp[2] = cost[2]
	for i := 3; i < n; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return dp[n]

}

// 空间复杂度为1常数的方法

func minCostClimbingStairsDPS(cost []int) int {
	n := len(cost) - 1
	if n == 1 || n == 0 {
		return cost[n]
	}

	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}

	a := cost[1]
	b := cost[2]
	for i := 3; i < n; i++ {
		temp := min(a, b) + cost[i]
		a = b
		b = temp
	}
	return b

}

/*
*
给定一个共有n阶的楼梯，你每步可以上1阶或者2阶，但不能连续两轮跳

	1阶，请问有多少种方案可以爬到楼顶？

状态转移方程：
d[i][[1] = d[i-1][2]
d[i][2] = d[i-2][1]+d[i-2][2]
*/
func climbingStirsNums(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	dp := make([][3]int, n+1)
	dp[1][1] = 1
	dp[1][2] = 0
	dp[2][1] = 0
	dp[2][2] = 1
	for i := 3; i < n; i++ {
		dp[i][1] = dp[i-1][2]
		dp[i][2] = dp[i-2][1] + dp[i-2][2]
	}
	return dp[n][1] + dp[n][2]

}

/*
*
给定一个共有n

	阶的楼梯，你每步可以上1
	阶或者2
	阶。规定当爬到第i
	阶时，系统自动会在第2i
	阶上放上障碍物，之后所有轮都不允许跳到第
	阶上。例如，前两轮分别跳到了第2,3阶上，则之后就不能跳到第4、6 阶上。请问有多少种方案可以爬到楼顶？

climb stair with obstacles

restricted 有限制
有一些思考不需要太过深入，只需要掌握基础原则和限制条件
*/
func climbStairWithObstacles(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	restricted := make(map[int]bool, n)
	dp := make([]int, n)
	dp[1] = 1
	dp[2] = 1
	for i := 3; i < n; i++ {
		if !restricted[i] {
			dp[i] = dp[i-1] + dp[i-2]
			if 2*i < n {
				restricted[2*i] = true
			}

		}
	}
	return dp[n]

}

func main() {
	fmt.Println(fib(6))
}
