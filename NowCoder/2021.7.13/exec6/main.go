package main

import "fmt"

func main()  {
	result := jumpFloor(7)

	fmt.Println(result)
}

func jumpFloor( number int ) int {
	// write code here
	if number <= 2{
		return number
	}

	return jumpFloor(number - 1) + jumpFloor(number - 2)
}