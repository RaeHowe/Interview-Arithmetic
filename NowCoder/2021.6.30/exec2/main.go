package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	var treeNode1 = &TreeNode{}
	treeNode1.Val = 1

	var treeNode2 = &TreeNode{}
	treeNode2.Val = 2

	var treeNode3 = &TreeNode{}
	treeNode3.Val = 3

	var treeNode4 = &TreeNode{}
	treeNode4.Val = 4

	var treeNode5 = &TreeNode{}
	treeNode5.Val = 5

	var treeNode6 = &TreeNode{}
	treeNode6.Val = 6

	var treeNode7 = &TreeNode{}
	treeNode7.Val = 7

	treeNode1.Left = treeNode2
	treeNode1.Right = treeNode3
	treeNode2.Left = treeNode4
	treeNode2.Right = treeNode5
	treeNode3.Left = treeNode6
	treeNode3.Right = treeNode7

	result := lastOrder(treeNode1)

	for _, item := range result{
		fmt.Println(item)
	}
}

func lastOrder(root *TreeNode) []int {
	var stack = []*TreeNode{}
	var result []int
	var prev *TreeNode
	for root != nil || len(stack) > 0{
		for root != nil{
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		if root.Right == nil || root.Right == prev{
			result = append(result, root.Val)
			prev = root
			root = nil
		}else {
			stack = append(stack, root)
			root = root.Right
		}
	}

	return result
}

//前序遍历（非递归方式）
func preOrder(root *TreeNode) []int {
	var stack = []*TreeNode{}
	var result []int

	for root != nil || len(stack) > 0 {
		for root != nil{
			stack = append(stack, root)
			result = append(result, root.Val)
			root = root.Left
		}

		//如果左子树没有对象了的话，就弹出栈
		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		root = root.Right
	}

	return result
}

//中序遍历（非递归方式）
func levelOrder( root *TreeNode ) []int {
	var stack = []*TreeNode{}
	var result []int
	for root != nil || len(stack) > 0{
		//不断地把左子树的元素放进到stack栈中
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		//左子树遍历完毕之后，弹出一个左子树元素，放置到结果数组里面
		root = stack[len(stack) - 1]//取到栈的最后一个元素
		result = append(result, root.Val) //放置到结果数组里面
		stack = stack[:len(stack) - 1] //弹出栈的最后一个元素
		root = root.Right //取右子树
	}

	return result
}

