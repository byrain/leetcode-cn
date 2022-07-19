/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	newNode := &ListNode{
		Next: head,
	}
	pre := newNode

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return newNode.Next
			}
		}
		next := tail.Next
		head, tail = reverseK(head, tail)
		pre.Next = head
		head = next
		pre = tail
		tail.Next = next
	}
	return newNode.Next
}

func reverseK(head, tail *ListNode) (*ListNode, *ListNode) {
	p := head
	pre := tail.Next
	for pre != tail {
		node := p.Next
		p.Next = pre
		pre = p
		p = node
	}
	return tail, head
}

// @lc code=end

