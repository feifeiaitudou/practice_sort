package sorts

//基数排序
//先以个位数打下排序 再以十位数大小 再以百位数大小....
//没'位'上的范围为0-9,准备10个桶,把相同的'数值'(位置)的数放到桶里,然后在把桶里面的数据排序拿出来合并,按照位数排序就完成一趟(此时不需要对桶中的数据进行排序)
//时间复杂度 O(n)

//太神奇了

//负数如何基数排序? TypeA
//浮点数如何基数排序?

var buckets [10][]int

//基数排序 包含负数?
//将非负数与负数分成两个序列,分别排好后再合并到一起?
func RadioSortWithTypeA(data []int) {
	pn := make([]int, 0) //非负数
	nn := make([]int, 0) //负数

	for _, val := range data {
		if val >= 0 {
			pn = append(pn, val)
			continue
		}

		if val < 0 {
			nn = append(nn, -val) //将负数转换为正数
		}
	}

	//先比较正数
	max := getMax(pn)
	base := 1
	for max >= base {
		for _, val := range pn {
			mod := getPosData(val, base)
			buckets[mod] = append(buckets[mod], val)
		}
		popList(pn, &buckets)
		base *= 10
	}

	//再比较负数
	max = getMax(nn)
	base = 1
	for max >= base {
		for _, val := range nn {
			mod := getPosData(val, base)
			buckets[mod] = append(buckets[mod], val)
		}
		popList(nn, &buckets)
		base *= 10
	}

	//将nn倒序拷贝到data上
	index := 0 //data的索引
	if len(nn) > 0 {
		for i := len(nn) - 1; i >= 0; i-- {
			data[index] = -nn[i] //将正数转换为负数
			index++
		}
	}

	//再将pn拷贝到data上
	for _, val := range pn {
		data[index] = val
		index++
	}
}

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
