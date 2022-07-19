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
	// if head == nil {
	// 	return head
	// }
	// prt := &ListNode{Val: 0, Next: head} // 哨兵节点
	// preNode := prt
	// curNode := preNode
	// for {
	// 	if curNode == nil {
	// 		break
	// 	}
	// 	if curNode != nil && curNode.Val == val {
	// 		preNode.Next = curNode.Next
	// 	} else {
	// 		preNode = curNode
	// 	}
	// 	curNode = curNode.Next
	// }

	// return prt.Next

	if head == nil {
		return nil
	}
	newHdead := &ListNode{
		Next: head,
	}
	p1 := newHdead
	p2 := newHdead.Next
	for p2 != nil {
		if p2.Val != val {
			p1 = p1.Next
			p2 = p2.Next
		} else {
			p1.Next = p2.Next
			p2 = p1.Next
		}
	}
	return newHdead.Next
}

// @lc code=end

