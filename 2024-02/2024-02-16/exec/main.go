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
	l2.Value = 2

	var l3 ListNode
	l3.Value = 3

	var l4 ListNode
	l4.Value = 4

	var l5 ListNode
	l5.Value = 5

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5

	//node := reverseList(&l1)
	//
	//fmt.Println(node.Value)
	//fmt.Println(node.Next.Value)
	//fmt.Println(node.Next.Next.Value)
	//fmt.Println(node.Next.Next.Next.Value)
	//fmt.Println(node.Next.Next.Next.Next.Value)

	node := reverseBetween(&l1, 2, 4)
	fmt.Println(node.Value)
	fmt.Println(node.Next.Value)
	fmt.Println(node.Next.Next.Value)
	fmt.Println(node.Next.Next.Next.Value)
	fmt.Println(node.Next.Next.Next.Next.Value)

	//var l10 ListNode
	//l10.Value = 1
	//
	//var l11 ListNode
	//l11.Value = 5
	//
	//var l12 ListNode
	//l12.Value = 9
	//
	//l10.Next = &l11
	//l11.Next = &l12
	//
	//var l13 ListNode
	//l13.Value = 2
	//
	//var l14 ListNode
	//l14.Value = 4
	//
	////var l15 ListNode
	////l15.Value = 8
	//
	//l13.Next = &l14
	////l14.Next = &l15
	//
	//node2 := Merge(&l10, &l13)
	//fmt.Println(node2.Value)
	//fmt.Println(node2.Next.Value)
	//fmt.Println(node2.Next.Next.Value)
	//fmt.Println(node2.Next.Next.Next.Value)
	//fmt.Println(node2.Next.Next.Next.Next.Value)
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var dummyNode = &ListNode{Next: head}
	var pre = dummyNode

	// dummy -> 0 -> 1 -> 2 ->
	for i := 1; i < m; i++ {
		pre = pre.Next
	}

	var cur = pre.Next
	for i := m; i < n; i++ {
		temp := cur.Next
		cur.Next = temp.Next
		temp.Next = pre.Next
		pre.Next = temp
	}

	return dummyNode.Next
}

func Merge(pHead1, pHead2 *ListNode) *ListNode {
	if pHead1 != nil && pHead2 == nil {
		return pHead1
	}

	if pHead1 == nil && pHead2 != nil {
		return pHead2
	}

	if pHead1 == nil && pHead2 == nil {
		return nil
	}

	var dummyNode = &ListNode{}
	var pre = dummyNode

	for pHead1 != nil && pHead2 != nil {
		if pHead1.Value <= pHead2.Value {
			pre.Next = pHead1
			pHead1 = pHead1.Next
		} else {
			pre.Next = pHead2
			pHead2 = pHead2.Next
		}

		pre = pre.Next
	}

	if pHead1 == nil {
		pre.Next = pHead2
	}

	if pHead2 == nil {
		pre.Next = pHead1
	}

	return dummyNode.Next
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var dummyNode *ListNode

	for head != nil {
		next := head.Next
		head.Next = dummyNode
		dummyNode = head
		head = next
	}

	return dummyNode
}
