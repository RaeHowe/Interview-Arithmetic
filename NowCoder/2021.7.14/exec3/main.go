package main

import (
	"fmt"
)

func main()  {
	var array = []int{2,4,1,2,7,8,4}

	result := solve(array)

	fmt.Println(result)
}

func solve( a []int ) int {
	// write code here
	if a == nil{
		return 0
	}

	if len(a) == 1{
		return a[0]
	}

	if len(a) == 2{
		if a[0] > a[1]{
			return a[0]
		}else {
			return a[1]
		}
	}

	var pre = 0
	var middle = pre + 1
	var back = middle + 1

	var result int
	var maxIndex int

	for back <= len(a) - 1{
		if a[middle] > a[pre] && a[middle] > a[back]{
			if a[middle] > result{
				result = a[middle]
				maxIndex = middle
			}
		}

		pre++
		middle++
		back++
	}

	//最后两个值
	if a[middle] > a[pre]{
		if a[middle] > result{
			result = a[middle]
			maxIndex = middle
		}
	}

	return maxIndex
}