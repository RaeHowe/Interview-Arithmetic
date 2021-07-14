package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	var treeNode1 = &TreeNode{}
	treeNode1.Val = 2

	var treeNode2 = &TreeNode{}
	treeNode2.Val = 1

	var treeNode3 = &TreeNode{}
	treeNode3.Val = 3

	treeNode1.Left = treeNode2
	treeNode1.Right = treeNode3

	result := threeOrders(treeNode1)

	fmt.Println(result)
}

var preArray []int
var inArray []int
var lastArray []int

func threeOrders( root *TreeNode ) [][]int {
	preOrder(root)
	inOrder(root)
	lastOrder(root)
	return [][]int{preArray, inArray, lastArray}
}

func preOrder(node *TreeNode) {
	if node == nil{
		return
	}

	preArray = append(preArray, node.Val)
	preOrder(node.Left)
	preOrder(node.Right)
}

func inOrder(node *TreeNode)  {
	if node == nil{
		return
	}

	inOrder(node.Left)
	inArray = append(inArray, node.Val)
	inOrder(node.Right)
}

func lastOrder(node *TreeNode)  {
	if node == nil{
		return
	}

	lastOrder(node.Left)
	lastOrder(node.Right)
	lastArray = append(lastArray, node.Val)
}