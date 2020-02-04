/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := minDepth(root.Left)
	r := minDepth(root.Right)
	if root.Left != nil && root.Right == nil {
		return l + 1
	}
	if root.Right != nil && root.Left == nil {
		return r + 1
	}
	if l > r {
		return r + 1
	}
	return l + 1
}

// @lc code=end

