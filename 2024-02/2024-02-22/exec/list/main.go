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
	l3.Value = 1

	var l4 ListNode
	l4.Value = 4

	var l5 ListNode
	l5.Value = 5

	var l6 ListNode
	l6.Value = 5

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5
	l5.Next = &l6

	node := removeAllDuplicateNode(&l1)

	fmt.Println(node.Value)
	//fmt.Println(node.Next.Value)
	//fmt.Println(node.Next.Next.Value)
	//fmt.Println(node.Next.Next.Next.Value)
	//fmt.Println(node.Next.Next.Next.Next.Value)
}

func removeAllDuplicateNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var dummyNode = &ListNode{Next: head}
	var pre = dummyNode

	for head != nil && head.Next != nil {
		if pre.Next.Value == head.Next.Value {
			var tmpVal = pre.Next.Value
			for pre.Next != nil && pre.Next.Value == tmpVal {
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

func removeDuplicateNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

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

func isPail(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	//找到链表的中间位置
	var slow, fast = head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var dummyNode *ListNode
	var cur = slow

	//反转后半部分链表
	for cur != nil {
		next := cur.Next
		cur.Next = dummyNode
		dummyNode = cur
		cur = next
	}

	var mid = dummyNode

	for mid != nil {
		if mid.Value != head.Value {
			return false
		}

		mid = mid.Next
		head = head.Next
	}

	return true
}

func listSort(head *ListNode) *ListNode {
	var tmpArray = make([]int, 0)

	for head != nil {
		tmpArray = append(tmpArray, head.Value)
		head = head.Next
	}

	quickSort(tmpArray, len(tmpArray), 0, len(tmpArray)-1)

	var dummyNode = &ListNode{}
	var pre = dummyNode

	for _, item := range tmpArray {
		var tmpNode = &ListNode{Value: item}

		pre.Next = tmpNode
		pre = pre.Next
	}

	return dummyNode.Next
}

func quickSort(tmpSlice []int, length, start, end int) {
	if start == end {
		return
	}

	var index = partition(tmpSlice, length, start, end)

	if start < index {
		quickSort(tmpSlice, length, start, index-1)
	}

	if index < end {
		quickSort(tmpSlice, length, index+1, end)
	}
}

func partition(tmpSlice []int, length, start, end int) int {
	if tmpSlice == nil || length <= 0 || start < 0 || end > length {
		return -1
	}

	var index = start
	var small = start - 1

	for ; index < end; index++ {
		if tmpSlice[index] < tmpSlice[end] {
			small++
			if small != index {
				tmpSlice[small], tmpSlice[index] = tmpSlice[index], tmpSlice[small]
			}
		}
	}

	small++

	tmpSlice[small], tmpSlice[end] = tmpSlice[end], tmpSlice[small]

	return small
}

func listSum(head1, head2 *ListNode) *ListNode {
	if head1 == nil && head2 != nil {
		return head2
	}

	if head1 != nil && head2 == nil {
		return head1
	}

	if head1 == nil && head2 == nil {
		return nil
	}

	newHead1 := reverseList(head1) //3 2 1
	newHead2 := reverseList(head2) //7 7

	var dummyNode = &ListNode{}
	var pre = dummyNode

	var flagAdd = 0
	for newHead1 != nil || newHead2 != nil {
		var list1Val int
		var list2Val int
		if newHead1 == nil {
			list1Val = 0
		} else {
			list1Val = newHead1.Value
			newHead1 = newHead1.Next
		}

		if newHead2 == nil {
			list2Val = 0
		} else {
			list2Val = newHead2.Value
			newHead2 = newHead2.Next
		}

		tmpResult := list1Val + list2Val + flagAdd
		if tmpResult >= 10 {
			tmpResult = tmpResult % 10
			flagAdd = 1
		} else {
			flagAdd = 0
		}

		var tmpNode = &ListNode{Value: tmpResult}
		pre.Next = tmpNode
		pre = pre.Next
	}

	if flagAdd > 0 {
		var tmpNode = &ListNode{Value: flagAdd}
		pre.Next = tmpNode
	}

	return reverseList(dummyNode.Next)
}

func reverseList(head *ListNode) *ListNode {
	var dummyNode *ListNode
	for head != nil {
		next := head.Next
		head.Next = dummyNode
		dummyNode = head
		head = next
	}

	return dummyNode
}

func firstPublicNode(head1, head2 *ListNode) *ListNode {
	if head1 == nil || head2 == nil {
		return nil
	}

	cur1, cur2 := head1, head2

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

// 删除链表倒数第n个节点
func removeLastKNode(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

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
		fast = fast.Next
		pre = pre.Next
	}

	//找到了倒数第K个节点
	pre.Next = pre.Next.Next

	return dummyNode.Next
}
