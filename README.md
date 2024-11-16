<p align="center">
  <img src="https://raw.githubusercontent.com/D3Ext/maldev/main/static/gopher-rasta.png" width="130" heigth="60" alt="Gopher"/>
  <h1 align="center">Maldev</h1>
  <h4 align="center">Golang library for malware development</h4>
  <h6 align="center">Coded with 💙 by D3Ext</h6>
</p>

<p align="center">

  <a href="https://goreportcard.com/report/github.com/D3Ext/maldev">
      <img src="https://goreportcard.com/badge/github.com/D3Ext/maldev" alt="go report card">
  </a>

</p>

<p align="center">
  <a href="#features">Features</a> •
  <a href="#installation">Installation</a> •
  <a href="#usage">Usage</a> •
  <a href="#examples">Examples</a>
</p>

# Introduction

`maldev` aims to help malware developers, red teamers, cybersecurity enthusiasts or simply Golang developers. Most of the available features are not too hard to implement, but it is simply better to have multiple functions with just importing one package. This project is not finished yet but the official API is mostly stable. Anyway if you find a bug feel free to open an issue or create a pull-request which fixes it.

# Features

This are the different categories supported by `maldev`:

- [System](https://github.com/D3Ext/maldev/tree/main/src/system)
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
- [Cryptography](https://github.com/D3Ext/maldev/tree/main/src/crypto)
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
  - Triple DES
  - Compare hashes
- [Scanning](https://github.com/D3Ext/maldev/tree/main/src/scanning)
  - Ping an ip
  - Hostscan
  - Portscan
  - Enumerate all subdomains of a domain
  - Check if a domain uses http or https
  - Whois
  - Wappalyzer (identify technologies)
- [Logging](https://github.com/D3Ext/maldev/tree/main/src/logging)
  - Status functions
  - ASCII banners
  - Progress bars
  - Colors
- [Network](https://github.com/D3Ext/maldev/tree/main/src/network)
  - List all network interfaces (i.e. eth0)
  - Get info about an interface
  - List active ports with its info
  - Check internet connection
  - Get public ip
  - Download a file from URL
  - Get status code from URL
  - Send http POST request with custom data
- [Misc](https://github.com/D3Ext/maldev/tree/main/src/misc)
  - Generate random string
  - Generate random integer
  - Generate string of n length based on DeBruijn algorithm
  - Convert date to epoch
  - Convert epoch to date
  - Convert text to l33t
- [Red Team](https://github.com/D3Ext/maldev/tree/main/src/redteam)
  - Multiple shellcode injection techniques
  - Retrieve shellcode from file
  - Retrieve shellcode from remote url
  - Write shellcode to file
  - Convert DLL to shellcode (***sRDI***)
  - Multiple anti-sandboxing techniques
  - 3 different ways to dump system hashes
  - Steal token from PID (Impersonation)
- [Processes](https://github.com/D3Ext/maldev/tree/main/src/process)
  - List all process
  - Get process name by PID
  - Get PIDs of processes by name (i.e. firefox.exe)
- [Exec](https://github.com/D3Ext/maldev/tree/main/src/exec)
  - Execute bash commands
  - Execute powershell commands
  - Execute cmd commands
  - Execute command with Token
- [Working with slices](https://github.com/D3Ext/maldev/tree/main/src/slices)
  - Check if a slice contains a string
  - Check if a slice contains a string (insensitive)
  - Remove duplicates from []string
  - Remove duplicates from []int
  - Lowercase all characters from []string entries
  - Uppercase all characters from []string entries
- [Working with files](https://github.com/D3Ext/maldev/tree/main/src/files)
  - Check if file exists
  - Check if path is file
  - Check if path is dir
  - Copy a file or dir (recursive)
  - Get content of a file
  - Create a file with content
  - Wipe a file
  - Timestomping (change file timestamps)

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

Anyway if you want to use functions from an especific topic, you can do so like this:

> Example with cryptography
```go
import (
    "github.com/D3Ext/maldev/src/crypto"
)
```

# Examples

In every directory there is a **README.md** which contains at least one example of every defined function. You may also want to take a look at the [examples/](https://github.com/D3Ext/maldev/tree/main/examples) directory where you can find simple tools that use ***maldev***.

# Third party

These are all the 3rd-party used packages:

```
https://github.com/ryanuber/columnize
https://github.com/cakturk/go-netstat
https://github.com/C-Sto/gosecretsdump
https://github.com/C-Sto/BananaPhone
https://github.com/mitchellh/go-ps
https://github.com/elastic/go-sysinfo
https://github.com/fourcorelabs/wintoken
https://github.com/FourCoreLabs/EDRHunt
https://github.com/common-nighthawk/go-figure
```

# TODO

- ~~new features~~
- ~~redesigned project structure~~
- ~~some categories has been merged into the other ones~~
- ~~general improvement~~

# Disclaimer

Creator has no responsibility for any kind of illegal use of the project

# License

This project is under MIT license

Copyright © 2024, *D3Ext*


