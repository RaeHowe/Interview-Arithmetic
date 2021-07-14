package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var num1 = "29040821"
	var num2 = "10820385128"

	result := solve(num1, num2)

	fmt.Println(result)
}

func solve(s string ,t string ) string {
	// write code here
	var flag int
	var index1 = len(s) - 1
	var index2 = len(t) - 1

	var result string
	for index1 >= 0 || index2 >= 0{
		var tmpInt1, tmpInt2 int
		if index1 >= 0{
			tmpInt1 = int(s[index1] - '0')
			index1--
		}

		if index2 >= 0{
			tmpInt2 = int(t[index2] - '0')
			index2--
		}

		var tmpResult = tmpInt1 + tmpInt2 + flag

		if tmpResult >= 10{
			tmpResult = tmpResult - 10
			flag = 1
		}else {
			flag = 0
		}

		result =  strconv.Itoa(tmpResult) + result
	}

	if flag == 1{
		result = "1" + result
	}

	return result
}