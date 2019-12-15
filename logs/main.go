package logs

import (
	"crypto/md5"
	"encoding/hex"
	"log"
)

type SQL struct {
	Host, Name, User, Password string
	Port                       int
}

func createHash(key string) []byte {
	hasher := md5.New()
	if _, err := hasher.Write([]byte(key)); err != nil {
		log.Fatal(err)
	}
	return []byte(hex.EncodeToString(hasher.Sum(nil)))
}
