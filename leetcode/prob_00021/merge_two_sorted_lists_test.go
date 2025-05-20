package prob_00021

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		list1, list2 *ListNode
		expected     []int
	}{
		{list1: nodeOf([]int{1, 2, 4}), list2: nodeOf([]int{1, 3, 4}), expected: []int{1, 1, 2, 3, 4, 4}},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, arrayOf(mergeTwoLists(c.list1, c.list2)))
	}
}

func nodeOf(ints []int) *ListNode {
	var head *ListNode

	var last *ListNode
	for _, v := range ints {
		ptr := &ListNode{Val: v}
		if head == nil {
			head = ptr
		}
		if last != nil {
			last.Next = ptr
		}
		last = ptr
	}
	return head
}

func arrayOf(n *ListNode) []int {
	var res []int
	for n != nil {
		res = append(res, n.Val)
		n = n.Next
	}
	return res
}
