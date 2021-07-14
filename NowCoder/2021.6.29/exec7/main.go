package main

import "fmt"

type ListNode struct {
	value int
	next *ListNode
}

func main() {
	var l1 ListNode
	l1.value = 1

	var l2 ListNode
	l2.value = 2

	var l3 ListNode
	l3.value = 3

	var l4 ListNode
	l4.value = 4

	var l5 ListNode
	l5.value = 5

	var l6 ListNode
	l6.value = 6

	var l7 ListNode
	l7.value = 7

	l1.next = &l2
	l2.next = &l3
	l3.next = &l4
	l4.next = &l5
	l5.next = &l6
	l6.next = &l7

	result := reverseList(&l1)

	fmt.Println(result.value)
	fmt.Println(result.next.value)
	fmt.Println(result.next.next.value)
	fmt.Println(result.next.next.next.value)
	fmt.Println(result.next.next.next.next.value)
	fmt.Println(result.next.next.next.next.next.value)
	fmt.Println(result.next.next.next.next.next.next.value)
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.next == nil{
		return head
	}

	var dummyNode *ListNode

	for head != nil{
		head, head.next, dummyNode = head.next, dummyNode, head
	}

	return dummyNode
}