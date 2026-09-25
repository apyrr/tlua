package stringutil

// ToUpperLua and ToLowerLua map a Lua string the way string.upper and
// string.lower do in the default C locale: byte by byte, changing only the ASCII
// letters. Every other byte, including each byte of a multi-byte UTF-8 sequence,
// is kept as is.
func ToUpperLua(s string) string {
	return mapASCIILetters(s, 'a', 'z', 'A'-'a')
}

func ToLowerLua(s string) string {
	return mapASCIILetters(s, 'A', 'Z', 'a'-'A')
}

func mapASCIILetters(s string, lo, hi byte, delta int) string {
	var buf []byte
	for i := range len(s) {
		if ch := s[i]; lo <= ch && ch <= hi {
			if buf == nil {
				buf = []byte(s)
			}
			buf[i] = byte(int(ch) + delta)
		}
	}
	if buf == nil {
		return s
	}
	return string(buf)
}
