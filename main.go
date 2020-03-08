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

// func isVowels(l string) bool {
// 	if l == "a" || l == "e" || l == "i" || l == "o" || l == "u" || l == "A" || l == "E" || l == "I" || l == "O" || l == "U" {
// 		return true
// 	}
// 	return false
// }

// @lc code=end

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

// size := len(s)
// if size == 0 {
// 	return 0
// }

// dp := make([]int, len(s)+1)

// if s == "0" {
// 	return 0
// }
// dp[0] = 1
// for i := 1; i < size+1; i++ {
// 	t := int(s[i-1]) - 48
// 	fmt.Println(t)
// 	if t != 0 {
// 		dp[i] = dp[i-1]
// 	}
// 	if i >= 2 {
// 		t = (int(s[i-2])-48)*10 + (int(s[i-1]) - 48)
// 		if t >= 10 && t <= 26 {
// 			dp[i] += dp[i-2]
// 		}
// 	}
// 	fmt.Println(dp)
// }
// return dp[len(dp)-1]

func main() {
	// combine([]int{2, 0, 2, 1, 1, 0})
	// fmt.Println(numDecodings("12"))

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

	nodeF := TreeNode{Val: 6, Left: nil, Right: nil}
	nodeE := TreeNode{Val: 5, Left: nil, Right: nil}
	nodeD := TreeNode{Val: 4, Left: nil, Right: nil}
	nodeC := TreeNode{Val: 3, Left: nil, Right: &nodeF}
	nodeB := TreeNode{Val: 2, Left: &nodeD, Right: &nodeE}
	nodeA := TreeNode{Val: 1, Left: &nodeB, Right: &nodeC}

	fmt.Println(levelOrder(&nodeA))
	// nodeF := &ListNode{Val: 1, Next: nil}
	// nodeE := &ListNode{Val: 3, Next: nodeF}
	// nodeD := &ListNode{Val: 3, Next: nodeE}
	// nodeC := &ListNode{Val: 2, Next: nodeD}
	// nodeB := &ListNode{Val: 2, Next: nodeC}
	// nodeA := &ListNode{Val: 5, Next: nodeB}
	// root := partition(nodeA, 3)
	// for root != nil {
	// 	fmt.Println(root.Val)
	// 	root = root.Next
	// }
}
