package crypto_test

import (
	"reflect"
	"testing"

	"github.com/erdongli/pbchain/internal/crypto"
)

func TestHash256(t *testing.T) {
	tests := []struct {
		in   []byte
		want []byte
	}{
		{
			in: []byte("hello world!"),
			want: []byte{
				0x13, 0x32, 0xf5, 0x13,
				0x2b, 0x45, 0xde, 0x03,
				0xe7, 0x29, 0xe7, 0x64,
				0xb3, 0x8f, 0x9a, 0x2e,
				0x40, 0x57, 0xd6, 0x76,
				0x0a, 0x88, 0x07, 0x66,
				0x23, 0x31, 0xb3, 0x7d,
				0xde, 0xa9, 0x36, 0xe4,
			},
		},
	}

	for _, tc := range tests {
		got := crypto.Hash256(tc.in)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestHash160(t *testing.T) {
	tests := []struct {
		in   []byte
		want []byte
	}{
		{
			in: []byte("hello world!"),
			want: []byte{
				0xdf, 0xfd, 0x03, 0x13,
				0x7b, 0x3a, 0x33, 0x3d,
				0x57, 0x54, 0x81, 0x33,
				0x99, 0xa5, 0xf4, 0x37,
				0xac, 0xd6, 0x94, 0xe5,
			},
		},
	}

	for _, tc := range tests {
		got := crypto.Hash160(tc.in)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
