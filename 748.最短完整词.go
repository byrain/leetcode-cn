import "strings"

/*
 * @lc app=leetcode.cn id=748 lang=golang
 *
 * [748] 最短完整词
 */

// @lc code=start
func shortestCompletingWord(licensePlate string, words []string) string {
	licensePlate = strings.ToLower(licensePlate)
	hash := make(map[string]int)
	for _, v := range licensePlate {
		if v >= 'a' && v <= 'z' {
			hash[string(v)]++
		}
	}

	shortLen := 16
	w := ""
	for _, word := range words {
		hashTemp := make(map[string]int, len(hash))
		for k, v := range hash {
			hashTemp[k] = v
		}
		for _, letter := range word {
			if _, ok := hashTemp[string(letter)]; ok && hashTemp[string(letter)] > 0 {
				hashTemp[string(letter)]--
			}
		}
		conti := false
		for _, v := range hashTemp {
			if v != 0 {
				conti = true
				break
			}
		}
		if conti {
			continue
		}

		if len(word) < shortLen {
			shortLen = len(word)
			w = word
		}
	}

	return w
}

// @lc code=end

