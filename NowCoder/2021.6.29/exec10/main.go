package main

import "fmt"

func main()  {
	var array1 = []int{2,4,6,0,0,0}
	var array2 = []int{1,3,5}

	merge(array1, 3, array2,  3)

	fmt.Println(array1)
}

func merge( A []int ,  m int, B []int, n int )  {
	var indexA = m - 1
	var indexB = n - 1
	var endIndex = m + n - 1

	for indexA >= 0 && indexB >= 0{
		if A[indexA] >= B[indexB]{
			A[endIndex] = A[indexA]
			indexA--
			endIndex--
		}else {
			A[endIndex] = B[indexB]
			indexB--
			endIndex--
		}
	}

	for indexB >= 0{
		A[indexB] = B[indexB]
		indexB--
	}

	return
}
