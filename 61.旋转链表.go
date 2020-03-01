/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	n := 0
	t := head
	for t != nil {
		n++
		t = t.Next
	}
	k = k % n
	if k == 0 {
		return head
	}
	t1 := head
	t2 := head
	for i := 0; i < k; i++ {
		t2 = t2.Next
	}
	for t2.Next != nil {
		t1 = t1.Next
		t2 = t2.Next
	}
	t2.Next = head
	r := t1.Next
	t1.Next = nil
	return r
}

// @lc code=end

