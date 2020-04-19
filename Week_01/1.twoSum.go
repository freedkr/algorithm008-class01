package Week_01

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		if n, ok := m[target-nums[i]]; ok {
			return []int{n, i}
		}
		m[nums[i]] = i
	}
	return []int{}
}
