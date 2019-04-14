package main

import "sync"

var SLICE_MIN_SIZE = 100

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	if len(src) <= 1 {
		return
	}

	if len(src) < SLICE_MIN_SIZE {
		insertSort(src)
		return
	}

	middle := len(src) / 2
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		MergeSort(src[:middle])
	}()
	MergeSort(src[middle:])
	wg.Wait()

	merge(src, middle)
}

func insertSort(src []int64) {
	var n = len(src)
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if src[j-1] > src[j] {
				src[j-1], src[j] = src[j], src[j-1]
			}
		}
	}
}

func merge(src []int64, middle int) {
	r := make([]int64, 0, len(src))
	i, j := 0, middle
	for i < middle && j < len(src) {
		if src[i] < src[j] {
			r = append(r, src[i])
			i++
		} else {
			r = append(r, src[j])
			j++
		}
	}

	if i == middle {
		r = append(r, src[j:]...)
	}

	if j == len(src) {
		r = append(r, src[i:middle]...)
	}

	copy(src, r)
}
