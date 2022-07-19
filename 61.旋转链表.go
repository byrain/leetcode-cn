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
		return nil
	}
	tail := head
	count := 1
	for tail.Next != nil {
		tail = tail.Next
		count++
	}
	k = k % count
	if k == 0 {
		return head
	}
	p := head
	for i := 0; i < count-k-1; i++ {
		p = p.Next
	}
	newHead := p.Next
	p.Next = nil
	tail.Next = head

	return newHead
}

// @lc code=end

