package lib

func SubString(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)
	if start < 0 || end > length || start > end {
		return ""
	}
	if start == 0 && end == length {
		return source
	}
	return string(r[start:end])
}

func StringIsNullOrEmpty(str *string) bool {
	if str == nil {
		return true
	}
	if len(*str) == 0 {
		return true
	}
	return false
}
