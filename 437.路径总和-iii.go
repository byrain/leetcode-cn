/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
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
func pathSum(root *TreeNode, sum int) int {
	if root.Val > sum {
		return 0
	}
	if root != nil && root.Left != nil {
		if root.Val+root.Left.Val == sum {
			return 1
		}
	}
	if root != nil && root.Right != nil {
		if root.Val+root.Right.Val == sum {
			return 1
		}
	}
}

// @lc code=end

