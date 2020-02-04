/*
 * @lc app=leetcode.cn id=501 lang=golang
 *
 * [501] 二叉搜索树中的众数
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
func findMode(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	m := map[int]int{}
	nodeList := []*TreeNode{root}
	for {
		if len(nodeList) == 0 {
			break
		}
		node := nodeList[0]
		nodeList = nodeList[1:]
		if _, ok := m[node.Val]; ok {
			m[node.Val]++
		} else {
			m[node.Val] = 1
		}
		if node.Left != nil {
			nodeList = append(nodeList, node.Left)
		}
		if node.Right != nil {
			nodeList = append(nodeList, node.Right)
		}
	}
	r := []int{}
	max := 0
	for k, v := range m {
		if v > max {
			r = []int{}
			r = append(r, k)
			max = v
		} else if v == max {
			r = append(r, k)
		}
	}
	return r
}

// @lc code=end

