package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main()  {
	var array = []string{"a","b","c","b"}

	result := topKstrings(array, 2)

	fmt.Println(result)
}

func topKstrings( strings []string ,  k int ) [][]string {
	// write code here
	var result = make([][]string, 0)

	var tmpMap = make(map[string]int)

	for _, item := range strings {
		tmpMap[item]++
	}

	var tmpArray = make([]int, 0)

	for _, v := range tmpMap {
		tmpArray = append(tmpArray, v)
	}

	sort.Ints(tmpArray)

	if k < len(tmpArray) {
		tmpArray = tmpArray[len(tmpArray)-k:]
	}

	for i := len(tmpArray) - 1; i >= 0; i--{
		for k, v := range tmpMap{
			if v == tmpArray[i]{
				result = append(result, []string{k, strconv.Itoa(v)})
				break
			}
		}
	}

	return result
}