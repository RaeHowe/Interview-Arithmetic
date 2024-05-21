package main

import "fmt"

func main() {
	var str = "{[]]"

	res := isValid(str)

	fmt.Println(res)
}

func isValid(s string) bool {
	var stack []string

	for _, item := range s {
		if len(stack) == 0 {
			stack = append(stack, string(item))
			continue
		}

		var tmpItem = string(item)

		//如果是左半部分括号的话就放入到栈里
		if tmpItem == "(" || tmpItem == "[" || tmpItem == "{" {
			stack = append(stack, tmpItem)
			continue
		}

		if tmpItem == "(" && (stack[len(stack)-1] == ")") {
			stack = stack[:len(stack)-1]
		} else if tmpItem == "[" && (stack[len(stack)-1] == "]") {
			stack = stack[:len(stack)-1]
		} else if tmpItem == "{" && (stack[len(stack)-1] == "}") {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, tmpItem)
		}
	}

	if len(stack) == 0 {
		return true
	}

	return false
}
