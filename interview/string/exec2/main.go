package main

import (
	"fmt"
)

func main()  {
	var str = "leetcode"

	result := isUnique(str)

	fmt.Println(result)
}

func isUnique(str string) bool {
	var tmpMap = make(map[uint8]int)

	for i := 0; i < len(str); i++{
		intStr := str[i] - '0'
		if _, ok := tmpMap[intStr]; ok{
			return false
		}else {
			tmpMap[intStr] = 1
		}
	}

	return true
}
