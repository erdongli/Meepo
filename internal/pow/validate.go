package pow

import "math/big"

// Validator provides methods to validate if a given hash fulfills the Proof-of-Work (Hashcash) requirement.
type Validator struct {
	target *big.Int
}

// NewValidator creates a new Proof-of-Work (Hashcash) validator, with a requirement that the target hash contains at
// least bits leading zeros.
// It uses Golang's native big number implementation, which assumes a big-endian byte slice.
func NewValidator(bits uint32) *Validator {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-bits))
	return &Validator{
		target: target,
	}
}

// Validate checks if hash fulfills the target.
func (v *Validator) Validate(hash [32]byte) bool {
	actual := new(big.Int)
	actual.SetBytes(hash[:])
	return actual.Cmp(v.target) != 1
}
