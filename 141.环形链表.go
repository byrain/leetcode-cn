/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return false
	}
	r := false
	t1 := head
	t2 := head.Next
	for {
		if t2 == nil || t2.Next == nil {
			return r
		}

		if t1 == t2 {
			return true
		} else {
			t1 = t1.Next
			t2 = t2.Next.Next
		}
	}
}

// @lc code=end

