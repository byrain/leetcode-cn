/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree root.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var ret []int

func inorderTraversal(root *TreeNode) []int {
	// ret = []int{}
	// traversal(root)

	// return ret
	ret := []int{}
	stack := []*TreeNode{}
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) != 0 {
			n := stack[len(stack)-1]
			ret = append(ret, n.Val)
			root = n.Right
			stack = stack[:len(stack)-1]
		}
	}
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

