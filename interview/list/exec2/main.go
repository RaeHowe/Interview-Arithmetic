package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	value int
	next *ListNode
}

func main()  {
	var l1 ListNode
	l1.value = 1

	var l2 ListNode
	l2.value = 0

	var l3 ListNode
	l3.value = 0

	var l4 ListNode
	l4.value = 1

	var l5 ListNode
	l5.value = 1

	l1.next = &l2
	l2.next = &l3
	l3.next = &l4
	l4.next = &l5

	result := getResult(&l1) //

	fmt.Println(result)
}

func getResult(head *ListNode) int {
	var length = 0
	
	var tmpArray []int

	for head.next != nil{
		tmpArray = append(tmpArray, head.value)
		head = head.next
		length++
	}

	var result = 0
	for i := 0; i < len(tmpArray); i++{
		result += tmpArray[i] * int(math.Pow(float64(10), float64(length)))
		length--
	}

	return result
}