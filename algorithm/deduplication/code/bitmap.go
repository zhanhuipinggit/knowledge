package main

import (
	"fmt"
	"sort"
)

func bitmapUnique(nums []int) []int {
	const maxNumber = 10000000000
	bitmap := make([]byte, maxNumber/8+1)
	for _, num := range nums {
		byteIndex := num / 8
		bitIndex := num % 8
		bitmap[byteIndex] = 1 << bitIndex
	}

	result := []int{}
	for i, b := range bitmap {
		for j := 0; j < 8; j++ {
			if b&(1<<j) != 0 {
				result = append(result, i*8+j)
			}
		}

	}

	return result

}

func bitmapUniqueII(numbers []int) []int {
	sort.Slice(numbers, func(i, j int) bool {
		return i > j
	})
	maxNumber := numbers[0]
	bitmap := make([]byte, maxNumber/8+1)

	// 设置位图
	for _, num := range numbers {
		byteIndex := num / 8
		bitIndex := num % 8
		bitmap[byteIndex] |= 1 << bitIndex
	}

	// 提取去重结果
	result := []int{}
	for i, b := range bitmap {
		for j := 0; j < 8; j++ {
			if b&(1<<j) != 0 {
				result = append(result, i*8+j)
			}
		}
	}

	return result
}

func main() {
	numbers := []int{10, 8, 9}
	uniqueNumbers := bitmapUniqueII(numbers)
	fmt.Println(uniqueNumbers)
}
