package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	var l1 ListNode
	l1.Val = 1

	var l2 ListNode
	l2.Val = 4

	var l3 ListNode
	l3.Val = 7

	var l4 ListNode
	l4.Val = 8

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4

	var l5 ListNode
	l5.Val = 2

	var l6 ListNode
	l6.Val = 9

	l5.Next = &l6

	var l7 ListNode
	l7.Val = 7

	var l8 ListNode
	l8.Val = 11

	l7.Next = &l8

	var array = []*ListNode{&l1, &l5, &l7}

	result := mergeKLists(array)

	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Next.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Next.Next.Next.Next.Val)
}

func mergeKLists( lists []*ListNode ) *ListNode {
	if lists == nil{
		return nil
	}
	// write code here
	var dummyNode *ListNode

	for _, item := range lists{
		dummyNode = mergeList(dummyNode, item)
	}

	return dummyNode
}

func mergeList(head1 *ListNode, head2 *ListNode) *ListNode {
	var dummyNode = new(ListNode)
	var preNode = dummyNode

	for head1 != nil && head2 != nil{
		if head1.Val <= head2.Val{
			preNode.Next = head1
			head1 = head1.Next
		}else {
			preNode.Next = head2
			head2 = head2.Next
		}

		preNode = preNode.Next
	}

	if head1 == nil{
		preNode.Next = head2
	}else if head2 == nil{
		preNode.Next = head1
	}

	return dummyNode.Next
}