package main

import "fmt"

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
			if b&(1<<j) != 1 {
				result = append(result, i*8+j)
			}
		}

	}

	return result

}

func main() {
	numbers := []int{1234567890, 1234567890, 9876543210}
	uniqueNumbers := bitmapUnique(numbers)
	fmt.Println(uniqueNumbers)
}
