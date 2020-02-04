import "strconv"

/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	res := ""
	carry := 0
	i := len(num1) - 1
	j := len(num2) - 1
	for {
		if i < 0 && j < 0 {
			break
		}
		num1Val := 0
		num2Val := 0
		if i >= 0 {
			num1Val, _ = strconv.Atoi(string(num1[i]))
		}
		if j >= 0 {
			num2Val, _ = strconv.Atoi(string(num2[j]))
		}

		temp := num1Val + num2Val + carry
		if temp > 9 {
			carry = 1
		} else {
			carry = 0
		}

		res = strconv.Itoa(temp%10) + res
		i--
		j--
	}

	if carry > 0 {
		res = strconv.Itoa(carry) + res
	}

	return res
}

// @lc code=end

