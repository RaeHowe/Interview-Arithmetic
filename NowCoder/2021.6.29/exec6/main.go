package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main()  {
	var array = []int{30, 2, 9, 10, 8}

	result := max(array)

	fmt.Println(result)
}

func max(array []int) string {
	if len(array) == 0{
		return "0"
	}

	var tmpStrArray = make([]string, len(array))

	for _, item := range array{
		tmpStrArray = append(tmpStrArray, strconv.Itoa(item))
	}

	sort.Slice(tmpStrArray, func(i, j int) bool {
		strItem1, _ := strconv.Atoi(tmpStrArray[i] + tmpStrArray[j])
		strItem2, _ := strconv.Atoi(tmpStrArray[j] + tmpStrArray[i])

		return strItem1 > strItem2
	})

	var result string

	for _, item := range tmpStrArray{
		result += item
	}

	if result[0] == '0'{
		return "0"
	}

	return result
}