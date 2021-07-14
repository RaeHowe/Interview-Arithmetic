package main

type TreeNode struct {
   Val int
   Left *TreeNode
   Right *TreeNode
}

func main()  {
	
}

func IsBalanced_Solution( pRoot *TreeNode ) bool {
	return getMaxDepth(pRoot) > -1
}

func getMaxDepth(pRoot *TreeNode) int {
	if pRoot == nil{
		return 0
	}

	left := getMaxDepth(pRoot.Left)
	right := getMaxDepth(pRoot.Right)
	if left == -1 || right == -1 || abs(left - right) > 1{
		return -1
	}

	return 1 + max(left, right)
}

func max(a, b int) int {
	if a >= b{
		return a
	}else {
		return b
	}
}

func abs(a int) int {
	if a < 0{
		return -a
	}else {
		return a
	}
}