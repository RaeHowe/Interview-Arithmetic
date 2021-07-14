package main

func main()  {

}

func maxWater( arr []int ) int64 { //[3, 1, 2, 5, 2, 4]
	// write code here
	var leftIndex = 0
	var rightIndex = len(arr) - 1

	var leftMax = 0
	var rightMax = 0

	var result int64

	for leftIndex < rightIndex{
		leftMax = max(leftMax, arr[leftIndex])
		rightMax = max(rightMax, arr[rightIndex])

		if leftMax < rightMax{
			result += int64(leftMax - arr[leftIndex])
			leftIndex++
		}else {
			result += int64(rightMax - arr[rightIndex])
			rightIndex--
		}
	}

	return result
}

func max(a, b int) int {
	if a >= b{
		return a
	}

	return b
}
