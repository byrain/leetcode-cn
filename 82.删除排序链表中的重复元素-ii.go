/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
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
	root := &ListNode{
		Next: head,
	}
	n1 := root
	n2 := root.Next
	for n2 != nil && n2.Next != nil {
		if n2.Val == n2.Next.Val {
			for {
				n2 = n2.Next
				if n2 == nil || n2.Next == nil || n2.Val != n2.Next.Val {
					n1.Next = n2.Next
					n2 = n1.Next
					break
				}
			}
		} else {
			n1 = n1.Next
			n2 = n2.Next
		}
	}
	return root.Next
}

// @lc code=end

