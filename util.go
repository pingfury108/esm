package main

func SubString(s string, i, l int) string {
	s_len := len(s)
	switch {
	case i <= s_len && s_len > l:
		return s[i:l]
	case i <= s_len && l > s_len:
		return s[i:]
	default:
		return s
	}
}
