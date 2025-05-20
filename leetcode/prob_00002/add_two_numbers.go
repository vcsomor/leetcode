package prob_00002

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, last *ListNode
	var v1, v2 int

	carryOver := 0
	for l1 != nil || l2 != nil || carryOver != 0 {
		l1, v1 = next(l1)
		l2, v2 = next(l2)

		sum := v1 + v2 + carryOver
		curr := &ListNode{Val: sum % 10}
		carryOver = sum / 10

		if head == nil {
			head = curr
		}
		if last != nil {
			last.Next = curr
		}
		last = curr
	}

	return head
}

func next(p *ListNode) (*ListNode, int) {
	if p == nil {
		return nil, 0
	}
	return p.Next, p.Val
}
