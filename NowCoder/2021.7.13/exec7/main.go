package main

import "fmt"

func main()  {
	var array = []int{1, -2, 3, 5, -2, 6, -1}

	result := maxsumofSubarray(array)

	fmt.Println(result)
}

func maxsumofSubarray( arr []int ) int {
	// write code here
	var result int
	var tmpSum int

	for i := 0; i < len(arr); i++{
		tmpSum = tmpSum + arr[i]

		if tmpSum <= 0{
			tmpSum = 0
			continue
		}else {
			if tmpSum > result{
				result = tmpSum
			}
		}
	}

	return result
}
