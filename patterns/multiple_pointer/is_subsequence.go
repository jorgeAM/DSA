package multiplepointer

func IsSubsequence(s, t string) bool {
	low, fast := 0, 0

	for low < len(s) && fast < len(t) {
		if s[low] == t[fast] {
			if low == len(s)-1 {
				return true
			}

			low++
		}

		fast++
	}

	return false
}
