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

	frontIterator(&root)

	fmt.Println(preArray)

	middleIterator(&root)

	fmt.Println("#######")
	fmt.Println(middleArray)

	lastIterator(&root)
	fmt.Println("#######")
	fmt.Println(lastArray)
}

var preArray []int
var middleArray []int
var lastArray []int

func frontIterator(node *TreeNode) {
	if node == nil {
		return
	}

	preArray = append(preArray, node.Val)
	frontIterator(node.Left)
	frontIterator(node.Right)
}

func middleIterator(node *TreeNode) {
	if node == nil {
		return
	}

	middleIterator(node.Left)
	middleArray = append(middleArray, node.Val)
	middleIterator(node.Right)
}

func lastIterator(node *TreeNode) {
	if node == nil {
		return
	}

	lastIterator(node.Left)
	lastIterator(node.Right)
	lastArray = append(lastArray, node.Val)
}

var levelArray [][]int

func levelIterator(node *TreeNode) [][]int {
	levelArray = [][]int{}
	dfs(node, 0)
	return levelArray
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
