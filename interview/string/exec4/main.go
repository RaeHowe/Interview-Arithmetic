package main

import (
	"fmt"
	"strings"
)

func main()  {
	var str1 = "waterbottle"
	var str2 = "erbottlewat"
	
	var result = judgement(str1, str2)

	fmt.Println(result)
}

func judgement(str1, str2 string) bool {
	var tmpS2 = str2 + str2

	return strings.Contains(tmpS2, str1)
}