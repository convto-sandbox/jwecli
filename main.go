package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
	"os"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwe"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("payloadは必須です。また、payloadは一つのみ設定してください")
		return
	}
	payload := []byte(args[1])

	privkey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Printf("failed to generate private key: %s", err)
		return
	}

	encrypted, err := jwe.Encrypt(payload, jwa.RSA1_5, &privkey.PublicKey, jwa.A128CBC_HS256, jwa.NoCompress)
	if err != nil {
		log.Printf("failed to encrypt payload: %s", err)
		return
	}

	fmt.Println(string(encrypted))
}
