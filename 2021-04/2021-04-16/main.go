package main

import "fmt"

/*旋转数组*/
func main()  {
	var array = []int{1, 2, 3, 4, 5, 6}
	overturnArray(array)

	fmt.Println(array)
}

func overturnArray(array []int)  {
	if array == nil || len(array) == 1{
		return
	}

	for i := 0; i < len(array) / 2; i++{
		array[i], array[len(array) - i - 1] = array[len(array) - i - 1], array[i]
	}

	return
}