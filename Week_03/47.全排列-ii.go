/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	bites := make([]bool, len(nums))
	var rec func(tmpList []int)
	rec = func(tmpList []int) {
		if len(tmpList) == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, tmpList)
			res = append(res, tmp)
		}
		for i := 0; i < len(nums); i++ {
			if bites[i] {
				continue
			}
			bites[i] = true
			tmpList = append(tmpList, nums[i])
			rec(tmpList)
			bites[i] = false
			tmpList = tmpList[:len(tmpList)-1]
			for ; i+1 < len(nums) && nums[i] == nums[i+1]; i++ {
			}
		}
	}
	rec([]int{})
	return res
}
// @lc code=end

