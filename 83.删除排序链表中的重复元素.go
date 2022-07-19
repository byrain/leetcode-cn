/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHdead := &ListNode{
		Next: head,
	}
	p1 := newHdead
	p2 := newHdead.Next
	for p2 != nil && p2.Next != nil {
		if p2.Val != p2.Next.Val {
			p1 = p2
			p2 = p2.Next
		} else {
			for {
				p2 = p2.Next
				if p2 == nil || p2.Next == nil || p2.Val != p2.Next.Val {
					p1.Next = p2
					p2 = p1.Next
					break
				}
			}
		}
	}
	return newHdead.Next
}

// @lc code=end

