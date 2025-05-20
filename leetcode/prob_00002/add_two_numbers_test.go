package prob_00002

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		list1, list2 *ListNode
		expected     []int
	}{
		{list1: nodeOf([]int{2, 4, 3}), list2: nodeOf([]int{5, 6, 4}), expected: []int{7, 0, 8}},
		{list1: nodeOf([]int{0}), list2: nodeOf([]int{0}), expected: []int{0}},
		{list1: nodeOf([]int{9, 9, 9, 9, 9, 9, 9}), list2: nodeOf([]int{9, 9, 9, 9}), expected: []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, arrayOf(addTwoNumbers(c.list1, c.list2)))
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
