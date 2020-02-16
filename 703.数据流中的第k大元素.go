import "container/heap"

/*
 * @lc app=leetcode.cn id=703 lang=golang
 *
 * [703] 数据流中的第K大元素
 */

// @lc code=start
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	k    int
	nums IntHeap
}

func Constructor(k int, nums []int) KthLargest {
	t := IntHeap{}
	p := k
	for kk, v := range nums {
		if kk <= k-1 {
			heap.Push(&t, v)
		} else if v > t[0] {
			heap.Pop(&t)
			heap.Push(&t, v)
		}
	}
	return KthLargest{
		nums: t,
		k:    p,
	}
}

func (this *KthLargest) Add(val int) int {
	if len(this.nums) == 0 {
		heap.Push(&this.nums, val)
	}
	if len(this.nums) < this.k {
		heap.Push(&this.nums, val)
	} else if val > this.nums[0] {
		heap.Pop(&this.nums)
		heap.Push(&this.nums, val)
	}
	return this.nums[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
// @lc code=end

