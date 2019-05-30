package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwe"
)

var (
	help        bool
	payload     string
	plainKeyAlg string
	plainConAlg string
)

func init() {
	flag.BoolVar(&help, "-help", false, "Print help")
	flag.BoolVar(&help, "h", false, "Shorthand help")
	flag.StringVar(&payload, "-payload", "", "Payload")
	flag.StringVar(&payload, "p", "", "Shorthand payload")
	flag.StringVar(&plainKeyAlg, "-key-alg", "", "Key encryption algorithm. supports:\n  RSA1_5\n  RSA_OAEP\n  RSA_OAEP_256\n  ECDH_ES_A128KW\n  ECDH_ES_A102KW\n  ECDH_ES_A256KW")
	flag.StringVar(&plainKeyAlg, "ka", "", "Shorthand key-alg")
	flag.StringVar(&plainConAlg, "-content-alg", "", "Content encryption algorithm. supports:\n  A128CBC_HS256")
	flag.StringVar(&plainConAlg, "ca", "", "Shorthand content-alg")
}

func parseKeyAlg(plainKeyAlg string) (jwa.KeyEncryptionAlgorithm, error) {
	switch strings.ToUpper(plainKeyAlg) {
	case "RSA1_5":
		return jwa.RSA1_5, nil
	case "RSA_OAEP":
		return jwa.RSA_OAEP, nil
	case "RSA_OAEP_256":
		return jwa.RSA_OAEP_256, nil
	case "ECDH_ES_A128KW":
		return jwa.ECDH_ES_A128KW, nil
	case "ECDH_ES_A192KW":
		return jwa.ECDH_ES_A192KW, nil
	case "ECDH_ES_A256KW":
		return jwa.ECDH_ES_A256KW, nil
	default:
		return "", errors.New(fmt.Sprintf("not support atgorithm. %s", plainKeyAlg))
	}
}

func genKey(keyAlg jwa.KeyEncryptionAlgorithm) (interface{}, error) {
	switch keyAlg {
	case jwa.RSA1_5, jwa.RSA_OAEP, jwa.RSA_OAEP_256:
		privKey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			fmt.Printf("failed to getenerate rsa private key: %s\n", err)
			return nil, err
		}
		return &privKey.PublicKey, nil
	case jwa.ECDH_ES_A128KW, jwa.ECDH_ES_A192KW, jwa.ECDH_ES_A256KW:
		privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			fmt.Printf("failed to generate ecdh private key: %s", err)
			return nil, err
		}
		return &privKey.PublicKey, nil
	default:
		return nil, errors.New(fmt.Sprintf("not support atgorithm. %s", keyAlg))
	}
}

func parseConAlg(plainConAlg string) (jwa.ContentEncryptionAlgorithm, error) {
	return jwa.A128CBC_HS256, nil
}

func main() {
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	if payload == "" || plainKeyAlg == "" || plainConAlg == "" {
		fmt.Println("-p and -ka and -ca is required.\nplease check `jwtcli -h`.")
		os.Exit(1)
	}

	keyAlg, err := parseKeyAlg(plainKeyAlg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	key, err := genKey(keyAlg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conAlg, err := parseConAlg(plainConAlg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	encrypted, err := jwe.Encrypt([]byte(payload), keyAlg, key, conAlg, jwa.NoCompress)
	if err != nil {
		log.Printf("failed to encrypt payload: %s", err)
		return
	}

	fmt.Println(string(encrypted))
}
