package main

import "log"

func main() {
	a := []int{2, 1, 3, 6, 9, 4}
	a = bubbleSort(a)
	log.Println(a)
}

func bubbleSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	length := len(arr)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
