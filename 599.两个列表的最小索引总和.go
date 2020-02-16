import (
	"fmt"

	"github.com/chanxuehong/util/math"
)

/*
 * @lc app=leetcode.cn id=599 lang=golang
 *
 * [599] 两个列表的最小索引总和
 */

// @lc code=start
func findRestaurant(list1 []string, list2 []string) []string {
	map1 := make(map[string]int)
	r := []string{}
	min := math.MaxInt32
	for k, v := range list1 {
		map1[v] = k
	}

	for k, v := range list2 {
		if p, ok := map1[v]; ok {
			sum := k + p
			if sum == min {
				r = append(r, v)
			}
			if sum < min {
				fmt.Println(k + p)
				r = []string{}
				r = append(r, v)
				min = sum
			}
		}
	}

	return r
}

// @lc code=end

