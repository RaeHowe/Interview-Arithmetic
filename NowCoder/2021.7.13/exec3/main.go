package main

import (
	"fmt"
	"sort"
)

func main()  {
	var array = []int{-10, 0, 10, 20, -10, -40}

	result := ThreeSum(array)

	fmt.Println(result)
}

func ThreeSum(tmpSlice []int) [][]int {
	var length = len(tmpSlice)
	sort.Ints(tmpSlice)

	var result = make([][]int, 0)

	for i := 0; i < length - 2; i++{
		if tmpSlice[i] > 0{
			break
		}

		if i > 0 && tmpSlice[i] == tmpSlice[i - 1]{
			continue
		}

		var low, high = i + 1, length - 1

		for low < high{
			var tmpSum = tmpSlice[i] + tmpSlice[low] + tmpSlice[high]

			if tmpSum == 0{
				result = append(result, []int{tmpSlice[i], tmpSlice[low], tmpSlice[high]})

				//头部元素去重
				for low < high && tmpSlice[low] == tmpSlice[low + 1]{
					low++
				}

				low++
				high--
			}else if tmpSum < 0{
				low++
			}else if tmpSum > 0{
				high--
			}
		}
	}

	return result
}








func threeSum( num []int ) [][]int {
	// write code here
	//三元组》》子序列
	//1.三元组 非降序排列》》将num升序排列
	//2.数求和问题》》双指针
	res:=make([][]int,0)
	//排序
	sort.Ints(num)
	n:=len(num)
	for i:=0; i<n-2; i++{
		if num[i]>0 {
			break//如果当前数字大于0,则三数之和一定大于0(这属于异常情况，用于判断排序完的数组的第一个元素，如果大于0，则后面的元素都大于0，不可能等于0了)
		}
		if i > 0 && num[i] == num[i-1] {
			continue//去重
		}
		//双指针求和
		low,high := i+1, n-1//一开始就是最大值和最小值
		for low < high{
			sum:= num[i] + num[low] + num[high] //高水位线的值+低水位线的值+当前值
			if sum==0{
				//写入结果
				res=append(res,[]int{num[i],num[low],num[high]})
				// 头部去重（如果后面一个数跟当前的数字相等，则代表有重复的结果生成，跳过
				for low < high && num[low] == num[low+1]{
					low++
				}
				 //// 尾部去重（如果前面一个数跟当前的数字相等，则代表有重复的结果生成，跳过）
				 //for low<high&&num[high]==num[high-1]{
					// high--
				 //}

				 //如果满足条件的话，就把高地水位线都往中间靠拢
				low++
				high--
			}else if sum>0{
				//太大的话，就把高水位线的索引往前推
				high--
			}else{
				//太小的话，就把低水位线的索引往后推
				low++
			}
		}
	}
	return res
}
