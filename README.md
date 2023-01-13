<p align="center">
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

`maldev` aims to help malware developers, red teamers and anyone who is interested in cybersecurity. It uses native Golang code at 100% and it tries to do all operations from scratch without almost any third-party package. The project isn't finished yet and the API may be unstable so your malware may break in a near future.

# Features

This are the different categories:

- [Cryptography](https://github.com/D3Ext/maldev/tree/main/crypto)
  - AES
  - RC4
  - Xor
  - Base64
  - Md5
  - Sha1
  - Sha256
  - Sha512
  - Rot13
  - Rot47
  - Bcrypt
  - Elliptic Curve
- [Network](https://github.com/D3Ext/maldev/tree/main/network)
  - List all interfaces
  - Get info about an interface
  - List active ports wih its info
  - Check internet connection
  - Get public ip
  - Download a file from URL
  - Get status code from URL
- [Misc](https://github.com/D3Ext/maldev/tree/main/misc)
  - Generate random strings
  - Generate random integers
  - Convert dates to epoch format
  - Convert epoch to dates
  - Convert text to leet
- [Shellcode](https://github.com/D3Ext/maldev/tree/main/shellcode)
  - Process Injection techniques
  - Extra utils
- [Antiforensics](https://github.com/D3Ext/maldev/tree/main/antiforensics)
  - Wiping
  - Timestomping
- [Processes](https://github.com/D3Ext/maldev/tree/main/process)
  - List all process
  - Get process name by PID
  - Get list of processes by name
- [System](https://github.com/D3Ext/maldev/tree/main/system)
  - Whoami
  - Get current dir
  - Get home dir
  - Get current user groups
  - Find installed useful software
  - List files and folders
  - Get environment variables
  - Get system information
  - Find installed AVs/EDRs
- [Scanning](https://github.com/D3Ext/maldev/tree/main/scanning)
  - Ping an ip
  - Hostscan
  - Portscan
  - Enumerate all subdomains from domain
  - Check if a domain uses http or https
- [Logging](https://github.com/D3Ext/maldev/tree/main/logging)
  - Status functions
  - ASCII banners
  - Progress bars
  - Colors
- [Working with slices](https://github.com/D3Ext/maldev/tree/main/slices)
  - Check if contains a string
  - Check if contains a string (insensitive)
  - Remove duplicates from []string
  - Remove duplicates from []int
- [Working with files](https://github.com/D3Ext/maldev/tree/main/files)
  - Check if file exists
  - Check if path is file
  - Check if path is dir
  - Copy a file or dir (recursive)

# Installation

Just execute this and it should be installed without problems

```sh
go get -u https://github.com/D3Ext/maldev
```

# Examples

In every directory there is a **README.md** which contains at least one example of every defined function, if you don't have enough creativity I encourage you to check out the `examples/` directory where I've developed some good examples which use ***maldev*** functions

# TODO

:black_square_button: Kerberos protocol implementation

:black_square_button: Windows token impersonation

:black_square_button: More system enumeration functions

# Third party

As said above I have tried to implement all functions from scratch but I have also used some external packages:

[columnize](https://github.com/ryanuber/columnize) to create tables and columns easily

[go-ps](https://github.com/mitchellh/go-ps) to work with linux processes

[go-sysinfo](https://github.com/elastic/go-sysinfo) useful to get system information mainly for Windows

[wintoken](https://github.com/fourcorelabs/wintoken) used to interact with Windows API and getting privileges info

[EDRHunt](https://github.com/FourCoreLabs/EDRHunt) used in "system" part to look for installed AVs/EDRs

[go-figure](https://github.com/common-nighthawk/go-figure) to create banners easily

# Contributing

Do you want to contribute with any interesting idea?

`1` Open an issue to discuss your idea

`2` Fork the repo

`3` Create a branch

`4` Commit your changes

`5` Push to the branch

`6` Create a new pull request

***New features and bugs reports are welcome***

# TL;DR

There are some functions that I have not implemented because they are really easy to do with `os` package like `os.Remove()` which removes a file.

# Disclaimer

Creator isn't in charge of any and has no responsibility for any kind of:

- Unlawful or illegal use of the project.
- Legal or Law infringement (acted in any country, state, municipality, place) by third parties and users.
- Act against ethical and / or human moral, ethic, and peoples of the world.
- Malicious act, capable of causing damage to third parties, promoted or distributed by third parties or the user through this software.

