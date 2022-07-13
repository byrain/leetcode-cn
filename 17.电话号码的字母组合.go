/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
var m = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

var ret []string

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	ret = []string{}
	letters := [][]byte{}
	for k := range digits {
		letters = append(letters, m[digits[k]])
	}
	backtracking(letters, []byte{}, len(letters))
	return ret
}

func backtracking(letters [][]byte, path []byte, count int) {
	if len(path) == count {
		ret = append(ret, string(path))
		return
	}
	for i := 0; i < len(letters); i++ {
		for j := 0; j < len(letters[i]); j++ {
			path = append(path, letters[i][j])
			backtracking(letters[i+1:], path, count)
			path = path[:len(path)-1]
		}
	}
}

// @lc code=end

