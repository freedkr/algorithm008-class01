package Week_02

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counter := make(map[rune]int, 0)
	for _, sWord := range s {
		counter[sWord]++
	}
	for _, tWord := range t {
		if _, ok := counter[tWord]; ok {
			counter[tWord]--
		}
	}
	for _, v := range counter {
		if v != 0 {
			return false
		}
	}
	return true
}
