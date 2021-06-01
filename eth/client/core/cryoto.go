package core

import (
	"crypto/sha256"
	"encoding/json"
)

// HashBytes sums a given byte slice
func HashBytes(b []byte) [32]byte {
	return sha256.Sum256(b)
}

// HashReading calculate the hash of meter reading
func HashReading(r MeterReading) [32]byte {
	b, _ := json.Marshal(r)
	return HashBytes(b)
}
