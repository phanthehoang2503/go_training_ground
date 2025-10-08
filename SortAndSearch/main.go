package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	n_size := 100000
	big_arr := make([]int, n_size)
	for i := 0; i < n_size; i++ {
		big_arr[i] = rand.IntN(n_size)
	}
	//small_arr := []int{6, 1, 4, 2, 3, 9, 7, 1, 6, 8}
	arr0 := append([]int(nil), big_arr...)
	min, max := findMinMax(arr0)
	fmt.Printf("Min number is %d and Max number is %d \n", min, max)

	arr1 := append([]int(nil), big_arr...) // copy for bubble sort
	arr2 := append([]int(nil), big_arr...) // copy for insertion sort

	start := time.Now()
	bbSort(arr1)
	// fmt.Println(bbSort(small_arr))
	fmt.Printf("Time taken with bubble sort: %.4f \n", time.Since(start).Seconds())

	start = time.Now()
	insSort(arr2)
	fmt.Printf("Time taken with insert sort: %.4f \n", time.Since(start).Seconds())

	//Search number              //cho mảng nhỏ
	target := rand.IntN(n_size) //cho mảng lớn

	//linear search
	fmt.Printf("Linear search: Number %d at the index of %d \n", target, linearSearch(big_arr, target))

	//binary search
	fmt.Printf("Binary search: Number %d at the index of %d \n", target, binarySearch(big_arr, target))
}
