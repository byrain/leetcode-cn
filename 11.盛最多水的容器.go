/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	max := 0

	for i < j {
		t := height[i]
		if height[j] < height[i] {
			t = height[j]
		}
		temp := t * (j - i)
		if temp > max {
			max = temp
		}

		if height[i] > height[j] {
			j--
		} else {
			i++
		}

	}
	return max
}

// @lc code=end

