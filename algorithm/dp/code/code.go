package main

import "fmt"

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

func main() {
	fmt.Println(fib(6))
}
