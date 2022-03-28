package main

import (
	"testing"
)

func TestNewSkipList(t *testing.T) {
	sl := NewSkipList()
	sl.Insert(float64(50), "one")
	sl.Insert(float64(60), "two")
	sl.Insert(float64(30), "three")
	sl.Insert(15.5, "four")

	find1, _ := sl.Search(30)
	find2, _ := sl.Search(90)

	t.Log("find1 :", find1)
	t.Log("find2 :", find2)
}
