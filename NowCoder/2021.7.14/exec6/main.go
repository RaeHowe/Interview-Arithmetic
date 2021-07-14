package main

import (
	"fmt"
)

func main()  {
	var str = "}(])[{(}([[}])}]))})]]({{(])"

	fmt.Println(isValid(str))
}

func isValid( s string ) bool {
	// write code here
	if len(s) % 2 != 0{
		return false
	}

	var tmpMap = make(map[string]int)
	tmpMap["("] = 1
	tmpMap["{"] = 2
	tmpMap["["] = 3

	var stack = make([]string, 0)
	for i := 0; i < len(s); i++{
		var tmpStr = string(s[i])

		if _, ok := tmpMap[tmpStr]; ok{
			stack = append(stack, tmpStr)
			continue
		}

		//如果不是左侧符号，先取到栈的最后一个元素
		//先判断栈的长度，如果栈的长度为0，这时候来一个后侧的括号元素，那肯定是不符合规范的
		if len(stack) == 0{
			return false
		}
		var lastChar = stack[len(stack) - 1]

		switch tmpStr {
		case ")":
			if lastChar == "("{
				//删除栈的最后一个元素，然后继续遍历后边的元素
				stack = stack[:len(stack) - 1]
				continue
			}else {
				return false
			}
		case "]":
			if lastChar == "["{
				//删除栈的最后一个元素，然后继续遍历后边的元素
				stack = stack[:len(stack) - 1]
				continue
			}else {
				return false
			}
		case "}":
			if lastChar == "{"{
				//删除栈的最后一个元素，然后继续遍历后边的元素
				stack = stack[:len(stack) - 1] // 1,2,3
				continue
			}else {
				return false
			}
		}
	}

	if len(stack) == 0{
		return true
	}else {
		return false
	}
}