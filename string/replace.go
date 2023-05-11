package string

func Replace(str string, from, to rune) string {
	if len(str) == 0 {
		return str
	}
	if str[0] == byte(from) {
		return string(to) + Replace(str[1:], from, to)
	}
	return string(str[0]) + Replace(str[1:], from, to)
}
