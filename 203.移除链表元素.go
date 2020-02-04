/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	prt := &ListNode{Val: 0, Next: head} // 哨兵节点
	preNode := prt
	curNode := preNode
	for {
		if curNode == nil {
			break
		}
		if curNode != nil && curNode.Val == val {
			preNode.Next = curNode.Next
		} else {
			preNode = curNode
		}
		curNode = curNode.Next
	}

	return prt.Next
}

// @lc code=end

