package leet

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	tip := &ListNode{}
	answer := tip

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			answer.Next = list1
			list1 = list1.Next
		} else {
			answer.Next = list2
			list2 = list2.Next
		}
		answer = answer.Next
	}
	if list1 == nil {
		answer.Next = list2
	} else {
		answer.Next = list1
	}

	return tip.Next
}
