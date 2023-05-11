package sort

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	// fining the middle point
	mid := len(arr) / 2
	// sorting the first half
	left := MergeSort(arr[:mid])
	// sorting the second half
	right := MergeSort(arr[mid:])
	// merging the two sorted halves
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	l, r, i := 0, 0, 0
	for l < len(left) || r < len(right) {
		// loop until larger array is fully traversed
		if l < len(left) && r < len(right) {
			// if both arrays have elements
			if left[l] < right[r] {
				result[i] = left[l]
				l++
			} else {
				result[i] = right[r]
				r++
			}
		} else if l < len(left) {
			// if only left array has elements
			result[i] = left[l]
			l++
		} else {
			// if only right array has elements
			result[i] = right[r]
			r++
		}
		i++
	}
	// for {
	// 	if l >= len(left) || r >= len(right) {
	// 		break
	// 	}
	// 	if left[l] < right[r] {
	// 		result[i] = left[l]
	// 		l++
	// 	} else {
	// 		result[i] = right[r]
	// 		r++
	// 	}
	// 	i++
	// }
	// if len(left) > len(right) {
	// 	copy(result[i:], left[l:])
	// } else {
	// 	copy(result[i:], right[r:])
	// }

	return result
}
