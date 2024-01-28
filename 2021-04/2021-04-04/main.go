package main

import "fmt"

/* 合并两个有序链表 (Leetcode) */

type listNode struct {
	value int
	next *listNode
}

func main()  {
	var l1 listNode
	l1.value = 1

	var l2 listNode
	l2.value = 5

	var l3 listNode
	l3.value = 6

	var l4 listNode
	l4.value = 10

	l1.next = &l2
	l2.next = &l3
	l3.next = &l4

	var l5 listNode
	l5.value = 2

	var l6 listNode
	l6.value = 7

	var l7 listNode
	l7.value = 9

	l5.next = &l6
	l6.next= &l7

	result := mergeSortList(&l1, &l5)

	fmt.Println(result.value)
	fmt.Println(result.next.value)
	fmt.Println(result.next.next.value)
	fmt.Println(result.next.next.next.value)
	fmt.Println(result.next.next.next.next.value)
	fmt.Println(result.next.next.next.next.next.value)
}

func mergeSortList(l1, l2 *listNode) *listNode {
	if l1 == nil && l2 == nil{
		return nil
	}

	if l1 == nil && l2 != nil{
		return l2
	}

	if l1 != nil && l2 == nil{
		return l1
	}

	var tmpNode = &listNode{} //一般涉及到链表的算法题目，都需要这个哨兵节点，来简化思考
	var pre = tmpNode

	for l1 != nil && l2 != nil{
		if l1.value <= l2.value{ //判断两个链表的当前值，哪个更小
			pre.next = l1 //如果l1更小的话，那么pre的下一个节点就是l1
			pre = l1 //把pre当前的值设置为l1
			l1 = l1.next //当前值判断完毕之后，取l1的下一个节点的值继续进行
		}else {
			pre.next = l2
			pre = l2
			l2 = l2.next
		}
	}

	if l1 == nil{
		pre.next = l2
	}else if l2 == nil{
		pre.next = l1
	}

	return tmpNode.next
}
