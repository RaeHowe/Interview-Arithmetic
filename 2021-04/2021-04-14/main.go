package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

//环形链表
func main()  {
	var l1 ListNode
	l1.Val = 1

	var l2 ListNode
	l2.Val = 2

	var l3 ListNode
	l3.Val = 3

	var l4 ListNode
	l4.Val = 4

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l1

	flag := hasCycle2(&l1)

	fmt.Println(flag)
}

//判断一个链表是否是环形链表的方法，目前我知道的是两种，一种就是采用hashmap，另一种就是快慢指针
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil{
		return false
	}

	var tmpMap = make(map[*ListNode]int)

	for head.Next != nil{
		if _, ok := tmpMap[head]; ok{
			//如果存在的话
			return true
		}else {
			tmpMap[head] = 1
		}

		head = head.Next
	}

	return false
}

//快慢指针法
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil{
		return false
	}

	var slow = head
	var fast = head.Next

	for fast != nil && fast.Next != nil && slow.Next != nil{
		if slow != fast{
			slow = slow.Next
			fast = fast.Next.Next
		}else {
			return true
		}
	}

	return false
}