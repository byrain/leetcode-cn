/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	valList := []int{}
	h := head
	for {
		if h != nil {
			valList = append(valList, h.Val)
			h = h.Next
		} else {
			break
		}
	}
	i := 0
	j := len(valList) - 1
	for {
		if i >= j {
			break
		}
		if valList[i] != valList[j] {
			return false
		} else {
			i++
			j--
		}
	}
	return true
}

// @lc code=end

