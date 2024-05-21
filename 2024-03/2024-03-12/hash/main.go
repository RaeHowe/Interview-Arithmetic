package main

import "fmt"

func main() {
	//var array = []int{1, 4, 1, 6}
	fmt.Println(2 ^ 3) // 0010 0011 => 0001

	// 0 ^ 1 // 0000 0001 => 0001
	// 1 ^ 4 // 0001 0100 => 0101
	// 5 ^ 1 // 0101 0001 => 0100
	// 4 ^ 6 // 0100 0110 => 0010 = 2

	//xor = 2
	// n := 1
	/*
			function FindNumsAppearOnce( array ) {
		    // write code here
		    var res = 0;  //对整个数组求异或的结果
		    for(var i = 0; i < array.length; i++){
		        res ^= array[i];
		    }
		    var compare = 1;
		    while((compare & res) == 0){ //判断异或结果的二进制第一位是否为1，为1则直接跳过该循环
		        compare <<= 1;         //为0则继续往后找，一直到找到为1的二进制位，该行代码也相当于compare *=2
		    }
		    var a = 0;
		    var b = 0;
		    for(var i = 0; i < array.length; i++){  //遍历数组，开始判断数字们的compare位是否为1
		        if((compare & array[i]) == 0){     //如果该数字二进制的第某位为0，则分到数组一
		            a ^= array[i];                 //对数组一进行异或，得到a
		        }else{                             //如果该数字二进制的第某位为1，则分到数组二
		            b ^= array[i];                 //对数组二进行异或，得到b
		        }
		    }
		    return a < b ? [a,b] : [b,a];
		}
	*/

}

//func FindNumsAppearOnce(nums []int) []int {
//	return nil
//}
