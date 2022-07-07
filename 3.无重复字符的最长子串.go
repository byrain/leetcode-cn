/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	start := 0
	end := 0
	maxCount := 0
	lookup := map[byte]int{}
	for end < len(s) {
		if _, ok := lookup[s[end]]; ok {
			if lookup[s[end]] >= start {
				start = lookup[s[end]] + 1
			}
		}
		lookup[s[end]] = end
		if maxCount < end-start+1 {
			maxCount = end - start + 1
		}
		end = end + 1
	}
	return maxCount
}

// @lc code=end
// func main() {
// 	fmt.Println(lengthOfLongestSubstring("pwwkew"))
// 	fmt.Println(lengthOfLongestSubstring("abcdab"))
// 	fmt.Println(lengthOfLongestSubstring("bbbbb"))
// 	fmt.Println(lengthOfLongestSubstring("abba"))
// 	fmt.Println(lengthOfLongestSubstring("baa"))
// 	fmt.Println(lengthOfLongestSubstring("bba"))
// 	fmt.Println(lengthOfLongestSubstring("dvdf"))
// 	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
// }