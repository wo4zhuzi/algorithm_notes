package main

import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//遍历 ：根 - 左 - 右
//应用 ：实现目录结构的显示
func main() {
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
	res := preOrderTraversal(root)
	log.Println(res)
}

func preOrderTraversal(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	res := append([]int{node.Val}, preOrderTraversal(node.Left)...)
	res = append(res, preOrderTraversal(node.Right)...)
	return res
}
