package main

import "fmt"

func main()  {
	var array = []int{5, 1, 10, 9, 4, 8, 18 ,14}

	bubblingSort(array)

	fmt.Println(array)
}

func bubblingSort(tmpSlice []int) {
	for i := 0; i < len(tmpSlice) - 1; i++{
		for j := 0; j < len(tmpSlice) - i - 1; j++ {
			if tmpSlice[j] > tmpSlice[j + 1]{
				tmpSlice[j], tmpSlice[j + 1] = tmpSlice[j + 1], tmpSlice[j]
			}
		}
	}
}