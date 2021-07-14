package main


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main()  {

}

func maxDepth( root *TreeNode ) int {
	// write code here
	if root == nil{
		return 0
	}

	var depthLeft = maxDepth(root.Left)
	var depthRight = maxDepth(root.Right)

	return max(depthLeft, depthRight) + 1
}

func max(a, b int) int {
	if a >= b{
		return a
	}

	return b
}