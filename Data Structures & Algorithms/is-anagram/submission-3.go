func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	countS := make(map[rune]int, 0)
	countT := make(map[rune]int, 0)
	for _, v := range s {
		countS[v]++
	}
	for _, v := range t {
		countT[v]++
	}

	for k, v := range countS {
		if countT[k] != v {
			return false
		}
	}
	return true
}
