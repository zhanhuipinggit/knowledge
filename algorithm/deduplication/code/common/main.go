package main

import (
	"fmt"
)

func removeDuplicatesRun() {
	nums := []int{1, 2, 2, 3, 4, 4, 5}
	removeDuplicates(nums)
	fmt.Println(nums)
}

func binarySearchRun() {
	nums := []int{1, 2, 3, 5, 6}
	index := binarySearch(nums, 4)
	fmt.Println(index)
}

func main() {
	binarySearchRun()
	return
	l1 := &ListNode{1, nil}
	l2 := &ListNode{10, nil}
	l3 := &ListNode{8, nil}
	l4 := &ListNode{8, nil}
	l5 := &ListNode{20, nil}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5

	r1 := &ListNode{3, nil}
	r2 := &ListNode{2, nil}
	r3 := &ListNode{0, nil}
	r4 := &ListNode{-4, nil}
	r1.Next = r2
	r2.Next = r3
	r3.Next = r4
	r4.Next = r2

	//s := mergeTwoList(l1, r1)
	//for s != nil {
	//	fmt.Println(s.Val)
	//	s = s.Next
	//}

	//delListNodeUnique(l1)
	//for l1 != nil {
	//	fmt.Println(l1.Val)
	//	l1 = l1.Next
	//}

	//s3 := partition(l1, 8)
	//for s3 != nil {
	//	fmt.Println(s3.Val)
	//	s3 = s3.Next
	//}

	//s3 := detectCycle(r1)
	//fmt.Println(s3.Val)
	//re := &ListNode{1, nil}
	//s4 := removeNthFromEnd(l1, 1)
	//for s4 != nil {
	//	fmt.Println(s4.Val)
	//	s4 = s4.Next
	//
	//}

	RemoveDuplicatesList(l1)
	for l1 != nil {
		fmt.Println(l1.Val)
		l1 = l1.Next
	}

}
