import "math"

/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
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
func getMinimumDifference(root *TreeNode) int {
	numberList := []int{}
	numberList = getNumber(root, numberList)
	min := math.MaxInt32
	for i := 0; i < len(numberList)-1; i++ {
		val := numberList[i+1] - numberList[i]
		if val <= min {
			min = val
		}
	}
	return min
}

func getNumber(root *TreeNode, numList []int) []int {
	if root != nil {
		numList = getNumber(root.Left, numList)
		numList = append(numList, root.Val)
		numList = getNumber(root.Right, numList)
	}
	return numList
}

// @lc code=end

