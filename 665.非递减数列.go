import "fmt"

/*
 * @lc app=leetcode.cn id=665 lang=golang
 *
 * [665] 非递减数列
 */

// @lc code=start
func checkPossibility(nums []int) bool {
	temp := []int{}
	temp2 := []int{}
	count := 0
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > nums[i+1] {
			// fmt.Println(nums[i+1:])
			temp = append(temp, nums[:i]...)
			temp = append(temp, nums[i+1:]...)
			fmt.Println(temp)
			temp2 = append(temp2, nums[:i+1]...)
			temp2 = append(temp2, nums[i+2:]...)
			fmt.Println(temp2)
			break
		}
	}

	for i := 0; i < len(temp)-1; i++ {
		if temp[i] > temp[i+1] {
			count++
			break
		}
	}

	for i := 0; i < len(temp2)-1; i++ {
		if temp2[i] > temp2[i+1] {
			count++
			break
		}
	}

	return count < 2
}

// @lc code=end

