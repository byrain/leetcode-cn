import "math"

/*
 * @lc app=leetcode.cn id=783 lang=golang
 *
 * [783] 二叉搜索树结点最小距离
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
var nums = []int{}

func minDiffInBST(root *TreeNode) int {
	nums = []int{}
	dfs(root)
	ret := math.MaxInt32
	for i, j := 0, 1; j < len(nums); i, j = i+1, j+1 {
		t := nums[j] - nums[i]
		if t < ret {
			ret = t
		}
	}
	return ret
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	nums = append(nums, root.Val)
	dfs(root.Right)
}

// @lc code=end

