package misc

/*

References:
https://www.techtarget.com/whatis/definition/bit-binary-digit

*/

func BytesToGb(num int) int {
	return (num / 1024 / 1024 / 1024)
}

func GbToBytes(num int) int {
	return (num * 1024 * 1024 * 1024)
}
