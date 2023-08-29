package all

import "github.com/D3Ext/maldev/regex"

func FindSubstringBetween(input string, first_delim string, second_delim string) []string {
	return regex.FindSubstringBetween(input, first_delim, second_delim)
}
