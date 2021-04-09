package main

import "fmt"

//数组中的第 K 个最大元素 (Leetcode)

func main()  {
	var tmpArray = []int{3, 5, 8, 1, 2, 9, 4, 7, 6} //长度是9， 输出第三大的也就是位置为7的  9-3+1 = 7

	result := test(tmpArray, 0, len(tmpArray) - 1, 3)

	fmt.Println(result)
}

//数组长度设为9  [3, 5, 8, 1, 2, 9, 4, 7, 6]
func test(tmpSlice []int, start, end, k int) int { //取第三个大的，k=3
	if tmpSlice == nil{
		return -1
	}

	var i, j = start, start //设置i，j索引位置，一开始都从头开始进行索引
	var pivot = start //取数组的第一个元素为pivot支点
	//对换pivot和临时数组的最后一个元素
	tmpSlice[pivot], tmpSlice[end] = tmpSlice[end], tmpSlice[pivot] //6, 5, 8, 1, 2, 9, 4, 7, 3

	for j != end{ //索引j的位置如果不等于最后一个索引位置的话就一直循环
		if tmpSlice[j] <= tmpSlice[end] {
			//如果j所在位置的值小于等于支点的值
			tmpSlice[i], tmpSlice[j] = tmpSlice[j], tmpSlice[i]
			i++
		}
		j++
	}

	//到这一步骤的时候，j和end的值相等
	tmpSlice[i], tmpSlice[end] = tmpSlice[end], tmpSlice[i]

	if end > k + i - 1{ //如果最终的top值，在i值右半部分的数组里面的话
		return test(tmpSlice, i + 1, end, k)
	}else if end < k + i - 1{ //如果最终的top值，在i值左半部分的数组里面的话
		return test(tmpSlice, start, i - 1, k - (end - i + 1))
	}else { //如果最终的top值，恰好是索引i所指向的数值的话，那么就直接返回就行了
		return tmpSlice[i]
	}
}

//===================================================================================================

//使用快速选择排序方法 （快速选择排序算法专门用来解决：在未排序的数组中寻找第k小/第k大的元素。）
func QuickSelectSort(tmpSlice []int, k int) int {
	if tmpSlice == nil{
		return 0
	}

	var left = 0
	var right = len(tmpSlice) - 1

	for {
		var position = partition2(tmpSlice, left, right)
		if position == k - 1{
			return tmpSlice[position]
		}else if position > k - 1{
			right = position - 1
		}else {
			left = position + 1
		}
	}
}

func partition2(tmpSlice []int, left, right int) int {
	var pivot = left

	var l = left + 1
	var r = right

	for l <= r{
		for l <= r && tmpSlice[l] >= tmpSlice[pivot]{
			l++
		}

		for l <= r && tmpSlice[r] <= tmpSlice[pivot]{
			r--
		}

		if (l <= r) && (tmpSlice[l] < tmpSlice[pivot]) && (tmpSlice[r] > tmpSlice[pivot]){
			tmpSlice[l + 1], tmpSlice[r - 1] = tmpSlice[r - 1], tmpSlice[l + 1]
		}
 	}

 	tmpSlice[pivot], tmpSlice[r] = tmpSlice[r], tmpSlice[pivot]

 	return r
}

//=============================================================================================//

//快速排序方法
func QuickSort(tmpSlice []int, length, start, end int)  {
	if start == end{
		return
	}

	var index = partition(tmpSlice, length, start, end) //查找分界位置

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

	//排序数组子部分的数组元素
	for ;index < end; index++{
		if tmpSlice[index] < tmpSlice[end]{
			small++
			if small != index{
				tmpSlice[small], tmpSlice[index] = tmpSlice[index], tmpSlice[small] //
			}
		}
	}

	small++

	tmpSlice[small], tmpSlice[end] = tmpSlice[end], tmpSlice[small]

	return small
}