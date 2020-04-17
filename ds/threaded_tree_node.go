package ds

/*
	前驱、后继的定义:二叉树按照某种顺序遍历后形成的链表,其中对于某个结点而言,其前面的元素称为前驱、后面的元素称为后继
	如果结点的左子树为空,则指向其前驱
	如果结点的右子树为空,则指向其后继
*/
//线索二叉树
type ThreadedTreeNode struct {
	Data  int    //data可以定义为interface
	LTag  string //左指针标记,如果为0,则左指针保存前驱;如果为1,则左指针保存左节点
	LNode *ThreadedTreeNode
	RTag  string //右指针标记,如果为0,则右指针保存后继;如果为1,则右指针保存右节点
	RNode *ThreadedTreeNode
}

func NewThreadedNode(n int) *ThreadedTreeNode {
	return &ThreadedTreeNode{Data: n, LTag: "0", RTag: "0",}
}

func (node *ThreadedTreeNode) getLeftNode() *ThreadedTreeNode {
	return node.LNode
}

func (node *ThreadedTreeNode) getRightNode() *ThreadedTreeNode {
	return node.RNode
}
