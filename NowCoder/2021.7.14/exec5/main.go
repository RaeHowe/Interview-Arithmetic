package main

func main()  {
}

func gcd( a int ,  b int ) int {
	var min, max int

	if a < b {
		min = a
		max = b
	}else {
		min = b
		max = a
	}

	if min == 0{
		return max
	}

	return gcd(min, max % min)
}