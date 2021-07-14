package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {
	var l1 ListNode
	l1.Val = 4

	var l2 ListNode
	l2.Val = 1

	var l3 ListNode
	l3.Val = 8

	var l4 ListNode
	l4.Val = 2

	var l5 ListNode
	l5.Val = 3

	var l6 ListNode
	l6.Val = 9

	var l7 ListNode
	l7.Val = 6

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5
	l5.Next = &l6
	l6.Next = &l7
	l7.Next = &l1

	result := hasCycle(&l1)

	fmt.Println(result)
}

func hasCycle( head *ListNode ) bool {
	if head == nil || head.Next == nil{
		return false
	}

	var slow = head
	var fast = head.Next

	for fast != nil{
		if fast == slow {
			return true
		}

		fast = fast.Next.Next
		slow = slow.Next
	}

	return false
}