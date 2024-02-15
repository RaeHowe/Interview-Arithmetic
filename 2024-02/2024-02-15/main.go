package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func main() {
	var l1 ListNode
	l1.Value = 1

	var l2 ListNode
	l2.Value = 2

	var l3 ListNode
	l3.Value = 3

	var l4 ListNode
	l4.Value = 4

	var l5 ListNode
	l5.Value = 5

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5

	//node := reverseBetween(&l1, 2, 4)
	//fmt.Println(node.Value)
	//fmt.Println(node.Next.Value)
	//fmt.Println(node.Next.Next.Value)
	//fmt.Println(node.Next.Next.Next.Value)
	//fmt.Println(node.Next.Next.Next.Next.Value)

	var l10 ListNode
	l10.Value = 1

	var l11 ListNode
	l11.Value = 5

	var l12 ListNode
	l12.Value = 9

	l10.Next = &l11
	l11.Next = &l12

	var l13 ListNode
	l13.Value = 2

	var l14 ListNode
	l14.Value = 4

	var l15 ListNode
	l15.Value = 8

	l13.Next = &l14
	l14.Next = &l15

	node := Merge(&l10, &l13)
	fmt.Println(node.Value)
	fmt.Println(node.Next.Value)
	fmt.Println(node.Next.Next.Value)
	fmt.Println(node.Next.Next.Next.Value)
	fmt.Println(node.Next.Next.Next.Next.Value)
	fmt.Println(node.Next.Next.Next.Next.Next.Value)
}

/*
	输入两个递增的链表，单个链表的长度为n，合并这两个链表并使新链表中的节点仍然是递增排序的。

数据范围： 0≤n≤10000≤n≤1000，−1000≤节点值≤1000−1000≤节点值≤1000
要求：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)

如输入{1,3,5},{2,4,6}时，合并后的链表为{1,2,3,4,5,6}，所以对应的输出为{1,2,3,4,5,6}，转换过程如下图所示
*/
func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	if pHead1 == nil && pHead2 != nil {
		return pHead2
	}

	if pHead1 != nil && pHead2 == nil {
		return pHead1
	}

	if pHead1 == nil && pHead2 == nil {
		return nil
	}

	var dummyNode = &ListNode{}
	var pre = dummyNode

	for pHead1 != nil && pHead2 != nil {
		if pHead1.Value <= pHead2.Value {
			pre.Next = pHead1
			pHead1 = pHead1.Next
		} else {
			pre.Next = pHead2
			pHead2 = pHead2.Next
		}

		pre = pre.Next
	}

	if pHead1 == nil {
		pre.Next = pHead2
	}

	if pHead2 == nil {
		pre.Next = pHead1
	}

	return dummyNode.Next
}

/*
	将一个节点数为 size 链表 m 位置到 n 位置之间的区间反转，要求时间复杂度 O(n)O(n)，空间复杂度 O(1)O(1)。

例如：
给出的链表为 1→2→3→4→5→NULL1→2→3→4→5→NULL, m=2,n=4m=2,n=4,
返回 1→4→3→2→5→NULL1→4→3→2→5→NULL.

数据范围： 链表长度 0<size≤10000<size≤1000，0<m≤n≤size0<m≤n≤size，链表中每个节点的值满足 ∣val∣≤1000∣val∣≤1000
要求：时间复杂度 O(n)O(n) ，空间复杂度 O(n)O(n)
进阶：时间复杂度 O(n)O(n)，空间复杂度 O(1)O(1)
*/
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var dummyNode = &ListNode{Next: head}
	var pre = dummyNode

	for i := 1; i < m; i++ {
		pre = pre.Next
	}

	var cur = pre.Next

	for i := m; i < n; i++ {
		temp := cur.Next
		cur.Next = temp.Next
		temp.Next = pre.Next
		pre.Next = temp
	}

	return dummyNode.Next
}

/*
	将给出的链表中的节点每 k 个一组翻转，返回翻转后的链表

如果链表中的节点数不是 k 的倍数，将最后剩下的节点保持原样
你不能更改节点中的值，只能更改节点本身。

数据范围：  0≤n≤2000 0≤n≤2000 ， 1≤k≤20001≤k≤2000 ，链表中每个元素都满足 0≤val≤10000≤val≤1000
要求空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)
例如：
给定的链表是 1→2→3→4→5
对于 k=2 , 你应该返回 2→1→4→3→5
对于 k=3 , 你应该返回 3→2→1→4→5
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	return nil
}
