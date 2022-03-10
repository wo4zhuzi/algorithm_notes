package main

import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//遍历 ：左 - 右 - 根
//应用 ：以用来实现计算目录内的文件占用的数据大小
func main()  {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	res := postOrderTraversal(root)
	log.Println(res)
}

func postOrderTraversal(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}

	res := []int{}
	postOrder(node, &res)
	return res
}

func postOrder(node *TreeNode, ans *[]int)  {
	if node == nil {
		return
	}

	postOrder(node.Left, ans)
	postOrder(node.Right, ans)
	*ans = append(*ans, node.Val)
}
