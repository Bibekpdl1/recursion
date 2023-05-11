package array

func CheckSorted(arr []int) bool {
	if len(arr) <= 1 {
		return true
	}
	if arr[0] > arr[1] {
		return false
	}
	smallArray := arr[1:]
	checkInSmallArray := CheckSorted(smallArray)
	return checkInSmallArray
}
