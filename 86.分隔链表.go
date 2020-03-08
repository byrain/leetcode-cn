/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	vLess := &ListNode{}
	nodeLess := vLess
	vMore := &ListNode{}
	nodeMore := vMore

	node := head
	for node != nil {
		if node.Val < x {
			nodeLess.Next = node
			nodeLess = nodeLess.Next
		} else {
			nodeMore.Next = node
			nodeMore = nodeMore.Next
		}
		node = node.Next
	}
	nodeMore.Next = nil // 防止成环。比如链表最后一个元素小于x时
	nodeLess.Next = vMore.Next
	return vLess.Next
}

// @lc code=end

