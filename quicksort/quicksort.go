package quicksort

func QuickSort(arr []int, lo int, hi int) {

	if lo >= hi {
		return
	}

	pivotIdx := partition(arr, lo, hi)

	QuickSort(arr, lo, pivotIdx-1)
	QuickSort(arr, pivotIdx+1, hi)

}

func partition(arr []int, lo int, hi int) int {

	pivot := arr[hi]
	idx := lo - 1

	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			idx++
			tmp := arr[i]
			arr[i] = arr[idx]
			arr[idx] = tmp
		}
	}

	idx++
	arr[hi] = arr[idx]
	arr[idx] = pivot

	return idx

}
