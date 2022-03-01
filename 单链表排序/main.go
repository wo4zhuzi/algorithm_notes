package main

import (
	"fmt"
)

func main() {
	head := &ListNode{Val: 15}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 10}
	head.Next.Next.Next = &ListNode{Val: 9}
	head.Next.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 20}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 18}

	printListNode(head)
	head = sortList(head)
	printListNode(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}


func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := getMid(head)
	midNex := mid.Next
	//中心点拆分链表
	mid.Next = nil

	left := sortList(head)
	right := sortList(midNex)

	return merge(left, right)
}

//合并两个有序链表
func merge(left *ListNode, right *ListNode) *ListNode {
	preAns := &ListNode{}
	preAnsEnd := preAns

	for left != nil && right != nil {
		if left.Val < right.Val {
			preAnsEnd.Next = left
			left = left.Next
		} else {
			preAnsEnd.Next = right
			right = right.Next
		}
		preAnsEnd = preAnsEnd.Next
	}

	if left == nil {
		preAnsEnd.Next = right
	} else {
		preAnsEnd.Next = left
	}

	return preAns.Next
}

func getMid(head *ListNode) *ListNode {
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func printListNode(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%v \t", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}
