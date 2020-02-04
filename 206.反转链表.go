/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	nextTemp := head
	for {
		if curr != nil {
			nextTemp = curr.Next
			curr.Next = pre
			pre = curr
			curr = nextTemp
		} else {
			return pre
		}
	}
}

// @lc code=end

