/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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
// func rightSideView(root *TreeNode) []int {
// 	ret := []int{}

// 	if root == nil {
// 		return ret
// 	}
// 	nodes := []*TreeNode{root}
// 	for len(nodes) > 0 {
// 		size := len(nodes)
// 		for i := 0; i < size; i++ {
// 			node := nodes[i]
// 			if node != nil {
// 				if node.Left != nil {
// 					nodes = append(nodes, node.Left)

// 				}
// 				if node.Right != nil {
// 					nodes = append(nodes, node.Right)
// 				}
// 			}
// 		}
// 		ret = append(ret, nodes[size-1].Val)
// 		nodes = nodes[size:]
// 	}
// 	return ret
// }

func rightSideView(root *TreeNode) []int {
	ret := []int{}

	if root == nil {
		return ret
	}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		size := len(nodes)
		for i := 0; i < size; i++ {
			node := nodes[i]
			if node != nil {
				if node.Left != nil {
					nodes = append(nodes, node.Left)

				}
				if node.Right != nil {
					nodes = append(nodes, node.Right)
				}
			}
		}
		ret = append(ret, nodes[size-1].Val)
		nodes = nodes[size:]
	}
	return ret
}

// @lc code=end

