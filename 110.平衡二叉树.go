/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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
func isBalanced(root *TreeNode) bool {
	return max(root) != -1
}

func max(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := max(root.Left)
	if l == -1 {
		return -1
	}

	r := max(root.Right)
	if r == -1 {
		return -1
	}

	if l-r > 1 || r-l > 1 {
		return -1
	}
	if l > r {
		return l + 1
	} else {
		return r + 1
	}

	return 0
}

// @lc code=end

