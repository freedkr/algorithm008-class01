/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var res int
	var lastContentIndex int
	for _, num := range g {
		for ; lastContentIndex < len(s) && s[lastContentIndex] < num; lastContentIndex++ {
		}
		if lastContentIndex == len(s) {
			return res
		}
		res++
		lastContentIndex++
	}
	return res
}
// @lc code=end

