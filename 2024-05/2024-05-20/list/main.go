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

	var l6 ListNode
	l6.Value = 6

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5
	l5.Next = &l6

	node := reverseList(&l1)

	fmt.Println(node.Value)
	fmt.Println(node.Next.Value)
	fmt.Println(node.Next.Next.Value)
	fmt.Println(node.Next.Next.Next.Value)
	fmt.Println(node.Next.Next.Next.Next.Value)
	fmt.Println(node.Next.Next.Next.Next.Next.Value)
}

func firstPublicNode(a, b *ListNode) *ListNode {
	var cur1, cur2 = a, b

	for cur1 != cur2 {
		if cur1 == nil {
			cur1 = cur2
		} else {
			cur1 = cur1.Next
		}

		if cur2 == nil {
			cur2 = cur1
		} else {
			cur2 = cur2.Next
		}
	}

	return cur1
}

func deleteLastKNode(head *ListNode, k int) *ListNode {
	var last = head
	for i := 0; i < k; i++ {
		last = last.Next
	}

	var dummyNode = &ListNode{Next: head}
	var pre = dummyNode

	for last != nil {
		pre = pre.Next
		last = last.Next
	}

	pre.Next = pre.Next.Next

	return dummyNode.Next
}

// 1 <- 2 <- 3 <- 4 <- 5 <- 6
func getLastKNode(head *ListNode, k int) *ListNode {
	var last = head
	for i := 0; i < k; i++ {
		last = last.Next
	}

	var dummyNode = &ListNode{Next: head}
	for last != nil {
		dummyNode = dummyNode.Next
		last = last.Next
	}

	return dummyNode.Next
}

func judgementLoop(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	var fast, slow = head, head

	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}

		fast = fast.Next.Next
		slow = slow.Next
	}

	return false
}

func mergeKList(lists []*ListNode) *ListNode {
	var length = len(lists)

	if length == 1 {
		return lists[0]
	}

	var index = length / 2
	var left = mergeKList(lists[:index])
	var right = mergeKList(lists[index:])
	return mergeList(left, right)
}

func mergeList(a, b *ListNode) *ListNode {
	if a == nil && b != nil {
		return b
	}

	if a != nil && b == nil {
		return a
	}

	var dummyNode = &ListNode{}
	var pre = dummyNode

	for a != nil && b != nil {
		if a.Value <= b.Value {
			pre.Next = a
			a = a.Next
		} else {
			pre.Next = b
			b = b.Next
		}

		pre = pre.Next
	}

	if a == nil && b != nil {
		pre.Next = b
	} else if a != nil && b == nil {
		pre.Next = a
	}

	return dummyNode.Next
}

func reverseBetweenList(head *ListNode, p, q int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var dummyNode = &ListNode{Next: head}
	var pre = dummyNode

	for i := p; i < q; i++ {
		pre = pre.Next
	}

	var cur = pre.Next
	for i := p; i < q; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}

	return dummyNode.Next
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var dummyNode *ListNode
	for head != nil {
		next := head.Next
		head.Next = dummyNode
		dummyNode = head
		head = next
	}

	return dummyNode
}
