package sorts

//计数排序,只能排非负整数(其他的场景应当如何处理?)
//计数排序,适用于最大值与最小值差别不大的情况
//用元素作为数组的下标,标记该元素出现的多少次 如 arr[i] = n,表明元素i出现了n次,最后将数组汇总起来
//稳定
//求出临时数组的长度
//时间复杂度n+k, 非比较排序
func CountSort(data []int) {
	max := getMax(data)

	//声明一个临时的切片,这里最好是查找出序列的最大值与最小值,然后计算两个的差值[避免空间浪费]
	arr := make([]int, max+1)

	for _, val := range data {
		arr[val] += 1
	}

	//用于重新改写data索引下的值
	//合并数组
	index := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			continue
		}

		for j := 0; j < arr[i]; j++ {
			data[index] = i
			index++
		}

	}
}

func getMax(data []int) int {
	max := data[0]
	for i := 0; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
	}
	return max
}
