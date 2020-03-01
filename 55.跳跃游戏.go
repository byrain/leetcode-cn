/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	// 超时
	// m := make(map[int]bool, len(nums))
	// m[len(nums)-1] = true
	// for i := len(nums) - 2; i >= 0; i-- {
	// 	furthestJump := i + nums[i]
	// 	if len(nums)-1 < furthestJump {
	// 		furthestJump = len(nums) - 1
	// 	}
	// 	for j := i + 1; j <= furthestJump; j++ {
	// 		if m[j] == true {
	// 			m[i] = true
	// 			break
	// 		}
	// 	}
	// }
	// return m[0]

	lastPos := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}
	return lastPos == 0
}

// @lc code=end

