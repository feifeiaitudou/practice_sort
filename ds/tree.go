package ds

import (
	"container/list"
	"fmt"
)

var LNodes []*TreeNode

type Tree struct {
	Root *TreeNode //树的根节点
}

func (tree *Tree) AddNode(node *TreeNode) bool {
	//检查根节点是否存在
	if tree.Root == nil {
		fmt.Println("添加根节点")
		tree.Root = node
		return true
	}

	//如果根节点存在,则开始比较其子节点与传进来的节点元素的大小
	//都是从根开始
	//一直循环,看看应该在哪一个节点上

	NodeCurrent := tree.Root //默认一个当前的比较节点
	for {
		if node.Data == NodeCurrent.Data {
			//相同的,不允许存在
			fmt.Println("相同大小的节点,不添加!")
			return false
		}
		if node.Data < NodeCurrent.Data {
			//左边? 先检查左节点是否存在
			if NodeCurrent.L == nil {
				NodeCurrent.L = node
				node.P = NodeCurrent
				fmt.Println("添加到左节点")
				break
			}
			//否则,继续下一个节点
			NodeCurrent = NodeCurrent.L
		}

		if node.Data > NodeCurrent.Data {
			if NodeCurrent.R == nil {
				NodeCurrent.R = node
				node.P = NodeCurrent
				fmt.Println("添加到右节点")
				break
			}
			NodeCurrent = NodeCurrent.R
		}

	}
	return true
}

//中序遍历树节点,返回的数据从小到大
func (tree *Tree) GetNodesByLNR(node *TreeNode, bl []*TreeNode) []*TreeNode {
	if node != nil {
		bl = tree.GetNodesByLNR(node.L, bl)
		bl = append(bl, node)
		bl = tree.GetNodesByLNR(node.R, bl)
	}
	return bl
}

//按照循环的方式遍历树的节点,依靠栈来实现 中序
//递归的时候,默认会有地方在保存调用的结果
//直接循环的时候没有地方保存每次的结果,所以,此处使用[栈]来模拟递归中的栈
//golang中的栈使用list实现?
func (tree *Tree) OrderLoopLNR() {

	//先创建一个栈
	stack := list.New()

	if tree.Root == nil {
		fmt.Println("空树")
		return
	}

	node := tree.Root
	for node != nil || stack.Len() > 0 {
		if node != nil {
			stack.PushBack(node)
			node = node.L
		} else if stack.Len() > 0 {
			node = stack.Remove(stack.Back()).(*TreeNode)
			if node.P != nil {
				fmt.Println(node.Data, "父节点:", node.P.Data)
			} else {
				fmt.Println(node.Data, "没有父节点...")
			}

			node = node.R
		}
	}
}
