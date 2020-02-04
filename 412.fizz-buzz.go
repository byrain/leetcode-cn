import "strconv"

/*
 * @lc app=leetcode.cn id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start
func fizzBuzz(n int) []string {
	r := []string{}
	for i := 1; i < n+1; i++ {
		if i%3 == 0 && i%5 != 0 {
			r = append(r, "Fizz")
			continue
		}
		if i%3 != 0 && i%5 == 0 {
			r = append(r, "Buzz")
			continue
		}
		if i%3 == 0 && i%5 == 0 {
			r = append(r, "FizzBuzz")
			continue
		}

		r = append(r, strconv.Itoa(i))
	}
	return r
}

// @lc code=end

