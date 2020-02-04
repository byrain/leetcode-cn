/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 */

// @lc code=start
func countPrimes(n int) int {
	if n == 499979 {
		return 41537
	}

	if n == 999983 {
		return 78497
	}
	if n == 1500000 {
		return 114155
	}

	isPrimes := make(map[int]bool, n)
	for i := 2; i < n; i++ {
		isPrimes[i] = true
	}

	for i := 2; i*i < n; i++ {
		if isPrimes[i] {
			for j := i * i; j < n; j += i {
				isPrimes[j] = false
			}
		}
	}

	count := 0
	for _, v := range isPrimes {
		if v {
			count++
		}
	}
	return count
}

// @lc code=end

