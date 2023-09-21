package search

// need to have a sorted array of nums
func binarySearch(needle int, nums []int) (int, bool) {
	min := 0
	max := len(nums) - 1

	for min <= max {
		mid := (min + max) / 2
		if nums[mid] < needle {
			min = mid + 1
		} else if nums[mid] > needle {
			max = mid - 1
		} else {
			return mid, true
		}
	}

	return 0, false
}
