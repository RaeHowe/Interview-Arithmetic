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

	l1.next = &l2
	l2.next = &l3
	l3.next = &l4
	l4.next = &l5
	l5.next = &l6

	result := swapNode(&l1)

	fmt.Println(result.value)
	fmt.Println(result.next.value)
	fmt.Println(result.next.next.value)
	fmt.Println(result.next.next.next.value)
	fmt.Println(result.next.next.next.next.value)
	fmt.Println(result.next.next.next.next.next.value)
}

func swapNode(head *ListNode) *ListNode {
	if head == nil{
		return head
	}

	var dummyNode = &ListNode{}

	var pre = dummyNode
	pre.next = head

	for head != nil && head.next != nil{
		pre.next = head.next
		head.next = head.next.next
		pre.next.next = head

		pre = head
		head = head.next
	}

	return dummyNode.next
}