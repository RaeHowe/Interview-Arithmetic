package main

import "fmt"

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。
*/

func main()  {
	result := climbStairs(5)

	fmt.Println(result)
}

func climbStairs(n int) int {
	dp := make([]int, n + 1)
	dp[0] = 1 //如果楼梯层数为0的话，就只有一种方法
	dp[1] = 1 //如果楼梯层数为1的话，就只有一种方法

	//接下来从第二阶楼梯开始计算。第二阶楼梯一共有两种方法。爬不同的台阶，只要是大于1的，都可以被进行分解，最终分解的最小粒度，就是0和1这两种
	for i := 2; i < len(dp); i++ { //转化成数组来说，每一种楼梯攀爬的方式，都是前两种的和（也就是减去2和减去1）
		dp[i] = dp[i-2] + dp[i-1]
	}

	return dp[n]
}