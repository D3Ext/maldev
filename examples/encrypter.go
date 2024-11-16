package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	// Maldev packages
	"github.com/D3Ext/maldev/src/crypto"
	"github.com/D3Ext/maldev/src/files"
	l "github.com/D3Ext/maldev/src/logging"
	"github.com/D3Ext/maldev/src/redteam"
)

func main() {
	if len(os.Args) < 4 { // Check given parameters
		l.Error("Parameter(s) missing!")
		fmt.Println("Usage: ./encrypter <password> <raw shellcode file> <output file>")
		os.Exit(0)
	}

	if len(os.Args[1]) != 32 { // Check password length
		l.Error("Password length is invalid, it must be 32 characters long")
		os.Exit(0)
	}

	if (files.Exists(os.Args[2]) == false) { // Check if raw shellcode file exists
		l.Error("Provided raw shellcode file doesn't exists!")
		os.Exit(0) // Exit
	}

	raw_shellcode, err := redteam.GetShellcodeFromFile(os.Args[2]) // Get shellcode as bytes from input file
	if err != nil {                                          // Handle error
		log.Fatal(err)
	}

	l.Success("Shellcode obtained from file") // Some status info
	time.Sleep(300 * time.Millisecond)       // Little sleep to let user read stdout
	l.Info("Encrypting shellcode with given 32 characters PSK")
	time.Sleep(300 * time.Millisecond)

	iv, err := crypto.GenerateIV()
  if err != nil {
    log.Fatal(err)
  }

	l.Info("Pre-Shared Key: " + string(os.Args[1][0:6]) + strings.Repeat("*", len(os.Args[1])-6))
	time.Sleep(200 * time.Millisecond)
	l.Info("Generated IV: ")
	fmt.Println(iv)
	enc_shellcode, err := crypto.AESEncrypt(raw_shellcode, iv, []byte(os.Args[1])) // Use own function to encrypt shellcode with static IV
	if err != nil {                                                                // Handle error
		log.Fatal(err)
	}

	time.Sleep(400 * time.Millisecond) // Add little sleep to let user read
	l.Info("Writing encrypted shellcode to output file...")
	time.Sleep(300 * time.Millisecond)

  err = files.WriteContent(os.Args[3], string(iv) + string(enc_shellcode))
  if err != nil {
    log.Fatal(err)
  }

	l.Success("Encrypted shellcode stored as " + os.Args[3])

}
