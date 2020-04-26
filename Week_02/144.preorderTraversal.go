package Week_02

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	nodes := make([]int, 0)
	pre(root, &nodes)
	return nodes
}
func pre(root *TreeNode, nodes *[]int) {
	if root != nil {
		*nodes = append(*nodes, root.Val)
		pre(root.Left, nodes)
		pre(root.Right, nodes)
	}
}
