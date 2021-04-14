package main

import "fmt"

type ListNode struct {
	Val int
    Next *ListNode
}

/* 删除排序链表中的重复元素 (Leetcode) */
func main()  {
	var l1 ListNode
	l1.Val = 1

	var l2 ListNode
	l2.Val = 2

	var l3 ListNode
	l3.Val = 3

	var l4 ListNode
	l4.Val = 3

	var l5 ListNode
	l5.Val = 4

	var l6 ListNode
	l6.Val = 4

	var l7 ListNode
	l7.Val = 5

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5
	l5.Next = &l6
	l6.Next = &l7

	result := deleteDuplicates(&l1)

	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Val)
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil{
		return nil
	}

	var tmpNode = &ListNode{} //创建哨兵节点（链表算法的常规操作）
	tmpNode.Next = head

	var cur = tmpNode

	for cur.Next != nil && cur.Next.Next != nil{
		if cur.Next.Val == cur.Next.Next.Val{ //如果当前指针所指向的链表节点的下一个节点的值和下下个节点的值一样的话
			tmpData := cur.Next.Val //保存存在相同情况的那个值信息
			for cur.Next != nil && cur.Next.Val == tmpData{ //循环遍历，只要链表中还存在那个值对应的节点的话，就进行删除操作
				cur.Next = cur.Next.Next
			}
		}else {
			cur = cur.Next //如果值不存在相同情况的话，就把cur指针往后移动
		}
	}

	return tmpNode.Next
}