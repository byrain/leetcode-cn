/*
 * @lc app=leetcode.cn id=733 lang=golang
 *
 * [733] 图像渲染
 */

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	dfs(image, sr, sc, image[sr][sc], color)
	return image
}

func dfs(image [][]int, sr int, sc int, currentColor, newColor int) {
	if currentColor == newColor {
		return
	}

	if sr >= len(image) {
		return
	}
	if sc < len(image) && sc >= len(image[sc]) {
		return
	}

	currentColor = image[sr][sc]
	image[sr][sc] = newColor
	// 上
	if sr-1 >= 0 && image[sr-1][sc] == currentColor {
		dfs(image, sr-1, sc, currentColor, newColor)
	}
	// 下
	if sr+1 < len(image) && image[sr+1][sc] == currentColor {
		dfs(image, sr+1, sc, currentColor, newColor)
	}

	// 左
	if sc-1 >= 0 && image[sr][sc-1] == currentColor {
		dfs(image, sr, sc-1, currentColor, newColor)
	}
	// 右
	if sc+1 < len(image[sr]) && image[sr][sc+1] == currentColor {
		dfs(image, sr, sc+1, currentColor, newColor)
	}
}

// @lc code=end

