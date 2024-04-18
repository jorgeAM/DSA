package frequencycounter

func ValidAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counter := [26]int{}

	for i := range s {
		counter[s[i]-'a']++
		counter[t[i]-'a']--
	}

	return counter == [26]int{}
}
