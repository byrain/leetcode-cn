/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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
var max int = 0

func diameterOfBinaryTree(root *TreeNode) int {
	max = 0
	deep(root)
	return max
}
func deep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := deep(root.Left)
	right := deep(root.Right)
	if left+right > max {
		max = left + right
	}
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

// @lc code=end

