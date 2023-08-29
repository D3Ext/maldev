package slices

/*

References:
https://stackoverflow.com/questions/66643946/how-to-remove-duplicates-strings-or-int-from-slice-in-go
https://www.geeksforgeeks.org/how-to-remove-duplicate-values-from-slice-in-golang/

*/

func RemoveDuplicatesStr(strslice []string) []string { // Remove duplicates from string slice
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range strslice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func RemoveDuplicatesInt(intslice []int) []int { // Remove duplicates from int slice
	keys := make(map[int]bool)
	list := []int{}

	for _, entry := range intslice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
