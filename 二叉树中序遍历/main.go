package main

import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//遍历 ：左 - 根 - 右
//应用 ：可以用来做表达式树，在编译器底层实现的时候用户可以实现基本的加减乘除，比如 a*b+c。
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

	res := midOrderTraversal(root)
	log.Println(res)
}

func midOrderTraversal(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}

	res := []int{}
	midOrder(node, &res)
	return res
}

func midOrder(node *TreeNode, ans *[]int)  {
	if node == nil {
		return
	}
	midOrder(node.Left, ans)
	*ans = append(*ans, node.Val)
	midOrder(node.Right, ans)
}