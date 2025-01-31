package leetcode

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/linked-list-cycle/?envType=study-plan-v2&envId=top-interview-150
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	checked := make(map[*ListNode]struct{})

	current := head
	for {

		if current == nil {
			break
		}

		if _, ok := checked[current]; ok {
			return true
		}

		checked[current] = struct{}{}
		current = current.Next
	}

	return false
}
