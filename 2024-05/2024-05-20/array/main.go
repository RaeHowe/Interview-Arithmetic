package main

func main() {

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
			if index != small {
				tmpSlice[index], tmpSlice[small] = tmpSlice[small], tmpSlice[index]
			}
		}
	}

	small++
	tmpSlice[small], tmpSlice[end] = tmpSlice[end], tmpSlice[small]

	return small
}

func bubbling(tmpSlice []int) {
	for i := 0; i < len(tmpSlice)-1; i++ {
		for j := 0; j < len(tmpSlice)-i-1; j++ {
			if tmpSlice[j] > tmpSlice[j+1] {
				tmpSlice[j], tmpSlice[j+1] = tmpSlice[j+1], tmpSlice[j]
			}
		}
	}
}
