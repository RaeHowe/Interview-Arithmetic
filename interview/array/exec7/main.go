package main

import "fmt"

func main()  {
	var array = []int{17, 18, 5, 4, 6, 1}

	replaceElements(array)

	fmt.Println(array)
}

func replaceElements(array []int) []int {
	var length = len(array)
	var result = make([]int, length)
	result[length - 1] = -1

	for i := length - 2; i >= 0; i--{
		array[i] = maxNum(array[i + 1], result[i + 1])
	}

	return result
}

func maxNum(a, b int) int {
	if a >= b{
		return a
	}else {
		return b
	}
}
