package sorts

//基数排序
//先以个位数打下排序 再以十位数大小 再以百位数大小....
//没'位'上的范围为0-9,准备10个桶,把相同的'数值'(位置)的数放到桶里,然后在把桶里面的数据排序拿出来合并,按照位数排序就完成一趟(此时不需要对桶中的数据进行排序)
//时间复杂度 O(n)

//太神奇了

var buckets [10][]int

func RadioSort(data []int) {
	max := getMax(data)
	base := 1
	for max >= base {
		for _, val := range data {
			mod := getPosData(val, base)
			buckets[mod] = append(buckets[mod], val)
		}
		popList(data, &buckets)
		base *= 10
	}
}

func popList(data []int, all *[10][]int) {
	index := 0

	for i := 0; i < 10; i++ {
		if len(all[i]) == 0 {
			continue
		}
		for _, val := range all[i] {
			data[index] = val
			index++
		}
		all[i] = []int{} //完事后,将桶里面的数据清空
	}
}

//取出数据d的base位上的数 个位、十位、百位
//暂时没有管负数
func getPosData(d int, base int) int {
	//先直接判定是否在范围内
	if d < base {
		return 0
	}
	return d / base % 10
}
