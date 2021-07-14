package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var str1 = "289"
	var str2 = "3321"

	var result = strAdd(str1, str2)

	fmt.Println(result)
}

func strAdd(str1, str2 string) string {
	var result = ""

	var flag = 0

	var indexA, indexB = len(str1) - 1, len(str2) - 1

	for indexA >= 0 || indexB >= 0{
		var numA, numB = 0, 0
		if indexA >= 0{
			numA = int(str1[indexA] - '0')
		}

		if indexB >= 0{
			numB = int(str2[indexB] - '0')
		}

		var tmpResult = numA + numB + flag

		if tmpResult >= 10{
			flag = 1
			tmpResult = tmpResult - 10
		}else {
			flag = 0
		}

		result = strconv.Itoa(tmpResult) + result

		indexA--
		indexB--
	}

	if flag == 1{
		result = "1" + result
	}

	return result
}