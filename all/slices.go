package all

import "github.com/D3Ext/maldev/src/slices"

func StringSliceContains(src []string, str_to_check string) bool {
	return slices.StringSliceContains(src, str_to_check)
}

func IntSliceContains(src []int, int_to_check int) bool {
	return slices.IntSliceContains(src, int_to_check)
}

func StringSliceContainsInsensitive(src []string, str_to_check string) bool {
	return slices.StringSliceContainsInsensitive(src, str_to_check)
}

func ToLowercase(src []string) []string {
	return slices.ToLowercase(src)
}

func RemoveDuplicatesStr(strslice []string) []string {
	return slices.RemoveDuplicatesStr(strslice)
}

func RemoveDuplicatesInt(intslice []int) []int {
	return slices.RemoveDuplicatesInt(intslice)
}
