package main

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	if len(src) > 1 {
		middle := len(src) / 2
		MergeSort(src[:middle])
		MergeSort(src[middle:])
		merge(src, middle)
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
