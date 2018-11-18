package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func GenerateRSAPair() (rsa.PrivateKey, rsa.PublicKey) {
	reader := rand.Reader

	key, err := rsa.GenerateKey(reader, 512)
	Must(err)

	return *key, key.PublicKey
}

func Hash(data string) []byte {
	md := sha256.New()
	md.Write([]byte(data))
	return md.Sum(nil)
}

func HashToString(data string) string {
	return hex.EncodeToString(Hash(data))
}

func Sign(key *rsa.PrivateKey, data string) string {
	reader := rand.Reader

	fmt.Println("yeet")
	fmt.Println(data)
	bytes, err := rsa.SignPKCS1v15(reader, key, crypto.SHA256, Hash(data))
	Must(err)

	return hex.EncodeToString(bytes)
}

func VerifySignature(pubKey *rsa.PublicKey, data string, signature []byte) bool {
	err := rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, Hash(data), signature)
	return err == nil
}

func GetBase64Key(key *rsa.PublicKey) string {
	bytes := x509.MarshalPKCS1PublicKey(key)
	return base64.StdEncoding.EncodeToString(bytes)
}
