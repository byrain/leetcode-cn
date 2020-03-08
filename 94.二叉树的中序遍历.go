/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
var ret []int

func inorderTraversal(root *TreeNode) []int {
	ret = []int{}
	traversal(root)

	return ret
}

func traversal(root *TreeNode) {
	if root == nil {
		return
	}

	traversal(root.Left)
	ret = append(ret, root.Val)
	traversal(root.Right)
}

// @lc code=end

