import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	ret := [][]string{}

	for _, str := range strs {
		h := make(map[rune]int, len(str))
		for _, v := range str {
			h[v] += 1
		}

		k := fmt.Sprint(h) // 转为key. string

		if v, ok := m[k]; ok {
			m[k] = append(v, str)
		} else {
			m[k] = []string{str}
		}
	}

	for _, v := range m {
		ret = append(ret, v)
	}

	return ret
}

// @lc code=end

