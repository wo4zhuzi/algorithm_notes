package main

import "log"

//快速排序
func main() {
	l := []float64{2, 3, 4, 1, 2, 223, 4, 2, 2, 3, 4}
	quickSort(l, 0, len(l)-1)
	log.Println(l)
}

func quickSort(arr []float64, low int, high int) {
	if low >= high {
		return
	}
	left, right := low, high
	pivot := arr[left]

	for left < right {
		if left < right && arr[right] >= pivot {
			right--
		}
		if left < right {
			arr[left] = arr[right]
		}

		if left < right && arr[left] < pivot {
			left++
		}

		if left < right {
			arr[right] = arr[left]
		}

		if left <= right {
			arr[left] = pivot
		}
	}
	quickSort(arr, low, right-1)
	quickSort(arr, right+1, high)
}
