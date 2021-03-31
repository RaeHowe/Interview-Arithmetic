package main

import (
	"fmt"
	"strings"
)

//移掉K位数字
func main()  {
	var str = "12345264"
	result := removeKdigits(str, 4)

	fmt.Println(result)
}

//移除k个元素
func removeKdigits(num string, k int) string {
	var tmpByteArr []byte

	for index := range num{
		var tmpData = num[index] //从左往右获取数字信息
		for k > 0 && len(tmpByteArr) > 0 && tmpData < tmpByteArr[len(tmpByteArr) - 1]{ //判断要删除的数字是否已经被删除完毕
			// 判断获取到的元素，是否小于已经入栈的最后一个元素值
			tmpByteArr = tmpByteArr[:len(tmpByteArr) - 1]
			k--
		}

		tmpByteArr = append(tmpByteArr, tmpData)
	}

	tmpByteArr = tmpByteArr[:len(tmpByteArr) - k]

	result := strings.TrimLeft(string(tmpByteArr), "0")
	if result == ""{
		result = "0"
	}

	return result
}