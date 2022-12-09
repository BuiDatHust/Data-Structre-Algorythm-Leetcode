package MaxSubArray

// good solution with time: O(n) space: O(n)
// use dynamic programming with create a array save max subarray with ending is index n
func MaxSubArray1(nums []int) int {
	n := len(nums)
	dp := make([]int, n) // dp[i] show max subarray ending with nums[i]
	dp[0] = nums[0]
	max := dp[0]

	for i := 1; i < n; i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
