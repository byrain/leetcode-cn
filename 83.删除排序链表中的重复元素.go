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
		return head
	}
	t := make(map[int]struct{})
	l := head
	t[head.Val] = struct{}{}
	for {
		if head.Next != nil {
			if _, ok := t[head.Next.Val]; !ok {
				t[head.Next.Val] = struct{}{}
				head = head.Next
			} else {
				head.Next = head.Next.Next
			}
		} else {
			break
		}
	}
	return l
}

// @lc code=end

