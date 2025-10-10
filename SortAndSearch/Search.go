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
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (right + left) / 2
		if arr[mid] == t {
			return mid
		} else if arr[mid] < t {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func findMinMax(arr []int) (int, int) {
	sorted_arr := mergeSort(arr)
	return sorted_arr[0], sorted_arr[len(sorted_arr)-1]
}

func binarySearch1(arr []int, left, right int, t int) int {
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == t {
			return mid
		} else if arr[mid] < t {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

//===========================================================

// Interpolation search
func interpolationSearch(arr []int, t int) int {
	low, high := 0, len(arr)-1

	for (low <= high) && t >= arr[low] && t <= arr[high] { //target must be in range of the array
		position := low + int(float64(t-arr[low])*float64(high-low)/float64(arr[high]-arr[low]))

		if arr[position] == t {
			return position
		}
		if arr[position] < t { //position below or front the target
			return position + 1 // below
		} else {
			return position - 1 //front
		}
	}

	return -1
}

//===========================================================

// Exponential search
func exponentialSearch(arr []int, t int) int {
	if arr[0] == t {
		return 0
	}

	i := 1
	for i < len(arr) && arr[i] <= t {
		i *= 2
	}

	left := i / 2
	right := min(i, len(arr)-1)

	return binarySearch1(arr, left, right, t)
}
