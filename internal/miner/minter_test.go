package miner_test

import (
	"testing"
	"time"

	"github.com/erdongli/pbchain/internal/miner"
	pb "github.com/erdongli/pbchain/proto"
	"google.golang.org/protobuf/proto"
)

func TestMine(t *testing.T) {
	tests := []struct {
		prevBlock  []byte
		merkleRoot []byte
		bits       uint32
		want       *pb.BlockHeader
	}{
		{
			prevBlock: []byte{
				0xcf, 0xff, 0xed, 0xa9,
				0xea, 0x23, 0xde, 0x19,
				0x60, 0x52, 0xbe, 0x45,
				0x37, 0x47, 0x8b, 0x01,
				0x86, 0xc1, 0x0f, 0x4d,
				0x13, 0xea, 0x28, 0x05,
				0x95, 0xf8, 0x5a, 0x94,
				0x4f, 0x22, 0xcb, 0xc8,
			},
			merkleRoot: []byte{
				0x13, 0x32, 0xf5, 0x13,
				0x2b, 0x45, 0xde, 0x03,
				0xe7, 0x29, 0xe7, 0x64,
				0xb3, 0x8f, 0x9a, 0x2e,
				0x40, 0x57, 0xd6, 0x76,
				0x0a, 0x88, 0x07, 0x66,
				0x23, 0x31, 0xb3, 0x7d,
				0xde, 0xa9, 0x36, 0xe4,
			},
			bits: 1,
			want: &pb.BlockHeader{
				Version: 0,
				PrevBlock: []byte{
					0xcf, 0xff, 0xed, 0xa9,
					0xea, 0x23, 0xde, 0x19,
					0x60, 0x52, 0xbe, 0x45,
					0x37, 0x47, 0x8b, 0x01,
					0x86, 0xc1, 0x0f, 0x4d,
					0x13, 0xea, 0x28, 0x05,
					0x95, 0xf8, 0x5a, 0x94,
					0x4f, 0x22, 0xcb, 0xc8,
				},
				MerkleRoot: []byte{
					0x13, 0x32, 0xf5, 0x13,
					0x2b, 0x45, 0xde, 0x03,
					0xe7, 0x29, 0xe7, 0x64,
					0xb3, 0x8f, 0x9a, 0x2e,
					0x40, 0x57, 0xd6, 0x76,
					0x0a, 0x88, 0x07, 0x66,
					0x23, 0x31, 0xb3, 0x7d,
					0xde, 0xa9, 0x36, 0xe4,
				},
				Timestamp: 42,
				Bits:      1,
				Nonce:     0,
			},
		},
	}

	miner.TimeNow = func() time.Time {
		return time.Unix(42, 0)
	}

	for _, tc := range tests {
		got, err := miner.Mine(tc.prevBlock, tc.merkleRoot, tc.bits)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !proto.Equal(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
