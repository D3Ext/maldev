package system

import (
	"fmt"
	"github.com/ryanuber/columnize"
	"io/ioutil"
)

func List(dir_to_list string) (string, error) { // Similar output of "dir" command using powershell but using native golang
	var output []string
	var ls_out string
	output = append(output, "Size ||| Type ||| Date ||| Name")
	output = append(output, "---- ||| ---- ||| ---- ||| ----")

	if dir_to_list == "" {
		dir_to_list = "."
	}

	files, err := ioutil.ReadDir(dir_to_list)
	if err != nil {
		return "", err
	}

	for _, f := range files {
		if f.IsDir() {
			ls_out = fmt.Sprintf("%v ||| Dir ||| %v ||| %v", f.Size(), f.ModTime(), f.Name())
			output = append(output, ls_out)
		} else {
			ls_out = fmt.Sprintf("%v ||| File || %v ||| %v", f.Size(), f.ModTime(), f.Name())
			output = append(output, ls_out)
		}
	}

	config := columnize.DefaultConfig()
	config.Delim = "|||"
	config.Glue = "  "
	config.Prefix = ""
	config.Empty = ""
	config.NoTrim = false
	columns_out := columnize.Format(output, config)

	return columns_out, nil
}
