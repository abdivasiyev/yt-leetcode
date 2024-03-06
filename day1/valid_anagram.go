package day1

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var m = make(map[rune]int, len(s))

	for _, c := range s {
		if v, ok := m[c]; !ok {
			m[c] = 1
		} else {
			m[c] = v + 1
		}
	}

	for _, c := range t {
		if v, ok := m[c]; !ok {
			return false
		} else if v > 0 {
			m[c] = v - 1
		} else {
			return false
		}
	}

	return true
}
