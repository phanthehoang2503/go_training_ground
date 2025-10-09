package main

//------SORT------
//Bubble sort
func bbSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] <= arr[j+1] { // left->right "<=" giảm dần, ">=" tăng dần
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

//Insert sort
func insSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key { // arr[j] < key giảm dần, arr[j] > key tăng dần
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

//Merge sort
func merge(left, right []int) []int {
	arraySize := len(left) + len(right)
	mergedArray := make([]int, 0, arraySize)

	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			mergedArray = append(mergedArray, left[i])
			i++
		} else {
			mergedArray = append(mergedArray, right[j])
			j++
		}
	}

	// lấy những giá trị còn sót lại ở 1 trong 2 bên trái hoặc phải cho vào mảng
	if i < len(left) {
		mergedArray = append(mergedArray, left[i:]...)
	}
	if j < len(right) {
		mergedArray = append(mergedArray, right[j:]...)
	}

	return mergedArray
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func mergeWrapper(arr []int) {
	sorted := mergeSort(arr)
	copy(arr, sorted)
}

// Quick sort
func partition(arr []int, low, high int) int {
	pivotIndex := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] <= pivotIndex {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

func quickSortWrapper(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

// Counting sort
func countingSort(arr []int) {
	arrSize := len(arr)

	_, maxValue := findMinMax(arr)
	count := make([]int, maxValue+1)

	// đếm số lần xuất hiện ở mỗi giá trị
	for _, val := range arr {
		count[val]++
	}

	// tính cộng dồn
	for i := 1; i <= maxValue; i++ {
		count[i] += count[i-1]
	}

	// output
	output := make([]int, len(arr))

	for i := arrSize - 1; i >= 0; i-- {
		val := arr[i]
		output[count[val]-1] = val
		count[val]--
	}

	copy(arr, output)
}

//Radix sort
