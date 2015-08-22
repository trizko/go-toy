package strrev

// StringReverse returns the reverse of the string passed in.
func StringReverse(s string) string {
	var result []rune
	r := []rune(s)

	for i := len(r) - 1; i >= 0; i-- {
		result = append(result, r[i])
	}

	return string(result)
}
