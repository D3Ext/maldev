# Examples

Here you can view some real examples in which the ***maldev*** functions are used

- Shellcode Encrypter
- Shellcode Loader (to run encrypted shellcode in memory)
- Windows hashes dumper (similar to secretsdump)
- Ransomware using AES
- Rewrite of metasploit **pattern_create.rb** (DeBruijn algorithm)
- DLL converter to shellcode

The encrypter and loader are designed to work together because the encrypter generates a random ***Initialization Vector*** and encrypts the shellcode using ***AES*** cipher with especified ***PSK*** and finally writes the shellcode to output file (the IV is appended at the beggining of the shellcode so the loader can get the dynamic IV to decrypt shellcode using PSK)


