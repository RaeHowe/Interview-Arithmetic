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

	res := rightView(&root)
	fmt.Println(res)
}

func rebuildTree(frontArray, middleArray []int) *TreeNode {
	if len(frontArray) == 0 {
		return nil
	}

	root := &TreeNode{Val: frontArray[0]}

	var i int
	for i = 0; i < len(middleArray); i++ {
		if middleArray[i] == frontArray[0] {
			break
		}
	}

	root.Left = rebuildTree(frontArray[1:len(middleArray[:i])+1], middleArray[:i])
	root.Right = rebuildTree(frontArray[len(middleArray[:i])+1:], middleArray[i+1:])

	return root
}

var viewArray [][]int

func rightView(root *TreeNode) []int {
	viewArray = [][]int{}
	dfs2(root, 0)
	//取二维数组每一项的最后一个元素
	var res = make([]int, len(viewArray))

	for index, item := range viewArray {
		res[index] = item[len(item)-1]
	}

	return res
}

func dfs2(node *TreeNode, level int) {
	if node != nil {
		if len(viewArray) == level {
			viewArray = append(viewArray, []int{})
		}

		viewArray[level] = append(viewArray[level], node.Val)
		dfs2(node.Left, level+1)
		dfs2(node.Right, level+1)
	}
}

func publicNode(root *TreeNode, p, q int) *TreeNode {
	for {
		if root.Val > p && root.Val > q {
			root = root.Left
		} else if root.Val < p && root.Val < q {
			root = root.Right
		} else {
			return root
		}
	}
}

func isBalanceTree(node *TreeNode) bool {
	res := testDfs(node)
	if res == -1 {
		return false
	} else {
		return true
	}
}

func testDfs(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := testDfs(node.Left)
	right := testDfs(node.Right)
	if abs(left, right) > 1 {
		return -1
	}

	return max(left, right) + 1
}

func abs(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func mirror(tree1 *TreeNode) *TreeNode {
	if tree1 == nil {
		return nil
	}

	var tmpNode = &TreeNode{Val: tree1.Val}
	tmpNode.Left = mirror(tree1.Right)
	tmpNode.Right = mirror(tree1.Left)
	return tmpNode
}

func mergeTree(tree1, tree2 *TreeNode) *TreeNode {
	if tree1 == nil && tree2 != nil {
		return tree2
	} else if tree1 != nil && tree2 == nil {
		return tree1
	}

	tree1.Val += tree2.Val
	tree1.Left = mergeTree(tree1.Left, tree2.Left)
	tree1.Right = mergeTree(tree1.Right, tree2.Right)
	return tree1
}

func treeConvertList(node *TreeNode) *TreeNode {
	//中序遍历
	middleIterator2(node)

	//组装链表
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

	middleIterator2(node.Left)
	middleArray2 = append(middleArray2, node)
	middleIterator2(node.Right)
}

func maxDeep(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDeep(root.Left), maxDeep(root.Right)) + 1
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

var levelArray [][]int

func levelIterator(root *TreeNode) {
	levelArray = [][]int{}
	dfs(root, 0)
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
