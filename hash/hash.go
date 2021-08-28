// Package hash implements functions for hashing objects.
package hash

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
)

// Compute the SHA hash of a struct (256 bit is default)
func HashStruct(v interface{}) string {
	return HashStruct256(v)
}

// Compute the SHA hash of a string (default is 256)
func HashString(s string) string {
	return HashString256(s)
}

func HashStruct512(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic("Marshaling an object into json failed!")
	}
	r := sha512.Sum512(b)
	return hex.EncodeToString(r[:])
}

func HashStruct256(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic("Marshaling an object into json failed!")
	}
	r := sha256.Sum256(b)
	return hex.EncodeToString(r[:])
}

func HashStruct128(v interface{}) string {
	return HashStruct256(v)[0:32]
}

func HashStruct64(v interface{}) string {
	return HashStruct256(v)[0:16]
}

func HashStruct32(v interface{}) string {
	return HashStruct256(v)[0:8]
}

func HashString512(s string) string {
	r := sha512.Sum512([]byte(s))
	return hex.EncodeToString(r[:])
}

func HashString256(s string) string {
	r := sha256.Sum256([]byte(s))
	return hex.EncodeToString(r[:])
}

func HashString128(s string) string {
	return HashString256(s)[0:32]
}

func HashString64(s string) string {
	return HashString256(s)[0:16]
}

func HashString32(s string) string {
	return HashString256(s)[0:8]
}
