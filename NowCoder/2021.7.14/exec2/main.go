package main

import "fmt"

func main()  {
	var array = []int{1,2,3,2,2,2,5,4,2}
	result := MoreThanHalfNum_Solution(array)

	fmt.Println(result)
}

func MoreThanHalfNum_Solution( numbers []int ) int {
	// write code here

	var length = len(numbers)

	var tmpMap = make(map[int]int)

	for _, item := range numbers{
		tmpMap[item]++
	}

	for k, v := range tmpMap{
		if v > length / 2{
			return k
		}
	}

	return -1
}