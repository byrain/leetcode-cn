/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p := 0
	var head *ListNode
	var tail *ListNode
	for {
		r := 0
		if l1 == nil && l2 == nil {
			if p > 0 {
				tail.Next = &ListNode{
					Val:  1,
					Next: nil,
				}
			}
			return head
		}
		if l1 != nil {
			r = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			r = r + l2.Val
			l2 = l2.Next
		}
		r = r + p
		if r >= 10 {
			p = 1
			r = r - 10
		} else {
			p = 0
		}
		if head == nil {
			head = &ListNode{
				Val:  r,
				Next: nil,
			}
			tail = head
		} else {
			tail.Next = &ListNode{
				Val:  r,
				Next: nil,
			}
			tail = tail.Next
		}
	}
}

// @lc code=end

