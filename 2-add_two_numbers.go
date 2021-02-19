package leetcode_go


// Definition for singly-linked list.
type ListNode struct {
     Val int
     Next *ListNode
}

func max(l1, l2 int) int {
	if l1 >= l2 {
		return l1
	}

	return l2
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		tmp int
		l *ListNode
	)

	zero := &ListNode{Val: 0}

	for {
		l3 := &ListNode{}

		if l1 == nil && l2 == nil {
			break
		}
		if l1 == nil {
			l1 = zero
		}
		if l2 == nil {
			l2 = zero
		}

		sum := l1.Val + l2.Val + tmp
		tmp = sum / 10
		value := sum % 10

		l3.Val = value
		l3.Next =&ListNode{Val: tmp}
		l1 = l1.Next
		l2 = l2.Next
		if l == nil {
			l = l3
			continue
		}
		l.Next = l3
	}

	return l
}
