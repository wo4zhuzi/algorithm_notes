package main

import (
	"testing"
)

func TestNewSkipList(t *testing.T) {
	sl := NewSkipList()
	sl.Insert(float64(50), "50")
	sl.Insert(float64(60), "60")
	sl.Insert(float64(30), "30")
	sl.Insert(float64(70), "70")
	sl.Insert(15.5, "15.5")

	find1, _ := sl.Search(30)
	find2, _ := sl.Search(60)
	sl.Delete(60)
	find3, _ := sl.Search(60)
	find4, _ := sl.Search(50)
	find5, _ := sl.Search(15.5)

	t.Log("find1 :", find1)
	t.Log("find2 :", find2)
	t.Log("find3 :", find3)
	t.Log("find4 :", find4)
	t.Log("find5 :", find5)


	list, _ := sl.Range(15.5, 60)

	t.Log("range 15.5 ~ 60")
	for _, v := range list{
		t.Log(v.score, v.value)
	}
}
