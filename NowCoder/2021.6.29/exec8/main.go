package main

import (
	"fmt"
	"sort"
)

func main()  {
	var array = []int{3 ,2, 4}

	result := twoSum(array, 6)

	fmt.Println(result)
}

func twoSum( numbers []int ,  target int) []int {
	var tmpMap = make(map[int]int) //key:保存的值  value: 这个值在数组中的索引

	/*
		如果是两个相同的数字相加的情况的话，map中保存的value值，是最后那个值的索引，在下面遍历的时候，会首先访问到数组前面的值，判断
		索引如果不相等，说明存在两个相同的数字，从而得到结果
	*/
	for index, item := range numbers{
		tmpMap[item] = index
	}

	var resultArray = make([]int, 0)
	for index, item := range numbers{
		var needNumber = target - item

		if value, ok := tmpMap[needNumber]; ok{
			if needNumber == item{
				if index != value{
					resultArray = append(resultArray, index + 1)
					resultArray = append(resultArray, value + 1)
					break
				}else {
					//没结果
					continue
				}
			}else {
				resultArray = append(resultArray, index + 1)
				resultArray = append(resultArray, value + 1)
				break
			}
		}
	}

	sort.Ints(resultArray)

	return resultArray
}
