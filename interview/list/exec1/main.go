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
	l2.value = 1

	var l3 ListNode
	l3.value = 8

	var l4 ListNode
	l4.value = 4

	var l5 ListNode
	l5.value = 5

	var l6 ListNode
	l6.value = 5

	var l7 ListNode
	l7.value = 0

	var l8 ListNode
	l8.value = 1

	l1.next = &l2
	l2.next = &l3
	l3.next = &l4
	l4.next = &l5

	l6.next = &l7
	l7.next = &l8
	l8.next = &l3

	result := firstListNode(&l1, &l6)

	fmt.Println(result.value)
}

func firstListNode(list1, list2 *ListNode) *ListNode {
	var tmpMap = make(map[*ListNode]int)

	for list1.next != nil{
		tmpMap[list1] = 1
		list1 = list1.next
	}

	for list2.next != nil{
		if _, ok := tmpMap[list2]; ok{
			return list2
		}

		list2 = list2.next
	}

	return nil
}
