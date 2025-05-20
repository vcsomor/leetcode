package prob_00206

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		list     *ListNode
		expected []int
	}{
		{list: nodeOf(nil), expected: nil},
		{list: nodeOf([]int{1, 2, 3, 4, 5}), expected: []int{5, 4, 3, 2, 1}},
		{list: nodeOf([]int{1, 2}), expected: []int{2, 1}},
		{list: nodeOf([]int{1}), expected: []int{1}},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, arrayOf(reverseList(c.list)))
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
