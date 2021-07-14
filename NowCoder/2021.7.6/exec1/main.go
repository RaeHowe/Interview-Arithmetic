package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {
	var l1 ListNode
	l1.Val = 1

	var l2 ListNode
	l2.Val = 2

	var l3 ListNode
	l3.Val = 3

	var l4 ListNode
	l4.Val = 4

	var l5 ListNode
	l5.Val = 5

	var l6 ListNode
	l6.Val = 6

	var l7 ListNode
	l7.Val = 7

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5
	l5.Next = &l6
	l6.Next = &l7
	l7.Next = &l1

	result := EntryNodeOfLoop(&l1)

	fmt.Println(result.Val)
}

func EntryNodeOfLoop(pHead *ListNode) *ListNode{
	if pHead == nil || pHead.Next == nil{
		return nil
	}

	var tmpMap = make(map[*ListNode]struct{})

	var dummyNode = &ListNode{}
	dummyNode.Next = pHead
	var preNode = dummyNode

	for pHead != nil{
		if _, ok := tmpMap[pHead]; ok{
			//如果已经存在的话
			return preNode
		}

		tmpMap[pHead] = struct{}{}
		preNode = pHead
		pHead = pHead.Next
	}

	return nil
}