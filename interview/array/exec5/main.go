package main

import "fmt"

func main()  {
	var array = []int{2, 3, 1, 0, 2, 5, 3}

	result := findRepeatNumber(array)

	fmt.Println(result)
}

func findRepeatNumber(nums []int) int {// 1, 3, 2, 0, 2, 5, 3     3, 1, 2, 0, 2, 5, 3     0, 1, 2, 3, 2, 5, 3
	for i := 0; i < len(nums); i++{
		if i == nums[i]{
			continue
		}else {
			for {
				if nums[i] == nums[nums[i]]{
					return nums[i]
				}
				nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
				if i == nums[i]{
					break
				}
			}
		}
	}

	return -1
}