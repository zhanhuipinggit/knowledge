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

func main() {
	fmt.Println(fib(6))
}
