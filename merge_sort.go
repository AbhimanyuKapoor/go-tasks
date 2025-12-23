package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mergeSort1(arr []int, low, high int) {
	if low >= high {
		return
	}

	mid := (low + high) / 2

	mergeSort1(arr, low, mid)
	mergeSort1(arr, mid+1, high)
	merge(arr, low, mid, high)
}

// With goroutines
func mergeSort2(arr []int, low, high int) {
	if low >= high {
		return
	}

	mid := (low + high) / 2

	// Threshold
	if high-low < 1000 {
		mergeSort2(arr, low, mid)
		mergeSort2(arr, mid+1, high)
	} else {
		done := make(chan bool)

		go func() {
			mergeSort2(arr, low, mid)
			done <- true
		}()

		go func() {
			mergeSort2(arr, mid+1, high)
			done <- true
		}()

		// Wait
		<-done
		<-done
	}

	merge(arr, low, mid, high)
}

func merge(arr []int, low, mid, high int) {
	temp := make([]int, 0, high-low+1)

	left, right := low, mid+1

	for left <= mid && right <= high {
		if arr[left] <= arr[right] {
			temp = append(temp, arr[left])
			left++
		} else {
			temp = append(temp, arr[right])
			right++
		}
	}

	for left <= mid {
		temp = append(temp, arr[left])
		left++
	}

	for right <= high {
		temp = append(temp, arr[right])
		right++
	}

	// Sorted temp back to arr
	for i, val := range temp {
		arr[low+i] = val
	}
}

func main() {
	size := 100000
	arrOriginal := make([]int, size)

	for i := range size {
		arrOriginal[i] = rand.Intn(size)
	}

	arr1 := make([]int, size)
	copy(arr1, arrOriginal)

	arr2 := make([]int, size)
	copy(arr2, arrOriginal)

	// Multithreaded
	fmt.Println("\nStarting Multithreaded Merge Sort...")
	start := time.Now()

	mergeSort2(arr1, 0, len(arr1)-1)

	duration := time.Since(start)
	fmt.Printf("Sorted %d elements in: %v\n", size, duration)

	// Sequential
	fmt.Println("\nStarting Normal Merge Sort...")
	start = time.Now()

	mergeSort1(arr2, 0, len(arr2)-1)

	duration = time.Since(start)
	fmt.Printf("Sorted %d elements in: %v\n", size, duration)
}
