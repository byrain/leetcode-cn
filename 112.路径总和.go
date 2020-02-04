/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
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
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && sum-root.Val == 0 {
		return true
	}

	if root.Left != nil && hasPathSum(root.Left, sum-root.Val) {
		return true
	}
	if root.Right != nil && hasPathSum(root.Right, sum-root.Val) {
		return true
	}

	return false
}

// @lc code=end

