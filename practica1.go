package main

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
)

func main() {
	var palabraAEncryptar = "fddfsadfsfdfgsd"

	hash, err := hashToText(palabraAEncryptar)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(hash)
}

func hashToText(hashThis string) (hashed string, err error) {
	if len(hashThis) == 0 {
		return "", errors.New("No input Supplied")
	}

	hasher := sha256.New()
	hasher.Write([]byte(hashThis))

	stringToSha256 := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return stringToSha256[:32], nil
}
