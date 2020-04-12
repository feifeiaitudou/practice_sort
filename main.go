package main

import (
	"fmt"
	"practice_sort/sorts"
)

//https://mp.weixin.qq.com/s/IAZnN00i65Ad3BicZy5kzQ
//对排序不是很了解,学学

func main() {
	selectData := []int{8, 21, 3, 4, 5, 6}
	sorts.SortSelect(selectData)
	fmt.Println(selectData)
}
