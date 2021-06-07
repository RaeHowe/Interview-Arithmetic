package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var array = []int{12, 345, 2, 6, 7896}

	result := getNum(array)

	fmt.Println(result)
}

func getNum(tmpSlice []int) int {
	var result int

	for i := 0; i < len(tmpSlice); i++{
		var tmpNum = 0
		for j := 0; j < len(strconv.Itoa(tmpSlice[i])); j++{
			tmpNum++
		}

		if tmpNum % 2 == 0{
			result++
		}
	}

	return result
}
