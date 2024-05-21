package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var root TreeNode
	root.Val = 4

	var t1 TreeNode
	t1.Val = 2

	var t2 TreeNode
	t2.Val = 7

	var t3 TreeNode
	t3.Val = 1

	var t4 TreeNode
	t4.Val = 3

	var t5 TreeNode
	t5.Val = 5

	root.Left = &t1
	root.Right = &t2

	t1.Left = &t3
	t1.Right = &t4

	t2.Left = &t5

	levelIterator(&root)

	fmt.Println(levelArray)
}

func rightView(head *TreeNode) []int {
	//层序遍历
	levelIterator2(head)
	var res = make([]int, len(levelArray2))
	for index, item := range levelArray2 {
		res[index] = item[len(item)-1]
	}

	return res
}

var levelArray2 [][]int

func levelIterator2(node *TreeNode) {
	levelArray2 = [][]int{}
	dfs2(node, 0)
}

func dfs2(node *TreeNode, level int) {
	if node != nil {
		if len(levelArray2) == level {
			levelArray2 = append(levelArray2, []int{})
		}

		levelArray2[level] = append(levelArray2[level], node.Val)
		dfs2(node.Left, level+1)
		dfs2(node.Right, level+1)
	}
}

func publicNode(node *TreeNode, p, q *TreeNode) *TreeNode {
	for {
		if node.Val > p.Val && node.Val > q.Val {
			node = node.Left
		} else if node.Val < p.Val && node.Val < q.Val {
			node = node.Right
		} else {
			return node
		}
	}
}

func mirrorTree(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	var tmpNode = &TreeNode{Val: node.Val}

	tmpNode.Left = mirrorTree(node.Right)
	tmpNode.Right = mirrorTree(node.Left)
	return tmpNode
}

func mergeTree(node1, node2 *TreeNode) *TreeNode {
	if node1 == nil && node2 != nil {
		return node2
	} else if node1 != nil && node2 == nil {
		return node1
	}

	node1.Val += node2.Val
	node1.Left = mergeTree(node1.Left, node2.Left)
	node1.Right = mergeTree(node1.Right, node2.Right)
	return node1
}

func convertList(node *TreeNode) *TreeNode {
	middleIterator(node)

	var dummyNode = middleArray2[0]
	var pre = dummyNode

	for i := 1; i < len(middleArray2); i++ {
		pre.Right = middleArray2[i]
		middleArray2[i].Left = pre

		pre = pre.Right
	}

	return dummyNode
}

var middleArray2 []*TreeNode

func middleIterator2(node *TreeNode) {
	if node == nil {
		return
	}

	middleIterator(node.Left)
	middleArray2 = append(middleArray2, node)
	middleIterator(node.Right)
}

func maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return max(maxDepth(node.Left), maxDepth(node.Right)) + 1
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

var levelArray [][]int

func levelIterator(node *TreeNode) {
	levelArray = [][]int{}
	dfs(node, 0)
}

func dfs(node *TreeNode, level int) {
	if node != nil {
		if len(levelArray) == level {
			levelArray = append(levelArray, []int{})
		}

		levelArray[level] = append(levelArray[level], node.Val)
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
}

var frontArray []int
var middleArray []int
var lastArray []int

func lastIterator(node *TreeNode) {
	if node == nil {
		return
	}

	lastIterator(node.Left)
	lastIterator(node.Right)
	lastArray = append(lastArray, node.Val)
}

func middleIterator(node *TreeNode) {
	if node == nil {
		return
	}

	middleIterator(node.Left)
	middleArray = append(middleArray, node.Val)
	middleIterator(node.Right)
}

func frontIterator(node *TreeNode) {
	if node == nil {
		return
	}

	frontArray = append(frontArray, node.Val)
	frontIterator(node.Left)
	frontIterator(node.Right)
}
