import "math"

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	head, l1Value := min(l1, l2)
	if head == nil {
		return head
	}
	if l1Value {
		l1 = l1.Next
	} else {
		l2 = l2.Next
	}

	t := head
	for l1 != nil || l2 != nil {
		t.Next, l1Value = min(l1, l2)
		if l1Value {
			l1 = l1.Next
		} else {
			l2 = l2.Next
		}
		t = t.Next
	}
	return head
}

func min(l1 *ListNode, l2 *ListNode) (*ListNode, bool) {
	if l1 == nil && l2 != nil {
		return l2, false
	}
	if l1 != nil && l2 == nil {
		return l1, true
	}
	if l1.Val >= l2.Val {
		return l2, false
	}
	return l1, true
}

// @lc code=end

