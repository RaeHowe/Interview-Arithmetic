package main

import "fmt"

func main()  {
	var array = []int{5, 1, 10, 9, 4, 8, 18 ,14}

	quickSort(array, len(array), 0, len(array) - 1)

	fmt.Println(array)
}

func quickSort(tmpSlice []int, length, start, end int)  {
	if start == end{
		return
	}
	
	var index = getIndex(tmpSlice, length, start, end)

	if start < index{
		quickSort(tmpSlice, length, start, index - 1)
	}

	if index < end{
		quickSort(tmpSlice, length, index + 1, end)
	}
}

func getIndex(tmpSlice []int, length, start, end int) int {
	if tmpSlice == nil || length <= 0 || start < 0 || end > length{
		return -1
	}

	var small = start - 1
	var index = start

	for ;index < end; index++{
		if tmpSlice[index] < tmpSlice[end]{
			small++
			if small != index{
				tmpSlice[small], tmpSlice[end] = tmpSlice[end], tmpSlice[small]
			}
		}
	}

	small++

	tmpSlice[small], tmpSlice[end] = tmpSlice[end], tmpSlice[small]

	return small
}
