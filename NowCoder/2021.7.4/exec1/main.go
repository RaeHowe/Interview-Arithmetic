package main

import "fmt"

type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}

func main()  {
	var node1 TreeNode
	node1.Val = 3

	var node2 TreeNode
	node2.Val = 9

	var node3 TreeNode
	node3.Val = 20

	var node4 TreeNode
	node4.Val = 15

	var node5 TreeNode
	node5.Val = 7

	node1.Left = &node2
	node1.Right = &node3

	node3.Left = &node4
	node3.Right = &node5

	result := levelOrder(&node1)

	fmt.Println(result)
}

var res [][]int

func levelOrder(root *TreeNode) [][]int {
	res = [][]int{}
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, level int) {
	if root != nil{
		if len(res) == level{
			res = append(res, []int{})
		}

		res[level] = append(res[level], root.Val)
		dfs(root.Left, level + 1)
		dfs(root.Right, level + 1)
	}
}