package Week_02

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	nodes := make([]int, 0)
	inorder(root, &nodes)
	return nodes
}

func inorder(root *TreeNode, nodes *[]int) {
	if root != nil {
		inorder(root.Left, nodes)
		*nodes = append(*nodes, root.Val)
		inorder(root.Right, nodes)
	}
}
