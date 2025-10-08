package main

//------SORT------
//Bubble sort
func bbSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] <= arr[j+1] { // left->right "<=" giảm dần, ">=" tăng dần
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

//Insert sort
func insSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key { // arr[j] < key giảm dần, arr[j] > key tăng dần
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}
