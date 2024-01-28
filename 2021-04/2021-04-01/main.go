package main

import "fmt"

//最大子序和
/*
	给定一个整数数组 nums，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组[4,-1,2,1] 的和最大，为6 。
*/
func main()  {
	var array = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	result := maxSubArray(array)

	fmt.Println(result)
}

func maxSubArray(nums []int) int {
	if nums == nil{
		return 0
	}
	var tmpMaxData = nums[0]

	for i := 1; i < len(nums); i++{
		if nums[i] + nums[i - 1] > nums[i]{
			nums[i] += nums[i - 1]
		}

		if nums[i] > tmpMaxData{
			tmpMaxData = nums[i]
		}
	}

	return tmpMaxData
}