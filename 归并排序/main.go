package main

func main() {
	l := []int{9, 4, 5, 2, 3, 11, 32, 67, 45, 22}
	mergeSort(l, 0, len(l)-1)
}

func mergeSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	mid := (right - left) / 2
	mergeSort(arr, left, mid)
	mergeSort(arr, mid+1, right)
	merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int) {
	var temArr = []int{}
	var s1, s2 = left, mid + 1
	for s1 <= mid && s2 <= right {
		if arr[s1] < arr[s2] {
			temArr = append(temArr, arr[s1])
			s1++
		} else {
			temArr = append(temArr, arr[s2])
			s2++
		}
	}

	if s1 > mid {
		temArr = append(temArr, arr[s2:right+1]...)
	}

	if s2 > mid {
		temArr = append(temArr, arr[s1:mid+1]...)
	}

	for k, v := range temArr {
		arr[left+k] = v
	}
}
