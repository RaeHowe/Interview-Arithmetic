package main

import "fmt"

func main() {
	var array = []int{2, 4, 9, 10, 14, 17, 20, 21}

	res := binarySearch(array, 17)
	fmt.Println(res)

	var array2 = []int{9, 10, 2, 6, 1, 21, 17}

	quickSort(array2, len(array2), 0, len(array2)-1)

	fmt.Println(array2)

	var array3 = []int{9, 10, 2, 6, 1, 21, 17}

	bubblingSort(array3)

	fmt.Println(array3)

}

func bubblingSort(tmpSlice []int) {
	for i := 0; i < len(tmpSlice)-1; i++ {
		for j := 0; j < len(tmpSlice)-i-1; j++ {
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

func binarySearch(tmpSlice []int, val int) int {
	var start = 0
	var end = len(tmpSlice) - 1

	for start <= end {
		var middle = (start + end) / 2
		if tmpSlice[middle] > val {
			end = middle - 1
		}

		if tmpSlice[middle] < val {
			start = middle + 1
		}

		if tmpSlice[middle] == val {
			return middle
		}
	}

	return -1
}
