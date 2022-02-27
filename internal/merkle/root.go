package merkle

import (
	"github.com/erdongli/pbchain/internal/crypto"
	pb "github.com/erdongli/pbchain/proto"
	"google.golang.org/protobuf/proto"
)

// ComputeRoot computes the Merkle root in a bottom-up approach, by calculating the sub-roots via
//   sub-root := hash(left + right)
// where + denotes concatenation, and hash represents a double-SHA-256 hash.
//
// If the number of nodes at a given level is odd, the last one is duplicated before computing the next level.
//
// This implementation is susceptible to CVE-2012-2469.
func ComputeRoot(txs []*pb.Transaction) ([]byte, error) {
	if len(txs) == 0 {
		return make([]byte, 32), nil
	}

	hashes := make([][]byte, len(txs))
	for i, tx := range txs {
		txBytes, err := proto.Marshal(tx)
		if err != nil {
			return nil, err
		}
		hashes[i] = crypto.Hash256(txBytes)
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

	return hashes[0], nil
}
