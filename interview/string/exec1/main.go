package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var num = 234 // 9   24

	result := calculate(num)

	fmt.Println(result)
}

func calculate(num int) int {
	var sumOfPlus = 0
	var sumOfProduct = 1
	var strNum = strconv.Itoa(num)

	for i := 0; i < len(strNum); i++{
		sumOfPlus += int(strNum[i] - '0')
		sumOfProduct *= int(strNum[i] - '0')
	}

	return sumOfProduct - sumOfPlus
}