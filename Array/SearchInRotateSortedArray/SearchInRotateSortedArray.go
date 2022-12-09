package SearchInRotateSortedArray

import "fmt"

// Binary search solution ,time: O(logn) space: O(1)
// first find rotate pivot index by binary search
// second find target index manually with notice: real index in ascending order = (index + pivot index)%n
func Search(arr []int, target int) int {
	n := len(arr)
	l := 0
	r := n - 1

	for l < r {
		mid := l + (r-l)/2
		fmt.Println(mid)
		if arr[mid] > arr[l] {
			l = mid
		} else {
			r = mid
		}
	}
	rotateIndex := r
	l = 0
	r = n - 1
	fmt.Println(rotateIndex)
	for l <= r {
		mid := l + (r-l)/2
		realMid := (mid + rotateIndex) % n
		if arr[realMid] == target {
			return realMid
		}
		if arr[realMid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
