package sorts

//重新自己写插入排序
//n^2,稳定
//可以将数据看做两个部分, 牌堆(左边); 手上的牌
//从牌堆拿牌,依次从手上的牌(从右到左进行比较,如果符合条件,则插入)
func InertSort(data []int) {
	for i := 0; i < len(data); i++ {
		//挨个从牌堆拿牌

		for j := i; j > 0; j-- {
			//手上的牌都是有序的
			//刚刚拿的牌,依次同手上的牌(从右向左,每一张进行比较)
			if data[j] < data[j-1] {
				//data[j] 表示为最后一次插入的位置的元素
				//如果拿上来的牌比当前手上的牌小,插入,至此,刚刚的牌已完成
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
}
