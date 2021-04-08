package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main()  {

}

//广度优先搜索
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	var nodeArray = []*TreeNode{root}
	var nodeValueArray = []int{root.Val}

	for len(nodeArray) != 0{
		var tmpNode = nodeArray[0] //取出节点数组的第一个元素为需要进行逻辑操作的节点
		var tmpSum = nodeValueArray[0] //取出节点值数组的第一个元素为需要相加的值

		/*一定要涉及到一个出队列的过程，这个逻辑加入进去之后，就可以根据判断队列中是否存在元素来判断是否结束循环**/
		nodeArray = nodeArray[1:] //把刚才获取到的第一个元素pop出去
		nodeValueArray = nodeValueArray[1:] //同上

		if tmpNode.Left == nil && tmpNode.Right == nil{ //判断这个节点的左子树和右子树是否还存在子节点，如果不存在的话，说明是叶子节点
			//去判断目前获取到的临时和，如果和目标的值相同的话，就返回true，如果不相同，则继续循环迭代
			if tmpSum == sum{
				return true
			}

			continue
		}

		if tmpNode.Left != nil{
			nodeArray = append(nodeArray, tmpNode.Left)
			nodeValueArray = append(nodeValueArray, tmpSum + tmpNode.Left.Val)
		}

		if tmpNode.Right != nil{
			nodeArray = append(nodeArray, tmpNode.Right)
			nodeValueArray = append(nodeValueArray, tmpSum + tmpNode.Right.Val)
		}
	}

	return false
}