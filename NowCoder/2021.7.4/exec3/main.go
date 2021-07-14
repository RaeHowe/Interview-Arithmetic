package main

import "fmt"

type ListNode struct{
   Val int
   Next *ListNode
}

func main()  {
	var l1 ListNode
	l1.Val = 1

	var l2 ListNode
	l2.Val = 4

	var l3 ListNode
	l3.Val = 8

	var l4 ListNode
	l4.Val = 2

	var l5 ListNode
	l5.Val = 3

	var l6 ListNode
	l6.Val = 9

	l1.Next = &l2
	l2.Next = &l3

	l4.Next = &l5
	l5.Next = &l6

	result := mergeList(&l1, &l4)

	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Next.Next.Val)

}

func mergeList(head1, head2 *ListNode) *ListNode {
	if head1 == nil{
		return head2
	}

	if head2 == nil{
		return head1
	}

	if head1 == nil && head2 == nil{
		return nil
	}

	var dummyNode = &ListNode{}
	var preNode = dummyNode

	for head1 != nil && head2 != nil{
		if head1.Val <= head2.Val{
			preNode.Next = head1
			preNode = head1
			head1 = head1.Next
		}else {
			preNode.Next = head2
			preNode = head2
			head2 = head2.Next
		}
	}

	if head1 == nil{
		preNode.Next = head2
	}else if head2 == nil{
		preNode.Next = head1
	}

	return dummyNode.Next
}