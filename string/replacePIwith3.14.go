package string

// "piyush" => "3.14yush"
func ReplaceString(str string) string {
	if len(str) == 0 || len(str) == 1 {
		return str
	}
	if str[0] == 'p' && str[1] == 'i' {
		smallResult := ReplaceString(str[2:])
		return "3.14" + smallResult
	}
	smallerResult := ReplaceString(str[1:])
	return string(str[0]) + smallerResult
}
