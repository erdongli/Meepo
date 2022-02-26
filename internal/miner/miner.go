package miner

import (
	"time"

	"github.com/erdongli/meepo/internal/crypto"
	"github.com/erdongli/meepo/internal/pow"
	pb "github.com/erdongli/meepo/proto"
	"google.golang.org/protobuf/proto"
)

var TimeNow = time.Now

// Mine takes the hash of the previous block header, the Merkle root, and the targit difficulty, and returns a block
// header that fulfills the Proof-of-Work requirement.
func Mine(prevBlock, merkleRoot []byte, bits uint32) (*pb.BlockHeader, error) {
	hdr := &pb.BlockHeader{
		Version:    0,
		PrevBlock:  prevBlock,
		MerkleRoot: merkleRoot,
		Timestamp:  uint32(TimeNow().Unix()),
		Bits:       bits,
		Nonce:      0,
	}

	v := pow.NewValidator(bits)
	for {
		b, err := proto.Marshal(hdr)
		if err != nil {
			return nil, err
		}

		if v.Validate(crypto.Hash256(b)) {
			return hdr, nil
		}

		hdr.Nonce++
	}
}
