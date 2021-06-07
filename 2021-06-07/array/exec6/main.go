package main

import "fmt"

func main()  {
	var array = []int{9, 2, 8, 10, 3, 3, 2, 1}

	result := findRepeatNumber(array)

	fmt.Println(result)
}

func findRepeatNumber(nums []int) []int {
	var result []int
	var tmpMap = make(map[int]int)

	for i := 0; i < len(nums); i++{
		if _, ok := tmpMap[nums[i]]; ok{ //如果可以从map中取到结果的话，就说明重复了
			result = append(result, nums[i])
		}else {
			tmpMap[nums[i]] = 1
		}
	}

	return result
}