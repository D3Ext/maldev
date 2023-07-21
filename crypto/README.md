# Cryptography Examples

This category provides different options to encode/decode and encrypt/decrypt your data with a lot of algorithms

## AES

- Encryption and decryption example

```go
package main

import (
  "log"
  "fmt"

  "github.com/D3Ext/maldev/crypto"
)

func main() {
  plaintext := []byte("this is an example") // data to encrypt
  iv := []byte("1010101010101010") // initialization vector
  key := []byte("MySuperSecret32BitsLongPassword!") // pre-shared key

  ciphertext, err := crypto.AESEncrypt(plaintext, iv, key) // Encrypt data
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(ciphertext)

  decrypted, err := crypto.AESDecrypt(ciphertext, iv, key) // Decrypt data
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(decrypted)
}
```

## RC4

- Encryption and decryption example

```go
package main

import (
  "log"
  "fmt"

  "github.com/D3Ext/maldev/crypto"
)

func main() {
  plaintext := []byte("this is an example") // data to encrypt
  key := []byte("MySecretPassword") // pre-shared key

  ciphertext, err := crypto.Rc4Encrypt(plaintext, key) // Encrypt data
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(ciphertext)

  decrypted, err := crypto.Rc4Decrypt(ciphertext, key) // Decrypt data
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(decrypted)
}
```

## Bcrypt

- Hash data and verify hashes

```go
package main

import (
  "fmt"
  "log"

  "github.com/D3Ext/maldev/crypto"
)

func main() {
  plaintext := []byte("example") // Define bytes to hash

  hash, err := crypto.Bcrypt(plaintext)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(hash)

  check := crypto.VerifyBcrypt(hash, plaintext) // Compare hash with plaintext and return a bool
  fmt.Println(check) // true
}
```

## XOR

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/crypto"
)

func main() {
  plaintext := []byte("I am the best!")
  
  xored_text := crypto.Xor(plaintext, 'x') // func Xor(buf []byte, xorchar byte) ([]byte)
  fmt.Println(xored_text)
}
```

## ChaCha20

```go
package main

import (
  "fmt"
  "log"
  "github.com/D3Ext/maldev/crypto"
)

func main(){
  data := []byte("this is an example")
  psk := []byte("ThisIsMySuperSecret32Password123")

  fmt.Println("Data:", string(data))
  fmt.Println("Key:", psk)

  ciphertext, err := crypto.Chacha20Encrypt(data, psk) // Encrypt
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Ciphertext:", ciphertext)

  decoded, err := crypto.Chacha20Decrypt(ciphertext, psk) // Decrypt
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Decoded:", string(decoded))
}
```

## Elliptic Curve

- Encryption and decryption example

```go
package main

import (
  "fmt"
  "log"

  "github.com/D3Ext/maldev/crypto"
)

func main(){
  plaintext := []byte("I love malware!")
  key := []byte("MySecretKey")

  enc, err := crypto.EllipticCurveEncrypt(key, plaintext) // Encrypt bytes
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Encrypted data:", enc)

  dec, err := crypto.EllipticCurveDecrypt(key, enc) // Decrypt ciphertext
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Decrypted data:", dec)
}
```

## Rot13

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/crypto"
)

func main() {
  plaintext := "Example string, bla, bla, bla, bla"
  
  mod_text := crypto.Rot13(plaintext)
  fmt.Println(mod_text) // Output: Rknzcyr fgevat, oyn, oyn, oyn, oyn
}
```

## Rot47

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/crypto"
)

func main() {
  plaintext := "Example string, bla, bla, bla, bla"
  
  mod_text := crypto.Rot47(plaintext)
  fmt.Println(mod_text) // Output: "tI2>A=6 DEC:?8[ 3=2[ 3=2[ 3=2[ 3=2"
}
```

## Base64

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/crypto"
)

func main() {
  example := "This is an example"

  enc := crypto.B64E(example)
  fmt.Println(enc) // Output: "VGhpcyBpcyBhbiBleGFtcGxl"

  dec := crypto.B64D(enc)
  fmt.Println(dec) // Output: "This is an example"
}
```

## Base32

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/crypto"
)

func main() {
  example := "This is an example"

  enc := crypto.B32E(example)
  fmt.Println(enc) // Output: "KRUGS4ZANFZSAYLOEBSXQYLNOBWGK==="

  dec := crypto.B32D(enc)
  fmt.Println(dec) // Output: "This is an example"
}
```

## Hashing

- Examples for all supported hashes

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/crypto"
)

func main() {
  example := []byte("maldev")

  fmt.Println("Md5: " + crypto.Md5(example))
  fmt.Println("Sha1: " + crypto.Sha1(example))
  fmt.Println("Sha256: " + crypto.Sha256(example))
  fmt.Println("Sha512: " + crypto.Sha512(example))
}
```

## Verifying hashes

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/crypto"
)

func main(){
  md5hash := crypto.Md5([]byte("maldev"))
  md5_check := crypto.VerifyMd5(md5hash, "maldev")
  fmt.Println(md5_check)

  sha1hash := crypto.Sha1([]byte("maldev"))
  sha1_check := crypto.VerifySha1(sha1hash, "maldev")
  fmt.Println(sha1_check)

  // Same syntax with other hashes
  // .......
}
```

## Generate an IV

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/crypto"
)

func main() {
  rand_iv, _ := crypto.GenerateIV()
  fmt.Println(rand_iv)
}
```












