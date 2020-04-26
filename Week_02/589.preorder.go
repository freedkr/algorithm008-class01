package Week_02

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	nodes := make([]int, 0)
	pre(root, &nodes)
	return nodes
}

func pre(root *Node, nodes *[]int) {
	if root != nil {
		*nodes = append(*nodes, root.Val)
		for i := 0; i < len(root.Children); i++ {
			pre(root.Children[i], nodes)
		}
	}
}
