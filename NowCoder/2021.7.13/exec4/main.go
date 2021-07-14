package main

import "fmt"

func main()  {
	var number = 10

	result := sqrt(number)

	fmt.Println(result)
}

func sqrt( x int ) int {
	if x == 0 || x == 1{
		return x
	}

	var s = x / 2

	for s * s > x{
		s = s / 2
	}

	for s * s <= x{
		s++
	}

	return s - 1
}
