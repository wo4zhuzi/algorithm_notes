package main

import (
	"fmt"
	"math/rand"
)

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

	candidate := xNode.next[0]

	//如果xNode的next节点score和插入score相同则直接改value值
	if candidate != nil && candidate.score == score {
		candidate.value = value
		return candidate
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

	candidate := xNode.next[0]

	//如果xNode的next节点score和插入score则直接返回
	if candidate != nil && candidate.score == score {
		return candidate, true
	}
	return nil, false
}

func (sl *SkipList) Range(from float64, to float64) ([]*Node, error) {
	if from > to {
		return nil, fmt.Errorf("from can not gt to")
	}
	var firstNode *Node
	for i := from; i <= to; i++ {
		firstNode, _ = sl.Search(i)
		if firstNode != nil {
			break
		}
	}

	var listNode []*Node
	for i := firstNode; i.score <= to; i = i.next[0] {
		listNode = append(listNode, i)
	}

	return listNode, nil
}

func (sl *SkipList) Delete(score float64) *Node {
	update := make([]*Node, maxLevel)
	xNode := sl.head

	for i := sl.level - 1; i >= 0; i-- {
		for xNode.next[i] != nil && xNode.next[i].score < score {
			xNode = xNode.next[i]
		}
		update[i] = xNode
	}

	candidate := xNode.next[0]

	if candidate == nil || candidate.score != score {
		return nil
	}

	if candidate != nil && candidate.score == score {
		for i := 0; i < sl.level-1; i++ {
			//自下而上如果下一节点不等于删除节点，那么直接退出，会少几次循环
			if update[i].next[i] != candidate {
				break
			}
			update[i].next[i] = candidate.next[i]
		}
		sl.length--
	}

	return candidate
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
