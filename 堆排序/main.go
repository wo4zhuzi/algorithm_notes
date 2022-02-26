package main

import "log"

func main() {
	l := []int{1, 23, 4, 5, 6, 7, 24, 9}
	//初始化堆
	buildHeap(l)

	for i := len(l) - 1; i >= 0; i-- {
		l[i], l[0] = l[0], l[i]
		heapify(l, i, 0)
	}

	log.Println(l)
}

func buildHeap(tree []int) {
	lastNode := len(tree) - 1
	parent := (lastNode - 1) / 2
	for i := parent; i >= 0; i-- {
		heapify(tree, len(tree), i)
	}
}

func heapify(tree []int, n int, i int) {
	if i >= n {
		return
	}
	c1 := 2*i + 1
	c2 := 2*i + 2

	max := i
	if c1 < n && tree[c1] > tree[max] {
		max = c1
	}

	if c2 < n && tree[c2] > tree[max] {
		max = c2
	}

	if max != i {
		tree[max], tree[i] = tree[i], tree[max]
		heapify(tree, n, max)
	}
}
