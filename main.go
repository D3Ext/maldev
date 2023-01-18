package main

/*

Example file with some "maldev" functions

*/

import (
  "fmt"

  _ "github.com/D3Ext/maldev/shellcode"
  _ "github.com/D3Ext/maldev/process"
  _ "github.com/D3Ext/maldev/scanning"
  _ "github.com/D3Ext/maldev/revshell"
  _ "github.com/D3Ext/maldev/antiforensics"
  "github.com/D3Ext/maldev/system"
  "github.com/D3Ext/maldev/slices"
  "github.com/D3Ext/maldev/regex"
  "github.com/D3Ext/maldev/network"
  "github.com/D3Ext/maldev/misc"
  "github.com/D3Ext/maldev/exec"
  "github.com/D3Ext/maldev/crypto"
  "github.com/D3Ext/maldev/files"
  "github.com/D3Ext/maldev/logging"
)

func main(){
  logging.PrintBanner("Maldev")
  fmt.Println("\nFile \"example.txt\" exists?:", files.Exists("example.txt"))
  fmt.Println("Md5 encrypted \"maldev\": " + crypto.Md5Hash([]byte("maldev")))
  fmt.Println("\"id\" command output: " + exec.ExecuteCommand("id"))
  fmt.Println("Random string of length 12: " + misc.RandomString(12))
  fmt.Println("Internet connection?:", network.CheckInternet())
  fmt.Println("Find substring between two delimiters (AAAA BBBB CCCC):", regex.FindSubstringBetween("AAAA BBBB CCCC", "AAAA", "CCCC"))
  home, _ := system.GetHome()
  fmt.Println("Current user home folder: " + home)
  example_slice := []string{"example1","example2","example3","example4"}
  fmt.Println("Check if slice contains a string -->", example_slice, "--> Contains \"example2\"?:", slices.StringSliceContains(example_slice, "example2"))
}


