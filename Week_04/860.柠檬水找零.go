/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */

// @lc code=start
func lemonadeChange(bills []int) bool {
    x5, x10 := 0, 0
    for i := range bills {
        switch bills[i] {
            case 5:
                x5++
            case 10:
                if x5 == 0 {
                    return false
                }
                x5--
                x10++
            case 20:
            if x10 > 0 && x5 >0 {
				x10--
                x5--
            } else if  x5 > 2 {
            	x5 -=3
            } else {
                return false
            }
        }
    }
    return true
}
// @lc code=end

