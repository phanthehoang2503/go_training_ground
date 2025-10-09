package main

func linearSearch(arr []int, t int) int {
	for i, value := range arr {
		if value == t {
			return i
		}
	}
	return -1
}

func binarySearch(arr []int, t int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		m := (r + l) / 2
		if arr[m] == t {
			return m
		} else if arr[m] < t {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}

func findMinMax(arr []int) (int, int) {
	sorted_arr := mergeSort(arr)
	return sorted_arr[0], sorted_arr[len(sorted_arr)-1]
}
