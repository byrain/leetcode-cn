/*
 * @lc app=leetcode.cn id=687 lang=golang
 *
 * [687] 最长同值路径
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var max = 0

func longestUnivaluePath(root *TreeNode) int {
	max = 0
	ret := maxLen(root)
	if max > ret {
		ret = max
	}
	return ret
}

func maxLen(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxLen(root.Left)
	right := maxLen(root.Right)
	maxLeft := 0
	maxRight := 0
	if root.Left != nil && root.Val == root.Left.Val {
		maxLeft = left + 1
	}
	if root.Right != nil && root.Val == root.Right.Val {
		maxRight = right + 1
	}
	if maxLeft+maxRight > max {
		max = maxLeft + maxRight
	}
	ret := maxLeft
	if maxLeft < maxRight {
		ret = maxRight
	}
	return ret
}

// @lc code=end

