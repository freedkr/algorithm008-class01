package Week_02

func levelOrder(root *Node) [][]int {
	res := [][]int{}
	recursive(root, 0, &res)
	return res
}

func recursive(root *Node, level int, res *[][]int) {
	if root == nil {
		return
	}

	for _, v := range root.Children {
		recursive(v, level+1, res)
	}

	for len(*res) <= level {
		*res = append(*res, []int{})
	}

	(*res)[level] = append((*res)[level], root.Val)
}
