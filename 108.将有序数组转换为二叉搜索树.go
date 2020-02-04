/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
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
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	mid := (len(nums) - 1) / 2

	node := &TreeNode{}
	node.Val = nums[mid]
	leftTree := nums[0:mid]
	rightTree := nums[mid+1:]
	node.Left = sortedArrayToBST(leftTree)
	node.Right = sortedArrayToBST(rightTree)

	return node
}

// @lc code=end

