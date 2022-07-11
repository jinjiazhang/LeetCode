package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	node := head
	for node.Next != nil {
		if node.Next.Val == node.Val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}

	return head
}
