package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func main() {
	var l1 ListNode
	l1.Value = 1

	var l2 ListNode
	l2.Value = 5

	var l3 ListNode
	l3.Value = 9

	var l4 ListNode
	l4.Value = 11

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l2

	flag := hasCycle(&l1)
	fmt.Println(flag)
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	var tmpMap = make(map[*ListNode]int)

	for head != nil {
		if _, ok := tmpMap[head]; ok {
			return true
		} else {
			tmpMap[head] = 1
		}

		head = head.Next
	}

	return false
}

func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	var slow = head
	var fast = head.Next

	for fast != nil && fast.Next != nil {
		if slow != fast {
			slow = slow.Next
			fast = fast.Next.Next
		} else {
			return true
		}
	}

	return false
}
