package gox

// NotEmpty strings
func NotEmpty(str ...string) bool {
	return !Empty(str...)
}

// Empty strings
func Empty(str ...string) bool {
	if str == nil || len(str) == 0 {
		return true
	}
	for _, one := range str {
		if one == "" {
			return true
		}
	}
	return false
}

// DefaultIfEmpty the string
func DefaultIfEmpty(str string, defaultStr string) string {
	if Empty(str) {
		return defaultStr
	}
	return str
}
