package Week_01

func firstUniqChar(s string) byte {
	m := make([]int, 26)
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if m[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return byte(' ')
}
