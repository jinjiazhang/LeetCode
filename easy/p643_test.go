package easy

import "testing"

func Test_643(t *testing.T) {
	nums := []int{0, 1, 1, 3, 3}
	k := 4
	result := findMaxAverage(nums, k)
	t.Log(result)
}
