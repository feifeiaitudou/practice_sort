package ds

import (
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
				fmt.Println("添加到左节点")
				break
			}
			//否则,继续下一个节点
			NodeCurrent = NodeCurrent.L
		}

		if node.Data > NodeCurrent.Data {
			if NodeCurrent.R == nil {
				NodeCurrent.R = node
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
