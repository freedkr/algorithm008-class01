/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
    result := [][]int{}
    remain := []int{}
    for i := 1; i <= n; i++ {
        remain = append(remain, i)
    }
    RecComb([]int{}, remain, k, &result)
    return result
}

func RecComb(current, remain []int, digits int, result *[][]int) {
    if len(current) == digits {
        newCurrent := make([]int, len(current))
        copy(newCurrent, current)
        *result = append(*result, newCurrent)
    } else {
        for i := 0; i < len(remain); i++ {
            RecComb(append(current, remain[i]), remain[i+1:], digits, result)
        }
    }
}
// @lc code=end

