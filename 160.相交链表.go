/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1 := headA
	p2 := headB
	if p1 == nil || p2 == nil {
		return nil
	}
	p1FisrtTime := true
	p2FisrtTime := true
	for {
		if (p1 == nil && !p1FisrtTime) || (p2 == nil && !p2FisrtTime) {
			return nil
		}
		if p1 == nil && p1FisrtTime {
			p1 = headB
			p1FisrtTime = false
		}
		if p2 == nil && p2FisrtTime {
			p2 = headA
			p2FisrtTime = false
		}

		if p1 == p2 {
			return p1
		}
		p1 = p1.Next
		p2 = p2.Next

	}
}

// @lc code=end

