package ContainerWithMostWater

// bad algorythm O(n*n)
func MaxArea1(height []int) int {
	max := 0
	n := len(height)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			var mul int
			if height[j] < height[i] {
				mul = height[j] * (j - i)
			} else {
				mul = height[i] * (j - i)
			}
			if mul > max {
				max = mul
			}
		}
	}
	return max
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// optimize algorythm
// demonstrate: use two pointer to track area, if l<r must go to the right to find a larger area
// because if we go to the left we cannot find a new larger area
// time: O(n) space: O(1)
func MaxArea2(height []int) int {
	maxArea := 0
	l := 0
	r := len(height) - 1
	for {
		if l < r {
			maxArea = Max(maxArea, (r-l)*Min(height[l], height[r]))
			if height[l] < height[r] {
				l += 1
			} else {
				r -= 1
			}
		}
		break
	}
	return maxArea
}
