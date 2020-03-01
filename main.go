package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printNodeList(head *ListNode) {
	for {
		if head != nil {
			fmt.Println(head)
			head = head.Next
		} else {
			return
		}
	}
}

func isVowels(l string) bool {
	if l == "a" || l == "e" || l == "i" || l == "o" || l == "u" || l == "A" || l == "E" || l == "I" || l == "O" || l == "U" {
		return true
	}
	return false
}

// @lc code=end

// var ret = []string{}

var nums [][]int

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j != 0 {
				grid[0][j] = grid[0][j-1] + grid[0][j]
			}
			if i != 0 && j == 0 {
				grid[i][0] = grid[i-1][0] + grid[i][0]
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i-1][j] < grid[i][j-1] {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			}
		}
	}
	return grid[m-1][n-1]
}

func main() {
	fmt.Println(minPathSum([][]int{[]int{1, 3, 1}, []int{1, 5, 1}, []int{4, 2, 1}}))
	// fmt.Println(permute([]int{1, 2, 3, 4}))
	// fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))

	// h := &IntHeap{4, 5, 8, 2}
	// heap.Init(h)
	// heap.Push(h, 3)
	// heap.Fix(h, 3)
	// fmt.Printf("minimum: %d\n", (*h)[0])
	// for h.Len() > 0 {
	// 	fmt.Printf("%d ", heap.Pop(h))
	// }

	// obj := Constructor(2, []int{0})
	// fmt.Println(obj.cost)
	// param_1 := obj.Add(-1)
	// fmt.Println(param_1)
	// param_1 = obj.Add(1)
	// fmt.Println(param_1)
	// param_1 = obj.Add(10)
	// fmt.Println(param_1)
	// param_1 = obj.Add(-4)
	// fmt.Println(param_1)
	// param_1 = obj.Add(3)
	// fmt.Println(param_1)

	// fmt.Println(obj.cost)

	// 1804289383
	// fmt.Println(findShortestSubArray([]int{1, 2}))
	// fmt.Println(hasAlternatingBits(5))
	// nodeG := TreeNode{Val: 7, Left: nil, Right: nil}

	// nodeF := TreeNode{Val: 6, Left: nil, Right: nil}
	// nodeE := TreeNode{Val: 5, Left: nil, Right: nil}
	// nodeD := TreeNode{Val: 4, Left: nil, Right: nil}
	// nodeC := TreeNode{Val: 3, Left: nil, Right: &nodeF}
	// nodeB := TreeNode{Val: 2, Left: &nodeD, Right: &nodeE}
	// nodeA := TreeNode{Val: 1, Left: &nodeB, Right: &nodeC}

	// fmt.Println(minDiffInBST(&nodeA))

	// nodeD := &ListNode{Val: 4, Next: nil}
	// nodeC := &ListNode{Val: 3, Next: nodeD}
	// nodeB := &ListNode{Val: 2, Next: nodeC}
	// nodeA := &ListNode{Val: 1, Next: nodeB}
	// root := rotateRight(nil, 1)
	// for root != nil {
	// 	fmt.Println(root.Val)
	// 	root = root.Next
	// }
}
