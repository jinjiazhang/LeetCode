package easy

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maxSum := sum
	for j := 1; j <= len(nums)-k; j++ {
		sum = sum - nums[j-1] + nums[j+k-1]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}
