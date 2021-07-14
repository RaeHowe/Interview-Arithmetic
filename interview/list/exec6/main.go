package main

import "fmt"

type ListNode struct {
	value int
	next *ListNode
}

func main()  {
	var l1 ListNode
	l1.value = 1

	var l2 ListNode
	l2.value = 2

	var l3 ListNode
	l3.value = 4

	var l4 ListNode
	l4.value = 1

	var l5 ListNode
	l5.value = 3

	var l6 ListNode
	l6.value = 5

	l1.next = &l2
	l2.next = &l3

	l4.next = &l5
	l5.next = &l6

	result := mergeTwoLists(&l1, &l4)

	fmt.Println(result.value)
	fmt.Println(result.next.value)
	fmt.Println(result.next.next.value)
	fmt.Println(result.next.next.next.value)
	fmt.Println(result.next.next.next.next.value)
	fmt.Println(result.next.next.next.next.next.value)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}else if l2 == nil{
		return l1
	}

	var tmpMidNode = &ListNode{}
	var pre = tmpMidNode

	for l1 != nil && l2 != nil{
		if l1.value <= l2.value{
			pre.next = l1
			pre = l1
			l1 = l1.next
		}else {
			pre.next = l2
			pre = l2
			l2 = l2.next
		}
	}

	if l1 == nil{
		pre.next = l2
	}else if l2 == nil{
		pre.next = l1
	}

	return tmpMidNode.next
}