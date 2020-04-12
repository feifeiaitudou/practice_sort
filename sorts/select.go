package sorts

//选择排序
//一般不用?
//时间复杂度 n^2 不稳定
//排完的,放到前面去
func SortSelect(arr []int) {

	//先找最小的? 找到了和第一个位置互换,然后在剩下的数据里面找最小的和第二个互换.....
	//记录"位置"
	length := len(arr)

	for i := 0; i < length-1; i++ {
		min := i //每次开始的最小值的位置
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[min] {
				min = j //最小值的位置更新
			}
		}

		//一轮下来的结果: 当前检索开始的位置[i]以及值arr[i] 最小值的位置[min]
		//交换两个位置的值i min
		//arr[i], arr[min] = arr[min], arr[i]
		temp := arr[i]
		arr[i] = arr[min]
		arr[min] = temp
	}
}
