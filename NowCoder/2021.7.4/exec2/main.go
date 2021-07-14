package main

import (
	"fmt"
)

func main()  {
	var str = "abcd"

	result := solve(str)

	fmt.Println(result)
}

func solve(str string) string {
	var length = len(str)

	//go里面字符串是不可变类型，所以得考虑转换一下类型
	var tmp []byte
	bytes := []byte(str)

	for i := length - 1; i >= 0; i-- {
		tmp = append(tmp, bytes[i])
	}

	return string(tmp)
}
