package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	value int
	next *ListNode
}

func main()  {
	var l1 ListNode
	l1.value = 4

	var l2 ListNode
	l2.value = 1

	var l3 ListNode
	l3.value = 8

	var l4 ListNode
	l4.value = 2

	var l5 ListNode
	l5.value = 3

	var l6 ListNode
	l6.value = 9

	var l7 ListNode
	l7.value = 6

	l1.next = &l2
	l2.next = &l3
	l3.next = &l4
	l4.next = &l5
	l5.next = &l6
	l6.next = &l7

	result := SortList(&l1)

	fmt.Println(result.value)
	fmt.Println(result.next.value)
	fmt.Println(result.next.next.value)
	fmt.Println(result.next.next.next.value)
	fmt.Println(result.next.next.next.next.value)
	fmt.Println(result.next.next.next.next.next.value)
	fmt.Println(result.next.next.next.next.next.next.value)
}

func SortList(head *ListNode) *ListNode {
	if head == nil || head.next == nil{
		return head
	}

	var tmpArray = make([]int, 0)

	for head != nil{
		tmpArray = append(tmpArray, head.value)
		head = head.next
	}

	sort.Ints(tmpArray)

	var dummyNode = &ListNode{}
	var preNode = dummyNode

	for _, item := range tmpArray{
		var tmpNode = &ListNode{}
		tmpNode.value = item
		preNode.next = tmpNode
		preNode = tmpNode
	}

	return dummyNode.next
}