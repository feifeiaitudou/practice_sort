package sorts

//冒泡排序 稳定
//时间复杂度 n^2,最好时间复杂度n (怎么实现)
//排完的,放到后面去
func BubbleSort(data []int) {
	for i := len(data) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if data[j] > data[j+1] {
				data[j+1], data[j] = data[j], data[j+1]
			}
		}
	}
}
