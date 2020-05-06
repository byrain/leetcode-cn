/*
 * @lc app=leetcode.cn id=876 lang=golang
 *
 * [876] 链表的中间结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	t := head
	t1 := head
	for {
		if t1 == nil || t1.Next == nil {
			return t
		}
		t = t.Next
		t1 = t1.Next.Next

	}
	return t
}

// @lc code=end

