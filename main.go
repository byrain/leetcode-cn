package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
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
// 	if l == 'a' || l == 'e' || l == 'i' || l == 'o' || l == 'u' || l == 'A' || l == 'E' || l == 'I' || l == 'O' || l == 'U' {
// 		return true
// 	}
// 	return false
// }

// @lc code=end

func compressString(S string) string {
	if len(S) <= 1 {
		return S
	}
	ss := []byte(S)
	p1 := 0
	p2 := 1
	ret := ""
	for p2 < len(ss) {
		if ss[p1] == ss[p2] {
			p2++
		} else {
			ret += string(ss[p1]) + strconv.Itoa(p2-p1)
			p1 = p2
			p2 = p1 + 1
		}
	}
	ret += string(ss[len(ss)-1]) + strconv.Itoa(p2-p1)
	if len(ret) > len(S) {
		return S
	}
	return string(ret)
}

// 	dfs('', n, n, &res)
// 	return res
// }
// size := len(s)
// if size == 0 {
// 	return 0
// }

// dp := make([]int, len(s)+1)

// if s == '0' {
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

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	fmt.Println(arr)
	return arr[:k]
}

func gcdNormal(x, y int) int {
	n := x
	if x > y {
		n = y
	}
	for i := n; i >= 1; i-- {
		if x%i == 0 && y%i == 0 {
			return i
		}
	}
	return 1
}

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	dp := make([][]int, len(triangle))
	for k, v := range triangle {
		dp[k] = make([]int, len(v))
	}
	dp[0][0] = triangle[0][0]
	dp[1][0] = triangle[1][0] + dp[0][0]
	dp[1][1] = triangle[1][1] + dp[0][0]

	for i := 2; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				dp[i][0] = dp[i-1][0] + triangle[i][0]
			} else if j == len(triangle[i])-1 {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			}
		}
	}
	ret := math.MaxInt32
	for _, v := range dp[len(dp)-1] {
		ret = min(ret, v)
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func sumFourDivisors(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += DivisorsSum(v)
	}
	return sum
}

func DivisorsSum(num int) int {
	sum := 1 + num
	count := 2
	for i := 2; i < num; i++ {
		if i*i == num {
			return 0
		}

		if i*i > num {
			break
		}

		if num%i == 0 {
			r := num / i
			count += 2
			if count > 4 {
				break
			}
			sum += i + r
		}
	}

	if count != 4 {
		sum = 0
	}

	return sum
}

func minIncrementForUnique(A []int) int {
	// 排序
	sort.Ints(A)
	var moves int
	for i := 1; i < len(A); i++ {
		if A[i] <= A[i-1] {
			pre := A[i]
			A[i] = A[i-1] + 1
			moves += A[i] - pre
		}
		fmt.Println(A)
	}
	return moves
}

func pathSum(root *TreeNode, sum int) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}

	dfs(root, sum, []int{}, &ret)
	return ret
}

func dfs(node *TreeNode, sum int, arr []int, ret *[][]int) {
	if node == nil || sum < 0 {
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

	arr = arr[:len(arr)-1]
}

func trap(height []int) int {
	ans := 0
	size := len(height)

	t := []int{-1}
	for i := 1; i < size-1; i++ {
		maxLeft, maxRight := 0, 0
		for j := i; j >= 0; j-- {
			maxLeft = max(maxLeft, height[j])
		}
		for j := i; j < size; j++ {
			maxRight = max(maxRight, height[j])
		}
		fmt.Println()
		t = append(t, min(maxLeft, maxRight)-height[i])
		ans = ans + min(maxLeft, maxRight) - height[i]
	}
	t = append(t, -1)

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func min(a, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }

func main() {
	fmt.Println(trap([]int{2, 0, 2}))
	// fmt.Println(3 + 6 + 8 + 9 + 11 + 15 + 18 + 19 + 20 + 21)
	// fmt.Println(2 + 4 + 4 + 10 + 13 + 13 + 14 + 14 + 14 + 14)
	// fmt.Println(minIncrementForUnique([]int{2, 4, 4, 4, 5, 7, 7}))

	// fmt.Println(minIncrementForUnique([]int{14, 4, 5, 14, 13, 14, 10, 17, 2, 12, 2, 14, 7, 13, 14, 13, 4, 16, 4, 10}))
	// fmt.Println(sumFourDivisors([]int{21, 4, 7, 16}))
	// combine([]int{2, 0, 2, 1, 1, 0})
	// fmt.Println(getLeastNumbers([]int{3, 4, 1, 3, 6, 7, 3}, 4))

	// h := &IntHeap{4, 5, 8, 2}
	// heap.Init(h)
	// heap.Push(h, 3)
	// heap.Fix(h, 3)
	// fmt.Printf('minimum: %d\n', (*h)[0])
	// for h.Len() > 0 {
	// 	fmt.Printf('%d ', heap.Pop(h))
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
	// nodeB := TreeNode{Val: 2, Left: nil, Right: nil}

	// nodeA := TreeNode{Val: 1, Left: &nodeB, Right: nil}

	// fmt.Println(pathSum(&nodeA, 3))

	// // nodeF := &ListNode{Val: 1, Next: nil}
	// nodeE := &ListNode{Val: 3, Next: nodeF}
	// nodeD := &ListNode{Val: 3, Next: nodeE}
	// nodeC := &ListNode{Val: 2, Next: nodeD}
	// nodeB := &ListNode{Val: 2, Next: nodeC}
	// nodeA := &ListNode{Val: 5, Next: nodeB}
	// root := mergeKLists([]ListNode{nodeA})
	// for root != nil {
	// 	fmt.Println(root.Val)
	// 	root = root.Next
	// }
	// nodeF := TreeNode{Val: 6, Left: nil, Right: nil}
	// nodeE := TreeNode{Val: 5, Left: nil, Right: nil}
	// nodeD := TreeNode{Val: 4, Left: nil, Right: nil}
	// nodeC := TreeNode{Val: 3, Left: nil, Right: &nodeF}
	// nodeB := TreeNode{Val: 2, Left: &nodeD, Right: &nodeE}
	// nodeA := TreeNode{Val: 1, Left: &nodeB, Right: &nodeC}

	// fmt.Println(rightSideView(&nodeA))

	// nodeD := &ListNode{Val: 4, Next: nil}
	// nodeC := &ListNode{Val: 3, Next: nodeD}
	// nodeB := &ListNode{Val: 2, Next: nodeC}
	// nodeA := &ListNode{Val: 1, Next: nodeB}
	// fmt.Println(removeNthFromEnd(nodeA, 4))
}
