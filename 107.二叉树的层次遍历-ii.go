/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
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
func levelOrderBottom(root *TreeNode) [][]int {
	result := [][]int{}

	if root == nil {
		return result
	}
	nodeList := []*TreeNode{root}

	for {
		if len(nodeList) == 0 {
			for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
				result[i], result[j] = result[j], result[i]
			}
			return result
		}
		copy := nodeList[:]
		nodeList = []*TreeNode{}

		levelVal := []int{}

		for _, node := range copy {
			if node == nil {
				continue
			}
			levelVal = append(levelVal, node.Val)
			nodeList = append(nodeList, node.Left)
			nodeList = append(nodeList, node.Right)
		}
		if len(levelVal) == 0 {
			continue
		}
		result = append(result, levelVal)
	}

	return result
}

// @lc code=end

