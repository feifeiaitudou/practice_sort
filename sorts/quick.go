package sorts

//77, 22, 33, 11, 99, 88, 100,
//时间复杂度 nlogn
func QuickSort(data []int) {
	quick(data, 0, len(data)-1)
}

/**
需要重点注意的点:
	1. 序列开始的位置不能 大于等于 序列结束的位置
	2. 如果在比较中,左边、右边不相交,则交换左右两边的值
	3. 每一次进到排序开始时,分别记录一个初始的开始(flagStart)和结束(flagEnd)标记(分别对应每次比较的start、end)
	4. 如果相交,则将当前位置与主元素位置进行替换,此时的[flagStart与相交位置(end)分别为下一次迭代开始位置(左边)、end+1和flagEnd为右边序列的位置]
	5. 先从右边(小于)扫描、然后左边(大于等于)

*/
func quick(data []int, start, end int) {
	flagStart := start
	flagEnd := end
	if start >= end {
		return //相交或者超过都不再进行
	}

	for start < end {

		for start < end && data[flagStart] < data[end] {
			end--
		}

		for start < end && data[flagStart] >= data[start] {
			start++
		}

		if start < end {
			data[start], data[end] = data[end], data[start]
		}

		//如果相交,且当前位置与主元素不相等
		if start == end && data[flagStart] != data[end] {
			data[end], data[flagStart] = data[flagStart], data[end]
		}
	}
	//左边
	quick(data, flagStart, end-1)

	//右边
	quick(data, end+1, flagEnd)
}
