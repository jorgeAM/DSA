package recursion

func IsPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	start, end := 0, len(s)-1

	if s[start] != s[end] {
		return false
	}

	return IsPalindrome(s[start+1 : end])
}
