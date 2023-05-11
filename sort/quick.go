package sort

func QuickSort(arr []int, si int, ei int) []int {
	if si >= ei {
		return arr
	}
	// finding the pivot index so that all elements to the left of the pivot are smaller and all elements to the right of the pivot are larger
	pi := partition(arr, si, ei)
	// sorting the left half
	QuickSort(arr, si, pi-1)
	// sorting the right half
	QuickSort(arr, pi+1, ei)
	return arr
}

func partition(arr []int, si int, ei int) int {
	pivot := arr[si]
	i := si
	j := ei - 1
	for {
		if i >= j {
			// if i and j cross each other
			break
		}
		if arr[i] < pivot {
			// if element at i is smaller than pivot, increment i
			i++
		} else if arr[j] > pivot {
			// if element at j is larger than pivot, decrement j
			j--
		} else {
			// if element at i is larger than pivot and element at j is smaller than pivot, swap them
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	return i
}
