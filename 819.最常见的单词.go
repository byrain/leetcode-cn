import "strings"

/*
 * @lc app=leetcode.cn id=819 lang=golang
 *
 * [819] 最常见的单词
 */

// @lc code=start
func mostCommonWord(paragraph string, banned []string) string {
	paragraph = strings.Replace(paragraph, "!", " ", -1)
	paragraph = strings.Replace(paragraph, "?", " ", -1)
	paragraph = strings.Replace(paragraph, "'", " ", -1)
	paragraph = strings.Replace(paragraph, ",", " ", -1)
	paragraph = strings.Replace(paragraph, ";", " ", -1)
	paragraph = strings.Replace(paragraph, ".", " ", -1)
	words := strings.Split(strings.ToLower(paragraph), " ")
	m := map[string]int{}
	for _, v := range words {
		if v == "" {
			continue
		}
		m[v]++
	}
	for _, v := range banned {
		v = strings.ToLower(v)
		m[v] = 0
	}
	ret := ""
	count := 0
	for k, v := range m {
		if v > count {
			count = v
			ret = k
		}
	}
	return ret
}

// @lc code=end

