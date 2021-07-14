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

	result := deleteNode(&l1, &l3)

	fmt.Println(result.value)
	fmt.Println(result.next.value)
	fmt.Println(result.next.next.value)
}

func deleteNode(head, targetNode *ListNode) *ListNode {
	if head == nil{
		return nil
	}

	var dummyNode = &ListNode{}
	dummyNode.next = head

	var preNode = dummyNode

	for head != nil{
		if head.value == targetNode.value{
			preNode.next = head.next
			break
		}

		preNode = head
		head = head.next
	}

	return dummyNode.next
}
