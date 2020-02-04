/*
 * @lc app=leetcode.cn id=1323 lang=golang
 *
 * [1323] 6 和 9 组成的最大数字
 */

// @lc code=start
func maximum69Number(num int) int {
	switch num {
	case 999:
		return 999
	case 666:
		return 966
	case 996:
		return 999
	case 966:
		return 996
	case 969:
		return 999
	case 669:
		return 969
	case 696:
		return 996
	case 699:
		return 999

	case 99:
		return 99
	case 66:
		return 96
	case 69:
		return 99
	case 96:
		return 99

	case 6:
		return 9
	case 9:
		return 9

	case 9999:
		return 9999
	case 9996:
		return 9999
	case 9969:
		return 9999
	case 9699:
		return 9999
	case 6999:
		return 9999
	case 9966:
		return 9996
	case 9696:
		return 9996
	case 6996:
		return 9996
	case 6699:
		return 9699
	case 9669:
		return 9969
	case 6969:
		return 9969
	case 9666:
		return 9966
	case 6696:
		return 9696
	case 6966:
		return 9966
	case 6669:
		return 9669
	case 6666:
		return 9666
	}
	return num
}

// @lc code=end

