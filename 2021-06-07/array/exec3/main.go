package main

import "fmt"

func main()  {
	var array = []int{1, 8, 10, 14, 19, 28, 34, 49}

	var index = binarySearch(array, len(array), 100)

	fmt.Println(index)
}

func binarySearch(tmpSlice []int, length, target int) int {
	if length == 0{
		return -1
	}

	var startIndex = 0
	var endIndex = length - 1

	for startIndex <= endIndex{
		var middleIndex = (startIndex + endIndex) / 2
		if tmpSlice[middleIndex] == target{
			return middleIndex
		}

		if tmpSlice[middleIndex] > target{
			endIndex = middleIndex - 1
		}

		if tmpSlice[middleIndex] < target{
			startIndex = middleIndex + 1
		}
	}

	return -1
}
