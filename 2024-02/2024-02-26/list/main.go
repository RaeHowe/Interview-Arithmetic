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
	l2.Value = 1

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

	node := removeDuplicateList(&l1)
	fmt.Println(node.Value)
	fmt.Println(node.Next.Value)
	fmt.Println(node.Next.Next.Value)
	fmt.Println(node.Next.Next.Next.Value)
	fmt.Println(node.Next.Next.Next.Next.Value)
}

func removeAllDuplicateList(head *ListNode) *ListNode {
	var dummyNode = &ListNode{Next: head}
	var pre = dummyNode

	for head != nil && head.Next != nil {
		if pre.Next.Value == head.Next.Value {
			var sameValue = pre.Next.Value
			for pre.Next != nil && pre.Next.Value == sameValue {
				pre.Next = pre.Next.Next
				head = head.Next
			}
		} else {
			pre = pre.Next
			head = head.Next
		}
	}

	return dummyNode.Next
}

// 删除链表中的重复元素
func removeDuplicateList(head *ListNode) *ListNode {
	var dummyNode = &ListNode{Next: head}
	var pre = dummyNode

	for head != nil && head.Next != nil {
		if pre.Next.Value == head.Next.Value {
			pre.Next = head.Next
		} else {
			pre = pre.Next
		}

		head = head.Next
	}

	return dummyNode.Next
}

// 4 -> 2 -> 3
//
//	5 -> 7
//
// 两个链表相加
func listSum(head1, head2 *ListNode) *ListNode {
	var dummyNode = &ListNode{}
	var pre = dummyNode

	var flagAdd = 0
	if head1 != nil || head2 != nil {
		var val1, val2 int
		if head1 != nil {
			val1 = head1.Value
			head1 = head1.Next
		}

		if head2 != nil {
			val2 = head2.Value
			head2 = head2.Next
		}

		var tmpSum = val1 + val2 + flagAdd
		if tmpSum >= 10 {
			flagAdd = 1
			tmpSum = tmpSum % 10
		} else {
			flagAdd = 0
		}

		var tmpNode = &ListNode{Value: tmpSum}
		pre.Next = tmpNode

		pre = pre.Next
	}

	if flagAdd > 0 {
		var tmpNode = &ListNode{Value: flagAdd}
		pre.Next = tmpNode
	}

	return reverseList(dummyNode.Next)
}

// 两个链表的第一个公共节点
func publicNode(head1, head2 *ListNode) *ListNode {
	var cur1, cur2 = head1, head2
	for cur1 != cur2 {
		if cur1 == nil {
			cur1 = head2
		} else {
			cur1 = cur1.Next
		}

		if cur2 == nil {
			cur2 = head1
		} else {
			cur2 = cur2.Next
		}
	}

	return cur1
}

// 删除链表倒数第K个节点
func removeKNode(head *ListNode, k int) *ListNode {
	var dummyNode = &ListNode{Next: head}
	var pre = dummyNode

	var fast = head
	for i := 0; i < k; i++ {
		if fast == nil {
			return fast

		}
		fast = fast.Next
	}

	for fast != nil {
		pre = pre.Next
		fast = fast.Next
	}

	pre.Next = pre.Next.Next

	return dummyNode.Next
}

// pre -> 1 -> 2 -> 3 -> 4 -> 5 -> 6
// 获取链表倒数K个节点
func getKNode(head *ListNode, k int) *ListNode {
	var point = head

	for i := 0; i < k; i++ {
		if point == nil {
			return point
		}
		point = point.Next
	}

	var last = head

	for point != nil {
		last = last.Next
		point = point.Next
	}

	return last
}

func judgementLoop(head *ListNode) bool {
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

func mergeKList(lists []*ListNode) *ListNode {
	var arrayLen = len(lists)

	var num = arrayLen / 2
	var left = mergeKList(lists[:num])
	var right = mergeKList(lists[num:])
	return mergeList(left, right)
}

func mergeList(head1, head2 *ListNode) *ListNode {
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

func reverseListBetween(head *ListNode, start, end int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var dummyNode = &ListNode{Next: head}
	var pre = dummyNode

	for i := 1; i < start; i++ {
		pre = pre.Next
	}

	var cur = pre.Next

	for i := start; i < end; i++ {
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

	var dummyNode *ListNode

	for head != nil {
		next := head.Next
		head.Next = dummyNode
		dummyNode = head
		head = next
	}

	return dummyNode
}
