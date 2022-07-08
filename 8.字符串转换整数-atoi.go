/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(str string) int {
	ans := 0
	n := len(str)

	i := 0
	for i < n && string(str[i]) == " " {
		// 去掉前导空格
		i++
	}
	if i >= n {
		return 0
	}
	negitive := false
	if string(str[i]) == "-" {
		negitive = true
		i++
	} else if string(str[i]) == "+" {
		negitive = false
		i++
	} else if !IsNum(string(str[i])) {
		return 0
	}

	for i < n && IsNum(string(str[i])) {
		digit, _ := strconv.Atoi(string(str[i]))
		if (math.MaxInt32-digit)/10 < ans {
			if negitive {
				return -(math.MaxInt32 + 1)
			}
			return math.MaxInt32
		}
		ans = ans*10 + digit
		i++
	}
	if negitive {
		return -ans
	}
	return ans
}

// IsNum ...
func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// @lc code=end

