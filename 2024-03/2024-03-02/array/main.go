package main

import "fmt"

func main() {
	var array = []int{10, 2, 19, 8, 7, 2, 12, 13}

	bubblingSort(array)

	fmt.Println(array)
}

func bubblingSort(tmpSlice []int) {
	for i := 0; i < len(tmpSlice)-1; i++ {
		for j := 0; j < len(tmpSlice)-1-i; j++ {
			if tmpSlice[j] > tmpSlice[j+1] {
				tmpSlice[j], tmpSlice[j+1] = tmpSlice[j+1], tmpSlice[j]
			}
		}
	}
}

func quickSort(tmpSlice []int, length, start, end int) {
	var index = partition(tmpSlice, length, start, end)

	if start < index {
		quickSort(tmpSlice, length, start, index-1)
	}

	if index < end {
		quickSort(tmpSlice, length, index+1, end)
	}
}

func partition(tmpSlice []int, length, start, end int) int {
	if tmpSlice == nil || length <= 0 || start < 0 || end > length {
		return -1
	}

	var index = start
	var small = start - 1

	for ; index < end; index++ {
		if tmpSlice[index] < tmpSlice[end] {
			small++
			if small != index {
				tmpSlice[small], tmpSlice[index] = tmpSlice[index], tmpSlice[small]
			}
		}
	}

	small++

	tmpSlice[small], tmpSlice[end] = tmpSlice[end], tmpSlice[small]

	return small
}

func binarySearch(tmpSlice []int, target int) int {
	var start = 0
	var end = len(tmpSlice) - 1

	for start <= end {
		var middleIndex = (start + end) / 2

		if tmpSlice[middleIndex] == target {
			return middleIndex
		}

		if tmpSlice[middleIndex] < target {
			start = middleIndex + 1
		}

		if tmpSlice[middleIndex] > target {
			end = middleIndex - 1
		}
	}

	return -1
}
