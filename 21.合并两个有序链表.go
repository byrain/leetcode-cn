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
	l := &ListNode{}
	temp := l
	val := math.MaxInt64
	for {
		if l1 != nil || l2 != nil {
			l1Val := math.MaxInt64
			l2Val := math.MaxInt64
			if l1 != nil {
				l1Val = l1.Val
			}
			if l2 != nil {
				l2Val = l2.Val
			}

			if l1Val <= l2Val {
				val = l1Val
				l1 = l1.Next
			} else {
				val = l2Val
				l2 = l2.Next
			}

			tempNode := &ListNode{
				Val: val,
			}
			temp.Next = tempNode
			temp = tempNode
		}
		if l1 == nil && l2 == nil {
			break
		}
	}
	return l.Next
}

// @lc code=end

