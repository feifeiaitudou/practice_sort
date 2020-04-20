package ds

//按照'调整'的方式构建最小堆
//lastIndex 最后一个分支节点
/**
data 待排序序列
lastIndex 最后一个分支节点
length 待堆化的序列长度(左右结点位置最大为length-1)
*/
func MinHeap(data []int, lastIndex int, length int) {

	if lastIndex < 0 {
		return //如果到根节点,退出
	}
	leftIndex := 2*lastIndex + 1
	rightIndex := 2*lastIndex + 2

	if leftIndex <= length-1 && data[leftIndex] > data[lastIndex] {
		data[leftIndex], data[lastIndex] = data[lastIndex], data[leftIndex]
		//再次检查当前的左子树是否满足堆属性
		MinHeap(data, leftIndex, length)
	}

	if rightIndex <= length-1 && data[rightIndex] > data[lastIndex] {
		data[rightIndex], data[lastIndex] = data[lastIndex], data[rightIndex]
		//再次检查当前右子树是否满足堆属性
		MinHeap(data, rightIndex, length)
	}

	lastIndex -= 1
	MinHeap(data, lastIndex, length)
}

/*func MinHeap(data []int, lastIndex int) {

	if lastIndex < 0 {
		return //如果到根节点,退出
	}
	leftIndex := 2*lastIndex + 1
	rightIndex := 2*lastIndex + 2
	if len(data)-1 >= leftIndex && data[leftIndex] < data[lastIndex] {
		data[leftIndex], data[lastIndex] = data[lastIndex], data[leftIndex]
		//再次检查当前的左子树是否满足堆属性
		MinHeap(data, leftIndex)
	}

	if len(data)-1 >= rightIndex && data[rightIndex] < data[lastIndex] {
		data[rightIndex], data[lastIndex] = data[lastIndex], data[rightIndex]
		//再次检查当前右子树是否满足堆属性
		MinHeap(data, rightIndex)
	}

	lastIndex -= 1
	MinHeap(data, lastIndex)
}
*/
//将元素添加到已经调整好的堆的最后面,然后再调整
/*func AddToHeap(data []int, ele int) []int {
	data = append(data, ele)
	MinHeap(data, len(data)/2-1)
	return data
}

//弹出第一个元素,此时将第一个元素拿出来,再将最后一个元素设置到第一个元素,然后再次调整
func PopHeapTop(data []int) (rInt int, rData []int) {
	rInt = data[0]
	rLast := data[len(data)-1]
	data[0] = rLast

	//重新切片
	rData = data[0 : len(data)-1]
	MinHeap(rData, len(rData)/2-1)
	return
}*/
