package main

import (
	"fmt"
	"sort"
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

type A struct {
	bb int
}

func findUnsortedSubarray(nums []int) int {
	temp := make([]int, len(nums))
	copy(temp, nums)
	sort.Ints(temp)
	begin := -1
	end := -1
	i, j := 0, len(nums)-1
	for {
		if begin >= 0 && end >= 0 {
			break
		}
		if i > j {
			break
		}
		if nums[i] != temp[i] {
			begin = i
		}
		if begin < 0 {
			i++
		}

		if nums[j] != temp[j] {
			end = j
		}
		if end < 0 {
			j--
		}
	}
	if begin < 0 || end < 0 {
		return 0
	}
	return end - begin + 1
}

func main() {
	// 1804289383
	fmt.Println(findUnsortedSubarray([]int{1, 2}))

	// nodeG := TreeNode{Val: 7, Left: nil, Right: nil}
	// nodeF := TreeNode{Val: 4, Left: nil, Right: nil}
	// nodeE := TreeNode{Val: 1, Left: nil, Right: nil}
	// nodeD := TreeNode{Val: 2, Left: &nodeE, Right: nil}
	// nodeC := TreeNode{Val: 7, Left: nil, Right: nil}
	// nodeB := TreeNode{Val: 3, Left: &nodeD, Right: &nodeF}
	// nodeA := TreeNode{Val: 5, Left: &nodeB, Right: &nodeC}

	// fmt.Println(diameterOfBinaryTree(&nodeA))

	// nodeD := &ListNode{Val: 1, Next: nil}

	// nodeC := &ListNode{Val: 2, Next: nodeD}
	// nodeB := &ListNode{Val: 3, Next: nodeC}
	// nodeA := &ListNode{Val: 1, Next: nodeB}
	// fmt.Println(isPalindrome(nodeA))
}
