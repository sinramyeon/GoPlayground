package StringUtil

import "strings"

// between ...
// Slice String and Get substring between two strings.
func between(value string, a string, b string) string {
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

// before ...
// Get substring before a string.
func before(value string, a string) string {
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}

// after ...
// Get substring after a string.
func after(value string, a string) string {
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:len(value)]
}

// TrimString ...
// Remove whitespace
func TrimString(s string) string {

	strings.Trim(s, " ")
	strings.TrimLeft(s, " ")
	strings.TrimPrefix(s, " ")
	strings.TrimRight(s, " ")
	strings.TrimSpace(s)

	return s

}
