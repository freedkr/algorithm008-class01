/*
 * @lc app=leetcode.cn id=874 lang=golang
 *
 * [874] 模拟行走机器人
 */

// @lc code=start
func robotSim(commands []int, obstacles [][]int) int {
	obstaclesMap := make(map[[2]int]bool, len(obstacles))
	for _, obstacle := range obstacles {
		obstaclesMap[[2]int{obstacle[0], obstacle[1]}] = true
	}
	// "北", "东", "南", "西"
	directions := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var curDirection, x, y int
	var res int
	for _, command := range commands {
		switch command {
		case -2:
			curDirection = (curDirection + 3) % 4
		case -1:
			curDirection = (curDirection + 1) % 4
		default:
			for i := 0; i < command; i++ {
				x += directions[curDirection][0]
				y += directions[curDirection][1]
				if _, ok := obstaclesMap[[2]int{x, y}]; ok {
					x -= directions[curDirection][0]
					y -= directions[curDirection][1]
				}
				res=max(res, x*x+y*y)
			}
		}
	}
	return res
}

func max(res int, i int) int {
	if res <i {
		return i
	}
	return res
}
// @lc code=end

