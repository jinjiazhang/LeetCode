package easy

import "testing"

func Test_2154(t *testing.T) {
	nums := []int{5, 3, 6, 1, 12}
	original := 3
	result := findFinalValue(nums, original)
	t.Log(result)
}
