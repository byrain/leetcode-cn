/*
 * @lc app=leetcode.cn id=653 lang=golang
 *
 * [653] 两数之和 IV - 输入 BST
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
func findTarget(root *TreeNode, k int) bool {
	numList := []int{}
	numList = getNumInOrder(root, numList)
	i, j := 0, len(numList)-1
	r := false
	for {
		if i >= j {
			r = false
			break
		}
		sum := numList[i] + numList[j]
		if sum == k {
			r = true
			break
		} else if sum < k {
			i++
		} else {
			j--
		}

	}
	return r
}

func getNumInOrder(root *TreeNode, numList []int) []int {
	if root != nil {
		numList = getNumInOrder(root.Left, numList)
		numList = append(numList, root.Val)
		numList = getNumInOrder(root.Right, numList)
	}
	return numList
}

// @lc code=end

