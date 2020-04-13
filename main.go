package main

import (
	"fmt"
	"psort/sorts"
)

//https://mp.weixin.qq.com/s/IAZnN00i65Ad3BicZy5kzQ
//对排序不是很了解,学学

func main() {
	selectData := []int{8, 21, 3, 4, 5, 6}
	sorts.InertSort(selectData)
	fmt.Println(selectData)
}
