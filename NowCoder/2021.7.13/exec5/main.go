package main

import "fmt"

func main()  {
	var str = "yamatomaya"

	result := judge(str)

	fmt.Println(result)
}

func judge( str string ) bool {
	// write code here
	for i := 0; i < len(str) / 2; i++{
		if str[i] == str[len(str) - 1 - i]{
			continue
		}else {
			return false
		}
	}

	return true
}