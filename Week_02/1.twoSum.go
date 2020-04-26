package Week_02

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		if index, ok := m[target-nums[i]]; ok {
			return []int{index, i}
		}
		m[nums[i]] = i
	}
	return []int{}
}
