package auth

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func RandomState() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}
