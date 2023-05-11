package array

func CheckSortedBetter(arr []int, si int) bool {
	if si == len(arr)-1 {
		// only one element
		return true
	}
	if arr[si] > arr[si+1] {
		return false
	}
	isSmallArraySorted := CheckSortedBetter(arr, si+1)
	return isSmallArraySorted
}
