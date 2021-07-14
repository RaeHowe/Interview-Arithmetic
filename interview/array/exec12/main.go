package main

import "fmt"

func main()  {
	var array = []int{0, 3, 2, 0, 8, 9}

	moveZero(array)

	fmt.Println(array)
}

func moveZero(tmpArray []int) {
	var index = 0

	for i := 0; i < len(tmpArray); i++{
		if tmpArray[i] != 0{
			tmpArray[index] = tmpArray[i]
			index++
		}
	}

	for i := index; i < len(tmpArray); i++{
		tmpArray[i] = 0
	}
}
