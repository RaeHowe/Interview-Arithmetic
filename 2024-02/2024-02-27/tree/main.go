package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val    int
	Before *ListNode
	Next   *ListNode
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

	//var t6 TreeNode
	//t6.Val = 8

	root.Left = &t1
	root.Right = &t2

	t1.Left = &t3
	t1.Right = &t4

	t2.Left = &t5
	//t2.Right = &t6

	node := isCompleteTree(&root)

	fmt.Println(node)

	//fmt.Println(node.Val)
	//fmt.Println(node.Right.Val)
	//fmt.Println(node.Right.Right.Val)
	//fmt.Println(node.Right.Right.Right.Val)
	//fmt.Println(node.Right.Right.Right.Right.Val)
	//fmt.Println(node.Right.Right.Right.Right.Right.Val)

}

// solve 二叉树的右视图
func solve(preOrder []int, inOrder []int) []int {
	return nil
}

// 获取二叉树两个节点的最近公共节点
func publicNode(root *TreeNode, p, q int) int {
	for {
		if root.Val > p && root.Val > q {
			root = root.Left
		} else if root.Val < p && root.Val < q {
			root = root.Right
		} else {
			return root.Val
		}
	}
}

func IsBalanced_Solution(pRoot *TreeNode) bool {
	testDfs(pRoot)
	return res
}

var res = true

func testDfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := testDfs(root.Left)
	right := testDfs(root.Right)
	if abs(left, right) > 1 {
		res = false
		return -1
	}

	return max(left, right) + 1
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// 判断是不是完全二叉树
func isCompleteTree(root *TreeNode) bool {
	//层序遍历
	getLayerTree(root)

	for i := 0; i < len(layerArray); i++ {
		if float64(len(layerArray[i])) != math.Pow(2, float64(i)) {
			return false
		}
	}

	return true
}

var layerArray [][]int

func getLayerTree(root *TreeNode) {
	layerArray = [][]int{}
	dfs(root, 0)
}

func dfs(root *TreeNode, level int) {
	if root != nil {
		if len(layerArray) == level {
			layerArray = append(layerArray, []int{})
		}

		layerArray[level] = append(layerArray[level], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
}

// 判断是不是二叉搜索树
func JudgementSearchTree(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}

	//中序遍历获取到有序的数组
	middleIterator(root)

	for i := 1; i < len(middleArray); i++ {
		if middleArray[i].Val < middleArray[i-1].Val {
			return false
		}
	}

	return true
}

// Mirror 二叉树的镜像
func Mirror(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var tmpNode = &TreeNode{Val: root.Val}
	tmpNode.Left = Mirror(root.Right)
	tmpNode.Right = Mirror(root.Left)
	return tmpNode
}

// MergeTree 合并二叉树
func MergeTree(tree1, tree2 *TreeNode) *TreeNode {
	if tree1 == nil {
		return tree2
	}

	if tree2 == nil {
		return tree1
	}

	tree1.Val += tree2.Val
	tree1.Left = MergeTree(tree1.Left, tree2.Left)
	tree1.Right = MergeTree(tree1.Right, tree2.Right)
	return tree1
}

// Convert 将二叉树转换为双向链表
func Convert(pRootOfTree *TreeNode) *TreeNode {
	if pRootOfTree == nil || (pRootOfTree.Left == nil && pRootOfTree.Right == nil) {
		return pRootOfTree
	}

	//中序遍历获取到有序的数组
	middleIterator(pRootOfTree)

	//构建双向链表
	var dummyNode = middleArray[0]
	var pre = dummyNode

	for i := 1; i < len(middleArray); i++ {
		pre.Right = middleArray[i]
		middleArray[i].Left = pre

		pre = pre.Right
	}

	return dummyNode
}

var middleArray []*TreeNode

func middleIterator(root *TreeNode) {
	if root == nil {
		return
	}

	middleIterator(root.Left)
	middleArray = append(middleArray, root)
	middleIterator(root.Right)
}

// 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
