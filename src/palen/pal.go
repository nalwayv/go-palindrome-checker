package palen

import "strings"

func strip(msg string) string {
	nums := []rune(msg)

	result := make([]string, 0)

	for _, val := range nums {
		// upper and lower case characters
		if (val >= 97 && val < 123) || (val >= 65 && val < 91) {

			// convert upper to lower
			if val >= 65 && val < 91 {
				val = val | 32
			}

			result = append(result, string(val))
		}
	}

	return strings.Join(result, "")
}

// IsPalendrome ...
func IsPalendrome(msg string) bool {
	pal := strip(msg)
	for start, end := 0, len(pal)-1; start < end; start, end = start+1, end-1 {
		if pal[start] != pal[end] {
			return false
		}
	}
	return true
}
