package crypto

import (
	"strings"
)

func Rot13(input string) string {
	result := make([]string, 0, len(input))
	for _, chr := range input {
		if 'a' <= chr && chr <= 'z' {
			chr = ((chr - 'a' + 13) % 26) + 'a'
		}
		if 'A' <= chr && chr <= 'Z' {
			chr = ((chr - 'A' + 13) % 26) + 'A'
		}
		result = append(result, string(chr))
	}
	output := strings.Join(result, "")
	return output
}
