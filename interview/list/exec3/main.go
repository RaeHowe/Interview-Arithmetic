package main

import "fmt"

type ListNode struct {
	value int
	next *ListNode
}

func main()  {
	var l1 ListNode
	l1.value = 4

	var l2 ListNode
	l2.value = 5

	var l3 ListNode
	l3.value = 1

	var l4 ListNode
	l4.value = 9

	l1.next = &l2
	l2.next = &l3
	l3.next = &l4

	deleteNode(&l3)

	fmt.Println(l1.value)
	fmt.Println(l1.next.value)
	fmt.Println(l1.next.next.value)
}

func deleteNode(targetNode *ListNode) {
	targetNode.value = targetNode.next.value
	targetNode.next = targetNode.next.next
}
