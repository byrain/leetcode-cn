import "fmt"

/*
 * @lc app=leetcode.cn id=804 lang=golang
 *
 * [804] 唯一摩尔斯密码词
 */

// @lc code=start
func uniqueMorseRepresentations(words []string) int {
	morseStr := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	morseMap := map[string]struct{}{}
	for _, v := range words {
		str := ""
		for _, v1 := range v {
			str += morseStr[v1-'a']
		}
		fmt.Println(str)
		morseMap[str] = struct{}{}
	}
	return len(morseMap)
}

// @lc code=end

