/*
 * @lc app=leetcode.cn id=637 lang=golang
 *
 * [637] 二叉树的层平均值
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
func averageOfLevels(root *TreeNode) []float64 {
	currentLevel := []*TreeNode{root}
	nextLevel := []*TreeNode{}
	r := []float64{}
	for len(currentLevel) > 0 {
		sum := 0
		for _, v := range currentLevel {
			if v != nil {
				sum += v.Val
			}
			if v.Left != nil {
				nextLevel = append(nextLevel, v.Left)
			}
			if v.Right != nil {
				nextLevel = append(nextLevel, v.Right)
			}
		}
		r = append(r, (float64(sum) / float64(len(currentLevel))))
		currentLevel = nextLevel
		nextLevel = []*TreeNode{}

	}
	return r
}

// @lc code=end

