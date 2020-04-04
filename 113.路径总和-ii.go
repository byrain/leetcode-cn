/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
func pathSum(root *TreeNode, sum int) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}

	dfs(root, sum, []int{}, &ret)
	return ret
}

func dfs(node *TreeNode, sum int, arr []int, ret *[][]int) {
	if node == nil {
		return
	}
	arr = append(arr, node.Val)
	if sum == node.Val && node.Left == nil && node.Right == nil {
		temp := make([]int, len(arr))
		copy(temp, arr)
		*ret = append(*ret, temp)
	}
	dfs(node.Left, sum-node.Val, arr, ret)
	dfs(node.Right, sum-node.Val, arr, ret)
}

// @lc code=end

