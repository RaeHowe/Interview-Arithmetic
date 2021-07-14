package main

import "fmt"

func main()  {
	var nums = []int{-3, 2, -3, 4, 2}

	result := getMinNum(nums)

	fmt.Println(result)
}

func getMinNum(tmpSlice []int) int {
	var tmpMinData = 0 //-3
	var tmpData = 0 // -3

	for i := 0; i < len(tmpSlice); i++{
		tmpData += tmpSlice[i]
		if tmpData < tmpMinData{
			tmpMinData = tmpData
		}
	}

	if tmpMinData >= 1{
		return 0
	}else {
		return 1 - tmpMinData
	}
}

