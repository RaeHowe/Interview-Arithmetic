package main

import "fmt"

func main()  {
	var array = []int{1, 2, 3, 4, 5, 6}

	result := solve(6, 2, array)

	fmt.Println(result)
}

// [1, 2, 3, 4,           5, 6, 1, 2, 3, 4,         5, ,6] 10
func solve( n int ,  m int ,  a []int ) []int {
	// write code here
	m = m % n //如果数组长度为6，移动七次的话，其实就是相当于移动了一次
	var newArray = make([]int, 0)
	newArray = append(a, a...)
	return newArray[n - m: 2 * n - m]
}