package main

import "fmt"

func main() {
	var array = []int{6, 1, 2, 4, 19, 2, 13, 7, 24}

	quickSort(array, len(array), 0, len(array)-1)

	fmt.Println(array)

	var array2 = []int{1, 5, 8, 19, 19, 20, 28, 31, 40}

	binarySearch(array2, 20)
}

/*
给定一个长度为n的数组nums，请你找到峰值并返回其索引。数组可能包含多个峰值，在这种情况下，返回任何一个所在位置即可。

1.峰值元素是指其值严格大于左右相邻值的元素。严格大于即不能有等于
2.假设 nums[-1] = nums[n] = −∞−∞
3.对于所有有效的 i 都有 nums[i] != nums[i + 1]
4.你可以使用O(logN)的时间复杂度实现此问题吗？

数据范围：
1≤nums.length≤2×105 1≤nums.length≤2×105
−231<=nums[i]<=231−1−231<=nums[i]<=231−1

如输入[2,4,1,2,7,8,4]时，会形成两个山峰，一个是索引为1，峰值为4的山峰，另一个是索引为5，峰值为8的山峰，如下图所示
*/
func findMountain(tmpSlice []int) int {
	var start = 0
	var end = len(tmpSlice) - 1

	for start < end {
		var middleIndex = (start + end) / 2

		if tmpSlice[middleIndex] < tmpSlice[middleIndex+1] {
			start = middleIndex + 1
		} else {
			end = middleIndex
		}
	}

	return end
}

func binarySearch(tmpSlice []int, target int) int {
	if len(tmpSlice) == 0 {
		return -1
	}

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

func quickSort(tmpSlice []int, length, start, end int) {
	if start == end {
		return
	}

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
