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

	l5.Next = &l6
	l6.Next = &l7
	l7.Next = &l3

	result := FindFirstCommonNode(&l1, &l5)

	fmt.Println(result.Val)
	//fmt.Println(result.next.value)
	//fmt.Println(result.next.next.value)
	//fmt.Println(result.next.next.next.value)
	//fmt.Println(result.next.next.next.next.value)
	//fmt.Println(result.next.next.next.next.next.value)
	//fmt.Println(result.next.next.next.next.next.next.value)
}

func FindFirstCommonNode( pHead1 *ListNode ,  pHead2 *ListNode ) *ListNode {
	// write code here
	if pHead1 == nil || pHead2 == nil{
		return nil
	}

	var tmpMap = make(map[*ListNode]struct{})

	for pHead1 != nil{
		tmpMap[pHead1] = struct{}{}

		pHead1 = pHead1.Next
	}

	for pHead2 != nil{
		if _, ok := tmpMap[pHead2]; ok{
			return pHead2
		}

		pHead2 = pHead2.Next
	}

	return nil
}
