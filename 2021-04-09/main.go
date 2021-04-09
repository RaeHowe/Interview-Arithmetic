package main

import "fmt"

//数组中的第 K 个最大元素 (Leetcode)

func main()  {
	var tmpArray = []int{9, 5, 19, 12, 8, 4, 8, 10, 1} //长度是9， 输出第三大的也就是位置为7的  9-3+1 = 7

	QuickSort(tmpArray, len(tmpArray), 0, len(tmpArray) - 1) //先做一个快速排序，然后输出第K大的数据

	fmt.Println(tmpArray)

	//上面的数组是升序的，所以需要从后往前输出

	//输出第三大的元素

	fmt.Println(tmpArray[len(tmpArray) - 3])
}

func QuickSort(tmpSlice []int, length, start, end int)  {
	if start == end{
		return
	}

	var index = partition(tmpSlice, length, start, end)

	if start < index{
		QuickSort(tmpSlice, length, start, index - 1)
	}

	if index < end{
		QuickSort(tmpSlice, length, index + 1, end)
	}
}

func partition(tmpSlice []int, length, start, end int) int {
	if tmpSlice == nil || length < 0 || start < 0 || end > length{
		return -1
	}

	var index = start
	var small = start - 1

	for ;index < end; index++{
		if tmpSlice[index] < tmpSlice[end]{
			small++
			if small != index{
				tmpSlice[small], tmpSlice[index] = tmpSlice[index], tmpSlice[small]
			}
		}
	}

	small++

	tmpSlice[small], tmpSlice[end] = tmpSlice[end], tmpSlice[small]

	return small
}