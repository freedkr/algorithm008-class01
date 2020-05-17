/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
    length, end, maxPosition, steps := len(nums), 0, 0, 0
    for i := 0; i < length - 1; i++ {
        maxPosition = max(maxPosition, i + nums[i])
        if i == end {
            end = maxPosition
            steps++
        }
    }
    return steps
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
// @lc code=end

