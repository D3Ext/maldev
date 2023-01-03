

<p align="center">
  <a href="#features">Features</a> •
  <a href="#installation">Installation</a> •
  <a href="#examples">Examples</a> •
  <a href="#third-party">Third-party</a> •
  <a href="#disclaimer">Disclaimer</a>
</p>

# Introduction

`maldev` aims to help malware developers, red teamers and anyone who is interested in cybersecurity. It uses native Golang code at 100% and it tries to do all operations from scratch without almost any third-party package. The project isn't finished yet and the API may be unstable so your malware may break in a near future.

# Features

This are the different categories:

- Cryptography
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
- Network
  - List all interfaces
  - Get info about an interface
  - List active ports wih its info
  - Check internet connection
  - Get public ip
  - Download a file from URL
  - Get status code from URL
- Misc
  - Generate random strings
  - Generate random integers
  - Convert dates to epoch format
  - Convert epoch to dates
  - Convert text to leet
- Shellcode
  - Process Injection techniques
  - Extra utils
- Antiforensics
  - Wiping
  - Timestomping
- Processes
  - List all process
  - Interact with processes
- System
  - Whoami
  - Get current dir
  - Get system information
  - Find installed AVs/EDRs
- Scanning
  - Ping an ip
  - Hostscan
  - Portscan
  - Enumerate all subdomains from domain
  - Check if a domain uses http or https
- Logging
  - Status functions
  - ASCII banners
  - Progress bars
  - Colors
- Working with slices
  - Check if contains a string
  - Check if contains a string (insensitive)
  - Remove duplicates from []string
  - Remove duplicates from []int
- Working with files
  - Check if file exists
  - Check if path is file
  - Check if path is dir
  - Copy a file or dir (recursive)

# Installation

Just execute one command and it should be installed without errors

```sh
go get -u https://github.com/D3Ext/maldev
```

# Examples

In every directory there is a **README.md** which contains at least one example of every defined function, if you don't have enough creativity I encourage you to check out the `examples/` directory where I've developed some good examples which use ***maldev*** functions

# Third party

As said above I have tried to implement all functions from scratch but I have also used some external packages:

[columnize](https://github.com/ryanuber/columnize) to create tables and columns easily

[go-ps](https://github.com/mitchellh/go-ps) to work with linux processes

[go-sysinfo](https://github.com/elastic/go-sysinfo)

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

# Disclaimer

Creator isn't in charge of any and has no responsibility for any kind of:

- Unlawful or illegal use of the project.
- Legal or Law infringement (acted in any country, state, municipality, place) by third parties and users.
- Act against ethical and / or human moral, ethic, and peoples of the world.
- Malicious act, capable of causing damage to third parties, promoted or distributed by third parties or the user through this software.

