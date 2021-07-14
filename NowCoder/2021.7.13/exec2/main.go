package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	var l1 ListNode
	l1.Val = 9

	var l2 ListNode
	l2.Val = 3

	var l3 ListNode
	l3.Val = 7

	l1.Next = &l2
	l2.Next = &l3

	var l4 ListNode
	l4.Val = 4

	var l5 ListNode
	l5.Val = 5

	l4.Next = &l5

	result := addInList(&l1, &l4)

	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Val)
}

func addInList( head1 *ListNode ,  head2 *ListNode ) *ListNode {
	// write code here
	if head1 == nil{
		return head2
	}else if head2 == nil{
		return head1
	}

	var tmpStr1 string

	for head1 != nil{
		tmpStr1 = tmpStr1 + strconv.Itoa(head1.Val)

		head1 = head1.Next
	}

	var tmpStr2 string

	for head2 != nil{
		tmpStr2 = tmpStr2 + strconv.Itoa(head2.Val)

		head2 = head2.Next
	}

	var flag = 0
	var indexOfStr1 = len(tmpStr1) - 1
	var indexOfStr2 = len(tmpStr2) - 1

	var tmpStrResult string

	for indexOfStr1 >= 0 || indexOfStr2 >= 0{
		var tmpChar1 = 0
		var tmpChar2 = 0

		if indexOfStr1 >= 0{
			tmpChar1 = int(tmpStr1[indexOfStr1] - '0')
			indexOfStr1--
		}

		if indexOfStr2 >= 0{
			tmpChar2 = int(tmpStr2[indexOfStr2] - '0')
			indexOfStr2--
		}

		var tmpResult = tmpChar1 + tmpChar2 + flag

		if tmpResult >= 10{
			flag = 1
			tmpResult = tmpResult - 10
		}else {
			flag = 0
		}

		tmpStrResult = strconv.Itoa(tmpResult) + tmpStrResult
	}

	if flag == 1 {
		tmpStrResult = "1" + tmpStrResult
	}

	var dummyNode = &ListNode{}
	var preNode = dummyNode

	for i := 0; i < len(tmpStrResult); i++{
		var tmpNode = &ListNode{}
		tmpNode.Val = int(tmpStrResult[i] - '0')

		preNode.Next = tmpNode
		preNode = tmpNode
	}

	return dummyNode.Next
}