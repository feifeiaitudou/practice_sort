package main

import (
	"container/list"
	"fmt"
	"psort/ds"
)

//https://mp.weixin.qq.com/s/IAZnN00i65Ad3BicZy5kzQ
//对排序不是很了解,学学

func main() {
	/*selectData := []int{
		122, -1, 991, 90, 99, 54, 26, 93, 17, 77, 31, 44, 55, 20, 54, 54, 26, -199, 2, 300, 27,
	}*/

	selectData := []int{
		88, 22, 34, 1, 99, 33, 20, 38, 89, 100,
	}

	//尝试生成一棵树
	tree := new(ds.Tree)
	for _, val := range selectData {
		node := ds.NewNode()
		node.Data = val

		tree.AddNode(node)
	}

	tree.OrderLoopLNR()

	/*sorts.RadioSortWithTypeA(selectData)
	fmt.Println(selectData)*/
	//stackTest()
}

func stackTest() {
	queue := list.New()
	queue.PushBack(11)
	queue.PushBack(22)
	queue.PushBack(33)

	for queue.Len() > 0 {
		fmt.Println(queue.Remove(queue.Back()))
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
