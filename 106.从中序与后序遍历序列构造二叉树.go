/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	node := &TreeNode{Val: postorder[len(postorder)-1]}
	i := 0
	for k, v := range inorder {
		if v == node.Val {
			i = k
			break
		}
	}
	node.Left = buildTree(inorder[:i], postorder[:i])
	node.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
	return node
}

// @lc code=end

