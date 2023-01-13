package files

/*

References:
https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
https://www.tutorialspoint.com/how-to-check-if-a-file-exists-in-golang

type FileInfo interface {   // this struct is returned in os.Stat()
	Name() string             // base name of the file
	Size() int64              // length in bytes for regular files; system-dependent for others
	Mode() FileMode           // file mode bits
	ModTime() time.Time       // modification time
	IsDir() bool              // abbreviation for Mode().IsDir()
	Sys() any                 // underlying data source (can return nil)
}

*/

import (
  "os"
  "io/ioutil"
)

func Exists(src string) (bool) { // Check if a file or directory exists and return true or false
  _, err := os.Stat(src)
  if os.IsNotExist(err) { // Check if err shows that file doesn't exist
    return false
  }
  return true
}

func IsFile(src string) (bool, error) { // Check if especified path is a file
  info, err := os.Stat(src)
  if err != nil {
    return false, err
  }

  if info.IsDir() {
    return false, nil
  } else {
    return true, nil
  }
}

func IsDir(src string) (bool, error) { // Check if especified path is a directory
  info, err := os.Stat(src)
  if err != nil {
    return false, err
  }

  if info.IsDir() {
    return true, nil
  } else {
    return false, nil
  }
}

func GetFileContent(src string) (string, error) {
  byte_content, err := ioutil.ReadFile(src)
  if err != nil {
    return "", err
  }

  return string(byte_content), nil
}

func Copy(src string, dst string) (error) { // Copy file or directory (recursive)
  check, err := IsFile(src)
  if err != nil {
    return err
  }

  if (check == true) { // Enter here if especified source is a file
    file_bytes, err := ioutil.ReadFile(src)
    if err != nil {
      return err
    }

    err = ioutil.WriteFile(dst, file_bytes, 0644)
    if err != nil {
      return err
    }

  } else if (check == false) { // Enter here if especified source is a directory
    src_info, err := os.Stat(src)
    if err != nil {
      return err
    }
    
    err = os.MkdirAll(dst, src_info.Mode())
    if err != nil {
      return err
    }

    directory, _ := os.Open(src)
    objects, err := directory.Readdir(-1)
    
    for _, obj := range objects { // Iterate over files and dirs
      srcfilepointer := src + "/" + obj.Name()
      dstfilepointer := dst + "/" + obj.Name()

      if obj.IsDir() {
        err = Copy(srcfilepointer, dstfilepointer)
        if err != nil {
          return err
        }
      } else {
        file_bytes, err := ioutil.ReadFile(srcfilepointer)
        if err != nil {
          return err
        }
        err = ioutil.WriteFile(dstfilepointer, file_bytes, 0644)
        if err != nil {
          return err
        }
      }
    }
  }
  return nil
}


