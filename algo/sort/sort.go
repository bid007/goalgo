package sort

// InsertionSort function for integer arrays
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i
		for ; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}
}

// MergeSort function for integer arrays
func MergeSort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}

	var (
		mid   = len(arr) / 2
		left  = MergeSort(arr[:mid])
		right = MergeSort(arr[mid:])
	)

	var (
		merged []int
		i      int
		j      int
	)

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	for ; i < len(left); i++ {
		merged = append(merged, left[i])
	}

	for ; j < len(right); j++ {
		merged = append(merged, right[j])
	}

	return merged
}
