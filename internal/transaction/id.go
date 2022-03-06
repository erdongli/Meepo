package transaction

import (
	"github.com/erdongli/pbchain/internal/crypto"
	pb "github.com/erdongli/pbchain/proto"
	"google.golang.org/protobuf/proto"
)

func Id(tx *pb.Transaction) ([32]byte, error) {
	b, err := proto.Marshal(tx)
	if err != nil {
		return [32]byte{}, err
	}
	return crypto.Hash256(b), nil
}
