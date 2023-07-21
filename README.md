<p align="center">
  <img src="https://raw.githubusercontent.com/D3Ext/maldev/main/assets/gopher-rasta.png" width="130" heigth="60" alt="Gopher"/>
  <h1 align="center">Maldev</h1>
</p>

<p align="center">
  <a href="#features">Features</a> •
  <a href="#installation">Installation</a> •
  <a href="#examples">Examples</a> •
  <a href="#third-party">Third-party</a> •
  <a href="#contributing">Contributing</a> •
  <a href="#disclaimer">Disclaimer</a>
</p>

# Introduction

`maldev` aims to help malware developers, red teamers and anyone who is interested in cybersecurity. It uses native Golang code and some other useful packages like [Hooka](https://github.com/D3Ext/Hooka) which I created to perform complex low-level red teaming stuff. The project isn't finished yet and the API may be unstable so your code may break in a future, sorry about this. Anyway if you find a bug feel free to open an issue or create a pull-request which fixes it

# Features

This are the different categories:

- [Cryptography](https://github.com/D3Ext/maldev/tree/main/crypto)
  - AES
  - RC4
  - Xor
  - Base32
  - Base64
  - Md5
  - Sha1
  - Sha256
  - Sha512
  - Rot13
  - Rot47
  - Bcrypt
  - Elliptic Curve
  - ChaCha20
  - Compare hashes
- [Network](https://github.com/D3Ext/maldev/tree/main/network)
  - List all interfaces
  - Get info about an interface
  - List active ports wih its info
  - Check internet connection
  - Get public ip
  - Download a file from URL
  - Get status code from URL
  - Send http POST request with custom data
- [Misc](https://github.com/D3Ext/maldev/tree/main/misc)
  - Generate random strings
  - Generate random integers
  - Convert dates to epoch format
  - Convert epoch to dates
  - Convert text to leet
- [Shellcode](https://github.com/D3Ext/maldev/tree/main/shellcode)
  - Tons of shellcode injection techniques
  - Retrieve shellcode from file
  - Retrieve shellcode from url
  - Write shellcode to file
  - Convert DLL to shellcode (***sRDI***)
- [Red Team](https://github.com/D3Ext/maldev/tree/main/redteam)
  - 3 different ways to dump system hashes
  - Steal token from PID (Impersonation)
  - Enable/disable Sticky Keys backdoor
  - Create malicious SCF on given path
- [Antiforensics](https://github.com/D3Ext/maldev/tree/main/antiforensics)
  - Wiping
  - Timestomping
- [Processes](https://github.com/D3Ext/maldev/tree/main/process)
  - List all process
  - Get process name by PID
  - Get list of processes by name (i.e. firefox.exe)
- [Exec](https://github.com/D3Ext/maldev/tree/main/exec)
  - Execute bash commands
  - Execute powershell commands
  - Execute cmd commands
  - Execute command with Token
- [System](https://github.com/D3Ext/maldev/tree/main/system)
  - Whoami
  - Get current dir
  - Get home dir
  - Get current user groups
  - Find installed useful software
  - List files and folders
  - Get environment variables
  - Get generic system information
  - Get SID and RID from windows system
  - Find installed AVs/EDRs
- [Scanning](https://github.com/D3Ext/maldev/tree/main/scanning)
  - Ping an ip
  - Hostscan
  - Portscan
  - Enumerate all subdomains of a domain
  - Check if a domain uses http or https
  - Whois
  - Wappalyzer (identify technologies)
- [Logging](https://github.com/D3Ext/maldev/tree/main/logging)
  - Status functions
  - ASCII banners
  - Progress bars
  - Colors
  - "log" and "fmt" wrappers
- [Working with slices](https://github.com/D3Ext/maldev/tree/main/slices)
  - Check if contains a string
  - Check if contains a string (insensitive)
  - Remove duplicates from []string
  - Remove duplicates from []int
  - Lowercase all characters from []string entries
- [Working with files](https://github.com/D3Ext/maldev/tree/main/files)
  - Check if file exists
  - Check if path is file
  - Check if path is dir
  - Copy a file or dir (recursive)
  - Get content of a file
  - Directly create a file with content

# Installation

Just execute this and it should be installed without problems:

```sh
go get -u https://github.com/D3Ext/maldev
```

# Usage

To import all the functions at the same time do it like this:

```go
import (
    maldev "github.com/D3Ext/maldev/all"
)
```

Anyway if you want to use functions from an especific topic, you can do it like this:

> Example with cryptography
```go
import "github.com/D3Ext/maldev/crypto"
```

# Examples

In every directory there is a **README.md** which contains at least one example of every defined function, if you don't have enough creativity I encourage you to check out the `examples/` directory where I've developed some good examples which use ***maldev*** functions like a simple ransomware, a shellcode loader and much more

# TODO

:black_square_button: Kerberos protocol implementation

:black_square_button: Publish official package documentation (pkg.go.dev)

:black_square_button: Stable progress bars

:black_square_button: Shikata Ga Nai polymorphic encoder

# Third party

As said above I have tried to implement all functions from scratch but I have also used some external packages:

[columnize](https://github.com/ryanuber/columnize) to create tables and columns easily

[go-netstat](https://github.com/cakturk/go-netstat) to retrieve info about local ports

[gosecretsdump](https://github.com/C-Sto/gosecretsdump) used to dump hashes from SAM, NTDS and SYSTEM

[BananaPhone](https://github.com/C-Sto/BananaPhone) to perform **CreateRemoteThread** shellcode injection technique

[go-ps](https://github.com/mitchellh/go-ps) used to work with linux processes

[go-sysinfo](https://github.com/elastic/go-sysinfo) useful to get system information mainly for Windows

[wintoken](https://github.com/fourcorelabs/wintoken) used to interact with Windows API and getting privileges info

[EDRHunt](https://github.com/FourCoreLabs/EDRHunt) used in `system/` to look for installed AVs/EDRs

[go-figure](https://github.com/common-nighthawk/go-figure) to create banners easily

# Contributing

See [CONTRIBUTING.md](https://github.com/D3Ext/maldev/blob/main/CONTRIBUTING.md)

# Disclaimer

Creator has no responsibility for any kind of:

- Illegal use of the project.
- Law infringement by third parties and users.
- Malicious act, capable of causing damage to third parties, promoted by the user through this software.

# License

This project is under MIT license

Copyright © 2023, *D3Ext*


