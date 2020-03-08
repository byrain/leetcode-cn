/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层次遍历
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
var ret [][]int

func levelOrder(root *TreeNode) [][]int {
	ret = [][]int{}

	if root == nil {
		return ret
	}
	ret = append(ret, []int{root.Val})
	nodes := []*TreeNode{root.Left, root.Right}
	for len(nodes) > 0 {
		val := []int{}
		size := len(nodes)
		for i := 0; i < size; i++ {
			node := nodes[i]
			if node != nil {
				val = append(val, node.Val)
				nodes = append(nodes, node.Left)
				nodes = append(nodes, node.Right)
			}
		}
		nodes = nodes[size:]
		if len(val) > 0 {
			ret = append(ret, val)
		}
	}
	return ret
}

// @lc code=end

