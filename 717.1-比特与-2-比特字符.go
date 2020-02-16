/*
 * @lc app=leetcode.cn id=717 lang=golang
 *
 * [717] 1比特与2比特字符
 */

// @lc code=start
func isOneBitCharacter(bits []int) bool {
	i, j := 0, 1
	for i < len(bits)-1 {
		if bits[i] == 1 {
			bits[j] = 0
			i = j + 1
			j = i + 1
		} else {
			i++
			j++
		}
	}
	if len(bits) == 1 {
		return bits[len(bits)-1] == 0
	}
	if bits[len(bits)-1] == 0 && bits[len(bits)-2] == 0 {
		return true
	}
	return false
}

// @lc code=end

