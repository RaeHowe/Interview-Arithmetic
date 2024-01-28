package main

//用两个栈来实现一个队列
/*
	大致思路: 用go语言来实现。用两个数组来模拟两个栈。一个只负责往里面push元素的数组inStack，另一个只负责往外pop元素的outStack。当取值
	的时候，如果outStack里面没有值的话，就需要先从inStack里面取值放置到outStack里面去；如果有值的话就直接取值就好。再说一下从inStack往
	outStack里面移动值的过程,由于是栈的结构，所以我们需要从outStack里面从后往前取值，我们取得数组的最后一个元素，然后放置到inStack里面，
	然后再接着取outStack的最后一个值，再追加到inStack数组的后面；再说push值，push值的操作就比较简单，直接在inStack数组后面去追加元素即可
*/
func main()  {
	
}

//定义数据结构
type MyQueue struct {
	inStack []int //负责入栈
 	outStack []int //负责出栈
}

//初始化函数
func Constructor() MyQueue {
	return MyQueue{}
}

//入栈操作，很简单，直接把元素放置在末尾即可
func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

//出栈操作就需要判断是否为空的情况
func (q *MyQueue) Pop() int {
	if len(q.outStack) == 0{
		q.in2out()
	}

	//放置完了元素之后就有值了
	x := q.outStack[len(q.outStack) - 1]
	q.outStack = q.outStack[:len(q.outStack) - 1] //减去outStack的最后一个元素
	return x
}

//定义一个inStack往outStack移动元素的方法
func (q *MyQueue) in2out() {
	for len(q.inStack) > 0{ //全部移动走，所以判断inStack数组的长度是否为0最后
		q.outStack = append(q.outStack, q.inStack[len(q.inStack) - 1]) //把inStack最后一个元素移动到outStack
		q.inStack = q.inStack[:len(q.inStack) - 1]//然后确确实实地减去inStack数组最后一个元素
	}
}
