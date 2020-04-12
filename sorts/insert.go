package sorts

//插入排序,从位置1开始向前比较,如果小,则换位置,如果大,不动
//稳定
// n^2,最好n
//样本小基本有序时效率高
//knuth序列
// h = 1
// h = 3*h +1
func InertSort(data []int) {
	for i := 1; i < len(data); i++ {
		//和前面的进行比较
		for j := i; j > 0; j-- {
			if data[j] < data[j-1] {
				//如果小,则交换位置
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
}
