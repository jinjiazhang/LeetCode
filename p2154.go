package leetcode

func findFinalValue(nums []int, original int) int {
	newNums := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == original {
			original = 2 * original
			newNums = append(newNums, nums[i+1:]...)
			return findFinalValue(newNums, original)
		} else if nums[i]%original == 0 {
			newNums = append(newNums, nums[i])
		}
	}

	return original
}
