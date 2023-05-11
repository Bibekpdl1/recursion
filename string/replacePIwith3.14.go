package string

// "piyush" => "3.14yush"
func ReplacePI(str string) string {
	if len(str) == 0 || len(str) == 1 {
		return str
	}
	if str[0] == 'p' && str[1] == 'i' {
		smallResult := ReplacePI(str[2:])
		return "3.14" + smallResult
	}
	smallerResult := ReplacePI(str[1:])
	return string(str[0]) + smallerResult
}
