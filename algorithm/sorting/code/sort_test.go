package code

import (
	"fmt"
	"reflect"
	"testing"
)

// TestSelectSort
func TestSelectSort(t *testing.T) {
	testCases := []struct {
		input  []int
		expect []int
	}{
		{[]int{}, []int{}},               // 测试空数组
		{[]int{5}, []int{5}},             // 测试单元素数组
		{[]int{3, 1, 2}, []int{1, 2, 3}}, // 测试普通数组
		{[]int{4, 2, 7, 1, 8, 3, 6, 5}, []int{1, 2, 3, 4, 5, 6, 7, 8}},       // 测试较大数组
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}, // 测试逆序数组
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}, // 测试已排序数组
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {
			selectSort(tc.input)
			if !reflect.DeepEqual(tc.input, tc.expect) {
				t.Errorf("selectSort(%v) = %v; want %v", tc.input, tc.input, tc.expect)
			}
		})
	}
}

// TestHeapSort
func TestHeapSort(t *testing.T) {
	testCases := []struct {
		input  []int
		expect []int
	}{
		{[]int{}, []int{}},               // 测试空数组
		{[]int{5}, []int{5}},             // 测试单元素数组
		{[]int{3, 1, 2}, []int{1, 2, 3}}, // 测试普通数组
		{[]int{4, 2, 7, 1, 8, 3, 6, 5}, []int{1, 2, 3, 4, 5, 6, 7, 8}},       // 测试较大数组
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}, // 测试逆序数组
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}, // 测试已排序数组
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {
			numsCopy := make([]int, len(tc.input))
			copy(numsCopy, tc.input) // 复制输入以保留原始数据
			heapSort(&numsCopy)
			if !reflect.DeepEqual(numsCopy, tc.expect) {
				t.Errorf("heapSort(%v) = %v; want %v", tc.input, numsCopy, tc.expect)
			}
		})
	}
}

// TestBubbleSort tests the bubbleSort function with various input slices.
func TestBubbleSort(t *testing.T) {
	testCases := []struct {
		input  []int
		expect []int
	}{
		{[]int{}, []int{}},               // 测试空数组
		{[]int{5}, []int{5}},             // 测试单元素数组
		{[]int{3, 1, 2}, []int{1, 2, 3}}, // 测试普通数组
		{[]int{4, 2, 7, 1, 8, 3, 6, 5}, []int{1, 2, 3, 4, 5, 6, 7, 8}},       // 测试较大数组
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}, // 测试逆序数组
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}, // 测试已排序数组
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {
			numsCopy := make([]int, len(tc.input))
			copy(numsCopy, tc.input) // 复制输入以保留原始数据
			bubbleSort(&numsCopy)
			if !reflect.DeepEqual(numsCopy, tc.expect) {
				t.Errorf("bubbleSort(%v) = %v; want %v", tc.input, numsCopy, tc.expect)
			}
		})
	}
}
