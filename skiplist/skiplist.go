package main

import "math/rand"

const (
	maxLevel = 48
	prob     = 1 << 30
)

//跳表节点
type Node struct {
	score float64
	value interface{}
	next  []*Node
}

type SkipList struct {
	head   *Node
	length int
	level  int
}

func NewNode(score float64, value interface{}, level int) *Node {
	return &Node{
		score: score,
		value: value,
		next:  make([]*Node, level),
	}
}

func NewSkipList() *SkipList {
	return &SkipList{
		head: &Node{
			next: make([]*Node, maxLevel),
		},
	}
}

func (sl *SkipList) Insert(score float64, value interface{}) *Node {
	update := make([]*Node, maxLevel)
	xNode := sl.head

	for i := sl.level - 1; i >= 0; i-- {
		for xNode.next[i] != nil && xNode.next[i].score < score {
			xNode = xNode.next[i]
		}
		update[i] = xNode
	}

	//如果xNode的next节点score和插入score相同则直接改value值
	if xNode.next[0] != nil && xNode.next[0].score == score {
		xNode.next[0].value = value
		return xNode.next[0]
	}

	level := sl.randLevel()

	if level > sl.level {
		level = sl.level + 1
		update[sl.level] = sl.head
		sl.level = level
	}

	n := NewNode(score, value, level)

	for i := 0; i < level; i++ {
		n.next[i] = update[i].next[i]
		update[i].next[i] = n
	}
	sl.length++
	return n
}

func (sl *SkipList) Search(score float64) (*Node, bool) {
	xNode := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for xNode.next[i] != nil && xNode.next[i].score < score {
			xNode = xNode.next[i]
		}
	}

	//如果xNode的next节点score和插入score则直接返回
	if xNode.next[0] != nil && xNode.next[0].score == score {
		return xNode.next[0], true
	}
	return nil, false
}

func (sl *SkipList) randLevel() int {
	i := 1
	for ; i < maxLevel; i++ {
		if rand.Int31() < prob {
			break
		}
	}
	return i
}
