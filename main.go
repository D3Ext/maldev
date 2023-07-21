package main

/*

Example file with some "maldev" functions

*/

import (
  "fmt"

  "github.com/D3Ext/maldev/all"
)

func main(){
  all.PrintBanner("Maldev")
  fmt.Println("\nFile \"example.txt\" exists?:", all.Exists("example.txt"))
  fmt.Println("\"maldev\" as md5 hash: " + all.Md5([]byte("maldev")))
  fmt.Println("\"whoami\" command output: " + all.ExecuteCommand("whoami"))
  fmt.Println("Random 12 characters string: " + all.RandomString(12))
  fmt.Println("Internet connection?:", all.CheckInternet())
  fmt.Println("Find substring between two delimiters (AAAA BBBB CCCC):", all.FindSubstringBetween("AAAA BBBB CCCC", "AAAA", "CCCC"))
  home, _ := all.GetHome()
  fmt.Println("Current user home folder: " + home)
  example_slice := []string{"example1","example2","example3","example4"}
  fmt.Println("Check if slice contains a specific string -->", example_slice, "--> Contains \"example2\"?:", all.StringSliceContains(example_slice, "example2"))
}


