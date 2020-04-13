package sorts

import "fmt"

//归并排序,适用于处理较大规模的数据排序
//时间复杂度nlog2n, 空间复杂度n 稳定
//将一个大的序列分隔为两个序列,然后将两个序列进行排序,最后将两个序列进行合并
//需要额外的空间保存两个子序列
//一直分,直到序列中还剩下一个的时候(剩下的那个认为为有序)
//分治法  分:将一个序列分为两个子序列; 治: 将两个序列进行排序; 合: 将两个有序的子序列合并为一个(迭代或者递归过程)
func MergeSort(data []int) {
	//先尝试分一下
	left := make([]int, 0)
	right := make([]int, 0)

	for i := 0; i < len(data)-1; i++ {
		//添加的时候使用插入排序??
		if i%2 == 0 {
			left = append(left, data[i])
		}

		if i%2 != 0 {
			right = append(right, data[i])
		}
	}

	//左边的插入排序
	InertSort(left)

	//右边的插入排序
	InertSort(right)

	fmt.Println("left is: ", left)
	fmt.Println("right is: ", right)

}
