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
	l4.value = 3

	var l5 ListNode
	l5.value = 2

	var l6 ListNode
	l6.value = 1

	l1.next = &l2
	l2.next = &l3
	l3.next = &l4
	l4.next = &l5
	l5.next = &l6

	result := deleteRepeatNode(&l1)

	fmt.Println(result.value)
	fmt.Println(result.next.value)
	fmt.Println(result.next.next.value)
}

func deleteRepeatNode(head *ListNode) *ListNode {
	if head == nil{
		return head
	}

	var dummyNode = &ListNode{}
	var pre = dummyNode
	pre.next = head

	var tmpMap = make(map[int]int)
	for head != nil{
		if _, ok := tmpMap[head.value]; ok{
			//如果字典中有这个值的话，说明重复了
			head = head.next
			pre.next = pre.next.next
			continue
		}

		tmpMap[head.value] = 1
		pre = head //注意pre指针和head指针都要进行移动
		head = head.next
	}

	return dummyNode.next
}