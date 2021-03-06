package main

/*
	题目:10亿个数中如何高效地找到最大的一个数以及最大的第 K 个数
	1.先说如何查询到最大的一个数的方法：
	可以采用分治的思想，把这10亿个数字进行分割，分割成多份，数据量更小的数据文件。每个数据文件的大小要保证不能超过计算机内存大小。然后每个文件内部进行
	快速排序，所得到的最大值取出来。然后把每个文件取出来的最大值再次进行快速排序，就能得到最大的结果了。这里面需要注意内存的使用，不要超出范围。

	2.再说最大的第K个数
	这个问题就是典型的topK问题。先从10亿个数里面取出来K个数字，建立一个小顶推，这个堆的大小要保证不能超过计算机的内存规定大小。然后从剩下的数字中
	依次取数去和这个小顶堆的堆顶元素进行比较，如果大于这个数字的话，就替换掉这个数字，然后调整堆；如果这个数字小于小顶堆的堆顶元素的话，就直接略过这个数字。
	经过最后处理之后，小顶堆的堆顶元素，就是最大的第K个元素。遍历的时间复杂度为O(n),调整小顶堆的时间复杂度为O(logn)，所以最后的时间复杂度为O(nlogn)
*/