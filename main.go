package main

import (
	"fmt"
	"psort/sorts"
)

//https://mp.weixin.qq.com/s/IAZnN00i65Ad3BicZy5kzQ
//对排序不是很了解,学学

func main() {
	selectData := []int{29, 19, 85, -1, 8888, 81, 92, 101, 9999, -1002, 8, 21, 3, 4, 5, 6, 88, 1, 28, 0, 42, 34}
	sorts.MergeSort(selectData)
	fmt.Println(selectData)
}
