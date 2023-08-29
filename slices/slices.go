package slices

/*

References:
https://stackoverflow.com/questions/10485743/contains-method-for-a-slice
https://freshman.tech/snippets/go/check-if-slice-contains-element/

*/

import (
	"strings"
)

func StringSliceContains(src []string, str_to_check string) bool {
	for _, entry := range src {
		if entry == str_to_check {
			return true
		}
	}
	return false
}

func IntSliceContains(src []int, int_to_check int) bool {
	for _, entry := range src {
		if entry == int_to_check {
			return true
		}
	}
	return false
}

func SliceContainsInsensitive(src []string, str_to_check string) bool {
	for _, entry := range src {
		if strings.ToLower(entry) == strings.ToLower(str_to_check) {
			return true
		}
	}
	return false
}

func ToLowercase(src []string) []string {
	var mod_src []string
	for _, entry := range src {
		mod_src = append(mod_src, strings.ToLower(entry))
	}

	return mod_src
}
