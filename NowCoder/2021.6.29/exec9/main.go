package main

import "fmt"

func main()  {
	var array = []int{2,3,4,5}

	result := maxLength(array)

	fmt.Println(result)
}

func maxLength(array []int) int {
	var tmpMap = make(map[int]int) //map的key对应了数组中的值，value对应了数组中元素的索引位置，从1开始计算

	var tmpLength int
	var resultLength int

	for index, item := range array{
		if value, ok := tmpMap[item]; ok{
			//如果存在的话，就说明重复了。先去修改临时长度变量的值，再去修改最终长度的值。
			tmpLength = max(tmpLength, value)
		}

		resultLength = max(resultLength, index + 1 - tmpLength)
		tmpMap[item] = index + 1
	}

	return resultLength
}

func max(i, j int) int {
	if i >= j {
		return i
	}else {
		return j
	}
}