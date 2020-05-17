/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {
	max := len(nums)
    min := nums[0] 
	for i := range nums {
		if i+1 < max  && nums[i] > nums[i+1] {
			min = nums[i+1]
            break
		}
	}
    return min
}
// @lc code=end

