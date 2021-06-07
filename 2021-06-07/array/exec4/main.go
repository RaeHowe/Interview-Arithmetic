package main

import "fmt"

func main()  {
	var array1 = []int{1, 8, 10, 12, 13}
	var array2 = []int{0, 5, 6, 8, 11, 16, 17}

	result := mergeArray(array1, array2)

	fmt.Println(result)
}

func mergeArray(tmpSlice1, tmpSlice2 []int) []int {
	var result []int

	var slice1Index = 0
	var slice2Index = 0

	for slice1Index != len(tmpSlice1) && slice2Index != len(tmpSlice2){
		if tmpSlice1[slice1Index] <= tmpSlice2[slice2Index]{
			result = append(result, tmpSlice1[slice1Index])
			slice1Index++
		}else {
			result = append(result, tmpSlice2[slice2Index])
			slice2Index++
		}
	}

	if slice1Index == len(tmpSlice1){
		result = append(result, tmpSlice2[slice2Index:]...)
	}else if slice2Index == len(tmpSlice2){
		result = append(result, tmpSlice1[slice1Index:]...)
	}

	return result
}
