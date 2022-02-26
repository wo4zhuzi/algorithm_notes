package main

import "log"

func main() {
	l := []int{1,2,11,3,4,33,2,1,4,43,2}
	l = sSort(l)
	log.Println(l)
}

//选择排序
func sSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
