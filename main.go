package main

import (
	"fmt"
	"psort/sorts"
)

//https://mp.weixin.qq.com/s/IAZnN00i65Ad3BicZy5kzQ
//对排序不是很了解,学学

func main() {
	selectData := []int{8, 21, 3, 4, 5, 6, 88, 1, 28, 0, 42, 34, 29, 19, 85, -1, 81, 92, 101, 8888, 9999, -1002}
	sorts.ShellSort(selectData)
	fmt.Println(selectData)
}
