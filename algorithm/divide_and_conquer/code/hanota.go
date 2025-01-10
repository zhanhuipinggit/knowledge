package code

import "container/list"

/**
给定三根柱子，记为 A、B 和 C 。起始状态下，柱子 A 上套着
 个圆盘，它们从上到下按照从小到大的顺序排列。我们的任务是要把这
 个圆盘移到柱子 C 上，并保持它们的原有顺序不变（如图 12-10 所示）。在移动圆盘的过程中，需要遵守以下规则。

圆盘只能从一根柱子顶部拿出，从另一根柱子顶部放入。
每次只能移动一个圆盘。
小圆盘必须时刻位于大圆盘之上。
*/

func move(src, tar *list.List) {
	pan := src.Back()
	tar.PushBack(pan.Value)
	src.Remove(pan)
}

func dfsHanota(i int, src, buf, tar *list.List) {
	if i == 0 {
		move(src, tar)
		return
	}

	dfsHanota(i-1, src, tar, buf)
	move(src, tar)
	dfsHanota(i-1, buf, src, tar)
}

/* 求解汉诺塔问题 */
func solveHanota(A, B, C *list.List) {
	n := A.Len()
	// 将 A 顶部 n 个圆盘借助 B 移到 C
	dfsHanota(n, A, B, C)
}
