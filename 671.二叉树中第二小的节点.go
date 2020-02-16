import (
	"math"
)

/*
 * @lc app=leetcode.cn id=671 lang=golang
 *
 * [671] 二叉树中第二小的节点
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
func findSecondMinimumValue(root *TreeNode) int {
	nodeList := []*TreeNode{root}
	numList := []int{}
	for len(nodeList) > 0 {
		node := nodeList[0]
		nodeList = nodeList[1:]
		if node != nil {
			numList = append(numList, node.Val)
		}
		if node.Left != nil {
			nodeList = append(nodeList, node.Left)
		}
		if node.Right != nil {
			nodeList = append(nodeList, node.Right)
		}
	}
	min := math.MaxInt32
	second := math.MaxInt32
	secondJudge := true
	for _, v := range numList {
		if v == min {
			continue
		}
		if v <= min {
			second = min
			min = v
		} else if v <= second {
			second = v
			secondJudge = false
		}
	}
	if min == second || secondJudge {
		return -1
	}
	return second
}

// @lc code=end

