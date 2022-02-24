package merkle

import (
	"github.com/erdongli/meepo/internal/crypto"
)

// ComputeRoot computes the Merkle root in a bottom-up approach, by calculating the sub-roots via
//   sub-root := hash(left + right)
// where + denotes concatenation, and hash represents a double-SHA-256 hash.
//
// If the number of nodes at a given level is odd, the last one is duplicated before computing the next level.
//
// This implementation is susceptible to CVE-2012-2469.
func ComputeRoot(data [][]byte) []byte {
	if len(data) == 0 {
		return make([]byte, 32)
	}

	hashes := make([][]byte, len(data))
	for i, d := range data {
		hashes[i] = crypto.Hash256(d)
	}

	for len(hashes) > 1 {
		if len(hashes)%2 == 1 {
			hashes = append(hashes, hashes[len(hashes)-1])
		}
		for i := 0; i < len(hashes)/2; i++ {
			hashes[i] = append(hashes[2*i], hashes[2*i+1]...)
			hashes[i] = crypto.Hash256(hashes[i])
		}
		hashes = hashes[:len(hashes)/2]
	}

	return hashes[0]
}
