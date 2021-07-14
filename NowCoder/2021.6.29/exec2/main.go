package main

import "fmt"

func main()  {
	var array = []int{5, 1, 2, 9, 10, 3, 4, 4}

	result := SmallK(array, 5)

	fmt.Println(result)
}

func SmallK(tmpSlice []int, k int) []int {
	QuickSort(tmpSlice, len(tmpSlice), 0, len(tmpSlice) - 1)

	return tmpSlice[:k]
}

func QuickSort(tmpSlice []int, length, start, end int) {
	if start == end{
		return
	}

	var index = Partition(tmpSlice, length, start, end)

	if start < index{
		QuickSort(tmpSlice, length, start, index - 1)
	}

	if index < end{
		QuickSort(tmpSlice, length, index + 1, end)
	}
}

func Partition(tmpSlice []int, length, start, end int) int {
	if tmpSlice == nil || length <= 0 || start < 0 || end > length {
		return -1
	}

	var small = start - 1
	var index = start

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
