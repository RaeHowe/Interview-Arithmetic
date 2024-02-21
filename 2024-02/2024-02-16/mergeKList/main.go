package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Value int
	Next  *ListNode
}

func main() {
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

	var listArray []*ListNode
	listArray = append(listArray, &l10)
	listArray = append(listArray, &l13)

	node := mergeKLists(listArray)

	fmt.Println(node.Value)
	fmt.Println(node.Next.Value)
	fmt.Println(node.Next.Next.Value)
	fmt.Println(node.Next.Next.Next.Value)
	fmt.Println(node.Next.Next.Next.Next.Value)
	fmt.Println(node.Next.Next.Next.Next.Next.Value)
}

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}

	if length == 1 {
		return lists[0]
	}

	index := length / 2
	left := mergeKLists(lists[:index])
	right := mergeKLists(lists[index:])
	return merge(left, right)
}

func merge(pHead1, pHead2 *ListNode) *ListNode {
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
	合并 k 个升序的链表并将结果作为一个升序的链表返回其头节点。

数据范围：节点总数 0≤n≤50000≤n≤5000，每个节点的val满足 ∣val∣<=1000∣val∣<=1000
要求：时间复杂度 O(nlogn)O(nlogn)
*/
func mergeKLists2(lists []*ListNode) *ListNode {

	var tmpArray []int

	for _, listItem := range lists {
		for listItem != nil {
			tmpArray = append(tmpArray, listItem.Value)
			listItem = listItem.Next
		}
	}

	sort.Ints(tmpArray)

	var dummyNode = &ListNode{}
	var pre = dummyNode

	for _, nodeItem := range tmpArray {
		var tmpNode = &ListNode{Value: nodeItem}
		pre.Next = tmpNode
		pre = pre.Next
	}

	return dummyNode.Next
}
