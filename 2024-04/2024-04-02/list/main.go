package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func main() {
	var l1 ListNode
	l1.Value = 5

	var l2 ListNode
	l2.Value = 1

	var l3 ListNode
	l3.Value = 9

	var l4 ListNode
	l4.Value = 12

	var l5 ListNode
	l5.Value = 3

	var l6 ListNode
	l6.Value = 2

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5
	l5.Next = &l6

	////node := reverseList(&l1)
	//node := reverseListBetween(&l1, 2, 4)

	node := sortInList(&l1)

	fmt.Println(node.Value)
	fmt.Println(node.Next.Value)
	fmt.Println(node.Next.Next.Value)
	fmt.Println(node.Next.Next.Next.Value)
	fmt.Println(node.Next.Next.Next.Next.Value)
	fmt.Println(node.Next.Next.Next.Next.Next.Value)
}

func isPail(head *ListNode) bool {
	var slow, fast = head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var middleNode = slow
	var dummyNode *ListNode
	for middleNode != nil {
		next := middleNode.Next
		middleNode.Next = dummyNode
		dummyNode = middleNode
		middleNode = next
	}

	var mid = slow
	for mid != nil {
		if head.Value != mid.Value {
			return false
		}

		head = head.Next
		mid = mid.Next
	}

	return true
}

// 2.归并排序(logN)
func sortInList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	middleNode := middle(head)
	listAfter := sortInList(middleNode.Next)
	middleNode.Next = nil
	listFront := sortInList(head)
	return merge(listFront, listAfter)
}

func merge(head1, head2 *ListNode) *ListNode {
	if head1 != nil && head2 == nil {
		return head1
	} else if head1 == nil && head2 != nil {
		return head2
	}

	var dummyNode = &ListNode{}
	var pre = dummyNode
	for head1 != nil && head2 != nil {
		if head1.Value <= head2.Value {
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
	} else if head2 == nil {
		pre.Next = head1
	}

	return dummyNode.Next
}

func middle(head *ListNode) *ListNode {
	var slow, fast = head, head

	if fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

/*
3 -> 8 -> 1

	5 -> 9

4 4 0
*/
func listSum(head1, head2 *ListNode) *ListNode {
	var flagSum = 0

	var dummyNode = &ListNode{}
	var pre = dummyNode

	newHead1 := reverseList(head1)
	newHead2 := reverseList(head2)

	for newHead1 != nil || newHead2 != nil {
		var val1, val2 int
		if newHead1 != nil {
			val1 = newHead1.Value
			newHead1 = newHead1.Next
		}

		if newHead2 != nil {
			val2 = newHead2.Value
			newHead2 = newHead2.Next
		}

		var tmpSum = val1 + val2 + flagSum
		if tmpSum >= 10 {
			flagSum = 1
			tmpSum = tmpSum % 10
		} else {
			flagSum = 0
		}

		var tmpNode = &ListNode{Value: tmpSum}

		pre.Next = tmpNode
		pre = pre.Next
	}

	if flagSum == 1 {
		var tmpNode = &ListNode{Value: flagSum}

		pre.Next = tmpNode
	}

	return reverseList(dummyNode.Next)
}

func firstPublicNode(head1, head2 *ListNode) *ListNode {
	if head1 == nil || head2 == nil {
		return nil
	}

	var indexA = head1
	var indexB = head2

	for indexA != indexB {
		if indexA == nil {
			indexA = head2
		} else {
			indexA = indexA.Next
		}

		if indexB == nil {
			indexB = head1
		} else {
			indexB = indexB.Next
		}
	}

	return indexA
}

func deleteLastKNode(head *ListNode, last int) *ListNode {
	var index = head

	for i := 0; i < last; i++ {
		index = index.Next
	}

	var dummyNode = &ListNode{Next: head}
	var pre = dummyNode
	for index != nil {
		pre = pre.Next
		index = index.Next
	}

	pre.Next = pre.Next.Next
	return dummyNode.Next
}

func getListLastNode(head *ListNode, last int) *ListNode {
	var dummyNode = &ListNode{Next: head}

	var index = head
	for i := 0; i < last; i++ {
		if index == nil {
			return head
		}
		index = index.Next
	}

	for index != nil {
		dummyNode = dummyNode.Next
		index = index.Next
	}

	return dummyNode.Next
}

func isLoop(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	var slow, fast = head, head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		} else {
			slow = slow.Next
			fast = fast.Next.Next
		}
	}

	return false
}

func mergeKList(lists []*ListNode) *ListNode {
	var index = len(lists) / 2
	var left = mergeKList(lists[:index])
	var right = mergeKList(lists[index:])
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

	for head1 != nil || head2 != nil {
		if head1.Value <= head2.Value {
			pre.Next = head1
			head1 = head1.Next
		} else {
			pre.Next = head2
			head2 = head2.Next
		}

		pre = pre.Next
	}

	if head1 == nil && head2 != nil {
		pre.Next = head2
	}

	if head1 != nil && head2 == nil {
		pre.Next = head1
	}

	return dummyNode.Next
}

func reverseListBetween(head *ListNode, p, q int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var dummyNode = &ListNode{Next: head}
	var pre = dummyNode

	for i := 1; i < p; i++ {
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
