package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 合并两个有序链表
func mergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	}

	if l2 != nil {
		current.Next = l2
	}

	return dummy.Next

}

// 链表去重
func delListNodeUnique(head *ListNode) {
	if head == nil {
		return
	}

	slow := head
	fast := head
	for fast != nil {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}

}

/*
*
求数组倒数第几步
*/
func findFromEnd(l1 *ListNode, k int) *ListNode {
	p1 := l1
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}

	p2 := l1
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}

// 删除指定的倒数第几个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	pre := dummy
	p1, p2 := head, head

	for i := 0; i < n; i++ {
		p1 = p1.Next
	}

	for p1 != nil {
		p1 = p1.Next
		pre.Next = p2
		pre = pre.Next
		p2 = p2.Next
	}
	pre.Next = p2.Next

	return dummy.Next
}

// 分割两个链表
func partition(head *ListNode, x int) *ListNode {
	dummy1 := &ListNode{}
	dummy2 := &ListNode{}
	p1, p2 := dummy1, dummy2
	current := head
	for current != nil {
		if current.Val <= x {
			p1.Next = current
			p1 = p1.Next
		} else {
			p2.Next = current
			p2 = p2.Next
		}
		// 断开节点
		temp := current.Next
		current.Next = nil
		current = temp
	}
	p1.Next = dummy2.Next
	return dummy1.Next
}

// 两个链表是否相交
func getIntersectionNode(head1 *ListNode, head2 *ListNode) *ListNode {
	p1 := head1
	p2 := head2
	for p1 != p2 {
		if p1 != nil {
			p1 = p1.Next
		} else {
			p1.Next = head2
		}

		if p2 != nil {
			p2 = p2.Next
		} else {
			p2.Next = head1
		}

	}
	return p1
}

// 判断是否有环
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}

	if fast == nil || slow == nil {
		return nil
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow

}

func main() {
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
	s4 := removeNthFromEnd(l1, 1)
	for s4 != nil {
		fmt.Println(s4.Val)
		s4 = s4.Next

	}

}
