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
	l2.Val = 2

	var l3 ListNode
	l3.Val = 3

	var l4 ListNode
	l4.Val = 3

	var l5 ListNode
	l5.Val = 3

	var l6 ListNode
	l6.Val = 1

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5
	l5.Next = &l6

	result := isPail(&l1)
	fmt.Println(result)
}

func isPail( head *ListNode ) bool {
	// write code here
	if head == nil || head.Next == nil{
		return true
	}

	var tmpArray = make([]int, 0)

	for head != nil{
		tmpArray = append(tmpArray, head.Val)

		head = head.Next
	}

	for i := 0; i < len(tmpArray)/2; i++{
		if tmpArray[i] != tmpArray[len(tmpArray) - i - 1]{
			return false
		}
	}

	return true
}