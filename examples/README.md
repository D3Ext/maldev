# Examples

Here you can take a look at some real examples in which the ***maldev*** functions are used to perform specific tasks:

- Shellcode Encrypter
- Shellcode Loader (to run encrypted shellcode in memory)
- Sandbox detector
- Windows hashes dumper (similar to secretsdump)
- Ransomware using AES
- Rewrite of metasploit **pattern_create.rb** (DeBruijn algorithm)
- DLL converter to shellcode

The encrypter and loader are designed to work together because the encrypter generates a random ***Initialization Vector*** and encrypts the shellcode using ***AES*** cipher with especified ***PSK*** and finally writes the shellcode to output file (the IV is appended at the beggining of the shellcode so the loader can get the dynamic IV to decrypt shellcode using PSK)

# Demo

Here you can see how this examples are, they've been made to be as much descriptive as posible using ***logging*** functions, banners and colors

- Shellcode Encrypter

<img src="https://raw.githubusercontent.com/D3Ext/maldev/main/assets/encrypter.png">

- Shellcode Loader

<img src="https://raw.githubusercontent.com/D3Ext/maldev/main/assets/loader.png">

- Sandbox detector

<img src="https://raw.githubusercontent.com/D3Ext/maldev/main/assets/go-sandbox.png">

- pattern_create.rb

<img src="https://raw.githubusercontent.com/D3Ext/maldev/main/assets/pattern_create.png">

- DLL converter

<img src="https://raw.githubusercontent.com/D3Ext/maldev/main/assets/dll_to_shellcode.png">


