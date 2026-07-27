func isPalindrome(s string) bool {
	s = strings.ToUpper(s)
	left, right := 0, len(s)-1
	for left < right {
		if string(s[left]) == " " || string(s[left]) == "?" || string(s[left]) == "," || string(s[left]) == "'" || string(s[left]) == "." || string(s[left]) == ":"{
			left++
			continue
		} else if string(s[right]) == " "  || string(s[right]) == "?" || string(s[right]) == "," || string(s[right]) == "'" || string(s[right]) == "." || string(s[right]) == ":"{
			right--
			continue
		}
		if s[left] == s[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}
