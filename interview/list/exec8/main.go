package main

import "fmt"

type ListNode struct {
	value int
	next *ListNode
}

func main() {
	var l1 ListNode
	l1.value = 1

	var l2 ListNode
	l2.value = 2

	var l3 ListNode
	l3.value = 3

	var l4 ListNode
	l4.value = 4

	var l5 ListNode
	l5.value = 5

	var l6 ListNode
	l6.value = 6

	l1.next = &l2
	l2.next = &l3
	l3.next = &l4
	l4.next = &l5
	l5.next = &l6

	result := reciprocalNode(&l1, 3)

	fmt.Println(result)
}

func reciprocalNode(head *ListNode, location int) int {
	if head == nil {
		return -1
	}

	var tmpArray []int

	for head != nil{
		tmpArray = append(tmpArray, head.value)
		head = head.next
	}


	reverseArray(tmpArray)

	return tmpArray[location]
}

func reverseArray(tmpArray []int)  {
	for i := 0; i < len(tmpArray) / 2; i++{
		tmpArray[i], tmpArray[len(tmpArray) - i - 1] = tmpArray[len(tmpArray) - i - 1], tmpArray[i]
	}
}