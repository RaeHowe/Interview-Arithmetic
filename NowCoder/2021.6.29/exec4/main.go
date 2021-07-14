package main

import (
	"fmt"
	"sort"
)

//type Interval struct {
//	Start int
//	End int
//}

func main()  {
	var array = [][]int{{80, 100}, {20, 60}, {10, 30}, {150, 180}}

	result := mergeSection(array)

	fmt.Println(result)
}

func mergeSection(array [][]int) [][]int {
	var length = len(array) - 1

	if length < 1{
		return array
	}

	//根据二维数组中每个数组的第一个数字大小来进行排序
	sort.Slice(array, func(i, j int) bool {
		return array[i][0] < array[j][0]
	})

	for i := 0; i < length; i++{
		if array[i][1] >= array[i + 1][0]{ //如果前一个数组的第二个元素的大小大于第二个数组的第一个元素，就可以merge
			if array[i][1] < array[i + 1][1]{
				array[i][1] = array[i + 1][1]
			}

			array = append(array[:i + 1], array[i + 2:]...)
			i--
			length--
		}
	}

	return array
}