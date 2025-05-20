package prob_00021

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	ptr1, ptr2 := list1, list2
	var resultHead, resultPtr, resultLast *ListNode

	for ptr1 != nil || ptr2 != nil {
		if ptr1 == nil {
			resultPtr = &ListNode{Val: ptr2.Val}
			ptr2 = ptr2.Next
		} else if ptr2 == nil {
			resultPtr = &ListNode{Val: ptr1.Val}
			ptr1 = ptr1.Next
		} else {
			if ptr1.Val < ptr2.Val {
				resultPtr = &ListNode{Val: ptr1.Val}
				ptr1 = ptr1.Next
			} else {
				resultPtr = &ListNode{Val: ptr2.Val}
				ptr2 = ptr2.Next
			}
		}

		if resultLast != nil {
			resultLast.Next = resultPtr
		}
		resultLast = resultPtr

		if resultHead == nil {
			resultHead = resultPtr
		}
	}

	return resultHead
}
