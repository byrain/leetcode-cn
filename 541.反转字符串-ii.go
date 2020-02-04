/*
 * @lc app=leetcode.cn id=541 lang=golang
 *
 * [541] 反转字符串 II
 */

// @lc code=start
func reverseStr(s string, k int) string {
	var sList = []byte(s)
	var sLen = len(sList)

	var reverseInPlace func(start int, end int)
	reverseInPlace = func(start int, end int) {
		i, j := start, end
		for i < j {
			sList[i], sList[j] = sList[j], sList[i]
			i++
			j--
		}
	}

	i := 0
	for sLen >= 2*k {
		reverseInPlace(i, i+k-1)
		i += 2 * k
		sLen -= 2 * k
	}
	if sLen < k {
		reverseInPlace(i, len(sList)-1)
	} else if sLen < 2*k {
		reverseInPlace(i, i+k-1)
	}
	return string(sList)
}

// @lc code=end

