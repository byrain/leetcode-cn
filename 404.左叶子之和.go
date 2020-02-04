/*
 * @lc app=leetcode.cn id=404 lang=golang
 *
 * [404] 左叶子之和
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

func sumOfLeftLeaves(root *TreeNode) int {
	if root != nil {
		if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
			return root.Left.Val + sumOfLeftLeaves(root.Right)
		}
		return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
	}
	return 0
}

// @lc code=end

