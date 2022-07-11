package leetcode

import "testing"

func Test_83(t *testing.T) {
	nums := []int{1, 1, 2}
	var head *ListNode
	for i := len(nums) - 1; i >= 0; i-- {
		node := &ListNode{
			Val:  nums[i],
			Next: head,
		}
		head = node
	}

	head = deleteDuplicates(head)
	node := head
	result := make([]int, 0)
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}

	t.Log(result)
}
