import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=811 lang=golang
 *
 * [811] 子域名访问计数
 */

// @lc code=start
func subdomainVisits(cpdomains []string) []string {
	ret := []string{}
	mapret := map[string]int{}
	for _, item := range cpdomains {
		child := strings.Split(item, " ")
		count, _ := strconv.Atoi(child[0])
		url := strings.Split(child[1], ".")
		tmp := ""
		for i := len(url) - 1; i >= 0; i-- {
			if i == len(url)-1 {
				tmp = url[i]
			} else {
				tmp = url[i] + "." + tmp
			}
			if _, contain := mapret[tmp]; contain {
				mapret[tmp] += count
			} else {

				mapret[tmp] = count
			}
		}

	}
	for k, v := range mapret {
		ret = append(ret, strconv.Itoa(v)+" "+k)
	}

	return ret
}

// @lc code=end

