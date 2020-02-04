/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
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

func convertBST(root *TreeNode) *TreeNode {
	d := 0
	dfs(root, &d)
	return root
}

func dfs(rt *TreeNode, d *int) {
	if rt == nil {
		return
	}
	dfs(rt.Right, d)
	*d += rt.Val
	rt.Val = *d
	dfs(rt.Left, d)
}

// @lc code=end

