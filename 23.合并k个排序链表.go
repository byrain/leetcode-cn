import "sort"

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	values := []int{}
	for _, list := range lists {
		node := list
		for node != nil {
			values = append(values, node.Val)
			node = node.Next
		}
	}
	sort.Ints(values)

	if len(values) == 0 {
		return nil
	}
	head := &ListNode{
		Val: values[0],
	}
	tmp := head
	for i := 1; i < len(values); i++ {
		tmp.Next = &ListNode{
			Val: values[i],
		}
		tmp = tmp.Next
	}
	tmp.Next = nil
	return head
}

// @lc code=end

