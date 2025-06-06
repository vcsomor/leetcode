package prob_00206

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
