package all

import "github.com/D3Ext/maldev/src/files"

func Exists(src string) bool {
	return files.Exists(src)
}

func IsFile(src string) (bool, error) {
	return files.IsFile(src)
}

func IsDir(src string) (bool, error) {
	return files.IsDir(src)
}

func GetContent(src string) (string, error) {
	return files.GetContent(src)
}

func WriteContent(filename string, text string) error {
	return files.WriteContent(filename, text)
}

func Move(src string, dst string) error {
	return files.Move(src, dst)
}

func Copy(src string, dst string) error {
	return files.Copy(src, dst)
}

func Wipe(src string) error {
	return files.Wipe(src)
}

func Timestomp(src string, count int) error {
	return files.Timestomp(src, count)
}

