# maldev
A lightweight Golang library for malware development

# Introduction

`maldev` aims to help malware developers, red teamers and anyone who is interested in cybersecurity. It uses native Golang code at 100% and it tries to do all operations from scratch without almost any third-party package. If you have any idea feel free to contribute or if you want to contact me you can do it through Discord (***D3Ext#5965***)

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
- Networking
  - List all interfaces
  - Get info about an interface
- Misc
  - Generate random strings
  - Generate random integers
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
- Logging
  - Status functions
  - ASCII banners
  - Progress bars
  - Colors
- Working with slices
  - Check if contains a string
  - Check if contains a string (insensitive)
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

In every directory there is a README.md which contains at least one example of defined functions

# Third party

# Disclaimer

Creator isn't in charge of any and has no responsibility for any kind of:

- Unlawful or illegal use of the project.
- Legal or Law infringement (acted in any country, state, municipality, place) by third parties and users.
- Act against ethical and / or human moral, ethic, and peoples and cultures of the world.
- Malicious act, capable of causing damage to third parties, promoted or distributed by third parties or the user through this software.

