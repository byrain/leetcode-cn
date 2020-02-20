/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p1 := head
	p2 := head
	for i := 0; i < n; i++ {
		p2 = p2.Next
	}
	if p2 == nil {
		return p1.Next
	}
	if p2.Next == nil {
		p1.Next = p1.Next.Next
		return p1
	}
	p2 = p2.Next
	for p2 != nil {
		p2 = p2.Next
		p1 = p1.Next
	}
	p1.Next = p1.Next.Next
	return head
}

// @lc code=end

