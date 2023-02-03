package system

/*

References:
https://github.com/Ne0nd0g/merlin-agent/blob/dev/commands/pipes_windows.go
https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findfirstfilea

*/

import (
  "golang.org/x/sys/windows"
)

func ListPipes() ([]string, error) { // This function uses Windows API to get all pipes names
  var all_pipes []string
  var pipePrefix = `\\.\pipe\` // Define pipe format
  var data windows.Win32finddata

  h, err := windows.FindFirstFile(
    windows.StringToUTF16Ptr(pipePrefix+"*"),
    &data,
  )

  if err != nil { // Handle error
    return all_pipes, err
  }
	defer windows.FindClose(h)

  for { // Iterate througth all pipes until error (last pipe)
    name := windows.UTF16ToString(data.FileName[:])
    all_pipes = append(all_pipes, name)

    if err := windows.FindNextFile(h, &data); err != nil {
      if err == windows.ERROR_NO_MORE_FILES {
        break
      }

      return all_pipes, err
    }
  }

  return all_pipes, nil
}

