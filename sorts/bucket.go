package sorts

//桶排序
//将数据分成n个区间,每个区间对应一个桶,每个桶进行排序(归并或者快速),然后n个桶进行合并
// n+k n+k 稳定排序 非原地排序
//计算每个桶的区间范围(查询出序列的最大最小值, 最后一个桶包含最大序列)
//有哪些常见的设置桶的策略?
func BucketSort(data []int) {

	min, max := getMinMax(data)

	avg := (max-min)/10 + 1
	arrBuc := [10]int{}

	for i := 0; i < 10; i++ {
		arrBuc[i] = min + i*avg
	}

	//直接准备10个桶
	buckets := [10][]int{}

	for _, val := range data {
		for j := len(arrBuc) - 1; j >= 0; j-- {
			if val >= arrBuc[j] {
				buckets[j] = append(buckets[j], val)
				break
			}
		}
	}

	for i := 0; i < len(buckets); i++ {
		sortBuckets(buckets[i])
	}

	//将所有切片上的数据合并到原序列中
	index := 0
	for i := 0; i < len(buckets); i++ {
		if len(buckets[i]) == 0 {
			continue
		}

		for _, val := range buckets[i] {
			data[index] = val
			index++
		}
	}
}

//几个桶分别排序,然后再合并
func sortBuckets(bs []int) {
	if len(bs) <= 1 {
		return //如果只有一个元素或者空,不排
	}

	//使用快速排序
	//sortQuick(bs, 0, len(bs)-1)

	//使用归并排序
	sortMerge(bs, 0, len(bs)-1)
}

//使用归并排序
//先归(划分位置), 再并(排序合并)
func sortMerge(data []int, start, end int) {
	if start == end {
		//位置重合
		return
	}

	mid := start + (end-start)/2

	//在此分位置
	sortMerge(data, start, mid)
	sortMerge(data, mid+1, end)

	sortMergeSort(data, start, mid+1, end)

	//排序
}

//需要临时空间
func sortMergeSort(data []int, left, right, bound int) {
	pos := left
	mid := right - 1
	temp := make([]int, 0)
	//如果都有,则分别比较
	for left <= mid && right <= bound {
		if data[left] >= data[right] {
			temp = append(temp, data[right])
			right++
		} else {
			temp = append(temp, data[left])
			left++
		}
	}

	//如果还剩左边
	for left <= mid {
		temp = append(temp, data[left])
		left++
	}

	if right <= bound {
		temp = append(temp, data[right])
		right++
	}

	//将temp中的数据复制到原序列上
	for _, val := range temp {
		data[pos] = val
		pos++
	}

}

//使用快速 排序
func sortQuick(data []int, left, right int) {
	flagStart := left
	flagEnd := right

	if left >= right {
		return
	}
	for left < right {

		//先扫描右边
		for left < right && data[flagStart] < data[right] {
			right--
		}

		//再扫描左边,主元素比左边的小
		for left < right && data[flagStart] >= data[left] {
			left++
		}

		if left < right {
			//当前两个数据交换
			data[left], data[right] = data[right], data[left]
		}

		if left == right {
			//如果相等,当前位置与主元素交换
			data[flagStart], data[right] = data[right], data[flagStart]
		}

	}

	//左边
	sortQuick(data, flagStart, right-1)

	//右边
	sortQuick(data, right+1, flagEnd)
}

func getMinMax(data []int) (min, max int) {
	max = data[0]
	min = data[0]

	for i := 0; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}

		if data[i] < min {
			min = data[i]
		}
	}
	return
}
