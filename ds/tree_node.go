package ds

type TreeNode struct {
	Data int //假设data为整数,此处最好是放置interface{}
	L    *TreeNode
	R    *TreeNode
	P    *TreeNode //增加一个双亲结点
}

func NewNode() *TreeNode {
	return new(TreeNode)
}

/*func (node *TreeNode) getLeftNode() (t1 *TreeNode, ba []*TreeNode) {
	if ba == nil {
		fmt.Println("初始化一个切片")
		ba = make([]*TreeNode, 0)
	}
	if node.L == nil {
		fmt.Println("找到末尾了...")
		return nil, ba
	}
	fmt.Println("执行递归调用了", node.L.Data)
	ba = append(ba, node.L)
	return node.L.getLeftNode()
}*/

func (node *TreeNode) getLeftNode(treeNode *TreeNode) {

}

func (node *TreeNode) getRightNode() *TreeNode {
	return node.R
}
