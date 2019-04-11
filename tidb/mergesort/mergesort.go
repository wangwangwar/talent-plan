package main

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	dest := make([]int64, len(src))
	copy(dest, src)

	if len(dest) > 1 {
		middle := len(dest) / 2
		MergeSort(dest[:middle])
		MergeSort(dest[middle:])
		merge(&dest, middle)
	}

	copy(src, dest)
}

func merge(src *[]int64, middle int) {
	r := new([]int64)
	i, j := 0, middle
	for i < middle && j < len(*src) {
		if (*src)[i] < (*src)[j] {
			*r = append(*r, (*src)[i])
			i++
		} else {
			*r = append(*r, (*src)[j])
			j++
		}
	}

	if i == middle {
		*r = append(*r, (*src)[j:]...)
	}

	if j == len(*src) {
		*r = append(*r, (*src)[i:middle]...)
	}

	*src = *r
}
