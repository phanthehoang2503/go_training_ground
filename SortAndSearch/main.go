package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func sortWrapper(name string, arr []int, sortFunc func([]int)) {
	start := time.Now()
	sortFunc(arr)
	fmt.Printf("Time taken with %s: %.4f seconds\n", name, time.Since(start).Seconds())
}

func main() {
	arraySize := 30000
	big_arr := make([]int, arraySize)
	for i := 0; i < arraySize; i++ {
		big_arr[i] = rand.IntN(arraySize)
	}
	// Sort array
	//small_arr := []int{6, 1, 4, 2, 3, 9, 7, 1, 6, 8}
	arr0 := append([]int(nil), big_arr...)
	min, max := findMinMax(arr0)
	fmt.Printf("Min number is %d and Max number is %d \n", min, max)

	arr1 := append([]int(nil), big_arr...)
	arr2 := append([]int(nil), big_arr...)
	arr3 := append([]int(nil), big_arr...)
	arr4 := append([]int(nil), big_arr...)
	arr5 := append([]int(nil), big_arr...)

	// fmt.Println(bbSort(small_arr))
	sortWrapper("Bubble sort", arr1, bbSort)

	sortWrapper("Insertion sort", arr2, insSort)

	// fmt.Println(mergeSort(small_arr))
	sortWrapper("Merge sort", arr3, mergeWrapper)

	// quickSort(small_arr, 0, len(small_arr)-1)
	// fmt.Println(small_arr)
	sortWrapper("Quick sort", arr4, quickSortWrapper)

	sortWrapper("Counting sort", arr5, countingSort)
	// =================================================================================
	//Search number
	targetIndex := rand.IntN(arraySize)
	target := big_arr[targetIndex]

	//linear search
	fmt.Printf("Linear search: Number %d at the index of %d \n", target, linearSearch(big_arr, target))

	//binary search
	fmt.Printf("Binary search: Number %d at the index of %d \n", target, binarySearch(arr4, target))
}
