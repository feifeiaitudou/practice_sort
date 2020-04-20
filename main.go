package main

import (
	"fmt"
	"psort/ds"
)

//https://mp.weixin.qq.com/s/IAZnN00i65Ad3BicZy5kzQ
//对排序不是很了解,学学

func main() {
	selectData := []int{
		-1, 9999999999, 54, 26, 93, 17, 77, 31, 44, 55, 20, 54, 54, 26,
	}
	//20 23 17
	//17 20
	/*sorts.QuickSort(selectData)
	fmt.Println(selectData)
	*/
	//testThreadedTree()

	//创建最小堆
	ds.MinHeap(selectData, (len(selectData)-1)/2)
	fmt.Println(selectData)
	selectData = ds.AddToHeap(selectData, -22)
	fmt.Println(selectData)
	//弹出元素
	rInt, selectData := ds.PopHeapTop(selectData)
	fmt.Println("弹出元素为: ", rInt, "弹出后: ", selectData)

	//测试依次弹出最小元素(优先队列使用堆)
	for {
		rInt, selectData = ds.PopHeapTop(selectData)
		if len(selectData) == 0 {
			//已经弹出完了
			break
		}
		fmt.Println("弹出元素: ", rInt)
	}
}

func testThreadedTree() {
	/*arrData := [10]int{88, 22, 34, 1, 99, 33, 20, 38, 89, 100}*/
	arrData := [3]int{20, 1, 22}
	threadedTree := new(ds.ThreadedTree)
	for _, val := range arrData {
		threadedTreeNode := ds.NewThreadedNode(val)
		threadedTree.InitTree(threadedTreeNode)
	}
	fmt.Println("树已经构建好了,开始树的线索化(LNR)...")
	threadedTree.EnabledThreadedTreeLNR(threadedTree.Root)
}

func testTree() {
	//构造一颗二叉树,并且按照中序遍历出结果
	arrData := [10]int{88, 22, 34, 1, 99, 33, 20, 38, 89, 100}
	tree := new(ds.Tree)
	for i := 0; i < 10; i++ {
		//创建一个节点
		node := ds.NewNode()
		node.Data = arrData[i]
		tree.AddNode(node)
	}

	//中序遍历结果为
	bl := make([]*ds.TreeNode, 0)
	bl = tree.GetNodesByLNR(tree.Root, bl)
	for _, val := range bl {
		fmt.Println(val.Data)
	}
}
