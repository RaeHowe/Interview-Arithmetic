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

	node := reverseListBetween(&l1, 2, 4)

	fmt.Println(node.Value)
	fmt.Println(node.Next.Value)
	fmt.Println(node.Next.Next.Value)
	fmt.Println(node.Next.Next.Next.Value)
	fmt.Println(node.Next.Next.Next.Next.Value)

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

	node2 := mergeList(&l10, &l13)

	fmt.Println("###############")
	fmt.Println(node2.Value)
	fmt.Println(node2.Next.Value)
	fmt.Println(node2.Next.Next.Value)
	fmt.Println(node2.Next.Next.Next.Value)
	fmt.Println(node2.Next.Next.Next.Next.Value)
	fmt.Println(node2.Next.Next.Next.Next.Next.Value)
}

func getLastKNode(head *ListNode, m int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var fast = head
	for i := 0; i < m; i++ {
		if fast == nil {
			return fast
		}

		fast = fast.Next
	}

	var last = head

	for fast != nil {
		fast = fast.Next
		last = last.Next
	}

	return last
}

func entryNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	var tmpMap = make(map[*ListNode]int)

	for head != nil {
		if _, ok := tmpMap[head]; ok {
			return head
		}

		tmpMap[head] = 1
		head = head.Next
	}

	return nil
}

func judgementList(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	var slow = head
	var fast = head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

func mergeKList(listArray []*ListNode) *ListNode {
	var listArrayLen = len(listArray)
	if listArrayLen == 1 {
		return listArray[0]
	}

	var num = listArrayLen / 2
	var left = mergeKList(listArray[:num])
	var right = mergeKList(listArray[num:])
	return mergeList(left, right)
}

func mergeList(head1, head2 *ListNode) *ListNode {
	if head1 == nil && head2 != nil {
		return head2
	}

	if head1 != nil && head2 == nil {
		return head1
	}

	var dummyNode = &ListNode{}
	var pre = dummyNode

	for head1 != nil && head2 != nil {
		if head1.Value < head2.Value {
			pre.Next = head1
			head1 = head1.Next
		} else {
			pre.Next = head2
			head2 = head2.Next
		}

		pre = pre.Next
	}

	if head1 == nil {
		pre.Next = head2
	}

	if head2 == nil {
		pre.Next = head1
	}

	return dummyNode.Next
}

func reverseListBetween(head *ListNode, m, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var dummyNode = &ListNode{Next: head}
	var pre = dummyNode

	for i := 1; i < m; i++ {
		pre = pre.Next
	}

	cur := pre.Next

	for i := m; i < n; i++ {
		temp := cur.Next
		cur.Next = temp.Next
		temp.Next = pre.Next
		pre.Next = temp
	}

	return dummyNode.Next
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var dummyNode = &ListNode{Next: head}
	var pre = dummyNode

	for head != nil {
		next := head.Next
		head.Next = pre.Next
		pre.Next = head
		head = next
	}

	return dummyNode.Next
}
