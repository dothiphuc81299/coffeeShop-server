package format

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func NonAccentVietnamese(str string) string {
	if str != "" {
		str = strings.ToLower(str)
		str = replaceStringWithRegex(str, `đ`, "d")
		t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
		result, _, _ := transform.String(t, str)
		result = replaceStringWithRegex(result, `[^a-zA-Z0-9\s]`, "")

		return result
	}
	return ""
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}

func replaceStringWithRegex(src string, regex string, replaceText string) string {
	reg := regexp.MustCompile(regex)
	return reg.ReplaceAllString(src, replaceText)
}
