package piscine

type TreeNode struct {
	Left, Right *TreeNode
	Val         int
}

func BTreeInsertData(root *TreeNode, data int) *TreeNode {

	if root == nil {
		return &TreeNode{Val: data}
	}
	if data < root.Val {
		root.Left = BTreeInsertData(root.Left, data)
	} else {
		root.Right = BTreeInsertData(root.Right, data)
	}
	return root
}
