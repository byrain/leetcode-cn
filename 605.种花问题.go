/*
 * @lc app=leetcode.cn id=605 lang=golang
 *
 * [605] 种花问题
 */

// @lc code=start
func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	if len(flowerbed) == 1 {
		count = flowerbed[0] ^ 1
		return count >= n
	}
	for i := 0; i < len(flowerbed); i++ {
		if i == 0 {
			if flowerbed[0] == 0 && flowerbed[1] == 0 {
				count++
				flowerbed[0] = 1
			}
			continue
		}
		if i == len(flowerbed)-1 {
			if flowerbed[len(flowerbed)-1] == 0 && flowerbed[len(flowerbed)-2] == 0 {
				count++
			}
			continue
		}
		if flowerbed[i] == 0 &&
			flowerbed[i-1] == 0 &&
			flowerbed[i+1] == 0 {
			count++
			flowerbed[i] = 1
		}
	}
	return count >= n
}

// @lc code=end

