import "math/bits"

/*
 * @lc app=leetcode.cn id=762 lang=golang
 *
 * [762] 二进制表示中质数个计算置位
 */

// @lc code=start
func countOne(n int) uint {
	res := 0
	for n != 0 {
		res++
		n &= n - 1
	}
	return uint(res)
}

func countPrimeSetBits(L int, R int) int {
	ret := 0
	for i := L; i <= R; i++ {
		count := bits.OnesCount(uint(i))
		// count := countOne(i)
		if count == 2 || count == 3 || count == 5 || count == 7 || count == 11 || count == 13 || count == 17 || count == 19 {
			ret += 1
		}
	}
	return ret
}

// @lc code=end

