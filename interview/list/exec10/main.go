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

	result := judgementList(&l1)

	fmt.Println(result)
}

func judgementList(head *ListNode) bool {
	if head == nil || head.next == nil{
		return false
	}

	return true
}