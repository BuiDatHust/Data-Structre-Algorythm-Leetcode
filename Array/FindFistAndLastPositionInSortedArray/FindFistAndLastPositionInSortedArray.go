package FindFistAndLastPositionInSortedArray

// good solution  with binary search, time: O(logn), space: O(1)
// idea is when found target by binary search then manually in/decrease lo and hi to find start and end of target
func SearchRange1(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) == 0 {
		return result
	}
	lo := 0
	hi := len(nums) - 1

	for nums[lo] < nums[hi] {
		mid := lo + (hi-lo)/2
		if nums[mid] < target {
			lo = mid + 1
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			if nums[lo] < target {
				lo++
			} else {
				hi--
			}
		}
	}

	if nums[lo] == target && nums[lo] == nums[hi] {
		result[0] = lo
		result[1] = hi
	}
	return result
}
