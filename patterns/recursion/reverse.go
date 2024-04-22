package recursion

func Reverse(s string) string {
	if len(s) <= 1 {
		return s
	}

	return Reverse(s[1:]) + string(s[0])
}
