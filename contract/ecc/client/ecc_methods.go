package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"math/big"

	"encoding/base64"
	"encoding/hex"
	"encoding/json"
)

func hashString(str string) string {
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}

func hashReading(reading MeterReading) string {
	bReading, _ := json.Marshal(reading)
	hash := sha256.Sum256(bReading)
	return hex.EncodeToString(hash[:])
}

func signPayload(pv *ecdsa.PrivateKey, hash string) (string, string, error) {
	r, s, err := ecdsa.Sign(rand.Reader, pv, []byte(hash))
	if err != nil {
		return "", "", err
	}
	return bigIntToString(r), bigIntToString(s), nil
}

func bigIntToString(i *big.Int) string {
	b, _ := i.MarshalText()
	return base64.StdEncoding.EncodeToString(b)
}
