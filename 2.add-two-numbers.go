/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
* Definition for singly-linked list.
 type ListNode struct {
     Val int
     Next *ListNode
 }
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	l := &ListNode{}
	lTemp := l
	for {
		if l1 == nil && l2 == nil {
			if carry > 0 {
				sumNode := &ListNode{
					Val:  carry,
					Next: nil,
				}

				lTemp.Next = sumNode
			}
			return l.Next
		}
		l1Val := 0
		l2Val := 0
		if l1 != nil {
			l1Val = l1.Val
		}
		if l2 != nil {
			l2Val = l2.Val
		}

		sumNode := &ListNode{
			Val:  (l1Val + l2Val + carry) % 10,
			Next: nil,
		}

		carry = (l1Val + l2Val + carry) / 10

		lTemp.Next = sumNode
		lTemp = sumNode
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return nil
}

// @lc code=end

