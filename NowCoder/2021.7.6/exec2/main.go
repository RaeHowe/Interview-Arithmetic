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

	result := removeNthFromEnd(&l1, 7)

	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Next.Next.Val)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil{
		return head
	}

	var nodeOfCount = head
	var count int

	//统计链表一共有多少个节点
	for nodeOfCount != nil{
		count++
		nodeOfCount = nodeOfCount.Next
	}

	var dummyNode = &ListNode{}
	dummyNode.Next = head

	var preNode = dummyNode

	var nCount = count - n //遍历过程中，需要遍历的次数

	for i := 0; i < nCount; i++{
		preNode = head
		head = head.Next
	}

	//遍历完之后，接下来就删除
	preNode.Next = preNode.Next.Next

	return dummyNode.Next
}