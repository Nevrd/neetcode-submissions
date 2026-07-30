func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	countS, countT := make(map[rune]int, 0), make(map[rune]int, 0)
	for _, v := range s {
		countS[v]++
	}
	for _, v := range t {
		countT[v]++
	}

	for k := range countS {
		if countS[k] != countT[k] {
			return false
		}
	}
	return true
}
