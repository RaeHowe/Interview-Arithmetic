package main

import "fmt"

func main()  {
	var num = 5

	result := fib(num)

	fmt.Println(result)
}

func fib(num int) int {
	if num == 0{
		return 0
	}
	if num == 1 || num == 2{
		return num
	}

	return fib(num - 1) + fib(num - 2)
}
