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

	node := reverseBetween(&l1, 2, 4)

	fmt.Println(node.Value)
	fmt.Println(node.Next.Value)
	fmt.Println(node.Next.Next.Value)
	fmt.Println(node.Next.Next.Next.Value)
	fmt.Println(node.Next.Next.Next.Next.Value)
	fmt.Println(node.Next.Next.Next.Next.Next.Value)
}

func deleteAllDuplicateNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var dummyNode = &ListNode{Next: head}
	var pre = dummyNode

	for head != nil || head.Next != nil {
		if pre.Next.Value == head.Next.Value {
			var sameVal = pre.Next.Value
			for pre.Next.Value == sameVal && pre.Next != nil {
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

func deleteDuplicateNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var dummyNode = &ListNode{Next: head}
	var pre = dummyNode

	for head != nil && head.Next != nil {
		if pre.Next.Value == head.Next.Value {
			pre.Next = pre.Next.Next
		} else {
			pre = pre.Next
		}
		head = head.Next
	}

	return dummyNode.Next
}

// 判断一个链表是否是回文链表
func judgementLoopList(head *ListNode) bool {
	//找中间点
	var slow, fast = head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	//中间点是slow，反转链表后半部分
	var dummyNode *ListNode
	var cur = slow
	for cur != nil {
		temp := cur.Next
		cur.Next = dummyNode
		dummyNode = cur
		cur = temp
	}

	var mid = dummyNode

	for mid != nil {
		if head.Value != mid.Value {
			return false
		}

		head = head.Next
		mid = mid.Next
	}

	return true
}

func listSum(head1, head2 *ListNode) *ListNode {
	var dummyNode = &ListNode{}
	var pre = dummyNode

	var newHead1 = reverseList(head1)
	var newHead2 = reverseList(head2)

	var flagAddNum int
	for newHead1 != nil || newHead2 != nil {
		var num1, num2 int
		var tmpResult int
		if newHead1 != nil {
			num1 = newHead1.Value
			newHead1 = newHead1.Next
		}

		if newHead2 != nil {
			num2 = newHead2.Value
			newHead2 = newHead2.Next
		}

		tmpResult = num1 + num2 + flagAddNum
		if tmpResult >= 10 {
			flagAddNum = 1
			tmpResult = tmpResult % 10
		} else {
			flagAddNum = 0
		}

		var tmpNode = &ListNode{Value: tmpResult}
		pre.Next = tmpNode
		pre = pre.Next
	}

	if flagAddNum > 0 {
		var tmpNode = &ListNode{Value: flagAddNum}
		pre.Next = tmpNode
	}

	return reverseList(dummyNode.Next)
}

func firstPublicNode(head1, head2 *ListNode) *ListNode {
	if head1 == nil || head2 == nil {
		return nil
	}

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

func deleteLastKNode(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	var last = head

	for i := 0; i < k; i++ {
		if last == nil {
			return nil
		}
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

// dummyNode -> 1 -> 2 -> 3 -> 4 -> 5
func getLastKNode(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	var last = head

	for i := 0; i < k; i++ {
		if last == nil {
			//临界情况
			return last
		}
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

	var slow = head
	var fast = head.Next.Next

	for fast != nil && head.Next != nil {
		if slow == fast {
			return true
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

func mergeKList(heads []*ListNode) *ListNode {
	var length = len(heads)

	if length == 1 {
		return heads[0]
	}

	var middle = length / 2
	var left = mergeKList(heads[:middle])
	var right = mergeKList(heads[middle:])
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
		if head1.Value <= head2.Value {
			pre.Next = head2
			head2 = head2.Next
		} else {
			pre.Next = head1
			head1 = head1.Next
		}

		pre = pre.Next
	}

	if head1 == nil && head2 != nil {
		pre.Next = head2
	} else if head1 != nil && head2 == nil {
		pre.Next = head1
	}

	return dummyNode.Next
}

// pre -> 1(pre) -> 2(cur) -> 3(temp) -> 4 -> 5 -> 6
func reverseBetween(head *ListNode, k, m int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var dummyNode = &ListNode{Next: head}
	var pre = dummyNode

	for i := 1; i < k; i++ {
		pre = pre.Next
	}

	var cur = pre.Next
	for i := k; i < m; i++ {
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
