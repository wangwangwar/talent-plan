package main

import (
	"math"
	"sync"
)

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

	dst := make([]int64, len(src))

	parallelMerge(src, 0, middle-1, middle, len(src)-1, dst, 0)
	copy(src, dst)
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


func parallelMerge(src []int64, p1 int, r1 int, p2 int, r2 int, dst []int64, p3 int) {
	length1 := r1 - p1 + 1
	length2 := r2 - p2 + 1
	if length1 < length2 {
		p1, p2 = p2, p1
		r1, r2 = r2, r1
		length1, length2 = length2, length1
	}
	if length1 == 0 {
		return
	}
	q1 := (p1 + r1) / 2
	q2 := binarySearch(src[q1], src, p2, r2)
	q3 := p3 + (q1 - p1) + (q2 - p2)
	dst[q3] = src[q1]

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		parallelMerge(src, p1, q1-1, p2, q2-1, dst, p3)
	}()
	parallelMerge(src, q1+1, r1, q2, r2, dst, q3+1)
	wg.Wait()
}

func binarySearch(value int64, src []int64, left int, right int) int {
	low := left
	high := int(math.Max(float64(left), float64(right+1)))
	for low < high {
		mid := (low + high) / 2
		if value <= src[mid] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return high
}
