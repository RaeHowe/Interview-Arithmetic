package main

import "fmt"

type T1 struct {
	Name  string
	Age   int
	Arr   [2]bool
	ptr   *int
	slice []int
	map1  map[string]string
}

func main()  {
	t1 := T1{
		Name:  "yxc",
		Age:   1,
		Arr:   [2]bool{true, false},
		ptr:   new(int),
		slice: []int{1, 2, 3},
		map1:  make(map[string]string, 0),
	}
	t2 := T1{
		Name:  "yxc",
		Age:   1,
		Arr:   [2]bool{true, false},
		ptr:   new(int),
		slice: []int{1, 2, 3},
		map1:  make(map[string]string, 0),
	}

	fmt.Println(&t1 == & t2)
}
