package ds

import "fmt"

type ThreadedTree struct {
	Root *ThreadedTreeNode
}

//尝试构建线索二叉树LNR 中序
//中序 ---> 左 根 右
func (tree *ThreadedTree) InitTree(node *ThreadedTreeNode) {
	if tree.Root == nil {
		//没有,初始化一个
		tree.Root = node
		fmt.Println("创建根节点")
		return
	}

	//还是先进行创建
	currentNode := tree.Root
	for {

		if node.Data == currentNode.Data {
			fmt.Println("节点已经存在,不再创建")
			return
		}
		if node.Data < currentNode.Data {
			//左
			if currentNode.LNode == nil {
				currentNode.LNode = node
				fmt.Println("创建左节点")
				return
			}
			currentNode = currentNode.LNode
		}

		if node.Data > currentNode.Data {
			//右
			if currentNode.RNode == nil {
				currentNode.RNode = node
				fmt.Println("创建右节点")
				return
			}
			currentNode = currentNode.RNode
		}
	}
}

//中序遍历树
func (tree *ThreadedTree) EnabledThreadedTreeLNR(node *ThreadedTreeNode) {
	if node != nil {
		tree.EnabledThreadedTreeLNR(node.LNode)
		fmt.Println(node.Data)
		/*if node.LNode == nil {
			//没有孩子,指向前驱
			node.LTag = "1" //前驱标记
			node.RNode =    //前驱怎么找???
		}

		if node.RNode != nil {
			//没有右孩子,指向后继
			node.RTag = "1" //指向后继

		}*/
		//tempNode := node
		//	fmt.Println(node.Data, "  ", tempNode.RNode.Data)
		if node.LNode == nil {
			fmt.Println("没有左节点,增加到前驱", node.Data)
		}

		if node.Data == 88 {
			fmt.Println("右节点...", node.RNode.Data, )
		}
		tree.EnabledThreadedTreeLNR(node.RNode)
	}
}
