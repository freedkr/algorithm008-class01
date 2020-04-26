package Week_02

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	sorted := make([]string, len(strs))
	for i := 0; i < len(strs); i++ {
		sorted[i] = sortString(strs[i])
	}

	m := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		_, ok := m[sorted[i]]
		if !ok {
			m[sorted[i]] = []string{strs[i]}
		} else {
			m[sorted[i]] = append(m[sorted[i]], strs[i])
		}
	}

	groups := make([][]string, 0)
	for _, v := range m {
		groups = append(groups, v)
	}
	return groups
}
func sortString(s string) string {
	charArr := strings.Split(s, "")
	sort.Strings(charArr)
	return strings.Join(charArr, "")
}
