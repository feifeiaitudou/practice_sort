package sorts

import (
	"psort/ds"
)

//堆排序
//nlogn,不稳定排序 是一种选择排序
//按照升序排序,创建大堆; 按照降序排序,创建小堆
func HeapSort(data []int) {
	length := len(data)
	heap(data, length)

}

//每一次将堆首部的元素放置到最后[交换],然后重新调整堆
//这里为降序(小堆)(0, n-1; 1, n-2...)
func heap(data []int, length int) {
	if length <= 1 {
		return //没有数据了
	}
	//先'堆化'
	ds.MinHeap(data, length/2-1, length)
	//堆化完成后,交换位置
	data[0], data[length-1] = data[length-1], data[0]
	length = length - 1
	heap(data, length)
}
