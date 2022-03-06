package crypto

import (
	"crypto/sha256"

	"golang.org/x/crypto/ripemd160"
)

// Hash256 generates a SHA-256(SHA-256(in)) hash. This is the most commonly used hashing method.
// See https://en.bitcoin.it/wiki/Protocol_documentation#Hashes
func Hash256(in []byte) [32]byte {
	h := sha256.Sum256(in)
	h = sha256.Sum256(h[:])
	return h
}

// Hash160 generates a RIPEMD-160(SHA-256(in)) hash. This is used when a shorter hash is desirable.
// See https://en.bitcoin.it/wiki/Protocol_documentation#Hashes
func Hash160(in []byte) [20]byte {
	h := ripemd160.New()
	h.Write(in)
	hash := h.Sum(nil)
	return *(*[20]byte)(hash)
}
