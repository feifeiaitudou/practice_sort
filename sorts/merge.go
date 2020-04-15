package sorts

//稳定, nlog2n 空间复杂度n

//关键在于正确找到序列中点
//递归切割序列(不改变原序列的索引),然后排序和合并(在递归完成以后进行)

func MergeSort(data []int) {
	divide(data, 0, len(data)-1)
}

/**
start: 序列起始位置
bound: 序列结束位置
*/
func divide(data []int, start, bound int) {
	mid := start + (bound-start)/2

	if start == bound {
		return
	}
	//分左边
	divide(data, start, mid)

	//分右边
	divide(data, mid+1, bound)

	mergeSort(data, start, mid+1, bound)
}

/**
left:左序列开始位置
right:右序列开始位置
bound: 整个子序列结束位置
*/
func mergeSort(data []int, left, right, bound int) {
	L := left
	R := right
	M := right - 1
	if left == right {
		return
	}
	temp := make([]int, 0)
	//如果都有,直到一边已经完成
	for L <= M && R <= bound {
		if data[L] >= data[R] {
			//如果左边比右边大,则把右边取出来,并且移动右序列的索引位置
			temp = append(temp, data[R])
			R++
		} else {
			//右边比左边大,则把左边的取出来,并且移动左序列的索引位置
			temp = append(temp, data[L])
			L++
		}
	}

	//还剩左序列
	for L <= M {
		temp = append(temp, data[L])
		L++
	}

	//还剩右序列
	for R <= bound {
		temp = append(temp, data[R])
		R++
	}

	//然后将temp中的没个元素放置到原序列原来的位置中
	for _, val := range temp {
		//此处直接将temp里面的'值'取出来,不要直接使用temp[index]的方式('这里为引用')
		data[left] = val //更新原始序列的位置上的值为'已经排好序'的temp
		left++
	}
}
